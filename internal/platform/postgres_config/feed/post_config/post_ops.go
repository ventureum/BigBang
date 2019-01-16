package post_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
	"time"
)

type PostExecutor struct {
	client_config.PostgresBigBangClient
}

func (postExecutor *PostExecutor) CreatePostTable() {
	postExecutor.CreateTimestampTrigger()
	postExecutor.CreateTable(TABLE_SCHEMA_FOR_POST, TABLE_NAME_FOR_POST)
	postExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST)
}

func (postExecutor *PostExecutor) DeletePostTable() {
	postExecutor.DeleteTable(TABLE_NAME_FOR_POST)
}

func (postExecutor *PostExecutor) UpsertPostRecordTx(postRecord *PostRecord) time.Time {
	res, err := postExecutor.Tx.NamedQuery(UPSERT_POST_COMMAND, postRecord)
	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", postRecord.PostHash, error_config.PostRecordLocation)
		log.Printf("Failed to upsert post record: %+v with error: %+v\n", postRecord, err)
		log.Panicln(errInfo.Marshal())
	}

	log.Printf("Sucessfully upserted post record for postHash %s\n", postRecord.PostHash)

	var createdTime time.Time
	for res.Next() {
		res.Scan(&createdTime)
	}
	return createdTime
}

func (postExecutor *PostExecutor) DeletePostRecordTx(postHash string) {
	_, err := postExecutor.Tx.Exec(DELETE_POST_COMMAND, postHash)
	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRecordLocation)
		log.Printf("Failed to delete post record for postHash %s with error: %+v\n", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted post record for postHash %s\n", postHash)
}

func (postExecutor *PostExecutor) GetPostRecordTx(postHash string) *PostRecord {
	var postRecord PostRecord
	err := postExecutor.Tx.Get(&postRecord, QUERY_POST_COMMAND, postHash)

	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRecordLocation)
		log.Printf("Failed to get post record for postHash %s with error: %+v\n", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	return &postRecord
}

func (postExecutor *PostExecutor) GetPostUpdateCountTx(postHash string) int64 {
	var updateCount sql.NullInt64
	err := postExecutor.Tx.Get(&updateCount, QUERY_POST_UPDATE_COUNT_COMMAND, postHash)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRecordLocation)
		log.Printf("Failed to get post update count for postHash %s with error: %+v\n", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	return updateCount.Int64
}

func (postExecutor *PostExecutor) VerifyPostRecordExistingTx(postHash string) {
	var existing bool
	err := postExecutor.Tx.Get(&existing, VERIFY_POSTHASH_EXISTING_COMMAND, postHash)
	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRecordLocation)
		log.Printf("Failed to verify Post Record existing for postHash %s with error: %+v\n", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoPostHashExisting,
			ErrorData: map[string]interface{}{
				"postHash": postHash,
			},
			ErrorLocation: error_config.PostRecordLocation,
		}
		log.Printf("No Post Record for postHash %s", postHash)
		log.Panicln(errorInfo.Marshal())
	}
}

func (postExecutor *PostExecutor) GetPostTypeTx(postHash string) feed_attributes.PostType {
	var postType feed_attributes.PostType
	err := postExecutor.Tx.Get(&postType, QUERY_POST_TYPE_COMMAND, postHash)

	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRecordLocation)
		log.Printf("Failed to get post type for postHash %s with error: %+v\n", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	return postType
}
