package post_reputations_record_config

import (
  "log"
  "database/sql"
  "time"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/app/feed_attributes"
)


type PostReputationsRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) CreatePostReputationsRecordTable( ) {
  postReputationsRecordExecutor.CreateTimestampTrigger()
  postReputationsRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_POST_REPUTATIONS_RECORD, TABLE_NAME_FOR_POST_REPUTATIONS_RECORD)
  postReputationsRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST_REPUTATIONS_RECORD)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) DeletePostReputationsRecordTable() {
  postReputationsRecordExecutor.DeleteTable(TABLE_NAME_FOR_POST_REPUTATIONS_RECORD)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) UpsertPostReputationsRecord (
    postReputationsRecord *PostReputationsRecord) *PostReputationsRecord {
  res, err := postReputationsRecordExecutor.C.NamedQuery(UPSERT_POST_REPUTATIONS_RECORD_COMMAND, postReputationsRecord)

  if err != nil {
    log.Panicf("Failed to upsert post reputations record: %+v with error:\n %+v", postReputationsRecord, err.Error())
  }

  log.Printf("Sucessfully upserted post reputations record for post_hash %s and actor %s with reputaions %d\n",
    postReputationsRecord.PostHash, postReputationsRecord.Actor, postReputationsRecord.Reputations)

  upsertedPostReputationsRecord := PostReputationsRecord{}
  for res.Next() {
    err = res.StructScan(&upsertedPostReputationsRecord)
    if err != nil {
      log.Panicf("Failed to scan upserted post reputations for post_hash %s and actor %s with error: %v\n",
        postReputationsRecord.PostHash, postReputationsRecord.Actor, err)
    }
  }

  return &upsertedPostReputationsRecord
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) DeletePostReputationsRecordsByPostHash(postHash string) {
  _, err := postReputationsRecordExecutor.C.Exec(DELETE_POST_REPUTATIONS_RECORDS_BY_POST_HASH_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete post reputations records for postHash %s with error:\n %+v", postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post reputations records for postHash %s\n", postHash)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) DeletePostReputationsRecordsByActor(actor string) {
  _, err := postReputationsRecordExecutor.C.Exec(DELETE_POST_REPUTATIONS_RECORDS_BY_ACTOR_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete post reputations records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted post reputations records for actor %s\n", actor)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) DeletePostReputationsRecordsByPostHashAndActor(
  postHash string, actor string) {
  _, err := postReputationsRecordExecutor.C.Exec(DELETE_POST_REPUTATIONS_RECORDS_BY_POST_HASH_AND_ACTOR_COMMAND,
    postHash, actor)
  if err != nil {
    log.Panicf("Failed to delete post reputations records for postHash %s and actor %s with error:\n %+v",
      postHash, actor, err.Error())
  }
  log.Printf("Sucessfully deleted post reputations records for postHash %s and actor %s\n", postHash, actor)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetTotalReputationsByPostHash(
  postHash string) feed_attributes.Reputation {
  var totalReputations sql.NullInt64
  err := postReputationsRecordExecutor.C.Get(&totalReputations, QUERY_TOTAL_REPUTATIONS_BY_POST_HASH_COMMAND, postHash)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total reputations for postHash %s with error:\n %+v", postHash, err.Error())
  }
  return feed_attributes.Reputation(totalReputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetReputationsByPostHashAndActor(
    postHash string, actor string) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := postReputationsRecordExecutor.C.Get(&reputations, QUERY_REPUTATIONS_BY_POST_HASH_AND_ACTOR_COMMAND,
    postHash, actor)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get reputations for postHash %s and actor %s with error:\n %+v",
      postHash, actor, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetReputationsByPostHashAndVoteType(
    postHash string, voteType feed_attributes.VoteType) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := postReputationsRecordExecutor.C.Get(&reputations, QUERY_REPUTATIONS_BY_POST_HASH_AND_VOTE_TYPE_COMMAND,
    postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total reputations for postHash %s and voteType %s with error:\n %+v",
      postHash, voteType, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetReputationsByPostHashAndActorWithLatestVoteTypeAndTimeCutOff(
    postHash string, actor string, voteType feed_attributes.VoteType, time time.Time) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := postReputationsRecordExecutor.Tx.Get(&reputations, QUERY_REPUTATIONS_BY_POST_HASH_AND_ACTOR_WITH_LATEST_VOTE_TYPE_AND_TIME_CUTOFF_COMMAND,
    postHash, actor, voteType, time)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get reputations for postHash %s, actor %s, lastestVoteType %s and cutOffTime %s with error:\n %+v",
      postHash, actor, voteType, time, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetTotalVotesCountByPostHashAndActorType(
    postHash string, actor string) int64 {
  var postVotes sql.NullInt64
  err := postReputationsRecordExecutor.C.Get(&postVotes, QUERY_TOTAL_VOTES_COUNT_BY_POST_HASH_AND_ACTOR_COMMAND, postHash, actor)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total votes count for postHash %s and actor %s with error:\n %+v",
      postHash, actor, err)
  }
  return postVotes.Int64
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetActorListByPostHashAndVoteType(
    postHash string, voteType feed_attributes.VoteType) *[]string {
  var actorList []string
  err := postReputationsRecordExecutor.C.Select(&actorList, QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND, postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get actor list for postHash %s and voteType %s with error:\n %+v", postHash,  voteType, err)
  }
  return &actorList
}


/*
 * Tx versions
 */
func (postReputationsRecordExecutor *PostReputationsRecordExecutor) UpsertPostReputationsRecordTx (
    postReputationsRecord *PostReputationsRecord) *PostReputationsRecord {
  res, err := postReputationsRecordExecutor.Tx.NamedQuery(UPSERT_POST_REPUTATIONS_RECORD_COMMAND, postReputationsRecord)

  if err != nil {
    log.Panicf("Failed to upsert post reputations record: %+v with error: %+v\n", postReputationsRecord, err)
  }

  log.Printf("Sucessfully upserted post reputations record for post_hash %s and actor %s with reputaions %d\n",
    postReputationsRecord.PostHash, postReputationsRecord.Actor, postReputationsRecord.Reputations)

  upsertedPostReputationsRecord := PostReputationsRecord{}
  for res.Next() {
    err = res.StructScan(&upsertedPostReputationsRecord)
    if err != nil {
      log.Panicf("Failed to scan upserted post reputations for post_hash %s and actor %s with error: %v\n",
        postReputationsRecord.PostHash, postReputationsRecord.Actor, err)
    }
  }

  return &upsertedPostReputationsRecord
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) DeletePostReputationsRecordsByActorTx(actor string) {
  _, err := postReputationsRecordExecutor.Tx.Exec(DELETE_POST_REPUTATIONS_RECORDS_BY_ACTOR_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete post reputations records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted post reputations records for actor %s\n", actor)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) DeletePostReputationsRecordsByPostHashAndActorTx(
    postHash string, actor string) {
  _, err := postReputationsRecordExecutor.Tx.Exec(DELETE_POST_REPUTATIONS_RECORDS_BY_POST_HASH_AND_ACTOR_COMMAND,
    postHash, actor)
  if err != nil {
    log.Panicf("Failed to delete post reputations records for postHash %s and actor %s with error:\n %+v",
      postHash, actor, err.Error())
  }
  log.Printf("Sucessfully deleted post reputations records for postHash %s and actor %s\n", postHash, actor)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetTotalReputationsByPostHashTx(
    postHash string) feed_attributes.Reputation {
  var totalReputations sql.NullInt64
  err := postReputationsRecordExecutor.Tx.Get(&totalReputations, QUERY_TOTAL_REPUTATIONS_BY_POST_HASH_COMMAND, postHash)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total reputations for postHash %s with error:\n %+v", postHash, err.Error())
  }
  return feed_attributes.Reputation(totalReputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetReputationsByPostHashAndActorTx(
    postHash string, actor string) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := postReputationsRecordExecutor.Tx.Get(&reputations, QUERY_REPUTATIONS_BY_POST_HASH_AND_ACTOR_COMMAND,
    postHash, actor)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get reputations for postHash %s and actor %s with error:\n %+v",
      postHash, actor, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetReputationsByPostHashAndVoteTypeTx(
    postHash string, voteType feed_attributes.VoteType) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := postReputationsRecordExecutor.Tx.Get(&reputations, QUERY_REPUTATIONS_BY_POST_HASH_AND_VOTE_TYPE_COMMAND,
    postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total reputations for postHash %s and voteType %s with error:\n %+v",
      postHash, voteType, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetReputationsByPostHashAndActorWithLatestVoteTypeAndTimeCutOffTx(
    postHash string, actor string, voteType feed_attributes.VoteType, time time.Time) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := postReputationsRecordExecutor.Tx.Get(&reputations, QUERY_REPUTATIONS_BY_POST_HASH_AND_ACTOR_WITH_LATEST_VOTE_TYPE_AND_TIME_CUTOFF_COMMAND,
    postHash, actor, voteType, time)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get reputations for postHash %s, actor %s, lastestVoteType %s and cutOffTime %s with error:\n %+v",
      postHash, actor, voteType, time, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetTotalVotesCountByPostHashAndActorTypeTx(
    postHash string, actor string) int64 {
  var postVotes sql.NullInt64
  err := postReputationsRecordExecutor.Tx.Get(&postVotes, QUERY_TOTAL_VOTES_COUNT_BY_POST_HASH_AND_ACTOR_COMMAND, postHash, actor)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total votes count for postHash %s and actor %s with error:\n %+v",
       postHash, actor, err)
  }
  return postVotes.Int64
}

func (postReputationsRecordExecutor *PostReputationsRecordExecutor) GetActorListByPostHashAndVoteTypeTx(
    postHash string, voteType feed_attributes.VoteType) *[]string {
  var actorList []string
  err := postReputationsRecordExecutor.Tx.Select(&actorList, QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND, postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get actor list for postHash %s and voteType %s with error:\n %+v", postHash,  voteType, err)
  }
  return &actorList
}
