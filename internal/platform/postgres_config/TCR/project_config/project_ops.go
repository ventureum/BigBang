package project_config

import (
  "log"
  "time"
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

func (projectExecutor *ProjectExecutor) UpsertProjectRecord(projectRecord *ProjectRecord) time.Time {
  res, err := projectExecutor.C.NamedQuery(UPSERT_PROJECT_COMMAND, projectRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectRecord.ProjectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to upsert project record: %+v with error:\n %+v", projectRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted project record for projectId %s\n", projectRecord.ProjectId)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
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
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to get project record for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }
  return &projectRecord
}

func (projectExecutor *ProjectExecutor) VerifyProjectRecordExisting (projectId string) {
  var existing bool
  err := projectExecutor.C.Get(&existing, VERIFY_PROJECT_EXISTING_COMMAND, projectId)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to verify project record existing for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }
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

func (projectExecutor *ProjectExecutor) GetProjectRecordsByCursor(cursor int64, limit int64) *[]ProjectRecord {
  var projectRecords []ProjectRecord
  var err error
  if cursor > 0 {
    err = projectExecutor.C.Select(
      &projectRecords, PAGINATION_QUERY_PROJECT_LIST_COMMAND, cursor, limit)
  } else {
    err = projectExecutor.C.Select(
      &projectRecords, QUERY_PROJECT_LIST_COMMAND, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get project records for cursor %d and limit %d with error: %+v\n", cursor, limit, err)
  }
  return &projectRecords
}

/*
 * Tx versions
 */
func (projectExecutor *ProjectExecutor) UpsertProjectRecordTx(projectRecord *ProjectRecord) time.Time {
  res, err := projectExecutor.Tx.NamedQuery(UPSERT_PROJECT_COMMAND, projectRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectRecord.ProjectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to upsert project record: %+v with error:\n %+v", projectRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted project record for projectId %s\n", projectRecord.ProjectId)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
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
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to get project record for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }
  return &projectRecord
}

func (projectExecutor *ProjectExecutor) VerifyProjectRecordExistingTx (projectId string) {
  var existing bool
  err := projectExecutor.Tx.Get(&existing, VERIFY_PROJECT_EXISTING_COMMAND, projectId)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.ProjectRecordLocation)
    log.Printf("Failed to verify project record existing for projectId %s with error: %+v\n", projectId, err)
    log.Panicln(errInfo.Marshal())
  }
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

func (projectExecutor *ProjectExecutor) GetProjectRecordsByCursorTx(cursor int64, limit int64) *[]ProjectRecord {
  var projectRecords []ProjectRecord
  var err error
  if cursor == 0 {
    err = projectExecutor.Tx.Select(
      &projectRecords, PAGINATION_QUERY_PROJECT_LIST_COMMAND, cursor, limit)
  } else {
    err = projectExecutor.Tx.Select(
      &projectRecords, QUERY_PROJECT_LIST_COMMAND, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get project records for cursor %d and limit %d with error: %+v\n", cursor, limit, err)
  }
  return &projectRecords
}
