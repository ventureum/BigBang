package post_rewards_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type PostRewardsRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) CreatePostRewardsRecordTable() {
	postRewardsRecordExecutor.CreateTimestampTrigger()
	postRewardsRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_POST_REWARDS_RECORD, TABLE_NAME_FOR_POST_REWARDS_RECORD)
	postRewardsRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST_REWARDS_RECORD)
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) DeletePostRewardsRecordTable() {
	postRewardsRecordExecutor.DeleteTable(TABLE_NAME_FOR_POST_REWARDS_RECORD)
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) UpsertPostRewardsRecordTx(postRewardsRecord *PostRewardsRecord) {
	_, err := postRewardsRecordExecutor.Tx.NamedExec(UPSERT_POST_REWARDS_RECORD_COMMAND, postRewardsRecord)
	if err != nil {
		log.Panicf("Failed to upsert post rewards record: %+v with error: %+v\n", postRewardsRecord, err)
	}
	log.Printf("Sucessfully upserted post rewards record for postHash %s\n", postRewardsRecord.PostHash)
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) DeletePostRewardsRecordsTx(postHash string) {
	_, err := postRewardsRecordExecutor.Tx.Exec(DELETE_POST_REWARDS_RECORD_COMMAND, postHash)
	if err != nil {
		log.Panicf("Failed to delete post rewards records for postHash %s with error: %+v\n", postHash, err)
	}
	log.Printf("Sucessfully deleted post rewards records for postHash %s\n", postHash)
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) UpdatePostRewardsRecordsByAggregationsTx() *[]PostRewardsForUpdate {
	var postRewardsForUpdate []PostRewardsForUpdate
	err := postRewardsRecordExecutor.Tx.Select(&postRewardsForUpdate, UPSERT_POST_REWARDS_RECORD_BY_AGGREGATION_COMMAND)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to update post rewards records by aggregations with error: %+v\n", err)
	}
	return &postRewardsForUpdate
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) GetPostRewardsRecordByPostHashTx(
	postHash string) *PostRewardsRecord {
	var postRewardsRecord PostRewardsRecord
	err := postRewardsRecordExecutor.Tx.Get(
		&postRewardsRecord, QUERY_POST_REWARDS_RECORD_COMMAND, postHash)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf("Failed to query post rewards record by postHash %s with error: %+v\n", postHash, err)
	}

	return &postRewardsRecord
}

func (postRewardsRecordExecutor *PostRewardsRecordExecutor) GetRecentPostRewardsRecordsByActorTx(
	actor string, postType feed_attributes.PostType, limit int64) *[]PostRewardsRecord {
	var postRewardsRecords []PostRewardsRecord
	err := postRewardsRecordExecutor.Tx.Select(
		&postRewardsRecords, QUERY_RECENT_POST_REWARDS_RECORDS_BY_ACTOR_COMMAND, actor, postType, limit)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get recent %d post rewards records for actor %s and postType %s with error:\n %+v", limit, actor, postType, err)
	}
	return &postRewardsRecords
}
