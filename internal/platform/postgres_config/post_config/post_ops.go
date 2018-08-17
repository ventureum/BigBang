package post_config

import (
  "log"
  "time"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
)


type PostExecutor struct {
  client_config.PostgresFeedClient
}

func (postExecutor *PostExecutor) CreatePostTable() {
  postExecutor.CreateTimestampTrigger()
  postExecutor.CreateTable(TABLE_SCHEMA_FOR_POST, TABLE_NAME_FOR_POST)
  postExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_POST)
}

func (postExecutor *PostExecutor) DeletePostTable() {
   postExecutor.DeleteTable(TABLE_NAME_FOR_POST)
}

func (postExecutor *PostExecutor) UpsertPostRecord(postRecord *PostRecord) time.Time {
  res, err := postExecutor.C.NamedQuery(UPSERT_POST_COMMAND, postRecord)
  if err != nil {
    log.Panicf("Failed to upsert post record: %+v with error:\n %+v", postRecord, err.Error())
  }

  log.Printf("Sucessfully upserted post record for postHash %s\n", postRecord.PostHash)

  var updatedTime time.Time
  for res.Next() {
    res.Scan(&updatedTime)
  }
  return updatedTime
}

func (postExecutor *PostExecutor) DeletePostRecord(postHash string) {
  _, err := postExecutor.C.Exec(DELETE_POST_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete post record for postHash %s with error:\n %+v", postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post record for postHash %s\n", postHash)
}

func (postExecutor *PostExecutor) GetPostRecord(postHash string) *PostRecord {
  var postRecord PostRecord
  err := postExecutor.C.Get(&postRecord, QUERY_POST_COMMAND, postHash)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get post record for postHash %s with error:\n %+v", postHash, err.Error())
  }
  return &postRecord
}

func (postExecutor *PostExecutor) GetPostUpdateCount(postHash string) int64 {
  var updateCount int64
  err := postExecutor.C.Get(&updateCount, QUERY_POST_UPDATE_COUNT_COMMAND, postHash)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get post update count for postHash %s with error:\n %+v", postHash, err.Error())
  }
  return updateCount
}

/*
 * Tx versions
 */
func (postExecutor *PostExecutor) UpsertPostRecordTx(postRecord *PostRecord) time.Time {
  res, err := postExecutor.Tx.NamedQuery(UPSERT_POST_COMMAND, postRecord)
  if err != nil {
    log.Panicf("Failed to upsert post record: %+v with error:\n %+v", postRecord, err.Error())
  }

  log.Printf("Sucessfully upserted post record for postHash %s\n", postRecord.PostHash)

  var updatedTime time.Time
  for res.Next() {
    res.Scan(&updatedTime)
  }
  return updatedTime
}

func (postExecutor *PostExecutor) DeletePostRecordTx(postHash string) {
  _, err := postExecutor.Tx.Exec(DELETE_POST_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete post record for postHash %s with error:\n %+v", postHash, err.Error())
  }
  log.Printf("Sucessfully deleted post record for postHash %s\n", postHash)
}

func (postExecutor *PostExecutor) GetPostRecordTx(postHash string) *PostRecord {
  var postRecord PostRecord
  err := postExecutor.Tx.Get(&postRecord, QUERY_POST_COMMAND, postHash)

  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get post record for postHash %s with error:\n %+v", postHash, err.Error())
  }
  return &postRecord
}

func (postExecutor *PostExecutor) GetPostUpdateCountTx(postHash string) int64 {
  var updateCount sql.NullInt64
  err := postExecutor.Tx.Get(&updateCount, QUERY_POST_UPDATE_COUNT_COMMAND, postHash)
  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get post update count for postHash %s with error:\n %+v", postHash, err.Error())
  }
  return updateCount.Int64
}
