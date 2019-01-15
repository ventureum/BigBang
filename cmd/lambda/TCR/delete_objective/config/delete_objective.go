package lambda_delete_objective_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
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
	ObjectiveId int64  `json:"objectiveId,required"`
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

	projectId := request.Body.ProjectId
	milestoneId := request.Body.MilestoneId
	objectiveId := request.Body.ObjectiveId
	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

	objectiveExecutor.VerifyObjectiveRecordExistingTx(projectId, milestoneId, objectiveId)
	objectiveExecutor.DeleteObjectiveRecordByIDsTx(projectId, milestoneId, objectiveId)
	milestoneExecutor.DecreaseNumObjectivesTx(projectId, milestoneId)

	postgresBigBangClient.Commit()

	log.Printf("Objective is deleted for projectId %s, milestoneId %d and objectiveId %d\n",
		projectId, milestoneId, objectiveId)

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
