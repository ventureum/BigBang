package purchase_reputations_record_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
)


type PurchaseReputationsRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (purchaseReputationsRecordExecutor *PurchaseReputationsRecordExecutor) CreatePurchaseReputationsRecordTable() {
  purchaseReputationsRecordExecutor.CreateTimestampTrigger()
  purchaseReputationsRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_PURCHASE_REPUTATIONS_RECORDS, TABLE_NAME_FOR_PURCHASE_REPUTATIONS_RECORDS)
  purchaseReputationsRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_PURCHASE_REPUTATIONS_RECORDS)
}

func (purchaseReputationsRecordExecutor *PurchaseReputationsRecordExecutor) DeletePurchaseReputationsRecordTable() {
  purchaseReputationsRecordExecutor.DeleteTable(TABLE_NAME_FOR_PURCHASE_REPUTATIONS_RECORDS)
}

func (purchaseReputationsRecordExecutor *PurchaseReputationsRecordExecutor) UpsertPurchaseReputationsRecord(purchaseReputationsRecord *PurchaseReputationsRecord) {
  _, err := purchaseReputationsRecordExecutor.C.NamedExec(UPSERT_PURCHASE_REPUTATIONS_RECORD_COMMAND, purchaseReputationsRecord)
  if err != nil {
    log.Panicf("Failed to upsert purchase reputations Record: %+v with error:\n %+v", purchaseReputationsRecord, err.Error())
  }
  log.Printf("Sucessfully upserted purchase reputations Record for purchaser %s", purchaseReputationsRecord.Purchaser)
}

/*
 * Tx versions
 */
func (purchaseReputationsRecordExecutor *PurchaseReputationsRecordExecutor) UpsertPurchaseReputationsRecordTx(
  purchaseReputationsRecord *PurchaseReputationsRecord) {
  _, err := purchaseReputationsRecordExecutor.Tx.NamedExec(UPSERT_PURCHASE_REPUTATIONS_RECORD_COMMAND, purchaseReputationsRecord)
  if err != nil {
    log.Panicf("Failed to upsert purchase reputations Record: %+v with error:\n %+v", purchaseReputationsRecord, err.Error())
  }
  log.Printf("Sucessfully upserted purchase reputations Record for purchaser %s", purchaseReputationsRecord.Purchaser)
}
