package lambda_delete_batch_objectives_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string      `json:"principalId,required"`
	Body        RequestBody `json:"body,required"`
}

type RequestBody struct {
	RequestList []RequestContent `json:"requestList,required"`
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

	requestList := request.Body.RequestList
	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}

	for _, singleRequest := range requestList {
		objectiveExecutor.VerifyObjectiveRecordExistingTx(
			singleRequest.ProjectId, singleRequest.MilestoneId, singleRequest.ObjectiveId)
	}

	for _, singleRequest := range requestList {
		projectId := singleRequest.ProjectId
		milestoneId := singleRequest.MilestoneId
		objectiveId := singleRequest.ObjectiveId
		objectiveExecutor.DeleteObjectiveRecordByIDsTx(projectId, milestoneId, objectiveId)
		milestoneExecutor.DecreaseNumObjectivesTx(projectId, milestoneId)
		log.Printf("Objective is deleted for projectId %s, milestoneId %d and objectiveId %d\n",
			projectId, milestoneId, objectiveId)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
