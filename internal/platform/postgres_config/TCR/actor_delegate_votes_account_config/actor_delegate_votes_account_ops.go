package actor_delegate_votes_account_config

import (
  "log"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
)


type ActorDelegateVotesAccountExecutor struct {
  client_config.PostgresBigBangClient
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) CreateActorRatingVoteAccountTable( ) {
  actorDelegateVotesAccountExecutor.CreateTimestampTrigger()
  actorDelegateVotesAccountExecutor.CreateTable(TABLE_SCHEMA_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT, TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT)
  actorDelegateVotesAccountExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) DeleteActorRatingVoteAccountTable() {
  actorDelegateVotesAccountExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) UpsertActorDelegateVotesAccountRecord (
    actorDelegateVotesAccountRecord *ActorDelegateVotesAccountRecord) {
  _, err := actorDelegateVotesAccountExecutor.C.NamedExec(UPSERT_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actorDelegateVotesAccountRecord)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actorDelegateVotesAccountRecord.Actor, error_config.ActorDelegateVotesAccountRecordLocation)
    log.Printf("Failed to upsert Actor Delegate Votes Account Record for actor %s with error: %+v\n", actorDelegateVotesAccountRecord.Actor, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Sucessfully upserted Actor Delegate Votes Account Record for actor %s\n", actorDelegateVotesAccountRecord.Actor)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) DeleteActorDelegateVotesAccountRecord(actor string) {
  _, err := actorDelegateVotesAccountExecutor.C.Exec(DELETE_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
    log.Printf("Failed to delete Actor Delegate Votes Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted Actor Delegate Votes Account Record for actor %s\n", actor)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) GetActorDelegateVotesAccountRecord(
  actor string) *ActorDelegateVotesAccountRecord {
  var actorDelegateVotesAccountRecord ActorDelegateVotesAccountRecord
  err := actorDelegateVotesAccountExecutor.C.Get(&actorDelegateVotesAccountRecord, QUERY_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
    log.Printf("Failed to get Actor Delegate Votes Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ActorDelegateVotesAccountRecordLocation,
    }
    log.Printf("No Actor Delegate Votes Account Record for actor %s\n", actor)
    log.Panicln(errorInfo.Marshal())
  }
  return &actorDelegateVotesAccountRecord
}

/*
 * Tx versions
 */
func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) UpsertActorDelegateVotesAccountRecordTx (
    actorDelegateVotesAccountRecord *ActorDelegateVotesAccountRecord) {
  _, err := actorDelegateVotesAccountExecutor.Tx.NamedExec(UPSERT_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actorDelegateVotesAccountRecord)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actorDelegateVotesAccountRecord.Actor, error_config.ActorDelegateVotesAccountRecordLocation)
    log.Printf("Failed to upsert Actor Delegate Votes Account Record for actor %s with error: %+v\n", actorDelegateVotesAccountRecord.Actor, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Sucessfully upserted Actor Delegate Votes Account Record for actor %s\n", actorDelegateVotesAccountRecord.Actor)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) DeleteActorDelegateVotesAccountRecordTx(actor string) {
  _, err := actorDelegateVotesAccountExecutor.Tx.Exec(DELETE_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
    log.Printf("Failed to delete Actor Delegate Votes Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted Actor Delegate Votes Account Record for actor %s\n", actor)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) GetActorDelegateVotesAccountRecordTx(
    actor string) *ActorDelegateVotesAccountRecord {
  var actorDelegateVotesAccountRecord ActorDelegateVotesAccountRecord
  err := actorDelegateVotesAccountExecutor.Tx.Get(&actorDelegateVotesAccountRecord, QUERY_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actor)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
    log.Printf("Failed to get Actor Delegate Votes Account Record for actor %s with error: %+v\n",
      actor, err)
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ActorDelegateVotesAccountRecordLocation,
    }
    log.Printf("No Actor Delegate Votes Account Record for actor %s\n", actor)
    log.Panicln(errorInfo.Marshal())
  }
  return &actorDelegateVotesAccountRecord
}