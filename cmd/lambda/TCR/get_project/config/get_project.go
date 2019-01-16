package lambda_get_project_config

import (
	"BigBang/cmd/lambda/TCR/common"
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	ProjectId string `json:"projectId,required"`
}

type Response struct {
	Project *tcr_attributes.Project `json:"project,omitempty"`
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Project = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	projectId := request.Body.ProjectId

	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

	projectExecutor.VerifyProjectRecordExistingTx(projectId)
	projectRecord := projectExecutor.GetProjectRecordTx(projectId)
	response.Project = common.ConstructProjectFromProjectRecordTx(projectRecord, postgresBigBangClient)

	log.Printf("Project Content is loaded for projectId %s\n", projectId)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
