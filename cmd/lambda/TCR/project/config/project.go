package lambda_project_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	ProjectId      string `json:"projectId,required"`
	Admin          string `json:"admin,required"`
	Content        string `json:"content,required"`
	BlockTimestamp int64  `json:"blockTimestamp,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToProjectRecord() (record *project_config.ProjectRecord) {
	projectRecord := &project_config.ProjectRecord{
		ID:             utils.GenerateIdByInt64AndStr(request.Body.BlockTimestamp, request.Body.ProjectId),
		ProjectId:      request.Body.ProjectId,
		Admin:          request.Body.Admin,
		Content:        request.Body.Content,
		BlockTimestamp: request.Body.BlockTimestamp,
	}
	return projectRecord
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

	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	projectId := request.Body.ProjectId
	admin := request.Body.Admin
	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

	existing := projectExecutor.VerifyAdminExistingTx(projectId, admin)

	if existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.ProjectAdminReassign,
			ErrorData: map[string]interface{}{
				"projectId": projectId,
				"admin":     admin,
			},
			ErrorLocation: error_config.ProjectRecordLocation,
		}
		log.Printf("Admin %s has been assigned to project with projectId %s", projectId, admin)
		log.Panicln(errorInfo.Marshal())
	}

	projectExecutor.UpsertProjectRecordTx(request.ToProjectRecord())

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
