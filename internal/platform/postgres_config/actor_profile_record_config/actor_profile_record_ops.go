package actor_profile_record_config

import (
  "log"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
)

type ActorProfileRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) CreateActorProfileRecordTable() {
  actorProfileRecordExecutor.CreateTimestampTrigger()
  actorProfileRecordExecutor.CreateTable(
    TABLE_SCHEMA_FOR_ACTOR_PROFILE_RECORD, TABLE_NAME_FOR_ACTOR_PROFILE_RECORD)
  actorProfileRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_PROFILE_RECORD)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecordTable() {
  actorProfileRecordExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_PROFILE_RECORD)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) UpsertActorProfileRecord(
  actorProfileRecord *ActorProfileRecord) {
  _, err := actorProfileRecordExecutor.C.NamedExec(
    UPSERT_ACTOR_PROFILE_RECORD_COMMAND, actorProfileRecord)
  if err != nil {
    log.Panicf("Failed to upsert profile record: %+v with error:\n %+v", actorProfileRecord, err.Error())
  }
  log.Printf("Sucessfully upserted profile record for actor %s\n", actorProfileRecord.Actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecords(actor string) {
  _, err := actorProfileRecordExecutor.C.Exec(DELETE_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete profile records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorProfileRecord(actor string) *ActorProfileRecord {
  var actorProfileRecord ActorProfileRecord
  err := actorProfileRecordExecutor.C.Get(&actorProfileRecord, QUERY_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    if err == sql.ErrNoRows {
      log.Panicf("No actor profile exists for actor %s", actor)
    }
    log.Panicf("Failed to get actor profile record for actor %s with error:\n %+v", actor, err.Error())
  }
  return &actorProfileRecord
}

/*
 * Tx Versions
 */
func (actorProfileRecordExecutor *ActorProfileRecordExecutor) UpsertActorProfileRecordTx(
    actorProfileRecord *ActorProfileRecord) {
  _, err := actorProfileRecordExecutor.Tx.NamedExec(
    UPSERT_ACTOR_PROFILE_RECORD_COMMAND, actorProfileRecord)
  if err != nil {
    log.Panicf("Failed to upsert profile record: %+v with error:\n %+v", actorProfileRecord, err.Error())
  }
  log.Printf("Sucessfully upserted profile record for actor %s\n", actorProfileRecord.Actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecordsTx(actor string) {
  _, err := actorProfileRecordExecutor.Tx.Exec(DELETE_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete profile records for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorProfileRecordTx(actor string) *ActorProfileRecord {
  var actorProfileRecord ActorProfileRecord
  err := actorProfileRecordExecutor.Tx.Get(&actorProfileRecord, QUERY_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    if err == sql.ErrNoRows {
      log.Panicf("No actor profile exists for actor %s", actor)
    }
    log.Panicf("Failed to get actor profile record for actor %s with error:\n %+v", actor, err.Error())
  }
  return &actorProfileRecord
}
