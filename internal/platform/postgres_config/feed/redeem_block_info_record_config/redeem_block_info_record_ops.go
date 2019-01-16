package redeem_block_info_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type RedeemBlockInfoRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) CreateRedeemBlockInfoRecordTable() {
	redeemBlockInfoRecordExecutor.CreateTimestampTrigger()
	redeemBlockInfoRecordExecutor.CreateTable(
		TABLE_SCHEMA_FOR_REDEEM_BLOCK_INFO_RECORD, TABLE_NAME_FOR_REDEEM_BLOCK_INFO_RECORD)
	redeemBlockInfoRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_REDEEM_BLOCK_INFO_RECORD)
	nextRedeemBlock := feed_attributes.MoveToNextNRedeemBlock(1)
	redeemBlockInfoRecordExecutor.InitRedeemBlockInfoTx(nextRedeemBlock)
	redeemBlockInfoRecordExecutor.UpdateTotalEnrolledMilestonePointsForRedeemBlockInfoRecordTx(nextRedeemBlock)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) DeleteRedeemBlockInfoRecordTable() {
	redeemBlockInfoRecordExecutor.DeleteTable(TABLE_NAME_FOR_REDEEM_BLOCK_INFO_RECORD)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) InitRedeemBlockInfoTx(nextRedeemBlock feed_attributes.RedeemBlock) {
	executedAt := nextRedeemBlock.ConvertToTime()
	_, err := redeemBlockInfoRecordExecutor.Tx.Exec(
		INIT_REDEEM_BLOCK_INFO_RECORD_COMMAND, nextRedeemBlock, executedAt)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", nextRedeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to init redeem request record for nextRedeemBlock %d with error: %+v\n", nextRedeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully init redeem request record for nextRedeemBlock %d\n", nextRedeemBlock)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) UpsertRedeemBlockInfoRecordTx(
	redeemBlockInfoRecord *RedeemBlockInfoRecord) {
	_, err := redeemBlockInfoRecordExecutor.Tx.NamedExec(
		UPSERT_REDEEM_BLOCK_INFO_RECORD_COMMAND, redeemBlockInfoRecord)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlockInfoRecord.RedeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to upsert redeem request record: %+v with error: %+v\n", redeemBlockInfoRecord, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully upserted redeem block info record for redeemBlock %d\n", redeemBlockInfoRecord.RedeemBlock)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) CheckRedeemBlockInfoExistingTx(redeemBlock feed_attributes.RedeemBlock) bool {
	var existing bool
	err := redeemBlockInfoRecordExecutor.Tx.Get(&existing, VERIFY_REDEEM_BLOCK_INFO_RECORD_EXISTING_COMMAND, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to verify redeem block info record existing for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	return existing
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) VerifyRedeemBlockInfoExistingTx(redeemBlock feed_attributes.RedeemBlock) {
	existing := redeemBlockInfoRecordExecutor.CheckRedeemBlockInfoExistingTx(redeemBlock)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoReDeemBlockInfoRecordExisting,
			ErrorData: map[string]interface{}{
				"redeemBlock": redeemBlock,
			},
			ErrorLocation: error_config.RedeemBlockInfoRecordLocation,
		}
		log.Printf("No redeem block info record for redeemBlock %d", redeemBlock)
		log.Panicln(errorInfo.Marshal())
	}
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) DeleteRedeemBlockInfoRecordTx(redeemBlock feed_attributes.RedeemBlock) {
	_, err := redeemBlockInfoRecordExecutor.Tx.Exec(DELETE_REDEEM_BLOCK_INFO_RECORD_COMMAND, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to delete redeem block info record for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully deleted redeem block info record for redeemBlock %d\n", redeemBlock)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) GetRedeemBlockInfoTx(
	redeemBlock feed_attributes.RedeemBlock) *feed_attributes.RedeemBlockInfo {
	var redeemBlockInfo feed_attributes.RedeemBlockInfo
	err := redeemBlockInfoRecordExecutor.Tx.Get(&redeemBlockInfo, QUERY_REDEEM_BLOCK_INFO_COMMAND, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to get redeem block info for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panic(errorInfo.Marshal())
	}
	return &redeemBlockInfo
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) UpdateExecuteAtForRedeemBlockInfoRecordTx(redeemBlock feed_attributes.RedeemBlock) {
	_, err := redeemBlockInfoRecordExecutor.Tx.Exec(UPDATE_EXECUTED_AT_COMMAND, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to update executedAt for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully updated executedAt for redeemBlock %d\n", redeemBlock)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) UpdateTokenPoolForRedeemBlockInfoRecordTx(
	redeemBlock feed_attributes.RedeemBlock, tokenPool int64) {
	_, err := redeemBlockInfoRecordExecutor.Tx.Exec(UPDATE_TOOKEN_POOL_COMMAND, redeemBlock, tokenPool)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to update tokenPool for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully updated tokenPool for redeemBlock %d\n", redeemBlock)
}

func (redeemBlockInfoRecordExecutor *RedeemBlockInfoRecordExecutor) UpdateTotalEnrolledMilestonePointsForRedeemBlockInfoRecordTx(redeemBlock feed_attributes.RedeemBlock) {
	_, err := redeemBlockInfoRecordExecutor.Tx.Exec(UPDATE_TOTAL_ENROLLED_MILESTONEPOINTS_COMMAND, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.RedeemBlockInfoRecordLocation)
		log.Printf("Failed to update totalEnrolledMilestonePoints for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully update totalEnrolledMilestonePoints for redeemBlock %d\n", redeemBlock)
}
