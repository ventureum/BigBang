package rating_vote_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "database/sql"
)

type RatingVoteExecutor struct {
  client_config.PostgresBigBangClient
}

func (ratingVoteExecutor *RatingVoteExecutor) CreateRatingVoteTable() {
  ratingVoteExecutor.CreateTimestampTrigger()
  ratingVoteExecutor.CreateTable(TABLE_SCHEMA_FOR_RATING_VOTE, TABLE_NAME_FOR_RATING_VOTE)
  ratingVoteExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_RATING_VOTE)
}

func (ratingVoteExecutor *RatingVoteExecutor) DeleteRatingVoteTable() {
  ratingVoteExecutor.DeleteTable(TABLE_NAME_FOR_RATING_VOTE)
}

func (ratingVoteExecutor *RatingVoteExecutor) ClearRatingVoteTable() {
  ratingVoteExecutor.ClearTable(TABLE_NAME_FOR_RATING_VOTE)
}

func (ratingVoteExecutor *RatingVoteExecutor) UpsertRatingVoteRecord(ratingVoteRecord *RatingVoteRecord) {

  res, err := ratingVoteExecutor.C.NamedExec(UPDATE_RATING_VOTE_COMMAND, ratingVoteRecord)

  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", ratingVoteRecord.ObjectiveId, error_config.RatingVoteRecordLocation)
    errInfo.ErrorData["projectId"] = ratingVoteRecord.ProjectId
    errInfo.ErrorData["milestoneId"] = ratingVoteRecord.MilestoneId
    errInfo.ErrorData["voter"] = ratingVoteRecord.Voter
    log.Printf("Failed to upsert rating vote Record: %+v with error:\n %+v", ratingVoteRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  count, err := res.RowsAffected()

  if count == 0 {
    _, err = ratingVoteExecutor.C.NamedExec(INSERT_RATING_VOTE_COMMAND, ratingVoteRecord)

    if err != nil {
      errInfo := error_config.MatchError(err, "objectiveId", ratingVoteRecord.ObjectiveId, error_config.RatingVoteRecordLocation)
      errInfo.ErrorData["projectId"] = ratingVoteRecord.ProjectId
      errInfo.ErrorData["milestoneId"] = ratingVoteRecord.MilestoneId
      errInfo.ErrorData["voter"] = ratingVoteRecord.Voter
      log.Printf("Failed to upsert rating vote Record: %+v with error:\n %+v", ratingVoteRecord, err)
      log.Panicln(errInfo.Marshal())
    }
  }

  log.Printf("Sucessfully upserted rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s\n",
    ratingVoteRecord.ProjectId, ratingVoteRecord.MilestoneId, ratingVoteRecord.ObjectiveId, ratingVoteRecord.Voter)
}

func (ratingVoteExecutor *RatingVoteExecutor) DeleteRatingVoteRecordsByIDs(
    projectId string, milestoneId int64, objectiveId int64) {
  _, err := ratingVoteExecutor.C.Exec(DELETE_RATING_VOTE_BY_IDS_COMMAND, projectId, milestoneId, objectiveId)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to delete rating vote Record for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted rating vote Record for projectId %s, milestoneId %d and objectiveId %d\n",
    projectId, milestoneId, objectiveId)
}

func (ratingVoteExecutor *RatingVoteExecutor) DeleteRatingVoteRecordByIDsAndVoter(
    projectId string, milestoneId int64, objectiveId int64, voter string) {
  _, err := ratingVoteExecutor.C.Exec(DELETE_RATING_VOTE_BY_IDS_AND_VOTER_COMMAND, projectId, milestoneId, objectiveId, voter)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to delete rating vote Record for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["voter"] = voter
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s\n",
    projectId, milestoneId, objectiveId, voter)
}

func (ratingVoteExecutor *RatingVoteExecutor) GetRatingVoteRecordByIDsAndVoter(
    projectId string, milestoneId int64, objectiveId int64, voter string) *RatingVoteRecord {
  var ratingVoteRecord RatingVoteRecord
  err := ratingVoteExecutor.C.Get(&ratingVoteRecord, QUERY_RATING_VOTE_BY_IDS_AND_VOTER_COMMAND, projectId, milestoneId, objectiveId, voter)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to get rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s with error: %+v\n",
      projectId, milestoneId, objectiveId, voter, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["voter"] = voter
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoRatingVoteVoterExisting,
      ErrorData: map[string]interface{} {
        "objectiveId": objectiveId,
        "milestoneId": milestoneId,
        "projectId": projectId,
        "voter": voter,
      },
      ErrorLocation: error_config.RatingVoteRecordLocation,
    }
    log.Printf("No rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s",
      projectId, milestoneId, objectiveId, voter)
    log.Panicln(errorInfo.Marshal())
  }
  return &ratingVoteRecord
}

