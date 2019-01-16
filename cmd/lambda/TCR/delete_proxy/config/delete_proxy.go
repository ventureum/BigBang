package lambda_delete_proxy_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/proxy_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Proxy string `json:"proxy,required"`
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

	proxy := request.Body.Proxy
	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(proxy)
	existing := proxyExecutor.VerifyProxyRecordExistingTx(proxy)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoProxyUUIDExisting,
			ErrorData: map[string]interface{}{
				"uuid": proxy,
			},
			ErrorLocation: error_config.ProxyRecordLocation,
		}
		log.Printf("No proxy record for uuid %s", proxy)
		log.Panicln(errorInfo.Marshal())
	}
	proxyExecutor.DeleteProxyRecordTx(proxy)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
