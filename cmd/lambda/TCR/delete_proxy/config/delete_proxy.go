package config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/proxy_config"
  "log"
)


type Request struct {
  Proxy   string  `json:"proxy,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()
  postgresBigBangClient.Begin()

  proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
  existing := proxyExecutor.VerifyProxyRecordExisting(request.Proxy)
  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoProxyUUIDExisting,
      ErrorData: map[string]interface{} {
        "uuid": request.Proxy,
      },
      ErrorLocation: error_config.ProxyRecordLocation,
    }
    log.Printf("No proxy record for uuid %s", request.Proxy)
    log.Panicln(errorInfo.Marshal())
  }
  proxyExecutor.DeleteProxyRecord(request.Proxy)

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
