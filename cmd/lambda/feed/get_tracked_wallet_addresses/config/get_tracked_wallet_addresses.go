package lambda_get_tracked_wallet_addresses_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/wallet_address_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor string `json:"actor,required"`
}

type Response struct {
	WalletAddressList *[]string               `json:"walletAddressList,omitempty"`
	Ok                bool                    `json:"ok"`
	Message           *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.WalletAddressList = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	walletAddressRecordExecutor := wallet_address_record_config.WalletAddressRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(actor)
	response.WalletAddressList = walletAddressRecordExecutor.GetWalletAddressListByActorTx(actor)

	postgresBigBangClient.Commit()

	log.Printf("WalletAddressList are sucessfully loaded for actor %s\n", actor)

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
