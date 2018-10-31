package wallet_address_record_config

import (
  "log"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
)


type WalletAddressRecordExecutor struct {
  client_config.PostgresBigBangClient
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) CreateWalletAddressRecordTable() {
  walletAddressRecordExecutor.CreateTimestampTrigger()
  walletAddressRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_WALLET_ADDRESS_RECORD, TABLE_NAME_FOR_WALLET_ADDRESS_RECORD)
  walletAddressRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_WALLET_ADDRESS_RECORD)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) DeleteWalletAddressRecordTable() {
  walletAddressRecordExecutor.DeleteTable(TABLE_NAME_FOR_WALLET_ADDRESS_RECORD)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) UpsertWalletAddressRecord(walletAddressRecord *WalletAddressRecord) {
  _, err := walletAddressRecordExecutor.C.NamedExec(UPSERT_WALLET_ADDRESS_RECORD_COMMAND, walletAddressRecord)
  if err != nil {
    log.Panicf("Failed to upsert wallet address record: %+v with error:\n %+v", walletAddressRecord, err)
  }
  log.Printf("Sucessfully upserted wallet address record for actor %s\n",
    walletAddressRecord.Actor)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) DeleteWalletAddressRecordsByActor(actor string) {
  _, err := walletAddressRecordExecutor.C.Exec(DELETE_ALL_WALLET_ADDRESS_RECORDS_BY_ACTOR_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete wallet address records for actor %s with error:\n %+v", actor, err)
  }
  log.Printf("Sucessfully deleted wallet address records for actor %s\n", actor)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) DeleteWalletAddressRecordByActorAndAddress(actor string, walletAddress string) {
  _, err := walletAddressRecordExecutor.C.Exec(DELETE_WALLET_ADDRESS_RECORDS_BY_ACTOR_AND_ADDRESS_COMMAND, actor, walletAddress)
  if err != nil {
    log.Panicf("Failed to delete wallet address record for actor %s and walletAddress %s with error:\n %+v", actor, walletAddress, err)
  }
  log.Printf("Sucessfully deleted wallet address record for actor %s and walletAddress %s \n", actor, walletAddress)
}


func (walletAddressRecordExecutor *WalletAddressRecordExecutor) GetWalletAddressListByActor(actor string) *[]string {
  var walletAddressList []string
  err := walletAddressRecordExecutor.C.Select(&walletAddressList, QUERY_WALLET_ADDRESS_LIST_BY_ACTOR_COMMAND)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get wallet address list for actor %s and voteType %s with error:\n %+v", actor, err)
  }
  return &walletAddressList
}

/*
 * Tx versions
 */

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) UpsertWalletAddressRecordTx(walletAddressRecord *WalletAddressRecord) {
  _, err := walletAddressRecordExecutor.Tx.NamedExec(UPSERT_WALLET_ADDRESS_RECORD_COMMAND, walletAddressRecord)
  if err != nil {
    log.Panicf("Failed to upsert wallet address record: %+v with error:\n %+v", walletAddressRecord, err)
  }
  log.Printf("Sucessfully upserted wallet address record for actor %s\n",
    walletAddressRecord.Actor)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) DeleteWalletAddressRecordsByActorTx(actor string) {
  _, err := walletAddressRecordExecutor.Tx.Exec(DELETE_ALL_WALLET_ADDRESS_RECORDS_BY_ACTOR_COMMAND, actor)
  if err != nil {
    log.Panicf("Failed to delete wallet address records for actor %s with error:\n %+v", actor, err)
  }
  log.Printf("Sucessfully deleted wallet address records for actor %s\n", actor)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) DeleteWalletAddressRecordByActorAndAddressTx(actor string, walletAddress string) {
  _, err := walletAddressRecordExecutor.Tx.Exec(DELETE_WALLET_ADDRESS_RECORDS_BY_ACTOR_AND_ADDRESS_COMMAND, actor, walletAddress)
  if err != nil {
    log.Panicf("Failed to delete wallet address record for actor %s and walletAddress %s with error:\n %+v", actor, walletAddress, err)
  }
  log.Printf("Sucessfully deleted wallet address record for actor %s and walletAddress %s \n", actor, walletAddress)
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) GetWalletAddressListByActorTx(actor string) *[]string {
  var walletAddressList []string
  err := walletAddressRecordExecutor.Tx.Select(&walletAddressList, QUERY_WALLET_ADDRESS_LIST_BY_ACTOR_COMMAND)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf(
      "Failed to get wallet address list for actor %s and voteType %s with error:\n %+v", actor, err)
  }
  return &walletAddressList
}
