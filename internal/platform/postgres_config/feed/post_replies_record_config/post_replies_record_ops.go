package post_replies_record_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type PostRepliesRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (postRepliesRecordExecutor *PostRepliesRecordExecutor) CreatePostRepliesRecordTable() {
	postRepliesRecordExecutor.CreateTimestampTrigger()
	postRepliesRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_POST_REPLIES_RECORD, TABLE_NAME_FOR_POST_REPLIES_RECORD)
	postRepliesRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST_REPLIES_RECORD)
}

func (postRepliesRecordExecutor *PostRepliesRecordExecutor) DeletePostRepliesRecordTable() {
	postRepliesRecordExecutor.DeleteTable(TABLE_NAME_FOR_POST_REPLIES_RECORD)
}

func (postRepliesRecordExecutor *PostRepliesRecordExecutor) UpsertPostRepliesRecordTx(postRepliesRecord *PostRepliesRecord) {
	_, err := postRepliesRecordExecutor.Tx.NamedExec(UPSERT_POST_REPLIES_RECORD_COMMAND, postRepliesRecord)
	if err != nil {
		errorInfo := error_config.MatchError(err, "postHash", postRepliesRecord.PostHash, error_config.PostRepliesRecordLocation)
		errorInfo.AddErrorData("replayHash", postRepliesRecord.ReplyHash)
		log.Printf("Failed to upsert post replies record: %+v with error: %+v\n", postRepliesRecord, err)
		log.Panic(errorInfo.Marshal())
	}
	log.Printf("Sucessfully upserted post replies record for postHash %s\n", postRepliesRecord.PostHash)
}

func (postRepliesRecordExecutor *PostRepliesRecordExecutor) DeletePostRepliesRecordsTx(postHash string) {
	_, err := postRepliesRecordExecutor.Tx.Exec(DELETE_POST_REPLIES_RECORD_COMMAND, postHash)
	if err != nil {
		errorInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRepliesRecordLocation)
		log.Printf("Failed to delete post replies records for postHash %s with error: %+v\n", postHash, err)
		log.Panic(errorInfo.Marshal())
	}
	log.Printf("Sucessfully deleted post replies records for postHash %s\n", postHash)
}

func (postRepliesRecordExecutor *PostRepliesRecordExecutor) GetPostRepliesTx(postHash string) *[]string {
	var postReplies []string
	err := postRepliesRecordExecutor.Tx.Select(&postReplies, QUERY_POST_REPLIES_COMMAND, postHash)
	if err != nil {
		errorInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRepliesRecordLocation)
		log.Printf("Failed to get post repliers for postHash %s with error: %+v\n", postHash, err)
		log.Panic(errorInfo.Marshal())
	}
	return &postReplies
}

func (postRepliesRecordExecutor *PostRepliesRecordExecutor) GetPostRepliesRecordCountTx(postHash string) int64 {
	var updateCount sql.NullInt64
	err := postRepliesRecordExecutor.Tx.Get(&updateCount, QUERY_POST_REPLIES_COUNT_COMMAND, postHash)
	if err != nil && err != sql.ErrNoRows {
		errorInfo := error_config.MatchError(err, "postHash", postHash, error_config.PostRepliesRecordLocation)
		log.Printf("Failed to read post repliers count for postHash %s with error: %+v\n", postHash, err)
		log.Panic(errorInfo.Marshal())
	}
	return updateCount.Int64
}
