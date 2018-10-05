package config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/proxy_config"
  "BigBang/internal/pkg/utils"
  "log"
)


type Request struct {
  Limit int64 `json:"limit,required"`
  Cursor string `json:"cursor,omitempty"`
}

type Response struct {
  Proxies *[]string `json:"proxies,omitempty"`
  NextCursor string `json:"nextCursor,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Proxies = nil
      response.NextCursor = ""
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  limit := request.Limit
  cursorStr := request.Cursor

  var cursor int64
  if cursorStr != "" {
    cursor = utils.Base64DecodeToInt64(cursorStr)
  }

  proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}

  proxyRecords := proxyExecutor.GetListOfProxyByCursor(cursor, limit + 1)

  response.NextCursor = ""

  var proxyUUIDList []string
  for index, proxyRecord := range *proxyRecords {
    if index < int(limit) {
      proxyUUIDList = append(proxyUUIDList, proxyRecord.UUID)
    } else {
      response.NextCursor = utils.Base64EncodeInt64(proxyRecord.ID)
    }
  }

  response.Proxies = &proxyUUIDList

  if cursorStr == "" {
    log.Printf("proxyUUIDList  is loaded for first query with limit %d\n", limit)
  } else {
    log.Printf("ProxyUUIDList  is loaded for query with cursor %s and limit %d\n", cursorStr, limit)
  }
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}