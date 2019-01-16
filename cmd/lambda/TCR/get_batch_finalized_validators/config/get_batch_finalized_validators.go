package lambda_get_batch_finalized_validators_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
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
	MilestoneValidatorsInfoKeyList []tcr_attributes.MilestoneValidatorsInfoKey `json:"milestoneValidatorsInfoKeyList,required"`
}

type Response struct {
	MilestoneValidatorsInfoList *[]tcr_attributes.MilestoneValidatorsInfo `json:"milestoneValidatorsInfoList,omitempty"`
	Ok                          bool                                      `json:"ok"`
	Message                     *error_config.ErrorInfo                   `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.MilestoneValidatorsInfoList = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	milestoneValidatorRecordExecutor := milestone_validator_record_config.MilestoneValidatorRecordExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

	milestoneValidatorsInfoKeyList := request.Body.MilestoneValidatorsInfoKeyList

	for _, milestoneValidatorsInfoKey := range milestoneValidatorsInfoKeyList {
		projectId := milestoneValidatorsInfoKey.ProjectId
		milestoneId := milestoneValidatorsInfoKey.MilestoneId
		milestoneExecutor.VerifyMilestoneRecordExistingTx(projectId, milestoneId)
	}

	var milestoneValidatorInfoList []tcr_attributes.MilestoneValidatorsInfo
	for _, milestoneValidatorsInfoKey := range milestoneValidatorsInfoKeyList {
		projectId := milestoneValidatorsInfoKey.ProjectId
		milestoneId := milestoneValidatorsInfoKey.MilestoneId
		validators := milestoneValidatorRecordExecutor.GetMilestoneValidatorListByProjectIdAndMilestoneIdTx(projectId, milestoneId)
		milestoneValidatorInfo := tcr_attributes.MilestoneValidatorsInfo{
			MilestoneValidatorsInfoKey: tcr_attributes.MilestoneValidatorsInfoKey{
				ProjectId:   projectId,
				MilestoneId: milestoneId,
			},
			Validators: validators,
		}
		milestoneValidatorInfoList = append(milestoneValidatorInfoList, milestoneValidatorInfo)
		log.Printf("Validators are sucessfully loaded for projectId %s and milestoneId %d\n", projectId, milestoneId)
	}

	response.MilestoneValidatorsInfoList = &milestoneValidatorInfoList

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
