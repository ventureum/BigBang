package post_votes_record_config

import (
  "log"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/app/feed_attributes"
)


type PostVotesRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) CreatePostVotesRecordTable() {
  postVotesRecordExecutor.CreateTimestampTrigger()
  postVotesRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_POST_VOTES_RECORD, TABLE_NAME_FOR_POST_VOTES_RECORD)
  postVotesRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST_VOTES_RECORD)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordTable() {
  postVotesRecordExecutor.DeleteTable(TABLE_NAME_FOR_POST_VOTES_RECORD)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) UpsertPostVotesRecord(postVotesRecord *PostVotesRecord) {
  _, err := postVotesRecordExecutor.C.NamedExec(UPSERT_POST_VOTES_RECORD_COMMAND, postVotesRecord)
  if err != nil {
    log.Panicf("Failed to upsert post votes record: %+v with error:\n %+v", postVotesRecord, err.Error())
  }
  log.Printf("Sucessfully upserted post votes record for actor %s and post_hash %s with vote_type %s\n",
    postVotesRecord.Actor, postVotesRecord.PostHash, postVotesRecord.VoteType)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordsByPostHash(postHash string) {
  _, err := postVotesRecordExecutor.C.Exec(DELETE_POST_VOTES_RECORDS_BY_POST_HASH_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete post votes records for postHash %s with error:\n %+v", postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post votes records for postHash %s\n", postHash)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordsByActor(actor string) {
  _, err := postVotesRecordExecutor.C.Exec(DELETE_POST_VOTES_RECORDS_BY_ACTOR_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete post votes records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted post votes records for actor %s\n", actor)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordsByActorAndPostHash(
    actor string, postHash string) {
  _, err := postVotesRecordExecutor.C.Exec(DELETE_POST_VOTES_RECORDS_BY_ACTOR_AND_POST_HASH_COMMAND, actor, postHash)
  if err != nil {
    log.Panicf("Failed to delete post votes records for actor %s and postHash %s with error:\n %+v",
      actor, postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post votes records for actor %s and postHash %s\n", actor, postHash)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) GetTotalPostVotesCount(
    actor string, postHash string) int64 {
  var totalPostVotes sql.NullInt64
  err := postVotesRecordExecutor.C.Get(&totalPostVotes, QUERY_TOTAL_VOTES_COUNT_COMMAND, actor, postHash)

  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total post votes for actor %s and postHash %s with error:\n %+v", actor, postHash, err.Error())
  }
  return totalPostVotes.Int64
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) GetVotesCountByVoteType(
    actor string, postHash string, voteType feed_attributes.VoteType) int64 {
  var postVotes sql.NullInt64
  err := postVotesRecordExecutor.C.Get(&postVotes, QUERY_VOTES_COUNT_BY_VOTE_TYPE_COMMAND, actor, postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get votes count for actor %s, postHash %s and voteType %s with error:\n %+v",
      actor, postHash,  voteType, err)
  }
  return postVotes.Int64
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) GetActorListByPostHashAndVoteType(
    postHash string, voteType feed_attributes.VoteType) *[]string {
  var actorList []string
  err := postVotesRecordExecutor.C.Select(&actorList, QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND, postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get actor list for postHash %s and voteType %s with error:\n %+v", postHash,  voteType, err)
  }
  return &actorList
}

/*
 * Tx versions
 */
func (postVotesRecordExecutor *PostVotesRecordExecutor) UpsertPostVotesRecordTx(postVotesRecord *PostVotesRecord) {
  _, err := postVotesRecordExecutor.Tx.NamedExec(UPSERT_POST_VOTES_RECORD_COMMAND, postVotesRecord)
  if err != nil {
    log.Panicf("Failed to upsert post votes record: %+v with error:\n %+v", postVotesRecord, err.Error())
  }
  log.Printf("Sucessfully upserted post votes record for actor %s and post_hash %s with vote_type %s\n",
    postVotesRecord.Actor, postVotesRecord.PostHash, postVotesRecord.VoteType)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordsByPostHashTx(postHash string) {
  _, err := postVotesRecordExecutor.Tx.Exec(DELETE_POST_VOTES_RECORDS_BY_POST_HASH_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete post votes records for postHash %s with error:\n %+v", postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post votes records for postHash %s\n", postHash)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordsByActorTx(actor string) {
  _, err := postVotesRecordExecutor.Tx.Exec(DELETE_POST_VOTES_RECORDS_BY_ACTOR_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete post votes records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted post votes records for actor %s\n", actor)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) DeletePostVotesRecordsByActorAndPostHashTx(
    actor string, postHash string) {
  _, err := postVotesRecordExecutor.Tx.Exec(DELETE_POST_VOTES_RECORDS_BY_ACTOR_AND_POST_HASH_COMMAND, actor, postHash)
  if err != nil {
    log.Panicf("Failed to delete post votes records for actor %s and postHash %s with error:\n %+v",
      actor, postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post votes records for actor %s and postHash %s\n", actor, postHash)
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) GetTotalPostVotesCountTx(
    actor string, postHash string) int64 {
  var totalPostVotes sql.NullInt64
  err := postVotesRecordExecutor.Tx.Get(&totalPostVotes, QUERY_TOTAL_VOTES_COUNT_COMMAND, actor, postHash)

  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total post votes for actor %s and postHash %s with error:\n %+v", actor, postHash, err.Error())
  }
  return totalPostVotes.Int64
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) GetVotesCountByVoteTypeTx(
    actor string, postHash string, voteType feed_attributes.VoteType) int64 {
  var postVotes sql.NullInt64
  err := postVotesRecordExecutor.Tx.Get(&postVotes, QUERY_VOTES_COUNT_BY_VOTE_TYPE_COMMAND, actor, postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get votes count for actor %s, postHash %s and voteType %s with error:\n %+v",
      actor, postHash,  voteType, err)
  }
  return postVotes.Int64
}

func (postVotesRecordExecutor *PostVotesRecordExecutor) GetActorListByPostHashAndVoteTypeTx(
  postHash string, voteType feed_attributes.VoteType) *[]string {
  var actorList []string
  err := postVotesRecordExecutor.Tx.Select(&actorList, QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND, postHash, voteType)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get actor list for postHash %s and voteType %s with error:\n %+v", postHash,  voteType, err)
  }
  return &actorList
}
