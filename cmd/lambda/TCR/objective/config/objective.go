package lambda_objective_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
)


type Request struct {
  ProjectId   string  `json:"projectId,required"`
  MilestoneId int64  `json:"milestoneId,required"`
  ObjectiveId int64 `json:"objectiveId,required"`
  Content     string  `json:"content,required"`
  BlockTimestamp  int64 `json:"blockTimestamp,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToObjectiveRecord() (record *objective_config.ObjectiveRecord) {
  objectiveRecord := &objective_config.ObjectiveRecord{
    ProjectId:     request.ProjectId,
    MilestoneId: request.MilestoneId,
    ObjectiveId: request.ObjectiveId,
    Content:       request.Content,
    BlockTimestamp: request.BlockTimestamp,
  }
  return objectiveRecord
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

  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
  inserted := objectiveExecutor.UpsertObjectiveRecordTx(request.ToObjectiveRecord())
  if inserted {
    milestoneExecutor.IncreaseNumObjectivesTx(request.ProjectId, request.MilestoneId)
  }
  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
