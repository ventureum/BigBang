package lambda_get_tracked_wallet_addresses_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/wallet_address_record_config"
)


type Request struct {
  Actor   string  `json:"actor,required"`
}

type Response struct {
  WalletAddressList *[]string `json:"walletAddressList,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.WalletAddressList = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()
  postgresBigBangClient.Begin()


  actor := request.Actor
  walletAddressRecordExecutor := wallet_address_record_config.WalletAddressRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExisting(actor)
  response.WalletAddressList = walletAddressRecordExecutor.GetWalletAddressListByActor(actor)


  log.Printf("WalletAddressList are sucessfully loaded for actor %s\n", actor)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
