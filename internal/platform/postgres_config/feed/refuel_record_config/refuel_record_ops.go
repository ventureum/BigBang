package refuel_record_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
	"time"
)

type RefuelRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (refuelRecordExecutor *RefuelRecordExecutor) CreateRefuelRecordTable() {
	refuelRecordExecutor.CreateTimestampTrigger()
	refuelRecordExecutor.CreateTable(
		TABLE_SCHEMA_FOR_REFUEL_RECORD, TABLE_NAME_FOR_REFUEL_RECORD)
	refuelRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_REFUEL_RECORD)
}

func (refuelRecordExecutor *RefuelRecordExecutor) DeleteRefuelRecordTable() {
	refuelRecordExecutor.DeleteTable(TABLE_NAME_FOR_REFUEL_RECORD)
}

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
	err := refuelRecordExecutor.Tx.Select(&refuelRecords, QUERY_REFUEL_RECORDS_COMMAND, actor)
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
