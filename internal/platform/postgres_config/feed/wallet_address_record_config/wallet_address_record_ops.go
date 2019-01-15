package wallet_address_record_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
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
	err := walletAddressRecordExecutor.Tx.Select(&walletAddressList, QUERY_WALLET_ADDRESS_LIST_BY_ACTOR_COMMAND, actor)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get wallet address list for actor %s with error: %+v\n", actor, err)
	}
	return &walletAddressList
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) CheckWalletAddressExistingTx(actor string, walletAddress string) bool {
	var existing bool
	err := walletAddressRecordExecutor.Tx.Get(&existing, VERIFY_WALLET_ADDRESS_EXISTING_COMMAND, actor, walletAddress)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.WalletAddressRecordLocation)
		log.Printf("Failed to check  wallet addreess %s existing for actor %s with error: %+v\n", walletAddress, actor, err)
		log.Panicln(errInfo.Marshal())
	}
	return existing
}

func (walletAddressRecordExecutor *WalletAddressRecordExecutor) VerifyActorExistingTx(actor string, walletAddress string) {
	existing := walletAddressRecordExecutor.CheckWalletAddressExistingTx(actor, walletAddress)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoActorExisting,
			ErrorData: map[string]interface{}{
				"actor": actor,
			},
			ErrorLocation: error_config.ProfileAccountLocation,
		}
		log.Printf("No Actor Profile Acount for actor %s", actor)
		log.Panicln(errorInfo.Marshal())
	}
}
