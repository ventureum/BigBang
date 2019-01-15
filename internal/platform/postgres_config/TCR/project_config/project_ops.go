package project_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type ProjectExecutor struct {
	client_config.PostgresBigBangClient
}

func (projectExecutor *ProjectExecutor) CreateProjectTable() {
	projectExecutor.CreateTimestampTrigger()
	projectExecutor.CreateTable(TABLE_SCHEMA_FOR_PROJECT, TABLE_NAME_FOR_PROJECT)
	projectExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_PROJECT)
}

func (projectExecutor *ProjectExecutor) DeleteProjectTable() {
	projectExecutor.DeleteTable(TABLE_NAME_FOR_PROJECT)
}

func (projectExecutor *ProjectExecutor) ClearProjectTable() {
	projectExecutor.ClearTable(TABLE_NAME_FOR_PROJECT)
}

func (projectExecutor *ProjectExecutor) UpsertProjectRecordTx(projectRecord *ProjectRecord) {

	res, err := projectExecutor.Tx.NamedExec(UPDATE_PROJECT_COMMAND, projectRecord)

	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectRecord.ProjectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to update project record: %+v with error:\n %+v", projectRecord, err)
		log.Panicln(errInfo.Marshal())
	}

	count, err := res.RowsAffected()

	if count == 0 {
		_, err = projectExecutor.Tx.NamedExec(INSERT_PROJECT_COMMAND, projectRecord)

		if err != nil {
			errInfo := error_config.MatchError(err, "projectId", projectRecord.ProjectId, error_config.ProjectRecordLocation)
			log.Printf("Failed to insert project record: %+v with error:\n %+v", projectRecord, err)
			log.Panicln(errInfo.Marshal())
		}
	}

	log.Printf("Sucessfully upserted project record for projectId %s\n", projectRecord.ProjectId)
}

func (projectExecutor *ProjectExecutor) DeleteProjectRecordTx(projectId string) {
	_, err := projectExecutor.Tx.Exec(DELETE_PROJECT_COMMAND, projectId)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to delete project record for projectId %s with error: %+v\n", projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted project record for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) GetProjectRecordTx(projectId string) *ProjectRecord {
	var projectRecord ProjectRecord
	err := projectExecutor.Tx.Get(&projectRecord, QUERY_PROJECT_COMMAND, projectId)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to get project record for projectId %s with error: %+v\n", projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	if err == sql.ErrNoRows {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoProjectIdExisting,
			ErrorData: map[string]interface{}{
				"projectId": projectId,
			},
			ErrorLocation: error_config.ProjectRecordLocation,
		}
		log.Printf("No project record for projectId %s", projectId)
		log.Panicln(errorInfo.Marshal())
	}
	return &projectRecord
}

func (projectExecutor *ProjectExecutor) VerifyProjectRecordExistingTx(projectId string) {
	existing := projectExecutor.CheckProjectRecordExistingTx(projectId)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoProjectIdExisting,
			ErrorData: map[string]interface{}{
				"projectId": projectId,
			},
			ErrorLocation: error_config.ProjectRecordLocation,
		}
		log.Printf("No project record for projectId %s", projectId)
		log.Panicln(errorInfo.Marshal())
	}
}

func (projectExecutor *ProjectExecutor) CheckProjectRecordExistingTx(projectId string) bool {
	var existing bool
	err := projectExecutor.Tx.Get(&existing, VERIFY_PROJECT_EXISTING_COMMAND, projectId)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to verify project record existing for projectId %s with error: %+v\n", projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	return existing
}

func (projectExecutor *ProjectExecutor) GetProjectRecordsByCursorTx(cursor string, limit int64) *[]ProjectRecord {
	var projectRecords []ProjectRecord
	var err error
	if cursor != "" {
		err = projectExecutor.Tx.Select(
			&projectRecords, PAGINATION_QUERY_PROJECT_LIST_COMMAND, cursor, limit)
	} else {
		err = projectExecutor.Tx.Select(
			&projectRecords, QUERY_PROJECT_LIST_COMMAND, limit)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Panicf("Failed to get project records by cursor %s and limit %d with error: %+v\n", cursor, limit, err)
	}
	return &projectRecords
}

func (projectExecutor *ProjectExecutor) AddRatingAndWeightForProjectTx(projectId string, deltaRating int64, deltaWeight int64) {
	_, err := projectExecutor.Tx.Exec(ADD_RATING_AND_WEIGHT_FOR_PROJECT_COMMAND, projectId, deltaRating*deltaWeight, deltaWeight)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to add rating and weight for projectId %s with error: %+v\n", projectId, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully added rating and weight for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) IncreaseNumMilestonesTx(projectId string) {
	_, err := projectExecutor.Tx.Exec(INCREASE_NUM_MILESTONES_COMMAND, projectId)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to increase numMilestones for projectId %s with error: %+v\n", projectId, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully increased numMilestones for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) DecreaseNumMilestonesTx(projectId string) {
	_, err := projectExecutor.Tx.Exec(DECREASE_NUM_MILESTONES_COMMAND, projectId)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to decrease numMilestones for projectId %s with error: %+v\n", projectId, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully decreased numMilestones for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) IncreaseNumMilestonesCompletedTx(projectId string) {
	_, err := projectExecutor.Tx.Exec(INCREASE_NUM_MILESTONES_COMPLETED_COMMAND, projectId)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to increase numMilestonesCompleted for projectId %s with error: %+v\n", projectId, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully increased numMilestonesCompleted for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) SetCurrentMilestoneTx(projectId string, milestoneId int64) {
	_, err := projectExecutor.Tx.Exec(SET_CURRENT_MILESTONE_COMMAND, projectId, milestoneId)

	if err != nil {
		errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		log.Printf("Failed to set current milestone %d for projectId %s with error: %+v\n", milestoneId, projectId, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Successfully set current milestone %d for projectId %s\n", milestoneId, projectId)
}

func (projectExecutor *ProjectExecutor) VerifyAdminExistingTx(projectId string, admin string) bool {
	var existing bool
	err := projectExecutor.Tx.Get(&existing, VERIFY_PROJECT_AND_ADMIN_EXISTING_COMMAND, projectId, admin)
	if err != nil {
		errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
		errInfo.ErrorData["admin"] = admin
		log.Printf("Failed to verify projectId %s and admin %s existing with error: %+v\n", projectId, admin, err)
		log.Panicln(errInfo.Marshal())
	}

	return existing
}

func (projectExecutor *ProjectExecutor) GetProjectIdByAdminTx(admin string) string {
	var projectId string
	err := projectExecutor.Tx.Get(&projectId, QUERY_PROJECT_ID_BY_ADMIN_COMMAND, admin)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "admin", admin, error_config.ProjectRecordLocation)
		log.Printf("Failed to get project id for admin %s with error: %+v\n", admin, err)
		log.Panicln(errInfo.Marshal())
	}

	if err == sql.ErrNoRows {
		projectId = ""
	}
	return projectId
}
