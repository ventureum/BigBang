package lambda_milestone_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
)


type Request struct {
  ProjectId   string                  `json:"projectId,required"`
  MilestoneId int64                   `json:"milestoneId,required"`
  Content     string                  `json:"content,required"`
  BlockTimestamp  int64               `json:"blockTimestamp,required"`
  StartTime  int64                    `json:"startTime,required"`
  EndTime    int64                    `json:"endTime,required"`
  State tcr_attributes.MilestoneState `json:"state,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToMilestoneRecord() (record *milestone_config.MilestoneRecord) {
  milestoneRecord := &milestone_config.MilestoneRecord{
    ProjectId:     request.ProjectId,
    MilestoneId: request.MilestoneId,
    Content:       request.Content,
    BlockTimestamp: request.BlockTimestamp,
    StartTime: request.StartTime,
    EndTime: request.EndTime,
    State: request.State,
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

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
  inserted := milestoneExecutor.UpsertMilestoneRecordTx(request.ToMilestoneRecord())

  if inserted {
    projectExecutor.IncreaseNumMilestonesTx(request.ProjectId)
  }

  switch state := request.State; state {
    case tcr_attributes.CompleteMilestoneState:
      projectExecutor.IncreaseNumMilestonesCompletedTx(request.ProjectId)
    case tcr_attributes.InProgressMilestoneState:
      projectExecutor.SetCurrentMilestoneTx(request.ProjectId, request.MilestoneId)
  }

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