func (ratingVoteExecutor *RatingVoteExecutor) GetRatingVoteRecordsByIDs(
    projectId string, milestoneId int64, objectiveId int64) *[]RatingVoteRecord {
  var ratingVoteRecords []RatingVoteRecord
  err := ratingVoteExecutor.C.Select(&ratingVoteRecords, QUERY_RATING_VOTES_BY_IDS_COMMAND,
    projectId, milestoneId, objectiveId)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to get rating vote Records for projectId %s, milestoneId and %d objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["objectiveId"] = objectiveId

    log.Panicln(errInfo.Marshal())
  }
  return &ratingVoteRecords
}

func (ratingVoteExecutor *RatingVoteExecutor) VerifyRatingVoteRecordExisting (
    projectId string, milestoneId int64, objectiveId int64, voter string) bool {
  var existing bool
  err := ratingVoteExecutor.C.Get(&existing, VERIFY_RATING_VOTE_EXISTING_COMMAND, projectId, milestoneId, objectiveId, voter)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to verify rating vote Record existing for projectId %s, milestoneId %d, objectiveId %d and voter %s with error: %+v\n",
      projectId, milestoneId, objectiveId, voter, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["voter"] = voter
    log.Panicln(errInfo.Marshal())
  }
  return existing
}

func (ratingVoteExecutor *RatingVoteExecutor) GetRatingVoteRecordsByCursor(
    projectId string, milestoneId int64, objectiveId int64, cursor string, limit int64) *[]RatingVoteRecord {
  var ratingVoteRecords []RatingVoteRecord
  var err error
  if cursor != "" {
    err = ratingVoteExecutor.C.Select(
      &ratingVoteRecords,
      PAGINATION_QUERY_RATING_VOTE_LIST_COMMAND,
      projectId, milestoneId, objectiveId, cursor, limit)
  } else {
    err = ratingVoteExecutor.C.Select(
      &ratingVoteRecords,
      QUERY_RATING_VOTE_LIST_COMMAND,
      projectId, milestoneId, objectiveId, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to get rating vote Records by cursor %s and limit %d for projectId %s, milestoneId and %d objectiveId %d with error: %+v\n",
      cursor, limit, projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["objectiveId"] = objectiveId
    errInfo.ErrorData["cursor"] = cursor
    errInfo.ErrorData["limit"] = limit
    log.Panicln(errInfo.Marshal())
  }
  return &ratingVoteRecords
}

/*
 * Tx versions
 */
func (ratingVoteExecutor *RatingVoteExecutor) UpsertRatingVoteRecordTx(ratingVoteRecord *RatingVoteRecord) {

  res, err := ratingVoteExecutor.Tx.NamedExec(UPDATE_RATING_VOTE_COMMAND, ratingVoteRecord)

  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", ratingVoteRecord.ObjectiveId, error_config.RatingVoteRecordLocation)
    errInfo.ErrorData["projectId"] = ratingVoteRecord.ProjectId
    errInfo.ErrorData["milestoneId"] = ratingVoteRecord.MilestoneId
    errInfo.ErrorData["voter"] = ratingVoteRecord.Voter
    log.Printf("Failed to upsert rating vote Record: %+v with error:\n %+v", ratingVoteRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  count, err := res.RowsAffected()

  if count == 0 {
    _, err = ratingVoteExecutor.Tx.NamedExec(INSERT_RATING_VOTE_COMMAND, ratingVoteRecord)

    if err != nil {
      errInfo := error_config.MatchError(err, "objectiveId", ratingVoteRecord.ObjectiveId, error_config.RatingVoteRecordLocation)
      errInfo.ErrorData["projectId"] = ratingVoteRecord.ProjectId
      errInfo.ErrorData["milestoneId"] = ratingVoteRecord.MilestoneId
      errInfo.ErrorData["voter"] = ratingVoteRecord.Voter
      log.Printf("Failed to upsert rating vote Record: %+v with error:\n %+v", ratingVoteRecord, err)
      log.Panicln(errInfo.Marshal())
    }
  }

  log.Printf("Sucessfully upserted rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s\n",
    ratingVoteRecord.ProjectId, ratingVoteRecord.MilestoneId, ratingVoteRecord.ObjectiveId, ratingVoteRecord.Voter)
}

func (ratingVoteExecutor *RatingVoteExecutor) DeleteRatingVoteRecordsByIDsTx(
    projectId string, milestoneId int64, objectiveId int64) {
  _, err := ratingVoteExecutor.Tx.Exec(DELETE_RATING_VOTE_BY_IDS_COMMAND, projectId, milestoneId, objectiveId)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to delete rating vote Record for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted rating vote Record for projectId %s, milestoneId %d and objectiveId %d\n",
    projectId, milestoneId, objectiveId)
}

func (ratingVoteExecutor *RatingVoteExecutor) DeleteRatingVoteRecordByIDsAndVoterTx(
    projectId string, milestoneId int64, objectiveId int64, voter string) {
  _, err := ratingVoteExecutor.Tx.Exec(DELETE_RATING_VOTE_BY_IDS_AND_VOTER_COMMAND, projectId, milestoneId, objectiveId, voter)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to delete rating vote Record for projectId %s, milestoneId %d and objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["voter"] = voter
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s\n",
    projectId, milestoneId, objectiveId, voter)
}

