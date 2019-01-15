package purchase_mps_record_config

import (
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type PurchaseMPsRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (purchaseMPsRecordExecutor *PurchaseMPsRecordExecutor) CreatePurchaseMPsRecordTable() {
	purchaseMPsRecordExecutor.CreateTimestampTrigger()
	purchaseMPsRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_PURCHASE_MPS_RECORDS, TABLE_NAME_FOR_PURCHASE_MPS_RECORDS)
	purchaseMPsRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_PURCHASE_MPS_RECORDS)
}

func (purchaseMPsRecordExecutor *PurchaseMPsRecordExecutor) DeletePurchaseReputationsRecordTable() {
	purchaseMPsRecordExecutor.DeleteTable(TABLE_NAME_FOR_PURCHASE_MPS_RECORDS)
}

func (purchaseMPsRecordExecutor *PurchaseMPsRecordExecutor) UpsertPurchaseMPsRecordTx(
	purchaseMPsRecord *PurchaseMPsRecord) {
	_, err := purchaseMPsRecordExecutor.Tx.NamedExec(UPSERT_PURCHASE_MPS_RECORD_COMMAND, purchaseMPsRecord)
	if err != nil {
		log.Panicf("Failed to upsert purchase mps record %+v with error: %+v\n", purchaseMPsRecord, err)
	}
	log.Printf("Sucessfully upserted purchase mps record for purchaser %s", purchaseMPsRecord.Purchaser)
}
