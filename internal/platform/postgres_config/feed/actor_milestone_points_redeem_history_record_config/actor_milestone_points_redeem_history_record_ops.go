package actor_milestone_points_redeem_history_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type ActorMilestonePointsRedeemHistoryRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) CreateActorMilestonePointsRedeemHistoryRecordTable() {
	actorMilestonePointsRedeemHistoryRecordExecutor.CreateTimestampTrigger()
	actorMilestonePointsRedeemHistoryRecordExecutor.CreateTable(
		TABLE_SCHEMA_FOR_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD, TABLE_NAME_FOR_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD)
	actorMilestonePointsRedeemHistoryRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD)
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) DeleteActorMilestonePointsRedeemHistoryRecordTable() {
	actorMilestonePointsRedeemHistoryRecordExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD)
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) UpsertActorMilestonePointsRedeemHistoryRecordTx(
	actorMilestonePointsRedeemHistoryRecord *ActorMilestonePointsRedeemHistoryRecord) {
	_, err := actorMilestonePointsRedeemHistoryRecordExecutor.Tx.NamedExec(
		UPSERT_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD_COMMAND, actorMilestonePointsRedeemHistoryRecord)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actorMilestonePointsRedeemHistoryRecord.Actor, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		log.Printf("Failed to upsert actor milestone points redeem history record: %+v with error: %+v\n", actorMilestonePointsRedeemHistoryRecord, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully upserted actor milestone points redeem history record for actor %s\n", actorMilestonePointsRedeemHistoryRecord.Actor)
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) VerifyRedeemBlockExistingTx(redeemBlock feed_attributes.RedeemBlock) bool {
	var existing bool
	err := actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Get(&existing, VERIFY_REDEEM_BLOCK_EXISTING_COMMAND, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		log.Printf("Failed to verify redeem block existing for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}

	return existing
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) VerifyRedeemBlockForActorExistingTx(actor string, redeemBlock feed_attributes.RedeemBlock) bool {
	var existing bool
	err := actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Get(&existing, VERIFY_REDEEM_BLOCK_FOR_ACTOR_EXISTING_COMMAND, actor, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		errorInfo.ErrorData["actor"] = actor
		log.Printf("Failed to verify redeem block for actor existing for actor %s and redeemBlock %d with error: %+v\n", actor, redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}

	return existing
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) DeleteActorMilestonePointsRedeemHistoryRecordByActorTx(actor string) {
	_, err := actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Exec(DELETE_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD_BY_ACTOR_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		log.Printf("Failed to delete actor milestone points redeem history records for actor %s with error: %+v\n", actor, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully deleted actor milestone points redeem history record for actor %s\n", actor)
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) GetActorMilestonePointsRedeemHistoryTx(
	actor string) *feed_attributes.MilestonePointsRedeemHistory {
	var milestonePointsRedeemHistory feed_attributes.MilestonePointsRedeemHistory
	err := actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Get(&milestonePointsRedeemHistory, QUERY_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORDS_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		log.Printf("Failed to get actor milestone points redeem history record for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}
	return &milestonePointsRedeemHistory
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) UpsertBatchActorMilestonePointsRedeemHistoryRecordByRedeemBlockTx(redeemBlock feed_attributes.RedeemBlock) {
	_, err := actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Exec(UPSERT_BATCH_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORDS_BY_REDEEM_BLOCK, redeemBlock)
	if err != nil {
		errorInfo := error_config.MatchError(err, "redeemBlock", redeemBlock, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		log.Printf("Failed to upsert batch actor milestone points redeem history records for redeemBlock %d with error: %+v\n", redeemBlock, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully upserted batch actor milestone points redeem history records for redeemBlock %d\n", redeemBlock)
}

func (actorMilestonePointsRedeemHistoryRecordExecutor *ActorMilestonePointsRedeemHistoryRecordExecutor) GetActorMilestonePointsRedeemHistoryByCursorTx(
	actor string, cursor string, limit int64) *[]feed_attributes.MilestonePointsRedeemHistory {
	var redeems []feed_attributes.MilestonePointsRedeemHistory
	var err error
	if cursor != "" {
		err = actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Select(
			&redeems,
			PAGINATION_QUERY_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_COMMAND,
			actor, cursor, limit)
	} else {
		err = actorMilestonePointsRedeemHistoryRecordExecutor.Tx.Select(
			&redeems,
			QUERY_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_WITH_LIMIT_COMMAND,
			actor, limit)
	}

	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorMilestonePointsRedeemHistoryRecordLocation)
		log.Printf("Failed to get actor milestone points redeem history by cursor %s and limit %d for actor %s with error: %+v\n",
			cursor, limit, actor, err)
		errInfo.ErrorData["cursor"] = cursor
		errInfo.ErrorData["limit"] = limit
		log.Panicln(errInfo.Marshal())
	}
	return &redeems
}