func (ratingVoteExecutor *RatingVoteExecutor) GetRatingVoteRecordByIDsAndVoterTx(
    projectId string, milestoneId int64, objectiveId int64, voter string) *RatingVoteRecord {
  var ratingVoteRecord RatingVoteRecord
  err := ratingVoteExecutor.Tx.Get(&ratingVoteRecord, QUERY_RATING_VOTE_BY_IDS_AND_VOTER_COMMAND, projectId, milestoneId, objectiveId, voter)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to get rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s with error: %+v\n",
      projectId, milestoneId, objectiveId, voter, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["voter"] = voter
    log.Panicln(errInfo.Marshal())
  }

  if err == sql.ErrNoRows {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoRatingVoteVoterExisting,
      ErrorData: map[string]interface{} {
        "objectiveId": objectiveId,
        "milestoneId": milestoneId,
        "projectId": projectId,
        "voter": voter,
      },
      ErrorLocation: error_config.RatingVoteRecordLocation,
    }
    log.Printf("No rating vote Record for projectId %s, milestoneId %d, objectiveId %d and voter %s",
      projectId, milestoneId, objectiveId, voter)
    log.Panicln(errorInfo.Marshal())
  }
  return &ratingVoteRecord
}

func (ratingVoteExecutor *RatingVoteExecutor) GetRatingVoteRecordsByIDsTx(
    projectId string, milestoneId int64, objectiveId int64) *[]RatingVoteRecord {
  var ratingVoteRecords []RatingVoteRecord
  err := ratingVoteExecutor.Tx.Select(&ratingVoteRecords, QUERY_RATING_VOTES_BY_IDS_COMMAND,
    projectId, milestoneId, objectiveId)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to get rating vote Records for projectId %s, milestoneId and %d objectiveId %d with error: %+v\n",
      projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["objectiveId"] = objectiveId

    log.Panicln(errInfo.Marshal())
  }
  return &ratingVoteRecords
}

func (ratingVoteExecutor *RatingVoteExecutor) VerifyRatingVoteRecordExistingTx (
    projectId string, milestoneId int64, objectiveId int64, voter string) bool {
  var existing bool
  err := ratingVoteExecutor.Tx.Get(&existing, VERIFY_RATING_VOTE_EXISTING_COMMAND, projectId, milestoneId, objectiveId, voter)
  if err != nil {
    errInfo := error_config.MatchError(err, "objectiveId", objectiveId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to verify rating vote Record existing for projectId %s, milestoneId %d, objectiveId %d and voter %s with error: %+v\n",
      projectId, milestoneId, objectiveId, voter, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["voter"] = voter
    log.Panicln(errInfo.Marshal())
  }
  return existing
}

func (ratingVoteExecutor *RatingVoteExecutor) GetRatingVoteRecordsByCursorTx(
    projectId string, milestoneId int64, objectiveId int64, cursor string, limit int64) *[]RatingVoteRecord {
  var ratingVoteRecords []RatingVoteRecord
  var err error
  if cursor != "" {
    err = ratingVoteExecutor.Tx.Select(
      &ratingVoteRecords,
      PAGINATION_QUERY_RATING_VOTE_LIST_COMMAND,
      projectId, milestoneId, objectiveId, cursor, limit)
  } else {
    err = ratingVoteExecutor.Tx.Select(
      &ratingVoteRecords,
      QUERY_RATING_VOTE_LIST_COMMAND,
      projectId, milestoneId, objectiveId, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "projectId", projectId, error_config.RatingVoteRecordLocation)
    log.Printf("Failed to get rating vote Records by cursor %s and limit %d for projectId %s, milestoneId and %d objectiveId %d with error: %+v\n",
      cursor, limit, projectId, milestoneId, objectiveId, err)
    errInfo.ErrorData["milestoneId"] = milestoneId
    errInfo.ErrorData["objectiveId"] = objectiveId
    errInfo.ErrorData["cursor"] = cursor
    errInfo.ErrorData["limit"] = limit
    log.Panicln(errInfo.Marshal())
  }
  return &ratingVoteRecords
}
