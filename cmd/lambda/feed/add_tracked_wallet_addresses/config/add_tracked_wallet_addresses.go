package lambda_add_tracked_wallet_addresses_config

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
	Actor             string   `json:"actor,required"`
	WalletAddressList []string `json:"walletAddressList,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
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

	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	walletAddressList := request.Body.WalletAddressList

	walletAddressRecordExecutor := wallet_address_record_config.WalletAddressRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(actor)

	for _, walletAddress := range walletAddressList {
		existing := walletAddressRecordExecutor.CheckWalletAddressExistingTx(actor, walletAddress)
		if existing {
			errorInfo := error_config.ErrorInfo{
				ErrorCode: error_config.WalletAddressAlreadyExisting,
				ErrorData: map[string]interface{}{
					"actor":         actor,
					"walletAddress": walletAddress,
				},
				ErrorLocation: error_config.WalletAddressRecordLocation,
			}
			log.Printf("Wallet Address  %s already exists for actor %s", walletAddress, actor)
			log.Panicln(errorInfo.Marshal())
		}
	}

	for _, walletAddress := range walletAddressList {
		walletAddressRecordExecutor.UpsertWalletAddressRecordTx(&wallet_address_record_config.WalletAddressRecord{
			Actor:         actor,
			WalletAddress: walletAddress,
		})
	}
	postgresBigBangClient.Commit()

	log.Printf("WalletAddressList %+v are sucessfully added for actor %s\n", walletAddressList, actor)

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
