package proxy_config

import (
  "log"
  "time"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "database/sql"
)

type ProxyExecutor struct {
  client_config.PostgresBigBangClient
}

func (proxyExecutor *ProxyExecutor) CreateProxyTable() {
  proxyExecutor.CreateTimestampTrigger()
  proxyExecutor.CreateTable(TABLE_SCHEMA_FOR_PROXY, TABLE_NAME_FOR_PROXY)
  proxyExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_PROXY)
}

func (proxyExecutor *ProxyExecutor) DeleteProxyTable() {
   proxyExecutor.DeleteTable(TABLE_NAME_FOR_PROXY)
}

func (proxyExecutor *ProxyExecutor) ClearProxyTable() {
  proxyExecutor.ClearTable(TABLE_NAME_FOR_PROXY)
}

func (proxyExecutor *ProxyExecutor) UpsertProxyRecord(proxyRecord *ProxyRecord) time.Time {
  res, err := proxyExecutor.C.NamedQuery(UPSERT_PROXY_COMMAND, proxyRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", proxyRecord.UUID, error_config.ProxyRecordLocation)
    log.Printf("Failed to upsert proxy record: %+v with error:\n %+v", proxyRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted proxy record for uuid %s\n", proxyRecord.UUID)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
}

func (proxyExecutor *ProxyExecutor) DeleteProxyRecord(uuid string) {
  _, err := proxyExecutor.C.Exec(DELETE_PROXY_COMMAND, uuid)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", uuid, error_config.ProxyRecordLocation)
    log.Printf("Failed to delete proxy record for uuid %s with error: %+v\n", uuid, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted proxy record for uuid %s\n", uuid)
}

func (proxyExecutor *ProxyExecutor) GetProxyRecord(uuid string) *ProxyRecord {
  var proxyRecord ProxyRecord
  err := proxyExecutor.C.Get(&proxyRecord, QUERY_PROXY_COMMAND, uuid)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", uuid, error_config.ProxyRecordLocation)
    log.Printf("Failed to get proxy record for uuid %s with error: %+v\n", uuid, err)
    log.Panicln(errInfo.Marshal())
  }
  return &proxyRecord
}

func (proxyExecutor *ProxyExecutor) VerifyProxyRecordExisting (uuid string) {
  var existing bool
  err := proxyExecutor.C.Get(&existing, VERIFY_PROXY_EXISTING_COMMAND, uuid)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", uuid, error_config.ProxyRecordLocation)
    log.Printf("Failed to verify proxy record existing for uuid %s with error: %+v\n", uuid, err)
    log.Panicln(errInfo.Marshal())
  }
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoProxyUUIDExisting,
      ErrorData: map[string]interface{} {
        "uuid": uuid,
      },
      ErrorLocation: error_config.ProxyRecordLocation,
    }
    log.Printf("No proxy record for uuid %s", uuid)
    log.Panicln(errorInfo.Marshal())
  }
}

func (proxyExecutor *ProxyExecutor) GetListOfProxyUUIDs() *[]string {
  var proxyUUIDs []string
  err := proxyExecutor.C.Select(&proxyUUIDs, QUERY_LIST_OF_PROXY_UUIDS_COMMAND)

  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get list of proxy UUIDs with error: %+v\n", err)
  }
  return &proxyUUIDs
}

/*
 * Tx versions
 */
func (proxyExecutor *ProxyExecutor) UpsertProxyRecordTx(proxyRecord *ProxyRecord) time.Time {
  res, err := proxyExecutor.Tx.NamedQuery(UPSERT_PROXY_COMMAND, proxyRecord)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", proxyRecord.UUID, error_config.ProxyRecordLocation)
    log.Printf("Failed to upsert proxy record: %+v with error:\n %+v", proxyRecord, err)
    log.Panicln(errInfo.Marshal())
  }

  log.Printf("Sucessfully upserted proxy record for uuid %s\n", proxyRecord.UUID)

  var createdTime time.Time
  for res.Next() {
    res.Scan(&createdTime)
  }
  return createdTime
}

func (proxyExecutor *ProxyExecutor) DeleteProxyRecordTx(uuid string) {
  _, err := proxyExecutor.Tx.Exec(DELETE_PROXY_COMMAND, uuid)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", uuid, error_config.ProxyRecordLocation)
    log.Printf("Failed to delete proxy record for uuid %s with error: %+v\n", uuid, err)
    log.Panicln(errInfo.Marshal())
  }
  log.Printf("Sucessfully deleted proxy record for uuid %s\n", uuid)
}

func (proxyExecutor *ProxyExecutor) GetProxyRecordTx(uuid string) *ProxyRecord {
  var proxyRecord ProxyRecord
  err := proxyExecutor.Tx.Get(&proxyRecord, QUERY_PROXY_COMMAND, uuid)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", uuid, error_config.ProxyRecordLocation)
    log.Printf("Failed to get proxy record for uuid %s with error: %+v\n", uuid, err)
    log.Panicln(errInfo.Marshal())
  }
  return &proxyRecord
}

func (proxyExecutor *ProxyExecutor) VerifyProxyRecordExistingTx (uuid string) {
  var existing bool
  err := proxyExecutor.Tx.Get(&existing, VERIFY_PROXY_EXISTING_COMMAND, uuid)
  if err != nil {
    errInfo := error_config.MatchError(err, "uuid", uuid, error_config.ProxyRecordLocation)
    log.Printf("Failed to verify proxy record existing for uuid %s with error: %+v\n", uuid, err)
    log.Panicln(errInfo.Marshal())
  }
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoProxyUUIDExisting,
      ErrorData: map[string]interface{} {
        "uuid": uuid,
      },
      ErrorLocation: error_config.ProxyRecordLocation,
    }
    log.Printf("No proxy record for uuid %s", uuid)
    log.Panicln(errorInfo.Marshal())
  }
}

func (proxyExecutor *ProxyExecutor) GetListOfProxyUUIDsTx() *[]string {
  var proxyUUIDs []string
  err := proxyExecutor.Tx.Select(&proxyUUIDs, QUERY_LIST_OF_PROXY_UUIDS_COMMAND)

  if err != nil && err != sql.ErrNoRows {
    log.Panicf("Failed to get list of proxy UUIDs with error: %+v\n", err)
  }
  return &proxyUUIDs
}
