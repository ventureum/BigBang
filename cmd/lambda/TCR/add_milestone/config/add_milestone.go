package lambda_add_milestone_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
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
	MilestoneId    int64  `json:"milestoneId,required"`
	Content        string `json:"content,required"`
	BlockTimestamp int64  `json:"blockTimestamp,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToMilestoneRecord() (record *milestone_config.MilestoneRecord) {
	milestoneRecord := &milestone_config.MilestoneRecord{
		ProjectId:      request.Body.ProjectId,
		MilestoneId:    request.Body.MilestoneId,
		Content:        request.Body.Content,
		BlockTimestamp: request.Body.BlockTimestamp,
		StartTime:      0,
		EndTime:        0,
		State:          tcr_attributes.PendingMilestoneState,
	}
	return milestoneRecord
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
	milestoneId := request.Body.MilestoneId

	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

	invalid := milestoneExecutor.ValidateMilestoneRecordUpdatingTx(projectId, milestoneId)

	if invalid {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.MilestoneInvalidForUpdating,
			ErrorData: map[string]interface{}{
				"milestoneId": milestoneId,
				"projectId":   projectId,
			},
			ErrorLocation: error_config.MilestoneRecordLocation,
		}
		log.Printf("Milestone is invalid for updating for projectId %s and milestoneId %d", projectId, milestoneId)
		log.Panicln(errorInfo.Marshal())
	}

	inserted := milestoneExecutor.UpsertMilestoneRecordTx(request.ToMilestoneRecord())

	if inserted {
		projectExecutor.IncreaseNumMilestonesTx(request.Body.ProjectId)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
