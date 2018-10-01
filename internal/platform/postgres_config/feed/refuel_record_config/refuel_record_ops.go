package refuel_record_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/feed/client_config"
  "time"
  "BigBang/internal/pkg/error_config"
)

type RefuelRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (refuelRecordExecutor *RefuelRecordExecutor) CreateRefuelRecordTable( ) {
  refuelRecordExecutor.CreateTimestampTrigger()
  refuelRecordExecutor.CreateTable(
    TABLE_SCHEMA_FOR_REFUEL_RECORD, TABLE_NAME_FOR_REFUEL_RECORD)
  refuelRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_REFUEL_RECORD)
}

func (refuelRecordExecutor *RefuelRecordExecutor) DeleteRefuelRecordTable() {
  refuelRecordExecutor.DeleteTable(TABLE_NAME_FOR_REFUEL_RECORD)
}

func (refuelRecordExecutor *RefuelRecordExecutor) UpsertRefuelRecord(
  refuelRecord *RefuelRecord) {
  _, err := refuelRecordExecutor.C.NamedExec(
    UPSERT_REFUEL_RECORD_COMMAND, refuelRecord)
  if err != nil {
    log.Panicf("Failed to upsert Refuel Record %+v with error: %+v\n", refuelRecord, err)
  }
  log.Printf("Sucessfully upserted Refuel Record for actor %s\n", refuelRecord.Actor)
}

func (refuelRecordExecutor *RefuelRecordExecutor) DeleteRefuelRecord(actor string) {
  _, err := refuelRecordExecutor.C.Exec(DELETE_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete Refuel Records  for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted Refuel Records for actor %s\n", actor)
}

func (refuelRecordExecutor *RefuelRecordExecutor) GetRefuelRecord(
  actor string) *[]RefuelRecord {
  var refuelRecords []RefuelRecord
  err := refuelRecordExecutor.C.Select(& refuelRecords, QUERY_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to get Refuel Records for actor %s with error: %+v\n", actor, err)
  }
  return &refuelRecords
}

func (refuelRecordExecutor *RefuelRecordExecutor) GetLastRefuelTime(
    actor string) time.Time {
  res, err := refuelRecordExecutor.C.Queryx(QUERY_LATEST_REFUEL_TIME_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.RefuelRecordLocation)
    log.Printf("Failed to get lastRefuelTime for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }

  var lastRefuelTime time.Time
  for res.Next() {
    res.Scan(&lastRefuelTime)
  }
  return lastRefuelTime
}


/*
 * Tx versions
 */
func (refuelRecordExecutor *RefuelRecordExecutor) UpsertRefuelRecordTx(
    refuelRecord *RefuelRecord) {
  _, err := refuelRecordExecutor.Tx.NamedExec(
    UPSERT_REFUEL_RECORD_COMMAND, refuelRecord)
  if err != nil {
    log.Panicf("Failed to upsert Refuel Record %+v with error: %+v\n", refuelRecord, err)
  }
  log.Printf("Sucessfully upserted Refuel Record for actor %s\n", refuelRecord.Actor)
}

func (refuelRecordExecutor *RefuelRecordExecutor) DeleteRefuelRecordTx(actor string) {
  _, err := refuelRecordExecutor.Tx.Exec(DELETE_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete Refuel Records  for actor %s with error:\n %+v", actor, err)
  }
  log.Printf("Sucessfully deleted Refuel Records for actor %s\n", actor)
}

func (refuelRecordExecutor *RefuelRecordExecutor) GetRefuelRecordTx(
    actor string) *[]RefuelRecord {
  var refuelRecords []RefuelRecord
  err := refuelRecordExecutor.Tx.Select(& refuelRecords, QUERY_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to get Refuel Records for actor %s with error: %+v\n", actor, err)
  }
  return &refuelRecords
}

func (refuelRecordExecutor *RefuelRecordExecutor) GetLastRefuelTimeTx(
    actor string) time.Time {
  res, err := refuelRecordExecutor.Tx.Queryx(QUERY_LATEST_REFUEL_TIME_COMMAND, actor)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.RefuelRecordLocation)
    log.Printf("Failed to get lastRefuelTime for actor %s with error: %+v\n", actor, err)
    log.Panicln(errInfo.Marshal())
  }

  var lastRefuelTime time.Time
  for res.Next() {
    res.Scan(&lastRefuelTime)
  }
  return lastRefuelTime
}
