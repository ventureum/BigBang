package milestone_config

import (
  "log"
  "time"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "database/sql"
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

func (milestoneExecutor *MilestoneExecutor) UpsertMilestoneRecord(milestoneRecord *MilestoneRecord) time.Time {
  res, err := milestoneExecutor.C.NamedQuery(UPSERT_MILESTONE_COMMAND, milestoneRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "milestoneId", milestoneRecord.MilestoneId, error_config.MilestoneRecordLocation)
    errInfo.ErrorData["projectId"] = milestoneRecord.ProjectId
    log.Printf("Failed to upsert milestone record: %+v with error:\n %+v", milestoneRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted milestone record for milestoneId %s\n", milestoneRecord.MilestoneId)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
}

func (milestoneExecutor *MilestoneExecutor) DeleteMilestoneRecordByIDs(
    projectId string, milestoneId int64) {
  _, err := milestoneExecutor.C.Exec(DELETE_MILESTONE_BY_IDS_COMMAND, projectId, milestoneId)
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

func (milestoneExecutor *MilestoneExecutor) DeleteMilestoneRecordsByProjectId(projectId string) {
  _, err := milestoneExecutor.C.Exec(DELETE_MILESTONES_BY_PROJECT_ID_COMMAND, projectId)
  if err != nil {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
    log.Printf("Failed to delete milestone records for projectId %s with error: %+v\n",
      projectId, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted milestone record for projectId %s\n", projectId)
}

func (milestoneExecutor *MilestoneExecutor) GetMilestoneRecordByIDs(
    projectId string, milestoneId int64) *MilestoneRecord {
  var milestoneRecord MilestoneRecord
  err := milestoneExecutor.C.Get(&milestoneRecord, QUERY_MILESTONE_BY_IDS_COMMAND, projectId, milestoneId)
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
      ErrorData: map[string]interface{} {
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.MilestoneRecordLocation,
    }
    log.Printf("No milestone record for projectId %s and milestoneId %d", projectId, milestoneId)
    log.Panicln(errorInfo.Marshal())
  }
  return &milestoneRecord
}

func (milestoneExecutor *MilestoneExecutor) GetMilestonesRecordsByProjectId(
    projectId string) *[]MilestoneRecord {
  var milestoneRecords []MilestoneRecord
  err := milestoneExecutor.C.Select(&milestoneRecords, QUERY_MILESTONES_BY_PROJECT_ID_COMMAND, projectId)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
    log.Printf("Failed to get milestone records for projectId %s with error: %+v\n",
      projectId, err)
    log.Panicln(errInfo.Marshal())
  }
  return &milestoneRecords
}

func (milestoneExecutor *MilestoneExecutor) VerifyMilestoneRecordExisting (
    projectId string, milestoneId int64) {
  var existing bool
  err := milestoneExecutor.C.Get(&existing, VERIFY_MILESTONE_EXISTING_COMMAND, projectId, milestoneId)
  if err != nil {
    errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
    log.Printf("Failed to verify milestone record existing for projectId %s and milestoneId %d with error: %+v\n",
      projectId, milestoneId, err)
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoMilestoneIdExisting,
      ErrorData: map[string]interface{} {
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.MilestoneRecordLocation,
    }
    log.Printf("No milestone record for projectId %s and milestoneId %d", projectId, milestoneId)
    log.Panicln(errorInfo.Marshal())
  }
}

func (milestoneExecutor *MilestoneExecutor) IncreaseNumObjectives(projectId string, milestoneId int64) {
  _, err := milestoneExecutor.C.Exec(INCREASE_NUM_OBJECTIVES_COMMAND, projectId, milestoneId)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
    log.Printf("Failed to increase numObjectives for projectId %s and milestoneId %d with error: %+v\n",
      projectId, milestoneId, err)
    errorInfo.ErrorData["milestoneId"] = milestoneId
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully increased numObjectives for projectId %s and and milestoneId %d\n", projectId, milestoneId)
}

func (milestoneExecutor *MilestoneExecutor) DecreaseNumObjectives(projectId string, milestoneId int64) {
  _, err := milestoneExecutor.C.Exec(DECREASE_NUM_OBJECTIVES_COMMAND, projectId, milestoneId)

  if err != nil {
    errorInfo := error_config.MatchError(err, "projectId", projectId, error_config.MilestoneRecordLocation)
    log.Printf("Failed to decrease numObjectives for projectId %s and milestoneId %d with error: %+v\n",
      projectId, milestoneId, err)
    errorInfo.ErrorData["milestoneId"] = milestoneId
    log.Panic(errorInfo.Marshal())
  }

  log.Printf("Successfully decreased numObjectives for projectId %s and and milestoneId %d\n", projectId, milestoneId)
}

/*
 * Tx versions
 */
func (milestoneExecutor *MilestoneExecutor) UpsertMilestoneRecordTx(milestoneRecord *MilestoneRecord) time.Time {
  res, err := milestoneExecutor.Tx.NamedQuery(UPSERT_MILESTONE_COMMAND, milestoneRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "milestoneId", milestoneRecord.MilestoneId, error_config.MilestoneRecordLocation)
    errInfo.ErrorData["projectId"] = milestoneRecord.ProjectId
    log.Printf("Failed to upsert milestone record: %+v with error:\n %+v", milestoneRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted milestone record for milestoneId %s\n", milestoneRecord.MilestoneId)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
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
      ErrorData: map[string]interface{} {
        "milestoneId": milestoneId,
        "projectId": projectId,
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

func (milestoneExecutor *MilestoneExecutor) VerifyMilestoneRecordExistingTx (
    projectId string, milestoneId int64) {
  var existing bool
  err := milestoneExecutor.Tx.Get(&existing, VERIFY_MILESTONE_EXISTING_COMMAND, projectId, milestoneId)
  if err != nil {
    errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
    log.Printf("Failed to verify milestone record existing for projectId %s and milestoneId %d with error: %+v\n",
      projectId, milestoneId, err)
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoMilestoneIdExisting,
      ErrorData: map[string]interface{} {
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.MilestoneRecordLocation,
    }
    log.Printf("No milestone record for projectId %s and milestoneId %d", projectId, milestoneId)
    log.Panicln(errorInfo.Marshal())
  }
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

  log.Printf("Successfully increased numObjectives for projectId %s and and milestoneId %d\n", projectId, milestoneId)
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
