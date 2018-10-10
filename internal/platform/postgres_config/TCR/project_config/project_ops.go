package project_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "database/sql"
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

func (projectExecutor *ProjectExecutor) UpsertProjectRecord(projectRecord *ProjectRecord) {

  res, err := projectExecutor.C.NamedExec(UPDATE_PROJECT_COMMAND, projectRecord)

  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectRecord.ProjectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to update project record: %+v with error:\n %+v", projectRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  count, err := res.RowsAffected()

  if count == 0 {
    _, err = projectExecutor.C.NamedExec(INSERT_PROJECT_COMMAND, projectRecord)

    if err != nil {
      errInfo := error_config.MatchError(err, "projectId", projectRecord.ProjectId, error_config.ProjectRecordLocation)
      log.Printf("Failed to insert project record: %+v with error:\n %+v", projectRecord, err)
      log.Panicln(errInfo.Marshal())
    }
  }

  log.Printf("Sucessfully upserted project record for projectId %s\n", projectRecord.ProjectId)
}

func (projectExecutor *ProjectExecutor) DeleteProjectRecord(projectId string) {
  _, err := projectExecutor.C.Exec(DELETE_PROJECT_COMMAND, projectId)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to delete project record for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted project record for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) GetProjectRecord(projectId string) *ProjectRecord {
  var projectRecord ProjectRecord
  err := projectExecutor.C.Get(&projectRecord, QUERY_PROJECT_COMMAND, projectId)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to get project record for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoProjectIdExisting,
      ErrorData: map[string]interface{} {
        "projectId": projectId,
      },
      ErrorLocation: error_config.ProjectRecordLocation,
    }
    log.Printf("No project record for projectId %s", projectId)
    log.Panicln(errorInfo.Marshal())
  }
  return &projectRecord
}

func (projectExecutor *ProjectExecutor) VerifyProjectRecordExisting (projectId string) {
  existing := projectExecutor.CheckProjectRecordExisting(projectId)
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoProjectIdExisting,
      ErrorData: map[string]interface{} {
        "projectId": projectId,
      },
      ErrorLocation: error_config.ProjectRecordLocation,
    }
    log.Printf("No project record for projectId %s", projectId)
    log.Panicln(errorInfo.Marshal())
  }
}

func (projectExecutor *ProjectExecutor) CheckProjectRecordExisting (projectId string) bool {
  var existing bool
  err := projectExecutor.C.Get(&existing, VERIFY_PROJECT_EXISTING_COMMAND, projectId)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to verify project record existing for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }
  return existing
}

func (projectExecutor *ProjectExecutor) GetProjectRecordsByCursor(cursor string, limit int64) *[]ProjectRecord {
  var projectRecords []ProjectRecord
  var err error
  if cursor !=  "" {
    err = projectExecutor.C.Select(
      &projectRecords, PAGINATION_QUERY_PROJECT_LIST_COMMAND, cursor, limit)
  } else {
    err = projectExecutor.C.Select(
      &projectRecords, QUERY_PROJECT_LIST_COMMAND, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get project records by cursor %s and limit %d with error: %+v\n", cursor, limit, err)
  }
  return &projectRecords
}

func (projectExecutor *ProjectExecutor) AddRatingAndWeight(projectId string, deltaRating int64, deltaWeight int64) {
  _, err := projectExecutor.C.Exec(ADD_RATING_AND_WEIGHT_COMMAND, projectId, deltaRating, deltaWeight)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to add rating and weight for projectId %s with error: %+v\n", projectId, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully added rating and weight for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) IncreaseNumMilestones(projectId string) {
  _, err := projectExecutor.C.Exec(INCREASE_NUM_MILESTONES_COMMAND, projectId)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to increase numMilestones for projectId %s with error: %+v\n", projectId, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully increased numMilestones for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) DecreaseNumMilestones(projectId string) {
  _, err := projectExecutor.C.Exec(DECREASE_NUM_MILESTONES_COMMAND, projectId)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to decrease numMilestones for projectId %s with error: %+v\n", projectId, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully decreased numMilestones for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) IncreaseNumMilestonesCompleted(projectId string) {
  _, err := projectExecutor.C.Exec(INCREASE_NUM_MILESTONES_COMPLETED_COMMAND, projectId)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to increase numMilestonesCompleted for projectId %s with error: %+v\n", projectId, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully increased numMilestonesCompleted for projectId %s\n", projectId)
}

func (projectExecutor *ProjectExecutor) SetCurrentMilestone(projectId string, milestoneId int64) {
  _, err := projectExecutor.C.Exec(SET_CURRENT_MILESTONE_COMMAND, projectId, milestoneId)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to set current milestone %d for projectId %s with error: %+v\n", milestoneId, projectId, err)
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully set current milestone %d for projectId %s\n", milestoneId, projectId)
}

/*
 * Tx versions
 */

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
      ErrorData: map[string]interface{} {
        "projectId": projectId,
      },
      ErrorLocation: error_config.ProjectRecordLocation,
    }
    log.Printf("No project record for projectId %s", projectId)
    log.Panicln(errorInfo.Marshal())
  }
  return &projectRecord
}

func (projectExecutor *ProjectExecutor) VerifyProjectRecordExistingTx (projectId string) {
  existing := projectExecutor.CheckProjectRecordExistingTx(projectId)
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoProjectIdExisting,
      ErrorData: map[string]interface{} {
        "projectId": projectId,
      },
      ErrorLocation: error_config.ProjectRecordLocation,
    }
    log.Printf("No project record for projectId %s", projectId)
    log.Panicln(errorInfo.Marshal())
  }
}

func (projectExecutor *ProjectExecutor) CheckProjectRecordExistingTx (projectId string) bool {
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

func (projectExecutor *ProjectExecutor) AddRatingAndWeightTx(projectId string, deltaRating int64, deltaWeight int64) {
  _, err := projectExecutor.Tx.Exec(ADD_RATING_AND_WEIGHT_COMMAND, projectId, deltaRating, deltaWeight)

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
