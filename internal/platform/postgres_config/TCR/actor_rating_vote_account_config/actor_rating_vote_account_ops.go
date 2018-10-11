package actor_rating_vote_account_config

import (
  "log"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
)


type ActorRatingVoteAccountExecutor struct {
  client_config.PostgresBigBangClient
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) CreateActorRatingVoteAccountTable( ) {
  actorRatingVoteAccountExecutor.CreateTimestampTrigger()
  actorRatingVoteAccountExecutor.CreateTable(TABLE_SCHEMA_FOR_ACTOR_RATING_VOTE_ACCOUNT, TABLE_NAME_FOR_ACTOR_RATING_VOTE_ACCOUNT)
  actorRatingVoteAccountExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_RATING_VOTE_ACCOUNT)
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) DeleteActorRatingVoteAccountTable() {
  actorRatingVoteAccountExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_RATING_VOTE_ACCOUNT)
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) UpsertActorRatingVoteAccountRecord (
    actorRatingVoteAccountRecord *ActorRatingVoteAccountRecord) {
  _, err := actorRatingVoteAccountExecutor.C.NamedExec(UPSERT_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND, actorRatingVoteAccountRecord)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actorRatingVoteAccountRecord.Actor, error_config.ActorRatingVoteAccountRecordLocation)
    log.Printf("Failed to upsert Actor Rating Vote Account Record for actor %s with error: %+v\n", actorRatingVoteAccountRecord.Actor, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Sucessfully upserted Actor Rating Vote Account Record for actor %s\n", actorRatingVoteAccountRecord.Actor)
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) DeleteActorRatingVoteAccountRecord(actor string) {
  _, err := actorRatingVoteAccountExecutor.C.Exec(DELETE_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRatingVoteAccountRecordLocation)
    log.Printf("Failed to delete Actor Rating Vote Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted Actor Rating Vote Account Record for actor %s\n", actor)
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) GetActorRatingVoteAccountRecord(
  actor string) *ActorRatingVoteAccountRecord {
  var actorRatingVoteAccountRecord ActorRatingVoteAccountRecord
  err := actorRatingVoteAccountExecutor.C.Get(&actorRatingVoteAccountRecord, QUERY_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRatingVoteAccountRecordLocation)
    log.Printf("Failed to get Actor Rating Vote Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ActorRatingVoteAccountRecordLocation,
    }
    log.Printf("No Actor Rating Vote Account Record for actor %s\n", actor)
    log.Panicln(errorInfo.Marshal())
  }
  return &actorRatingVoteAccountRecord
}

/*
 * Tx versions
 */
func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) UpsertActorRatingVoteAccountRecordTx (
    actorRatingVoteAccountRecord *ActorRatingVoteAccountRecord) {
  _, err := actorRatingVoteAccountExecutor.Tx.NamedExec(UPSERT_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND, actorRatingVoteAccountRecord)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actorRatingVoteAccountRecord.Actor, error_config.ActorRatingVoteAccountRecordLocation)
    log.Printf("Failed to upsert Actor Rating Vote Account Record for actor %s with error: %+v\n", actorRatingVoteAccountRecord.Actor, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Sucessfully upserted Actor Rating Vote Account Record for actor %s\n", actorRatingVoteAccountRecord.Actor)
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) DeleteActorRatingVoteAccountRecordTx(actor string) {
  _, err := actorRatingVoteAccountExecutor.Tx.Exec(DELETE_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRatingVoteAccountRecordLocation)
    log.Printf("Failed to delete Actor Rating Vote Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted Actor Rating Vote Account Record for actor %s\n", actor)
}

func (actorRatingVoteAccountExecutor *ActorRatingVoteAccountExecutor) GetActorRatingVoteAccountRecordTx(
    actor string) *ActorRatingVoteAccountRecord {
  var actorRatingVoteAccountRecord ActorRatingVoteAccountRecord
  err := actorRatingVoteAccountExecutor.Tx.Get(&actorRatingVoteAccountRecord, QUERY_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRatingVoteAccountRecordLocation)
    log.Printf("Failed to get Actor Rating Vote Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ActorRatingVoteAccountRecordLocation,
    }
    log.Printf("No Actor Rating Vote Account Record for actor %s\n", actor)
    log.Panicln(errorInfo.Marshal())
  }
  return &actorRatingVoteAccountRecord
}
