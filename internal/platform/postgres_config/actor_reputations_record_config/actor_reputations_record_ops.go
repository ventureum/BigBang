package actor_reputations_record_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/feed_attributes"
)

type ActorReputationsRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) CreateActorReputationsRecordTable() {
  actorReputationsRecordExecutor.CreateTimestampTrigger()
  actorReputationsRecordExecutor.CreateTable(
    TABLE_SCHEMA_FOR_ACTOR_REPUTATIONS_RECORD, TABLE_NAME_FOR_ACTOR_REPUTATIONS_RECORD)
  actorReputationsRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_REPUTATIONS_RECORD)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) DeleteActorReputationsRecordTable() {
  actorReputationsRecordExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_REPUTATIONS_RECORD)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) UpsertActorReputationsRecord(
    actorReputationsRecord *ActorReputationsRecord) {
  _, err := actorReputationsRecordExecutor.C.NamedExec(
    UPSERT_ACTOR_REPUTATIONS_RECORD_COMMAND, actorReputationsRecord)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actorReputationsRecord.Actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to upsert reputaions record: %+v with error:\n %+v", actorReputationsRecord, err)
    log.Panicln(errorInfo.Marshal())
  }
  log.Printf("Sucessfully upserted reputaions record for actor %s\n", actorReputationsRecord.Actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) VerifyActorExisting (actor string) {
  var existing bool
  err := actorReputationsRecordExecutor.C.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to verify actor existing for actor %s with error: %+v\n", actor, err)
    log.Panicln(errorInfo.Marshal())
  }

  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation:  error_config.ReputationsAccountLocation,
    }
    log.Printf("No Actor Reputations Acount for actor %s", actor)
    log.Panicln(errorInfo.Marshal())
  }
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) DeleteActorReputationsRecords(actor string) {
  _, err := actorReputationsRecordExecutor.C.Exec(DELETE_ACTOR_REPUTATIONS_RECORD_COMMAND, actor)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to delete reputaions records for actor %s with error: %+v\n", actor, err)
    log.Panicln(errorInfo.Marshal())
  }
  log.Printf("Sucessfully deleted reputaions records for actor %s\n", actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetActorReputations(
    actor string) feed_attributes.Reputation {
  var reputations int64
  err := actorReputationsRecordExecutor.C.Get(&reputations, QUERY_ACTOR_REPUTATIONS_COMMAND, actor)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to get reputaions for actor %s with error: %+v\n", actor, err)
    log.Panic(errorInfo.Marshal())
  }
  return feed_attributes.Reputation(reputations)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) AddActorReputations(
    actor string, reputationToAdd feed_attributes.Reputation) {
  _, err := actorReputationsRecordExecutor.C.Exec(ADD_ACTOR_REPUTATIONS_COMMAND, actor, reputationToAdd)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to add reputaions for actor %s with error: %+v\n", actor, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully added reputaions %d for actor %s", reputationToAdd, actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) SubActorReputations(
    actor string, reputationToSub feed_attributes.Reputation) {
  var diff int64
  err := actorReputationsRecordExecutor.C.Get(&diff, SUB_ACTOR_REPUTATIONS_COMMAND, actor, reputationToSub)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to substract reputaions from actor %s with error:\n %+v", actor, err)
    log.Panic(errorInfo.Marshal())
  }

  if diff > 0 {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.InsufficientReputaionsAmount,
      ErrorData: map[string]interface{} {
        "diff": diff,
      },
      ErrorLocation: error_config.ReputationsAccountLocation,
    }
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully substracted reputaions %d from actor %s", reputationToSub, actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetTotalActorReputations() int64 {
  var totalReputations int64
  err := actorReputationsRecordExecutor.C.Get(&totalReputations, QUARY_TOTAL_REPUTATIONS_COMMAND)
  if err != nil {
    errorInfo := error_config.MatchError(err, "", "", error_config.ReputationsAccountLocation)
    log.Printf("Failed to get total reputations for all actors with error:\n %+v", err)
    log.Panic(errorInfo.Marshal())
  }
  return totalReputations
}

