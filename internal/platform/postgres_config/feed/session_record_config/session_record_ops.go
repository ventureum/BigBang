package session_record_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
	"time"
)

type SessionRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (sessionRecordExecutor *SessionRecordExecutor) CreateSessionRecordTable() {
	sessionRecordExecutor.CreateTimestampTrigger()
	sessionRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_SESSION_RECORDS, TABLE_NAME_FOR_SESSION_RECORDS)
	sessionRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_SESSION_RECORDS)
}

func (sessionRecordExecutor *SessionRecordExecutor) DeleteSessionRecordTable() {
	sessionRecordExecutor.DeleteTable(TABLE_NAME_FOR_SESSION_RECORDS)
}

func (sessionRecordExecutor *SessionRecordExecutor) UpsertSessionRecordTx(sessionRecord *SessionRecord) time.Time {
	res, err := sessionRecordExecutor.Tx.NamedQuery(UPSERT_SESSION_RECORD_COMMAND, sessionRecord)
	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", sessionRecord.PostHash, error_config.SessionRecordLocation)
		errInfo.ErrorData["actor"] = sessionRecord.Actor
		log.Printf("Failed to upsert session record: %+v with error: %+v", sessionRecord, err)
		log.Panicln(errInfo.Marshal())
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
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.SessionRecordLocation)
		log.Printf("Failed to delete session record for postHash %s with error: %+v", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted session record for postHash %s\n", postHash)
}

func (sessionRecordExecutor *SessionRecordExecutor) GetSessionRecordTx(postHash string) *SessionRecord {
	var sessionRecord SessionRecord
	err := sessionRecordExecutor.Tx.Get(&sessionRecord, QUERY_SESSION_RECORD_COMMAND, postHash)
	if err != nil {
		errInfo := error_config.MatchError(err, "postHash", postHash, error_config.SessionRecordLocation)
		log.Printf("Failed to get seesion record for postHash %s with error: %+v", postHash, err)
		log.Panicln(errInfo.Marshal())
	}
	return &sessionRecord
}
