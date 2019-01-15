package objective_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type ObjectiveExecutor struct {
	client_config.PostgresBigBangClient
}

func (objectiveExecutor *ObjectiveExecutor) CreateObjectiveTable() {
	objectiveExecutor.CreateTimestampTrigger()
	objectiveExecutor.CreateTable(TABLE_SCHEMA_FOR_OBJECTIVE, TABLE_NAME_FOR_OBJECTIVE)
	objectiveExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_OBJECTIVE)
}

func (objectiveExecutor *ObjectiveExecutor) DeleteObjectiveTable() {
	objectiveExecutor.DeleteTable(TABLE_NAME_FOR_OBJECTIVE)
}

func (objectiveExecutor *ObjectiveExecutor) ClearObjectiveTable() {
	objectiveExecutor.ClearTable(TABLE_NAME_FOR_OBJECTIVE)
}

func (objectiveExecutor *ObjectiveExecutor) UpsertObjectiveRecordTx(objectiveRecord *ObjectiveRecord) bool {
	res, err := objectiveExecutor.Tx.NamedQuery(UPSERT_OBJECTIVE_COMMAND, objectiveRecord)
	if err != nil {
		errInfo := error_config.MatchError(err, "objectiveId", objectiveRecord.ObjectiveId, error_config.ObjectiveRecordLocation)
		errInfo.ErrorData["milestoneId"] = objectiveRecord.MilestoneId
		errInfo.ErrorData["projectId"] = objectiveRecord.ProjectId
		log.Printf("Failed to upsert objective record: %+v with error:\n %+v", objectiveRecord, err)
		log.Panicln(errInfo.Marshal())
	}

	log.Printf("Sucessfully upserted objective record for objectiveId %s\n", objectiveRecord.ObjectiveId)

	var inserted sql.NullBool
	for res.Next() {
		err = res.Scan(&inserted)
	}
	return inserted.Bool
}

func (objectiveExecutor *ObjectiveExecutor) DeleteObjectiveRecordByIDsTx(
	projectId string, milestoneId int64, objectiveId int64) {
	_, err := objectiveExecutor.Tx.Exec(DELETE_OBJECTIVE_BY_IDS_COMMAND, projectId, milestoneId, objectiveId)
	if err != nil {
		errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to delete objective record for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
			projectId, milestoneId, objectiveId, err)
		errInfo.ErrorData["milestoneId"] = milestoneId
		errInfo.ErrorData["projectId"] = projectId
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted objective record for projectId %s, milestoneId %d and objectiveId %d\n",
		projectId, milestoneId, objectiveId)
}

func (objectiveExecutor *ObjectiveExecutor) DeleteObjectiveRecordsByProjectIdAndMilestoneIdTx(
	projectId string, milestoneId int64) {
	_, err := objectiveExecutor.Tx.Exec(DELETE_OBJECTIVES_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND, projectId, milestoneId)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to delete objective records for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errInfo.ErrorData["milestoneId"] = milestoneId
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted objective record for projectId %s and milestoneId %d\n",
		projectId, milestoneId)
}

func (objectiveExecutor *ObjectiveExecutor) GetObjectiveRecordByIDsTx(
	projectId string, milestoneId int64, objectiveId int64) *ObjectiveRecord {
	var objectiveRecord ObjectiveRecord
	err := objectiveExecutor.Tx.Get(&objectiveRecord, QUERY_OBJECTIVE_BY_IDS_COMMAND, projectId, milestoneId, objectiveId)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to get objective record for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
			projectId, milestoneId, objectiveId, err)
		errInfo.ErrorData["milestoneId"] = milestoneId
		errInfo.ErrorData["projectId"] = projectId

		log.Panicln(errInfo.Marshal())
	}

	if err == sql.ErrNoRows {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoObjectiveIdExisting,
			ErrorData: map[string]interface{}{
				"objectiveId": objectiveId,
				"milestoneId": milestoneId,
				"projectId":   projectId,
			},
			ErrorLocation: error_config.ObjectiveRecordLocation,
		}
		log.Printf("No objective record for projectId %s, milestoneId %d and objectiveId %d", projectId, milestoneId, objectiveId)
		log.Panicln(errorInfo.Marshal())
	}
	return &objectiveRecord
}

func (objectiveExecutor *ObjectiveExecutor) GetObjectiveRecordsByProjectIdAndMilestoneIdTx(
	projectId string, milestoneId int64) *[]ObjectiveRecord {
	var objectiveRecords []ObjectiveRecord
	err := objectiveExecutor.Tx.Select(&objectiveRecords, QUERY_OBJECTIVES_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND,
		projectId, milestoneId)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to get objective records for projectId %s and milestoneId %d with error: %+v\n",
			projectId, milestoneId, err)
		errInfo.ErrorData["milestoneId"] = milestoneId

		log.Panicln(errInfo.Marshal())
	}
	return &objectiveRecords
}

func (objectiveExecutor *ObjectiveExecutor) VerifyObjectiveRecordExistingTx(
	projectId string, milestoneId int64, objectiveId int64) {
	existing := objectiveExecutor.CheckObjectiveRecordExistingTx(projectId, milestoneId, objectiveId)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoObjectiveIdExisting,
			ErrorData: map[string]interface{}{
				"objectiveId": objectiveId,
				"milestoneId": milestoneId,
				"projectId":   projectId,
			},
			ErrorLocation: error_config.ObjectiveRecordLocation,
		}
		log.Printf("No objective record for projectId %s, milestoneId %d and objectiveId %d", projectId, milestoneId, objectiveId)
		log.Panicln(errorInfo.Marshal())
	}
}

func (objectiveExecutor *ObjectiveExecutor) CheckObjectiveRecordExistingTx(
	projectId string, milestoneId int64, objectiveId int64) bool {
	var existing bool
	err := objectiveExecutor.Tx.Get(&existing, VERIFY_OBJECTIVE_EXISTING_COMMAND, projectId, milestoneId, objectiveId)
	if err != nil {
		errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to verify objective record existing for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
			projectId, milestoneId, objectiveId, err)
		errInfo.ErrorData["milestoneId"] = milestoneId
		errInfo.ErrorData["projectId"] = projectId
		log.Panicln(errInfo.Marshal())
	}

	return existing
}

func (objectiveExecutor *ObjectiveExecutor) AddRatingAndWeightForObjectiveTx(
	projectId string, milestoneId int64, objectiveId int64, deltaRating int64, deltaWeight int64) {
	_, err := objectiveExecutor.Tx.Exec(
		ADD_RATING_AND_WEIGHT_FOR_OBJECTIVE_COMMAND, projectId, milestoneId, objectiveId, deltaRating*deltaWeight, deltaWeight)

	if err != nil {
		errorInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.ObjectiveRecordLocation)
		log.Printf("Failed to add rating and weight for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
			projectId, milestoneId, objectiveId, err)
		errorInfo.ErrorData["milestoneId"] = milestoneId
		errorInfo.ErrorData["projectId"] = projectId
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added rating and weight for projectId %s, milestoneId %d and objectiveId %d\n",
		projectId, milestoneId, objectiveId)
}
