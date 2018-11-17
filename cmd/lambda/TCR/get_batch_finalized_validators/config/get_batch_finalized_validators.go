package lambda_get_batch_finalized_validators_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_validator_record_config"
  "BigBang/internal/app/tcr_attributes"
)

type Request struct {
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
    }
    postgresBigBangClient.Close()
  }()

  milestoneValidatorRecordExecutor := milestone_validator_record_config.MilestoneValidatorRecordExecutor{*postgresBigBangClient}
  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

  milestoneValidatorsInfoKeyList := request.MilestoneValidatorsInfoKeyList

  for _ , milestoneValidatorsInfoKey := range milestoneValidatorsInfoKeyList {
    projectId := milestoneValidatorsInfoKey.ProjectId
    milestoneId := milestoneValidatorsInfoKey.MilestoneId
    milestoneExecutor.VerifyMilestoneRecordExisting(projectId, milestoneId)
  }

  var milestoneValidatorInfoList []tcr_attributes.MilestoneValidatorsInfo
  for _ , milestoneValidatorsInfoKey := range milestoneValidatorsInfoKeyList {
    projectId := milestoneValidatorsInfoKey.ProjectId
    milestoneId := milestoneValidatorsInfoKey.MilestoneId
    validators := milestoneValidatorRecordExecutor.GetMilestoneValidatorListByProjectIdAndMilestoneId(projectId, milestoneId)
    milestoneValidatorInfo := tcr_attributes.MilestoneValidatorsInfo{
      MilestoneValidatorsInfoKey: tcr_attributes.MilestoneValidatorsInfoKey{
        ProjectId: projectId,
        MilestoneId: milestoneId,
      },
      Validators: validators,
    }
    milestoneValidatorInfoList = append(milestoneValidatorInfoList, milestoneValidatorInfo)
    log.Printf("Validators are sucessfully loaded for projectId %s and milestoneId %d\n", projectId, milestoneId)
  }

  response.MilestoneValidatorsInfoList = &milestoneValidatorInfoList

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
