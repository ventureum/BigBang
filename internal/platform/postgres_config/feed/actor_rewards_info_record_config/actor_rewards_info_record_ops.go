package actor_rewards_info_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type ActorRewardsInfoRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) CreateActorRewardsInfoRecordTable() {
	actorRewardsInfoRecordExecutor.CreateTimestampTrigger()
	actorRewardsInfoRecordExecutor.CreateTable(
		TABLE_SCHEMA_FOR_ACTOR_REWARDS_INFO_RECORD, TABLE_NAME_FOR_ACTOR_REWARDS_INFO_RECORD)
	actorRewardsInfoRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_REWARDS_INFO_RECORD)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) DeleteActorRewardsInfoRecordTable() {
	actorRewardsInfoRecordExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_REWARDS_INFO_RECORD)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) VerifyActorExistingTx(actor string) {
	var existing bool
	err := actorRewardsInfoRecordExecutor.Tx.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Panicf("Failed to verify actor existing for actor %s with error: %+v\n", actor, err)
		log.Panicln(errorInfo.Marshal())
	}

	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoActorExisting,
			ErrorData: map[string]interface{}{
				"actor": actor,
			},
			ErrorLocation: error_config.ActorRewardsInfoRecordLocation,
		}
		log.Printf("No actor rewards info acount for actor %s", actor)
		log.Panicln(errorInfo.Marshal())
	}
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) DeleteActorRewardsInfoRecordsTx(actor string) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(DELETE_ACTOR_REWARDS_INFO_RECORD_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to delete rewards info records for actor %s with error: %+v\n", actor, err)
		log.Panicln(errorInfo.Marshal())
	}
	log.Printf("Sucessfully deleted rewards info records for actor %s\n", actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) GetActorRewardsInfoTx(
	actor string) *feed_attributes.RewardsInfo {
	var rewardsInfo feed_attributes.RewardsInfo
	err := actorRewardsInfoRecordExecutor.Tx.Get(&rewardsInfo, QUERY_ACTOR_REWARDS_INFO_COMMAND, actor)
	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to get rewards info for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}
	return &rewardsInfo
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) AddActorMilestonePointsTx(
	actor string, mpsToAdd feed_attributes.MilestonePoint) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(ADD_ACTOR_MILESTONE_POINTS_COMMAND, actor, mpsToAdd)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to add milestone points for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added milestone points %d for actor %s", mpsToAdd, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) AddActorMilestonePointsFromVotesTx(
	actor string, mpsToAdd feed_attributes.MilestonePoint) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(ADD_ACTOR_MILESTONE_POINTS_FROM_VOTES_COMMAND, actor, mpsToAdd)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to add milestone points from votes for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added milestone points from votes %d for actor %s", mpsToAdd, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) AddActorMilestonePointsFromPostsTx(
	actor string, mpsToAdd feed_attributes.MilestonePoint) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(ADD_ACTOR_MILESTONE_POINTS_FROM_POSTS_COMMAND, actor, mpsToAdd)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to add milestone points from posts for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added milestone points from posts %d for actor %s", mpsToAdd, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) AddActorMilestonePointsFromOthersTx(
	actor string, mpsToAdd feed_attributes.MilestonePoint) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(ADD_ACTOR_MILESTONE_POINTS_FROM_OTHERS_COMMAND, actor, mpsToAdd)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to add milestone points from others for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added milestone points from others %d for actor %s", mpsToAdd, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) AddActorReputationTx(
	actor string, reputationToAdd feed_attributes.Reputation) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(ADD_ACTOR_REPUTATIONS_COMMAND, actor, reputationToAdd)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to add reputaions for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added reputaions %d for actor %s", reputationToAdd, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) AddActorFuelTx(
	actor string, fuelToAdd feed_attributes.Fuel) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(ADD_ACTOR_FUEL_COMMAND, actor, fuelToAdd)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to add fuel for actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added fuel %d for actor %s", fuelToAdd, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) SubActorReputationTx(
	actor string, reputationToSub feed_attributes.Reputation) {
	var diff int64
	err := actorRewardsInfoRecordExecutor.Tx.Get(&diff, SUB_ACTOR_REPUTATION_COMMAND, actor, reputationToSub)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to substract reputaions from actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	if diff > 0 {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.InsufficientReputaionsAmount,
			ErrorData: map[string]interface{}{
				"diff": diff,
			},
			ErrorLocation: error_config.ActorRewardsInfoRecordLocation,
		}
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully substracted reputaions %d from actor %s", reputationToSub, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) GetTotalActorReputationTx() int64 {
	var totalReputations int64
	err := actorRewardsInfoRecordExecutor.Tx.Get(&totalReputations, QUARY_TOTAL_REPUTATION_COMMAND)
	if err != nil {
		errorInfo := error_config.MatchError(err, "", "", error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to get total reputation for all actors with error: %+v\n", err)
		log.Panic(errorInfo.Marshal())
	}
	return totalReputations
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) UpsertActorRewardsInfoRecordTx(
	actorRewardsInfoRecord *ActorRewardsInfoRecord) {
	_, err := actorRewardsInfoRecordExecutor.Tx.NamedExec(
		UPSERT_ACTOR_REWARDS_INFO_RECORD_COMMAND, actorRewardsInfoRecord)
	if err != nil {
		errorInfo := error_config.MatchError(err, "", "", error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to upsert rewards info record: %+v with error: %+v\n", actorRewardsInfoRecord, err)
		log.Panic(errorInfo.Marshal())
	}
	log.Printf("Sucessfully upserted rewards info record for actor %s\n", actorRewardsInfoRecord.Actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) SubActorFuelTx(
	actor string, fuelToSub feed_attributes.Fuel) {
	var diff int64
	err := actorRewardsInfoRecordExecutor.Tx.Get(&diff, SUB_ACTOR_FUEL_COMMAND, actor, fuelToSub)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actor, error_config.ActorRewardsInfoRecordLocation)
		log.Printf("Failed to substract fuel from actor %s with error: %+v\n", actor, err)
		log.Panic(errorInfo.Marshal())
	}

	if diff > 0 {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.InsufficientFuelAmount,
			ErrorData: map[string]interface{}{
				"diff": diff,
			},
			ErrorLocation: error_config.ActorRewardsInfoRecordLocation,
		}
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully substracted fuel %d from actor %s", fuelToSub, actor)
}

func (actorRewardsInfoRecordExecutor *ActorRewardsInfoRecordExecutor) ResetAllActorFuelTx(maxFuel feed_attributes.Fuel) {
	_, err := actorRewardsInfoRecordExecutor.Tx.Exec(RESET_ALL_ACTOR_FUEL_COMMAND, maxFuel)

	if err != nil {
		errorInfo := error_config.MatchError(err, "maxFuel", maxFuel, error_config.ActorRewardsInfoRecordLocation)
		log.Panicf("Failed to reset all actor fuel to be %d with error: %+v\n", maxFuel, err)
		log.Panicln(errorInfo.Marshal())
	}
}
