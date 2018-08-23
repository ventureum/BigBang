package actor_reputations_record_config

import (
  "log"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/error_config"
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
    log.Panicf("Failed to upsert reputaions record: %+v with error:\n %+v", actorReputationsRecord, err.Error())
  }
  log.Printf("Sucessfully upserted reputaions record for actor %s\n", actorReputationsRecord.Actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) DeleteActorReputationsRecords(actor string) {
  _, err := actorReputationsRecordExecutor.C.Exec(DELETE_ACTOR_REPUTATIONS_RECORD_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete reputaions records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted reputaions records for actor %s\n", actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetActorReputations(
  actor string) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := actorReputationsRecordExecutor.C.Get(&reputations , QUERY_ACTOR_REPUTATIONS_COMMAND, actor)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get reputaions for actor %s with error:\n %+v", actor, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) AddActorReputations(
    actor string, reputationToAdd feed_attributes.Reputation) {
  _, err := actorReputationsRecordExecutor.C.Exec(ADD_ACTOR_REPUTATIONS_COMMAND, actor, reputationToAdd)

  if err != nil {
    log.Panicf("Failed to add reputaions for actor %s with error:\n %+v", actor, err.Error())
  }

  log.Printf("Successfully added reputaions %d for actor %s", reputationToAdd, actor)
}


func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) SubActorReputations(
    actor string, reputationToSub feed_attributes.Reputation) {
  var diff int64
  err := actorReputationsRecordExecutor.C.Get(&diff, SUB_ACTOR_REPUTATIONS_COMMAND, actor, reputationToSub)

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
    }
    log.Panic(errorInfo.ToJsonText())
  }

  if diff > 0 {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.InsufficientReputaionsAmount,
      ErrorData: map[string]interface{} {
        "diff": diff,
      },
    }
    log.Panic(errorInfo.ToJsonText())
  }

  if err != nil {
    log.Panicf("Failed to substract reputaions from actor %s with error:\n %+v", actor, err)
  }

  log.Printf("Successfully substracted reputaions %d from actor %s", reputationToSub, actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetTotalActorReputations() int64 {
  var totalReputations int64
  err := actorReputationsRecordExecutor.C.Get(&totalReputations, QUARY_TOTAL_REPUTATIONS_COMMAND)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total reputations for all actors with error:\n %+v", err.Error())
  }
  return totalReputations
}


/*
 * Tx Versions
 */
func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) UpsertActorReputationsRecordTx(
    actorReputationsRecord *ActorReputationsRecord) {
  _, err := actorReputationsRecordExecutor.Tx.NamedExec(
    UPSERT_ACTOR_REPUTATIONS_RECORD_COMMAND, actorReputationsRecord)
  if err != nil {
    log.Panicf("Failed to upsert reputaions record: %+v with error:\n %+v", actorReputationsRecord, err.Error())
  }
  log.Printf("Sucessfully upserted reputaions record for actor %s\n", actorReputationsRecord.Actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) DeleteActorReputationsRecordsTx(actor string) {
  _, err := actorReputationsRecordExecutor.Tx.Exec(DELETE_ACTOR_REPUTATIONS_RECORD_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete reputaions records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted reputaions records for actor %s\n", actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetActorReputationsTx(
    actor string) feed_attributes.Reputation {
  var reputations sql.NullInt64
  err := actorReputationsRecordExecutor.Tx.Get(&reputations , QUERY_ACTOR_REPUTATIONS_COMMAND, actor)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get reputaions for actor %s with error:\n %+v", actor, err.Error())
  }
  return feed_attributes.Reputation(reputations.Int64)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) AddActorReputationsTx(
    actor string, reputationToAdd feed_attributes.Reputation) {
  _, err := actorReputationsRecordExecutor.Tx.Exec(ADD_ACTOR_REPUTATIONS_COMMAND, actor, reputationToAdd)

  if err != nil {
    log.Panicf("Failed to add reputaions for actor %s with error:\n %+v", actor, err.Error())
  }

  log.Printf("Successfully added reputaions %d for actor %s", reputationToAdd, actor)
}


func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) SubActorReputationsTx(
    actor string, reputationToSub feed_attributes.Reputation) {
  var diff int64
  err := actorReputationsRecordExecutor.Tx.Get(&diff, SUB_ACTOR_REPUTATIONS_COMMAND, actor, reputationToSub)

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
    }
    log.Panic(errorInfo.Marshal())
  }

  if diff > 0 {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.InsufficientReputaionsAmount,
      ErrorData: map[string]interface{} {
        "diff": diff,
      },
    }
    log.Panic(errorInfo.Marshal())
  }

  if err != nil {
    log.Panicf("Failed to substract reputaions from actor %s with error:\n %+v", actor, err)
  }

  log.Printf("Successfully substracted reputaions %d from actor %s", reputationToSub, actor)
}

func (actorReputationsRecordExecutor *ActorReputationsRecordExecutor) GetTotalActorReputationsTx() feed_attributes.Reputation {
  var totalReputations sql.NullInt64
  err := actorReputationsRecordExecutor.Tx.Get(&totalReputations, QUARY_TOTAL_REPUTATIONS_COMMAND)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get total reputations for all actors with error:\n %+v", err.Error())
  }
  return feed_attributes.Reputation(totalReputations.Int64)
}
