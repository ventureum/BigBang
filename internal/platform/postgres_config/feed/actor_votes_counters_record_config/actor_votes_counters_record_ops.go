package actor_votes_counters_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
	"time"
)

type ActorVotesCountersRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) CreateActorVotesCountersRecordTable() {
	actorVotesCountersRecordExecutor.CreateTimestampTrigger()
	actorVotesCountersRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_ACTOR_VOTES_COUNTERS_RECORD, TABLE_NAME_FOR_ACTOR_VOTES_COUNTERS_RECORD)
	actorVotesCountersRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_VOTES_COUNTERS_RECORD)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) DeleteActorVotesCountersRecordTable() {
	actorVotesCountersRecordExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_VOTES_COUNTERS_RECORD)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) UpsertActorVotesCountersRecordTx(
	postReputationsRecord *ActorVotesCountersRecord) *ActorVotesCountersRecord {
	res, err := actorVotesCountersRecordExecutor.Tx.NamedQuery(UPSERT_ACTOR_VOTES_COUNTERS_RECORD_COMMAND, postReputationsRecord)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", postReputationsRecord.Actor, error_config.ActorVotesCountersRecordLocation)
		errorInfo.ErrorData["postHash"] = postReputationsRecord.PostHash
		log.Printf("Failed to upsert actor votes counters record: %+v with error: %+v\n", postReputationsRecord, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Sucessfully upserted actor votes counters record for post_hash %s and actor %s\n",
		postReputationsRecord.PostHash, postReputationsRecord.Actor)

	upsertedActorVotesCountersRecord := ActorVotesCountersRecord{}
	for res.Next() {
		err = res.StructScan(&upsertedActorVotesCountersRecord)
		if err != nil {
			log.Panicf("Failed to scan upserted actor votes counters for post_hash %s and actor %s with error: %v\n",
				postReputationsRecord.PostHash, postReputationsRecord.Actor, err)
		}
	}

	return &upsertedActorVotesCountersRecord
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) DeleteActorVotesCountersRecordsByActorTx(actor string) {
	_, err := actorVotesCountersRecordExecutor.Tx.Exec(DELETE_ACTOR_VOTES_COUNTERS_RECORDS_BY_ACTOR_COMMAND, actor)
	if err != nil {
		log.Panicf("Failed to delete actor votes counters records for actor %s with error: %+v\n", actor, err)
	}
	log.Printf("Sucessfully deleted actor votes counters records for actor %s\n", actor)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) DeleteActorVotesCountersRecordsByPostHashAndActorTx(
	postHash string, actor string) {
	_, err := actorVotesCountersRecordExecutor.Tx.Exec(DELETE_ACTOR_VOTES_COUNTERS_RECORDS_BY_POST_HASH_AND_ACTOR_COMMAND,
		postHash, actor)
	if err != nil {
		log.Panicf("Failed to delete actor votes counters records for postHash %s and actor %s with error: %+v\n",
			postHash, actor, err)
	}
	log.Printf("Sucessfully deleted actor votes counters records for postHash %s and actor %s\n", postHash, actor)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetTotalReputationByPostHashTx(
	postHash string) feed_attributes.Reputation {
	var totalReputation sql.NullInt64
	err := actorVotesCountersRecordExecutor.Tx.Get(&totalReputation, QUERY_TOTAL_REPUTATION_BY_POST_HASH_COMMAND, postHash)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get total reputation for postHash %s with error: %+v\n", postHash, err)
	}
	return feed_attributes.Reputation(totalReputation.Int64)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetReputationByPostHashAndActorTx(
	postHash string, actor string) feed_attributes.Reputation {
	var reputation sql.NullInt64
	err := actorVotesCountersRecordExecutor.Tx.Get(&reputation, QUERY_REPUTATION_BY_POST_HASH_AND_ACTOR_COMMAND,
		postHash, actor)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get reputation for postHash %s and actor %s with error: %+v\n",
			postHash, actor, err)
	}
	return feed_attributes.Reputation(reputation.Int64)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetReputationByPostHashAndVoteTypeTx(
	postHash string, voteType feed_attributes.VoteType) feed_attributes.Reputation {
	var reputation sql.NullInt64
	err := actorVotesCountersRecordExecutor.Tx.Get(&reputation, QUERY_REPUTATION_BY_POST_HASH_AND_VOTE_TYPE_COMMAND,
		postHash, voteType)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get total reputation for postHash %s and voteType %s with error: %+v\n",
			postHash, voteType, err)
	}
	return feed_attributes.Reputation(reputation.Int64)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetReputationByPostHashAndActorWithLatestVoteTypeAndTimeCutOffTx(
	postHash string, actor string, voteType feed_attributes.VoteType, time time.Time) feed_attributes.Reputation {
	var reputation sql.NullInt64
	err := actorVotesCountersRecordExecutor.Tx.Get(&reputation, QUERY_REPUTATION_BY_POST_HASH_AND_ACTOR_WITH_LATEST_VOTE_TYPE_AND_TIME_CUTOFF_COMMAND,
		postHash, actor, voteType, time)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get reputation for postHash %s, actor %s, lastestVoteType %s and cutOffTime %s with error: %+v\n",
			postHash, actor, voteType, time, err)
	}
	return feed_attributes.Reputation(reputation.Int64)
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetTotalVotesCountByPostHashAndActorTypeTx(
	postHash string, actor string) int64 {
	var postVotes sql.NullInt64
	err := actorVotesCountersRecordExecutor.Tx.Get(&postVotes, QUERY_TOTAL_VOTES_COUNT_BY_POST_HASH_AND_ACTOR_COMMAND, postHash, actor)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get total votes count for postHash %s and actor %s with error: %+v\n",
			postHash, actor, err)
	}
	return postVotes.Int64
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetActorListByPostHashAndVoteTypeTx(
	postHash string, voteType feed_attributes.VoteType) *[]string {
	var actorList []string
	err := actorVotesCountersRecordExecutor.Tx.Select(&actorList, QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND, postHash, voteType)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get actor list for postHash %s and voteType %s with error: %+v\n", postHash, voteType, err)
	}
	return &actorList
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetActorVotesCountersRecordByPostHashAndActorTx(
	postHash string, actor string) *ActorVotesCountersRecord {
	var postReputationsRecord ActorVotesCountersRecord
	err := actorVotesCountersRecordExecutor.Tx.Get(
		&postReputationsRecord, QUERY_ACTOR_VOTES_COUNTERS_RECORD_BY_POST_HASH_AND_ACTOR_COMMAND, postHash, actor)

	if err != nil && err != sql.ErrNoRows {
		log.Panicf("Failed to query actor votes counters record by postHash %s and actor %s with error: %+v\n",
			postHash, actor, err)
	}

	return &postReputationsRecord
}

func (actorVotesCountersRecordExecutor *ActorVotesCountersRecordExecutor) GetTotalActorReputationByPostHashTx(postHash string) int64 {
	var totalReputations int64
	err := actorVotesCountersRecordExecutor.Tx.Get(&totalReputations, QUARY_TOTAL_REPUTATION_BY_POSTHASH_COMMAND, postHash)
	if err != nil {
		errorInfo := error_config.MatchError(err, "", "", error_config.ActorVotesCountersRecordLocation)
		log.Printf("Failed to get total reputation for postHash %s with error: %+v\n", postHash, err)
		log.Panic(errorInfo.Marshal())
	}
	return totalReputations
}
