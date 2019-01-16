package lambda_get_actor_uuid_from_public_key_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"log"
	"strings"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	PublicKey string `json:"publicKey,required"`
}

type Response struct {
	Actor   string                  `json:"actor,omitempty"`
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

	publicKey := request.Body.PublicKey
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	if publicKey == "" {
		errorInfo := error_config.ErrorInfo{
			ErrorCode:     error_config.EmptyPublicKey,
			ErrorLocation: error_config.ProfileAccountLocation,
		}
		log.Printf("Invalid Empty Public Key\n")
		log.Panicln(errorInfo.Marshal())
	}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{
		*postgresBigBangClient}
	response.Actor = actorProfileRecordExecutor.GetActorUuidFromPublicKeyTx(strings.ToLower(publicKey))

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
