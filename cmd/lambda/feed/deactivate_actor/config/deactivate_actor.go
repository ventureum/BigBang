package config

import (
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/feed/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
)


type Request struct {
  Actor string `json:"actor,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()


  actor := request.Actor

  postgresFeedClient.Begin()
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(actor)

  actorProfileRecordExecutor.DeactivateActorProfileRecordsTx(actor)
  log.Printf("Deactivated Profile Account for actor %s\n", actor)

  postgresFeedClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
