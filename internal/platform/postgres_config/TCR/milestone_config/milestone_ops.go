package milestone_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type MilestoneExecutor struct {
	client_config.PostgresBigBangClient
}

func (milestoneExecutor *MilestoneExecutor) CreateMilestoneTable() {
	milestoneExecutor.LoadMilestoneStateEnum()
	milestoneExecutor.CreateTimestampTrigger()
	milestoneExecutor.CreateTable(TABLE_SCHEMA_FOR_MILESTONE, TABLE_NAME_FOR_MILESTONE)
	milestoneExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_MILESTONE)
}

func (milestoneExecutor *MilestoneExecutor) DeleteMilestoneTable() {
	milestoneExecutor.DeleteTable(TABLE_NAME_FOR_MILESTONE)
	milestoneExecutor.DropMilestoneStateEnum()
}

func (milestoneExecutor *MilestoneExecutor) ClearMilestoneTable() {
	milestoneExecutor.ClearTable(TABLE_NAME_FOR_MILESTONE)
}

func (milestoneExecutor *MilestoneExecutor) UpsertMilestoneRecordTx(milestoneRecord *MilestoneRecord) bool {
	res, err := milestoneExecutor.Tx.NamedQuery(UPSERT_MILESTONE_COMMAND, milestoneRecord)
	if err != nil {
		errInfo := error_config.MatchError(err, "milestoneId", milestoneRecord.MilestoneId, error_config.MilestoneRecordLocation)
		errInfo.ErrorData["projectId"] = milestoneRecord.ProjectId
		log.Printf("Failed to upsert milestone record for projectId %s and milestoneId %d with error:\n %+v",
			milestoneRecord.ProjectId, milestoneRecord.MilestoneId, err)
		log.Panicln(errInfo.Marshal())
	}

	log.Printf("Sucessfully upserted milestone record for projectId %s and milestoneId %d\n",
		milestoneRecord.ProjectId, milestoneRecord.MilestoneId)

	var inserted sql.NullBool
	for res.Next() {
		err = res.Scan(&inserted)
	}
	return inserted.Bool
}

func (milestoneExecutor *MilestoneExecutor) DeleteMilestoneRecordByIDsTx(
	projectId string, milestoneId int64) {
	_, err := milestoneExecutor.Tx.Exec(DELETE_MILESTONE_BY_IDS_COMMAND, projectId, milestoneId)
	if err != nil {
		errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to delete milestone record for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errInfo.ErrorData["projectId"] = projectId
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted milestone record for projectId %s and milestoneId %d\n",
		projectId, milestoneId)
}

func (milestoneExecutor *MilestoneExecutor) DeleteMilestoneRecordsByProjectIdTx(projectId string) {
	_, err := milestoneExecutor.Tx.Exec(DELETE_MILESTONES_BY_PROJECT_ID_COMMAND, projectId)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to delete milestone records for projectId %s with error: %+v\n",
			projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted milestone record for projectId %s\n", projectId)
}

func (milestoneExecutor *MilestoneExecutor) GetMilestoneRecordByIDsTx(
	projectId string, milestoneId int64) *MilestoneRecord {
	var milestoneRecord MilestoneRecord
	err := milestoneExecutor.Tx.Get(&milestoneRecord, QUERY_MILESTONE_BY_IDS_COMMAND, projectId, milestoneId)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to get milestone record for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errInfo.ErrorData["projectId"] = projectId

		log.Panicln(errInfo.Marshal())
	}

	if err == sql.ErrNoRows {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoMilestoneIdExisting,
			ErrorData: map[string]interface{}{
				"milestoneId": milestoneId,
				"projectId":   projectId,
			},
			ErrorLocation: error_config.MilestoneRecordLocation,
		}
		log.Printf("No milestone record for projectId %s and milestoneId %d", projectId, milestoneId)
		log.Panicln(errorInfo.Marshal())
	}
	return &milestoneRecord
}

func (milestoneExecutor *MilestoneExecutor) GetMilestonesRecordsByProjectIdTx(
	projectId string) *[]MilestoneRecord {
	var milestoneRecords []MilestoneRecord
	err := milestoneExecutor.Tx.Select(&milestoneRecords, QUERY_MILESTONES_BY_PROJECT_ID_COMMAND, projectId)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to get milestone records for projectId %s with error: %+v\n",
			projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	return &milestoneRecords
}

