package objective_config

import (
  "log"
  "time"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "database/sql"
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

func (objectiveExecutor *ObjectiveExecutor) UpsertObjectiveRecord(objectiveRecord *ObjectiveRecord) time.Time {
  res, err := objectiveExecutor.C.NamedQuery(UPSERT_OBJECTIVE_COMMAND, objectiveRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveRecord.ObjectiveId, error_config.ObjectiveRecordLocation)
    errInfo.ErrorData["milestoneId"] = objectiveRecord.MilestoneId
    errInfo.ErrorData["projectId"] = objectiveRecord.ProjectId
    log.Printf("Failed to upsert objective record: %+v with error:\n %+v", objectiveRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted objective record for objectiveId %s\n", objectiveRecord.ObjectiveId)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
}

func (objectiveExecutor *ObjectiveExecutor) DeleteObjectiveRecordByIDs(
    projectId string, milestoneId int64, objectiveId int64) {
  _, err := objectiveExecutor.C.Exec(DELETE_OBJECTIVE_BY_IDS_COMMAND, projectId, milestoneId, objectiveId)
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

func (objectiveExecutor *ObjectiveExecutor) DeleteObjectiveRecordsByProjectIdAndMilestoneId(
    projectId string, milestoneId int64) {
  _, err := objectiveExecutor.C.Exec(DELETE_OBJECTIVES_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND, projectId, milestoneId)
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

func (objectiveExecutor *ObjectiveExecutor) GetObjectiveRecordByIDs(
    projectId string, milestoneId int64, objectiveId int64) *ObjectiveRecord {
  var objectiveRecord ObjectiveRecord
  err := objectiveExecutor.C.Get(&objectiveRecord, QUERY_OBJECTIVE_BY_IDS_COMMAND, projectId, milestoneId, objectiveId)
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
      ErrorData: map[string]interface{} {
        "objectiveId": objectiveId,
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.ObjectiveRecordLocation,
    }
    log.Printf("No objective record for projectId %s, milestoneId %d and objectiveId %d", projectId, milestoneId, objectiveId)
    log.Panicln(errorInfo.Marshal())
  }
  return &objectiveRecord
}

func (objectiveExecutor *ObjectiveExecutor) GetObjectiveRecordsByProjectIdAndMilestoneId(
    projectId string, milestoneId int64) *[]ObjectiveRecord {
  var objectiveRecords []ObjectiveRecord
  err := objectiveExecutor.C.Select(&objectiveRecords, QUERY_OBJECTIVES_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND,
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

func (objectiveExecutor *ObjectiveExecutor) VerifyObjectiveRecordExisting (
    projectId string, milestoneId int64, objectiveId int64) {
  var existing bool
  err := objectiveExecutor.C.Get(&existing, VERIFY_OBJECTIVE_EXISTING_COMMAND, projectId, milestoneId, objectiveId)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.ObjectiveRecordLocation)
    log.Printf("Failed to verify objective record existing for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoObjectiveIdExisting,
      ErrorData: map[string]interface{} {
        "objectiveId": objectiveId,
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.ObjectiveRecordLocation,
    }
    log.Printf("No objective record for projectId %s, milestoneId %d and objectiveId %d", projectId, milestoneId, objectiveId)
    log.Panicln(errorInfo.Marshal())
  }
}

func (objectiveExecutor *ObjectiveExecutor) AddRatingAndWeight(
  projectId string, milestoneId int, objectiveId int, deltaRating int64, deltaWeight int64) {
  _, err := objectiveExecutor.C.Exec(
    ADD_RATING_AND_WEIGHT_COMMAND, projectId, milestoneId, objectiveId, deltaRating, deltaWeight)

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

/*
 * Tx versions
 */

func (objectiveExecutor *ObjectiveExecutor) UpsertObjectiveRecordTx(objectiveRecord *ObjectiveRecord) time.Time {
  res, err := objectiveExecutor.Tx.NamedQuery(UPSERT_OBJECTIVE_COMMAND, objectiveRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveRecord.ObjectiveId, error_config.ObjectiveRecordLocation)
    errInfo.ErrorData["milestoneId"] = objectiveRecord.MilestoneId
    errInfo.ErrorData["projectId"] = objectiveRecord.ProjectId
    log.Printf("Failed to upsert objective record: %+v with error:\n %+v", objectiveRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted objective record for objectiveId %s\n", objectiveRecord.ObjectiveId)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
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
      ErrorData: map[string]interface{} {
        "objectiveId": objectiveId,
        "milestoneId": milestoneId,
        "projectId": projectId,
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

func (objectiveExecutor *ObjectiveExecutor) VerifyObjectiveRecordExistingTx (
    projectId string, milestoneId int64, objectiveId int64) {
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
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoObjectiveIdExisting,
      ErrorData: map[string]interface{} {
        "objectiveId": objectiveId,
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.ObjectiveRecordLocation,
    }
    log.Printf("No objective record for projectId %s, milestoneId %d and objectiveId %d", projectId, milestoneId, objectiveId)
    log.Panicln(errorInfo.Marshal())
  }
}

func (objectiveExecutor *ObjectiveExecutor) AddRatingAndWeightTx(
    projectId string, milestoneId int, objectiveId int, deltaRating int64, deltaWeight int64) {
  _, err := objectiveExecutor.Tx.Exec(
    ADD_RATING_AND_WEIGHT_COMMAND, projectId, milestoneId, objectiveId, deltaRating, deltaWeight)

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
