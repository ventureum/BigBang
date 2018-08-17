package reputations_refuel_record_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
)

type ReputationsRefuelRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) CreateReputationsRefuelRecordTable( ) {
  reputationsRefuelRecordExecutor.CreateTimestampTrigger()
  reputationsRefuelRecordExecutor.CreateTable(
    TABLE_SCHEMA_FOR_REPUTATION_REFUEL_RECORD, TABLE_NAME_FOR_REPUTATION_REFUEL_RECORD)
  reputationsRefuelRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_REPUTATION_REFUEL_RECORD)
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) DeleteReputationsRefuelRecordTable() {
  reputationsRefuelRecordExecutor.DeleteTable(TABLE_NAME_FOR_REPUTATION_REFUEL_RECORD)
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) UpsertReputationsRefuelRecord(
  reputationsRefuelRecord *ReputationsRefuelRecord) {
  _, err := reputationsRefuelRecordExecutor.C.NamedExec(
    UPSERT_REPUTATION_REFUEL_RECORD_COMMAND, reputationsRefuelRecord)
  if err != nil {
    log.Panicf("Failed to upsert Reputations Refuel Record %+v with error:\n %+v", reputationsRefuelRecord, err.Error())
  }
  log.Printf("Sucessfully upserted Reputations Refuel Record for actor %s\n", reputationsRefuelRecord.Actor)
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) DeleteReputationsRefuelRecord(actor string) {
  _, err := reputationsRefuelRecordExecutor.C.Exec(DELETE_REPUTATION_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete Reputations Refuel Records  for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted Reputations Refuel Records for actor %s\n", actor)
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) GetReputationsRefuelRecord(
  actor string) *[]ReputationsRefuelRecord {
  var reputationsRefuelRecords []ReputationsRefuelRecord
  err := reputationsRefuelRecordExecutor.C.Select(& reputationsRefuelRecords, QUERY_REPUTATION_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to get Reputations Refuel Records for actor %s with error:\n %+v", actor, err.Error())
  }
  return &reputationsRefuelRecords
}

/*
 * Tx versions
 */
func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) UpsertReputationsRefuelRecordTx(
    reputationsRefuelRecord *ReputationsRefuelRecord) {
  _, err := reputationsRefuelRecordExecutor.Tx.NamedExec(
    UPSERT_REPUTATION_REFUEL_RECORD_COMMAND, reputationsRefuelRecord)
  if err != nil {
    log.Panicf("Failed to upsert Reputations Refuel Record %+v with error:\n %+v", reputationsRefuelRecord, err.Error())
  }
  log.Printf("Sucessfully upserted Reputations Refuel Record for actor %s\n", reputationsRefuelRecord.Actor)
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) DeleteReputationsRefuelRecordTx(actor string) {
  _, err := reputationsRefuelRecordExecutor.Tx.Exec(DELETE_REPUTATION_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete Reputations Refuel Records  for actor %s with error:\n %+v", actor, err.Error())
  }
  log.Printf("Sucessfully deleted Reputations Refuel Records for actor %s\n", actor)
}

func (reputationsRefuelRecordExecutor *ReputationsRefuelRecordExecutor) GetReputationsRefuelRecordTx(
    actor string) *[]ReputationsRefuelRecord {
  var reputationsRefuelRecords []ReputationsRefuelRecord
  err := reputationsRefuelRecordExecutor.Tx.Select(& reputationsRefuelRecords, QUERY_REPUTATION_REFUEL_RECORDS_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to get Reputations Refuel Records for actor %s with error:\n %+v", actor, err.Error())
  }
  return &reputationsRefuelRecords
}