package principal_proxy_votes_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "database/sql"
)

type PrincipalProxyVotesExecutor struct {
  client_config.PostgresBigBangClient
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) CreatePrincipalProxyVotesTable() {
  principalProxyVotesExecutor.CreateTimestampTrigger()
  principalProxyVotesExecutor.CreateTable(TABLE_SCHEMA_FOR_PRINCIPAL_PROXY_VOTES, TABLE_NAME_FOR_PRINCIPAL_PROXY_VOTES)
  principalProxyVotesExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_PRINCIPAL_PROXY_VOTES)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) DeletePrincipalProxyVotesTable() {
  principalProxyVotesExecutor.DeleteTable(TABLE_NAME_FOR_PRINCIPAL_PROXY_VOTES)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) ClearPrincipalProxyVotesTable() {
  principalProxyVotesExecutor.ClearTable(TABLE_NAME_FOR_PRINCIPAL_PROXY_VOTES)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) UpsertPrincipalProxyVotesRecord(principalProxyVotesRecord *PrincipalProxyVotesRecord) {

  res, err := principalProxyVotesExecutor.C.NamedExec(UPDATE_PRINCIPAL_PROXY_VOTES_COMMAND, principalProxyVotesRecord)

  if err != nil {
    errInfo := error_config.MatchError(err, "actor", principalProxyVotesRecord.Actor, error_config.PrincipalProxyVotesRecordLocation)
    errInfo.ErrorData["projectId"] = principalProxyVotesRecord.ProjectId
    errInfo.ErrorData["proxy"] = principalProxyVotesRecord.Proxy
    log.Printf("Failed to upsert Principal Proxy Votes Record: %+v with error:\n %+v", principalProxyVotesRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  count, err := res.RowsAffected()

  if count == 0 {
    _, err = principalProxyVotesExecutor.C.NamedExec(INSERT_PRINCIPAL_PROXY_VOTES_COMMAND, principalProxyVotesRecord)

    if err != nil {
      errInfo := error_config.MatchError(err, "actor", principalProxyVotesRecord.Actor, error_config.PrincipalProxyVotesRecordLocation)
      errInfo.ErrorData["projectId"] = principalProxyVotesRecord.ProjectId
      errInfo.ErrorData["proxy"] = principalProxyVotesRecord.Proxy
      log.Printf("Failed to upsert Principal Proxy Votes Record: %+v with error:\n %+v", principalProxyVotesRecord, err)
      log.Panicln(errInfo.Marshal())
    }
  }

  log.Printf("Sucessfully upserted Principal Proxy Votes Record for actor %s, projectId %s and proxy %s\n",
    principalProxyVotesRecord.Actor, principalProxyVotesRecord.ProjectId, principalProxyVotesRecord.Proxy)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) DeletePrincipalProxyVotesRecordByIDs(
    actor string, projectId string, proxy string) {
  _, err := principalProxyVotesExecutor.C.Exec(DELETE_PRINCIPAL_PROXY_VOTES_BY_IDS_COMMAND, actor, projectId, proxy)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to delete Principal Proxy Votes Record for actor %s, projectId %s and proxy %s with error: %+v\n",
      actor, projectId, proxy, err)
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["proxy"] = proxy
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted Principal Proxy Votes Record for actor %s, projectId %s and proxy %s\n",
    actor, projectId, proxy)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) GetPrincipalProxyVotesRecordByIDs(
    actor string, projectId string, proxy string) *[]PrincipalProxyVotesRecord {
  var principalProxyVotesRecords []PrincipalProxyVotesRecord
  err := principalProxyVotesExecutor.C.Select(&principalProxyVotesRecords, QUERY_PRINCIPAL_PROXY_VOTES_BY_IDS_COMMAND,
    actor, projectId, proxy)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to get Principal Proxy Votes Record for actor %s, projectId %s and proxy %s with error: %+v\n",
      actor, projectId, proxy, err)
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["proxy"] = proxy
    log.Panicln(errInfo.Marshal())
  }
  return &principalProxyVotesRecords
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) GetPrincipalProxyVotesRecordsByActorAndProjectId(
    actor string, projectId string) *[]PrincipalProxyVotesRecord {
  var principalProxyVotesRecords []PrincipalProxyVotesRecord
  err := principalProxyVotesExecutor.C.Select(&principalProxyVotesRecords, QUERY_PRINCIPAL_PROXY_VOTES_BY_ACTOR_AND_PROJECT_ID_COMMAND,
    actor, projectId)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to get Principal Proxy Votes Records for actor %s and projectId %s with error: %+v\n",
      actor, projectId, err)
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  return &principalProxyVotesRecords
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) GetPrincipalProxyVotesRecordListByCursor(
    actor string, projectId string, cursor string, limit int64) *[]PrincipalProxyVotesRecord {
  var principalProxyVotesRecords []PrincipalProxyVotesRecord
  var err error
  if cursor != "" {
    err = principalProxyVotesExecutor.C.Select(
      &principalProxyVotesRecords,
      PAGINATION_QUERY_PRINCIPAL_PROXY_VOTES_LIST_COMMAND,
      actor, projectId, cursor, limit)
  } else {
    err = principalProxyVotesExecutor.C.Select(
      &principalProxyVotesRecords,
      QUERY_PRINCIPAL_PROXY_VOTES_LIST_COMMAND,
      actor, projectId, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to get Principal Proxy Votes Records by cursor %s and limit %d  for actor %s and projectId %s with error: %+v\n",
      cursor, limit, actor, projectId, err)
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["cursor"] = cursor
    errInfo.ErrorData["limit"] = limit
    log.Panicln(errInfo.Marshal())
  }
  return &principalProxyVotesRecords
}


