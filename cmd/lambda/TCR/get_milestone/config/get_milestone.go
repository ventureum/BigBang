package lambda_get_milestone_config

import (
	"BigBang/cmd/lambda/TCR/common"
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	ProjectId   string `json:"projectId,required"`
	MilestoneId int64  `json:"milestoneId,required"`
}

type Response struct {
	Milestone *tcr_attributes.Milestone `json:"milestone,omitempty"`
	Ok        bool                      `json:"ok"`
	Message   *error_config.ErrorInfo   `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Milestone = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()

	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	projectId := request.Body.ProjectId
	milestoneId := request.Body.MilestoneId

	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

	milestoneExecutor.VerifyMilestoneRecordExistingTx(projectId, milestoneId)
	milestoneRecord := milestoneExecutor.GetMilestoneRecordByIDsTx(projectId, milestoneId)
	response.Milestone = common.ConstructMilestoneFromMilestoneRecordTx(milestoneRecord, postgresBigBangClient)

	log.Printf("Milestone Content is loaded for projectId %s and milestoneId %d\n",
		projectId, milestoneId)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
