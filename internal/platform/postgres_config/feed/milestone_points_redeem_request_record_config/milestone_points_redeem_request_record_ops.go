package milestone_points_redeem_request_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type MilestonePointsRedeemRequestRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) CreateMilestonePointsRedeemRequestRecordTable() {
	milestonePointsRedeemRequestRecordExecutor.CreateTimestampTrigger()
	milestonePointsRedeemRequestRecordExecutor.CreateTable(
		TABLE_SCHEMA_FOR_MILESTONE_POINTS_REDEEM_REQUEST_RECORD, TABLE_NAME_FOR_MILESTONE_POINTS_REDEEM_REQUEST_RECORD)
	milestonePointsRedeemRequestRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_MILESTONE_POINTS_REDEEM_REQUEST_RECORD)
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) DeleteMilestonePointsRedeemRequestRecordTable() {
	milestonePointsRedeemRequestRecordExecutor.DeleteTable(TABLE_NAME_FOR_MILESTONE_POINTS_REDEEM_REQUEST_RECORD)
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) UpsertMilestonePointsRedeemRequestRecordTx(
	milestonePointsRedeemRequestRecord *MilestonePointsRedeemRequestRecord) {
	_, err := milestonePointsRedeemRequestRecordExecutor.Tx.NamedExec(
		UPSERT_MILESTONE_POINTS_REDEEM_REQUEST_RECORD_COMMAND, milestonePointsRedeemRequestRecord)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", milestonePointsRedeemRequestRecord.Actor, error_config.MilestonePointsRedeemRequestRecordLocation)
		log.Printf("Failed to upsert milestone points redeem request record: %+v with error: %+v\n", milestonePointsRedeemRequestRecord, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully upserted milestone points redeem request record for actor %s\n", milestonePointsRedeemRequestRecord.Actor)
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) VerifyMilestonePointsRedeemRequestExistingTx(actor string) {
	var existing bool
	err := milestonePointsRedeemRequestRecordExecutor.Tx.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.MilestonePointsRedeemRequestRecordLocation)
		log.Printf("Failed to verify actor milestone points redeem request record existing for actor %s with error: %+v\n", actor, err)
		log.Panicln(errorInfo.Marshal())
	}

	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoActorExisting,
			ErrorData: map[string]interface{}{
				"actor": actor,
			},
			ErrorLocation: error_config.MilestonePointsRedeemRequestRecordLocation,
		}
		log.Printf("No milestone points redeem request record for actor %s", actor)
		log.Panicln(errorInfo.Marshal())
	}
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) DeleteMilestonePointsRedeemRequestRecordTx(actor string) {
	_, err := milestonePointsRedeemRequestRecordExecutor.Tx.Exec(DELETE_MILESTONE_POINTS_REDEEM_REQUEST_RECORD_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.MilestonePointsRedeemRequestRecordLocation)
		log.Printf("Failed to delete milestone points redeem request record for actor %s with error: %+v\n", actor, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully deleted milestone points redeem request record for actor %s\n", actor)
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) GetMilestonePointsRedeemRequestTx(
	actor string) *feed_attributes.MilestonePointsRedeemRequest {
	var milestonePointsRedeemRequest feed_attributes.MilestonePointsRedeemRequest
	err := milestonePointsRedeemRequestRecordExecutor.Tx.Get(&milestonePointsRedeemRequest, QUERY_MILESTONE_POINTS_REDEEM_REQUEST_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.MilestonePointsRedeemRequestRecordLocation)
		log.Printf("Failed to get ActualMilestonePoints Redeem Request for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}
	return &milestonePointsRedeemRequest
}

func (milestonePointsRedeemRequestRecordExecutor *MilestonePointsRedeemRequestRecordExecutor) GetTotalEnrolledMilestonePointsTx(
	nextRedeemBlock int64) int64 {
	var totalEnrolledMilestonePoints int64
	err := milestonePointsRedeemRequestRecordExecutor.Tx.Get(&totalEnrolledMilestonePoints, QUERY_TOTAL_ENROLLED_MILESTONE_POINTS_COMMAND, nextRedeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "nextRedeemBlock", nextRedeemBlock, error_config.MilestonePointsRedeemRequestRecordLocation)
		log.Printf("Failed to get total enrolled ActualMilestonePoints for nextRedeemBlock %d with error: %+v\n", nextRedeemBlock, err)
		log.Panic(errorInfo.Marshal())
	}
	return totalEnrolledMilestonePoints
}
