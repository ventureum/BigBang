package actor_profile_record_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "database/sql"
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
    actorProfileRecord *ActorProfileRecord) bool {
  var inserted bool
  rows, err := actorProfileRecordExecutor.C.NamedQuery(
    UPSERT_ACTOR_PROFILE_RECORD_COMMAND, actorProfileRecord)
  for rows.Next() {
    err = rows.Scan(&inserted)
  }
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actorProfileRecord.Actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to upsert profile record: %+v with error: %+v\n", actorProfileRecord, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully upserted profile record for actor %s\n", actorProfileRecord.Actor)
  return inserted
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecords(actor string) {
  _, err := actorProfileRecordExecutor.C.Exec(DELETE_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to delete profile records for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeactivateActorProfileRecords(actor string) {
  _, err := actorProfileRecordExecutor.C.Exec(DEACTIVATE_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to deactivate profile records for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deactivated profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorProfileRecord(actor string) *ActorProfileRecord {
  var actorProfileRecord ActorProfileRecord
  err := actorProfileRecordExecutor.C.Get(&actorProfileRecord, QUERY_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to get actor profile record for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }
  return &actorProfileRecord
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) VerifyActorExisting (actor string) {
  var existing bool
  err := actorProfileRecordExecutor.C.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to verify actor existing for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }

  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ProfileAccountLocation,
    }
    log.Printf("No Actor Reputations Acount for actor %s", actor)
    log.Panicln(errorInfo.Marshal())
  }
}

/*
 * Tx Versions
 */
func (actorProfileRecordExecutor *ActorProfileRecordExecutor) UpsertActorProfileRecordTx(
    actorProfileRecord *ActorProfileRecord) bool {
  var inserted sql.NullBool
  rows, err := actorProfileRecordExecutor.Tx.NamedQuery(
    UPSERT_ACTOR_PROFILE_RECORD_COMMAND, actorProfileRecord)
  for rows.Next() {
    err = rows.Scan(&inserted)
  }
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actorProfileRecord.Actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to upsert profile record: %+v with error: %+v\n", actorProfileRecord, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully upserted profile record for actor %s\n", actorProfileRecord.Actor)
  return inserted.Bool
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecordsTx(actor string) {
  _, err := actorProfileRecordExecutor.Tx.Exec(DELETE_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to delete profile records for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeactivateActorProfileRecordsTx(actor string) {
  _, err := actorProfileRecordExecutor.Tx.Exec(DEACTIVATE_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to deactivate profile records for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deactivated profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorProfileRecordTx(actor string) *ActorProfileRecord {
  var actorProfileRecord ActorProfileRecord
  err := actorProfileRecordExecutor.Tx.Get(&actorProfileRecord, QUERY_ACTOR_PROFILE_RECORD_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to get actor profile record for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }
  return &actorProfileRecord
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) VerifyActorExistingTx (actor string) {
  var existing bool
  err := actorProfileRecordExecutor.Tx.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
    log.Printf("Failed to verify actor existing for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }

  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoActorExisting,
      ErrorData: map[string]interface{} {
        "actor": actor,
      },
      ErrorLocation: error_config.ProfileAccountLocation,
    }
    log.Printf("No Actor Reputations Acount for actor %s", actor)
    log.Panicln(errorInfo.Marshal())
  }
}
