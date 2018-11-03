package lambda_get_finalized_validators_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_validator_record_config"
)


type Request struct {
  ProjectId   string  `json:"projectId,required"`
  MilestoneId  int64  `json:"milestoneId,required"`
}

type Response struct {
  Validators *[]string `json:"validators,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Validators = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId
  milestoneId := request.MilestoneId
  milestoneValidatorRecordExecutor := milestone_validator_record_config.MilestoneValidatorRecordExecutor{*postgresBigBangClient}
  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
  milestoneExecutor.VerifyMilestoneRecordExisting(projectId, milestoneId)

  response.Validators= milestoneValidatorRecordExecutor.GetMilestoneValidatorListByProjectIdAndMilestoneId(projectId, milestoneId)

  log.Printf("Validators are sucessfully loaded for projectId %s and milestoneId %d\n", projectId, milestoneId)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
