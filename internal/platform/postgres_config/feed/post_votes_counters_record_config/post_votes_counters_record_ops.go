package post_votes_counters_record_config

import (
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type PostVotesCountersRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (postVotesCountersRecordExecutor *PostVotesCountersRecordExecutor) CreatePostVotesCountersRecordTable() {
	postVotesCountersRecordExecutor.CreateTimestampTrigger()
	postVotesCountersRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_POST_VOTES_COUNTERS_RECORDS, TABLE_NAME_FOR_POST_VOTES_COUNTERS_RECORD)
	postVotesCountersRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST_VOTES_COUNTERS_RECORD)
}

func (postVotesCountersRecordExecutor *PostVotesCountersRecordExecutor) DeletePostVotesCountersRecordTable() {
	postVotesCountersRecordExecutor.DeleteTable(TABLE_NAME_FOR_POST_VOTES_COUNTERS_RECORD)
}

func (postVotesCountersRecordExecutor *PostVotesCountersRecordExecutor) UpsertPostVotesCountersRecordTx(
	postVotesCountersRecord *PostVotesCountersRecord) *PostVotesCountersRecord {
	res, err := postVotesCountersRecordExecutor.Tx.NamedQuery(UPSERT_POST_VOTES_COUNTRS_RECORD_COMMAND, postVotesCountersRecord)

	if err != nil {
		log.Panicf("Failed to upsert post votes counters record: %+v with error:\n %+v", postVotesCountersRecord)
	}

	log.Printf("Sucessfully upserted post votes counters record for post_hash %s", postVotesCountersRecord.PostHash)

	upsertedPostVotesCountersRecord := PostVotesCountersRecord{}
	for res.Next() {
		err = res.StructScan(&upsertedPostVotesCountersRecord)
		if err != nil {
			log.Panicf("Failed to scan upserted post reputations for post_hash %s with error: %v\n",
				postVotesCountersRecord.PostHash, err)
		}
	}

	return &upsertedPostVotesCountersRecord
}

func (postVotesCountersRecordExecutor *PostVotesCountersRecordExecutor) DeletePostVotesCountersRecordsByPostHashTx(postHash string) {
	_, err := postVotesCountersRecordExecutor.Tx.Exec(DELETE_POST_VOTES_COUNTRS_RECORDS_BY_POST_HASH_COMMAND, postHash)
	if err != nil {
		log.Panicf("Failed to delete post votes counters records for postHash %s with error: %+v\n", postHash, err)
	}
	log.Printf("Sucessfully deleted post votes counters records for postHash %s\n", postHash)
}

func (postVotesCountersRecordExecutor *PostVotesCountersRecordExecutor) GetPostVotesCountersRecordByPostHashTx(
	postHash string) *PostVotesCountersRecord {
	var postVotesCountersRecord PostVotesCountersRecord
	err := postVotesCountersRecordExecutor.Tx.Get(
		&postVotesCountersRecord, QUERY_POST_VOTES_COUNTRS_RECORDS_BY_POST_HASH_COMMAND, postHash)

	if err != nil && err != sql.ErrNoRows {
		log.Panicf("Failed to query post votes counters record by postHash %s with error: %+v\n", postHash, err)
	}

	return &postVotesCountersRecord
}

func (postVotesCountersRecordExecutor *PostVotesCountersRecordExecutor) GetPostVotesCountersRecordByPostHashForUpdateTx(
	postHash string) *PostVotesCountersRecord {
	var postVotesCountersRecord PostVotesCountersRecord
	err := postVotesCountersRecordExecutor.Tx.Get(
		&postVotesCountersRecord, QUERY_POST_VOTES_COUNTRS_RECORDS_BY_POST_HASH_FOR_UPDATE_COMMAND, postHash)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf("Failed to query post votes counter record for update by postHash %s with error: %+v\n", postHash, err)
	}

	return &postVotesCountersRecord
}
