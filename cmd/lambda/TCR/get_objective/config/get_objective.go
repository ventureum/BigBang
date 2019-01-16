package lambda_get_objective_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
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
	Objective *tcr_attributes.Objective `json:"objective,omitempty"`
	Ok        bool                      `json:"ok"`
	Message   *error_config.ErrorInfo   `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Objective = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	projectId := request.Body.ProjectId
	milestoneId := request.Body.MilestoneId
	objectiveId := request.Body.ObjectiveId

	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}

	objectiveExecutor.VerifyObjectiveRecordExistingTx(projectId, milestoneId, objectiveId)

	objectiveRecord := objectiveExecutor.GetObjectiveRecordByIDsTx(projectId, milestoneId, objectiveId)
	objective := &tcr_attributes.Objective{
		ProjectId:      objectiveRecord.ProjectId,
		MilestoneId:    objectiveRecord.MilestoneId,
		ObjectiveId:    objectiveRecord.ObjectiveId,
		Content:        objectiveRecord.Content,
		BlockTimestamp: objectiveRecord.BlockTimestamp,
		AvgRating:      objectiveRecord.AvgRating,
	}

	response.Objective = objective

	log.Printf("Objective Content is loaded for projectId %s, milestoneId %d and objectiveId %d\n",
		projectId, milestoneId, objectiveId)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
