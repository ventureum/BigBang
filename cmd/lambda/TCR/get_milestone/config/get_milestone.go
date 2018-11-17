package lambda_get_milestone_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/cmd/lambda/TCR/common"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
)


type Request struct {
  ProjectId string `json:"projectId,required"`
  MilestoneId int64  `json:"milestoneId,required"`
}

type Response struct {
  Milestone *tcr_attributes.Milestone `json:"milestone,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Milestone = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId
  milestoneId := request.MilestoneId

  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

  milestoneExecutor.VerifyMilestoneRecordExisting(projectId, milestoneId)
  milestoneRecord := milestoneExecutor.GetMilestoneRecordByIDs(projectId, milestoneId)
  response.Milestone = common.ConstructMilestoneFromMilestoneRecord(milestoneRecord, postgresBigBangClient)

  log.Printf("Milestone Content is loaded for projectId %s and milestoneId %d\n",
    projectId, milestoneId)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