/*
 * Tx Versions
 */
func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) VerifyActorExistingTx (actor string) {
  var existing bool
  err := actorReputationsRecordExecutor.Tx.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Panicf("Failed to verify actor existing for actor %s with error: %+v\n", actor, err)
    log.Panicln(errorInfo.Marshal())
  }

  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ReputationsAccountLocation,
    }
    log.Printf("No Actor Reputations Acount for actor %s", actor)
    log.Panicln(errorInfo.Marshal())
  }
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) DeleteActorReputationsRecordsTx(actor string) {
  _, err := actorReputationsRecordExecutor.Tx.Exec(DELETE_ACTOR_REPUTATIONS_RECORD_COMMAND, actor)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to delete reputaions records for actor %s with error: %+v\n", actor, err)
    log.Panicln(errorInfo.Marshal())
  }
  log.Printf("Sucessfully deleted reputaions records for actor %s\n", actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetActorReputationsTx(
    actor string) feed_attributes.Reputation {
  var reputations int64
  err := actorReputationsRecordExecutor.Tx.Get(&reputations, QUERY_ACTOR_REPUTATIONS_COMMAND, actor)
  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to get reputaions for actor %s with error: %+v\n", actor, err)
    log.Panic(errorInfo.Marshal())
  }
  return feed_attributes.Reputation(reputations)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) AddActorReputationsTx(
    actor string, reputationToAdd feed_attributes.Reputation) {
  _, err := actorReputationsRecordExecutor.Tx.Exec(ADD_ACTOR_REPUTATIONS_COMMAND, actor, reputationToAdd)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to add reputaions for actor %s with error: %+v\n", actor, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully added reputaions %d for actor %s", reputationToAdd, actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) SubActorReputationsTx(
    actor string, reputationToSub feed_attributes.Reputation) {
  var diff int64
  err := actorReputationsRecordExecutor.Tx.Get(&diff, SUB_ACTOR_REPUTATIONS_COMMAND, actor, reputationToSub)

  if err != nil {
    errorInfo := error_config.MatchError(err, "actor", actor, error_config.ReputationsAccountLocation)
    log.Printf("Failed to substract reputaions from actor %s with error:\n %+v", actor, err)
    log.Panic(errorInfo.Marshal())
  }

  if diff > 0 {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.InsufficientReputaionsAmount,
      ErrorData: map[string]interface{} {
        "diff": diff,
      },
      ErrorLocation: error_config.ReputationsAccountLocation,
    }
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully substracted reputaions %d from actor %s", reputationToSub, actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetTotalActorReputationsTx() int64 {
  var totalReputations int64
  err := actorReputationsRecordExecutor.Tx.Get(&totalReputations, QUARY_TOTAL_REPUTATIONS_COMMAND)
  if err != nil {
    errorInfo := error_config.MatchError(err, "", "", error_config.ReputationsAccountLocation)
    log.Printf("Failed to get total reputations for all actors with error: %+v\n", err)
    log.Panic(errorInfo.Marshal())
  }
  return totalReputations
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) UpsertActorReputationsRecordTx(
    actorReputationsRecord *ActorReputationsRecord) {
  _, err := actorReputationsRecordExecutor.Tx.NamedExec(
    UPSERT_ACTOR_REPUTATIONS_RECORD_COMMAND, actorReputationsRecord)
  if err != nil {
    errorInfo := error_config.MatchError(err, "", "", error_config.ReputationsAccountLocation)
    log.Printf("Failed to upsert reputaions record: %+v with error: %+v\n", actorReputationsRecord, err)
    log.Panic(errorInfo.Marshal())
  }
  log.Printf("Sucessfully upserted reputaions record for actor %s\n", actorReputationsRecord.Actor)
}