func (milestoneExecutor *MilestoneExecutor) VerifyMilestoneRecordExistingTx(
	projectId string, milestoneId int64) {
	existing := milestoneExecutor.CheckMilestoneRecordExistingTx(projectId, milestoneId)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoMilestoneIdExisting,
			ErrorData: map[string]interface{}{
				"milestoneId": milestoneId,
				"projectId":   projectId,
			},
			ErrorLocation: error_config.MilestoneRecordLocation,
		}
		log.Printf("No milestone record for projectId %s and milestoneId %d", projectId, milestoneId)
		log.Panicln(errorInfo.Marshal())
	}
}

func (milestoneExecutor *MilestoneExecutor) CheckMilestoneRecordExistingTx(
	projectId string, milestoneId int64) bool {
	var existing bool
	err := milestoneExecutor.Tx.Get(&existing, VERIFY_MILESTONE_EXISTING_COMMAND, projectId, milestoneId)
	if err != nil {
		errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to verify milestone record existing for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errInfo.ErrorData["projectId"] = projectId
		log.Panicln(errInfo.Marshal())
	}
	return existing
}

func (milestoneExecutor *MilestoneExecutor) ValidateMilestoneRecordUpdatingTx(
	projectId string, milestoneId int64) bool {
	var invalid bool
	err := milestoneExecutor.Tx.Get(&invalid, INVALID_MILESTONE_UPDATE_IF_EXISTING_COMMAND, projectId, milestoneId)
	if err != nil {
		errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to validate milestone record update for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errInfo.ErrorData["projectId"] = projectId
		log.Panicln(errInfo.Marshal())
	}
	return invalid
}

func (milestoneExecutor *MilestoneExecutor) IncreaseNumObjectivesTx(projectId string, milestoneId int64) {
	_, err := milestoneExecutor.Tx.Exec(INCREASE_NUM_OBJECTIVES_COMMAND, projectId, milestoneId)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to increase numObjectives for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errorInfo.ErrorData["milestoneId"] = milestoneId
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully increased numObjectives for projectId %s and milestoneId %d\n", projectId, milestoneId)
}

func (milestoneExecutor *MilestoneExecutor) DecreaseNumObjectivesTx(projectId string, milestoneId int64) {
	_, err := milestoneExecutor.Tx.Exec(DECREASE_NUM_OBJECTIVES_COMMAND, projectId, milestoneId)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to decrease numObjectives for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errorInfo.ErrorData["milestoneId"] = milestoneId
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully decreased numObjectives for projectId %s and and milestoneId %d\n", projectId, milestoneId)
}

func (milestoneExecutor *MilestoneExecutor) AddRatingAndWeightForMilestoneTx(
	projectId string, milestoneId int64, deltaRating int64, deltaWeight int64) {
	_, err := milestoneExecutor.Tx.Exec(
		ADD_RATING_AND_WEIGHT_FOR_MILESTONE_COMMAND, projectId, milestoneId, deltaRating*deltaWeight, deltaWeight)

	if err != nil {
		errorInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to add rating and weight for projectId %s, milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errorInfo.ErrorData["projectId"] = projectId
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added rating and weight for projectId %s, milestoneId %d\n",
		projectId, milestoneId)
}

func (milestoneExecutor *MilestoneExecutor) ActivateMilestoneTx(
	projectId string, milestoneId int64, blockTimestamp int64, startTime int64) {
	_, err := milestoneExecutor.Tx.Exec(ACTIVATE_MILESTONE_COMMAND, projectId, milestoneId, blockTimestamp, startTime)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
		errInfo.ErrorData["milestoneId"] = milestoneId
		errInfo.ErrorData["blockTimestamp"] = blockTimestamp
		errInfo.ErrorData["startTime"] = startTime
		log.Printf("Failed to activate milestone for projectId %s and milstoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully activate milestone for projectId %s and  milstoneId %d\n", projectId, milestoneId)
}

func (milestoneExecutor *MilestoneExecutor) FinalizeMilestoneTx(
	projectId string, milestoneId int64, blockTimestamp int64, endTime int64) {
	_, err := milestoneExecutor.Tx.Exec(FINALIZE_MILESTONE_COMMAND, projectId, milestoneId, blockTimestamp, endTime)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
		errInfo.ErrorData["milestoneId"] = milestoneId
		errInfo.ErrorData["blockTimestamp"] = blockTimestamp
		errInfo.ErrorData["endTime"] = endTime
		log.Printf("Failed to finalize milestone for projectId %s and milstoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully finalize milestone for projectId %s and  milstoneId %d\n", projectId, milestoneId)
}
