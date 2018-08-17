package session_record_config

import (
  "log"
  "time"
  "database/sql"
  "BigBang/internal/platform/postgres_config/client_config"
)


type SessionRecordExecutor struct {
  client_config.PostgresFeedClient
}

func (sessionRecordExecutor *SessionRecordExecutor) CreateSessionRecordTable() {
  sessionRecordExecutor.CreateTimestampTrigger()
  sessionRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_SESSION_RECORDS, TABLE_NAME_FOR_SESSION_RECORDS)
  sessionRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_SESSION_RECORDS)
}

func (sessionRecordExecutor *SessionRecordExecutor) DeleteSessionRecordTable() {
   sessionRecordExecutor.DeleteTable(TABLE_NAME_FOR_SESSION_RECORDS)
}

func (sessionRecordExecutor *SessionRecordExecutor) UpsertSessionRecord(sessionRecord *SessionRecord) time.Time {
  res, err := sessionRecordExecutor.C.NamedQuery(UPSERT_SESSION_RECORD_COMMAND, sessionRecord)
  if err != nil {
    log.Panicf("Failed to upsert session record: %+v with error: %+v", sessionRecord, err)
  }

  log.Printf("Sucessfully upserted session record for postHash %s\n", sessionRecord.PostHash)

  var updatedTime time.Time
  for res.Next() {
    res.Scan(&updatedTime)
  }
  return updatedTime
}

func (sessionRecordExecutor *SessionRecordExecutor) DeleteSessionRecord(postHash string) {
  _, err := sessionRecordExecutor.C.Exec(DELETE_SESSION_RECORD_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete session record for postHash %s with error: %+v", postHash, err)
  }
  log.Printf("Sucessfully deleted session record for postHash %s\n", postHash)
}

func (sessionRecordExecutor *SessionRecordExecutor) GetSessionRecord(postHash string) *SessionRecord {
  var sessionRecord SessionRecord
  err := sessionRecordExecutor.C.Get(&sessionRecord, QUERY_SESSION_RECORD_COMMAND, postHash)
  if err != nil  {
    if err == sql.ErrNoRows {
      log.Panicf("No session exists for postHash %s", postHash)
    }
    log.Panicf("Failed to get seesion record for postHash %s with error: %+v", postHash, err.Error())
  }
  return &sessionRecord
}

/*
 * Tx versions
 */
func (sessionRecordExecutor *SessionRecordExecutor) UpsertSessionRecordTx(sessionRecord *SessionRecord) time.Time {
  res, err := sessionRecordExecutor.Tx.NamedQuery(UPSERT_SESSION_RECORD_COMMAND, sessionRecord)
  if err != nil {
    log.Panicf("Failed to upsert session record: %+v with error: %+v", sessionRecord, err)
  }

  log.Printf("Sucessfully upserted session record for postHash %s\n", sessionRecord.PostHash)

  var updatedTime time.Time
  for res.Next() {
    res.Scan(&updatedTime)
  }
  return updatedTime
}

func (sessionRecordExecutor *SessionRecordExecutor) DeleteSessionRecordTx(postHash string) {
  _, err := sessionRecordExecutor.Tx.Exec(DELETE_SESSION_RECORD_COMMAND, postHash)
  if err != nil {
    log.Panicf("Failed to delete session record for postHash %s with error: %+v", postHash, err)
  }
  log.Printf("Sucessfully deleted session record for postHash %s\n", postHash)
}

func (sessionRecordExecutor *SessionRecordExecutor) GetSessionRecordTx(postHash string) *SessionRecord {
  var sessionRecord SessionRecord
  err := sessionRecordExecutor.Tx.Get(&sessionRecord, QUERY_SESSION_RECORD_COMMAND, postHash)
  if err != nil  {
    if err == sql.ErrNoRows {
      log.Panicf("No session exists for postHash %s", postHash)
    }
    log.Panicf("Failed to get seesion record for postHash %s with error: %+v", postHash, err.Error())
  }
  return &sessionRecord
}
