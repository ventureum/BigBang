package lambda_add_proxy_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/proxy_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
)


type Request struct {
  Proxy   string  `json:"proxy,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()
  postgresBigBangClient.Begin()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(request.Proxy)
  existing := proxyExecutor.VerifyProxyRecordExistingTx(request.Proxy)
  if existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.ProxyUUIDAlreadyExisting,
      ErrorData: map[string]interface{} {
        "uuid": request.Proxy,
      },
      ErrorLocation: error_config.ProxyRecordLocation,
    }
    log.Printf("Proxy record already exists for uuid %s", request.Proxy)
    log.Panicln(errorInfo.Marshal())
  }

  proxyExecutor.UpsertProxyRecordTx(&proxy_config.ProxyRecord{
    UUID: request.Proxy,
  })

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
