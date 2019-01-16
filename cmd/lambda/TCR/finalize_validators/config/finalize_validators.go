package lambda_finalize_validators_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_validator_record_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	ProjectId   string   `json:"projectId,required"`
	MilestoneId int64    `json:"milestoneId,required"`
	Validators  []string `json:"validators,required"`
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

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	projectId := request.Body.ProjectId
	milestoneId := request.Body.MilestoneId
	validators := request.Body.Validators
	milestoneValidatorRecordExecutor := milestone_validator_record_config.MilestoneValidatorRecordExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
	milestoneExecutor.VerifyMilestoneRecordExistingTx(projectId, milestoneId)

	for _, validator := range validators {
		existing := milestoneValidatorRecordExecutor.CheckMilestoneValidatorRecordExistingTx(projectId, milestoneId, validator)
		if existing {
			errorInfo := error_config.ErrorInfo{
				ErrorCode:     error_config.MilestoneValidatorAlreadyExisting,
				ErrorLocation: error_config.MilestoneValidatorRecordLocation,
				ErrorData: error_config.ErrorData{
					"projectId":   projectId,
					"milestoneId": milestoneId,
					"validator":   validator,
				},
			}
			log.Printf("Milestone Validator Already Exists for projectId %s and milestonrId %d\n", projectId, milestoneId)
			log.Panicln(errorInfo.Marshal())
		}
	}

	for _, validator := range validators {
		milestoneValidatorRecordExecutor.UpsertMilestoneValidatorRecordTx(&milestone_validator_record_config.MilestoneValidatorRecord{
			ProjectId:   projectId,
			MilestoneId: milestoneId,
			Validator:   validator,
		})
	}
	postgresBigBangClient.Commit()

	log.Printf("Validators %+v are sucessfully added for projectId %s and milestoneId %d\n", validators, projectId, milestoneId)

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