/*
 * Tx versions
 */

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) UpsertPrincipalProxyVotesRecordTx(principalProxyVotesRecord *PrincipalProxyVotesRecord) {

  res, err := principalProxyVotesExecutor.Tx.NamedExec(UPDATE_PRINCIPAL_PROXY_VOTES_COMMAND, principalProxyVotesRecord)

  if err != nil {
    errInfo := error_config.MatchError(err, "actor", principalProxyVotesRecord.Actor, error_config.PrincipalProxyVotesRecordLocation)
    errInfo.ErrorData["projectId"] = principalProxyVotesRecord.ProjectId
    errInfo.ErrorData["proxy"] = principalProxyVotesRecord.Proxy
    log.Printf("Failed to upsert Principal Proxy Votes Record: %+v with error:\n %+v", principalProxyVotesRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  count, err := res.RowsAffected()

  if count == 0 {
    _, err = principalProxyVotesExecutor.Tx.NamedExec(INSERT_PRINCIPAL_PROXY_VOTES_COMMAND, principalProxyVotesRecord)

    if err != nil {
      errInfo := error_config.MatchError(err, "actor", principalProxyVotesRecord.Actor, error_config.PrincipalProxyVotesRecordLocation)
      errInfo.ErrorData["projectId"] = principalProxyVotesRecord.ProjectId
      errInfo.ErrorData["proxy"] = principalProxyVotesRecord.Proxy
      log.Printf("Failed to upsert Principal Proxy Votes Record: %+v with error:\n %+v", principalProxyVotesRecord, err)
      log.Panicln(errInfo.Marshal())
    }
  }

  log.Printf("Sucessfully upserted Principal Proxy Votes Record for actor %s, projectId %s and proxy %s\n",
    principalProxyVotesRecord.Actor, principalProxyVotesRecord.ProjectId, principalProxyVotesRecord.Proxy)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) DeletePrincipalProxyVotesRecordByIDsTx(
    actor string, projectId string, proxy string) {
  _, err := principalProxyVotesExecutor.Tx.Exec(DELETE_PRINCIPAL_PROXY_VOTES_BY_IDS_COMMAND, actor, projectId, proxy)
  if err != nil {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to delete Principal Proxy Votes Record for actor %s, projectId %s and proxy %s with error: %+v\n",
      actor, projectId, proxy, err)
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["proxy"] = proxy
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted Principal Proxy Votes Record for actor %s, projectId %s and proxy %s\n",
    actor, projectId, proxy)
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) GetPrincipalProxyVotesRecordByIDsTx(
    actor string, projectId string, proxy string) *[]PrincipalProxyVotesRecord {
  var principalProxyVotesRecords []PrincipalProxyVotesRecord
  err := principalProxyVotesExecutor.Tx.Select(&principalProxyVotesRecords, QUERY_PRINCIPAL_PROXY_VOTES_BY_IDS_COMMAND,
    actor, projectId, proxy)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to get Principal Proxy Votes Record for actor %s, projectId %s and proxy %s with error: %+v\n",
      actor, projectId, proxy, err)
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["proxy"] = proxy
    log.Panicln(errInfo.Marshal())
  }
  return &principalProxyVotesRecords
}


func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) GetPrincipalProxyVotesRecordsByActorAndProjectIdTx(
    actor string, projectId string) *[]PrincipalProxyVotesRecord {
  var principalProxyVotesRecords []PrincipalProxyVotesRecord
  err := principalProxyVotesExecutor.Tx.Select(&principalProxyVotesRecords, QUERY_PRINCIPAL_PROXY_VOTES_BY_ACTOR_AND_PROJECT_ID_COMMAND,
    actor, projectId)
  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to get Principal Proxy Votes Records for actor %s and projectId %s with error: %+v\n",
      actor, projectId, err)
    errInfo.ErrorData["projectId"] = projectId
    log.Panicln(errInfo.Marshal())
  }
  return &principalProxyVotesRecords
}

func (principalProxyVotesExecutor *PrincipalProxyVotesExecutor) GetPrincipalProxyVotesRecordListByCursorTx(
    actor string, projectId string, cursor string, limit int64) *[]PrincipalProxyVotesRecord {
  var principalProxyVotesRecords []PrincipalProxyVotesRecord
  var err error
  if cursor != "" {
    err = principalProxyVotesExecutor.Tx.Select(
      &principalProxyVotesRecords,
      PAGINATION_QUERY_PRINCIPAL_PROXY_VOTES_LIST_COMMAND,
      actor, projectId, cursor, limit)
  } else {
    err = principalProxyVotesExecutor.Tx.Select(
      &principalProxyVotesRecords,
      QUERY_PRINCIPAL_PROXY_VOTES_LIST_COMMAND,
      actor, projectId, limit)
  }

  if err != nil && err != sql.ErrNoRows {
    errInfo := error_config.MatchError(err, "actor", actor, error_config.PrincipalProxyVotesRecordLocation)
    log.Printf("Failed to get Principal Proxy Votes Records by cursor %s and limit %d  for actor %s and projectId %s with error: %+v\n",
      cursor, limit, actor, projectId, err)
    errInfo.ErrorData["projectId"] = projectId
    errInfo.ErrorData["cursor"] = cursor
    errInfo.ErrorData["limit"] = limit
    log.Panicln(errInfo.Marshal())
  }
  return &principalProxyVotesRecords
}
