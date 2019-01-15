package lambda_get_session_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/post_config"
	"BigBang/internal/platform/postgres_config/feed/session_record_config"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	PostHash string `json:"postHash,required"`
}

type Response struct {
	Session *session_record_config.SessionRecordResult `json:"session,omitempty"`
	Ok      bool                                       `json:"ok"`
	Message *error_config.ErrorInfo                    `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Session = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)
	postHash := request.Body.PostHash

	postExecutor := post_config.PostExecutor{*postgresBigBangClient}
	sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresBigBangClient}

	postExecutor.VerifyPostRecordExistingTx(postHash)

	response.Session = sessionRecordExecutor.GetSessionRecordTx(postHash).ToSessionRecordResult()

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
