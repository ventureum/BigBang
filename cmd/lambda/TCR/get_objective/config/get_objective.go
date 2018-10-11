package lambda_get_objective_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
)


type Request struct {
  ProjectId   string  `json:"projectId,required"`
  MilestoneId int64  `json:"milestoneId,required"`
  ObjectiveId int64 `json:"objectiveId,required"`
}

type Response struct {
  Objective *tcr_attributes.Objective `json:"objective,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Objective = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId
  milestoneId := request.MilestoneId
  objectiveId := request.ObjectiveId

  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}

  objectiveExecutor.VerifyObjectiveRecordExisting(projectId, milestoneId, objectiveId)

  objectiveRecord := objectiveExecutor.GetObjectiveRecordByIDs(projectId, milestoneId, objectiveId)
  objective := &tcr_attributes.Objective{
    ProjectId: objectiveRecord.ProjectId,
    MilestoneId: objectiveRecord.MilestoneId,
    ObjectiveId: objectiveRecord.ObjectiveId,
    Content: objectiveRecord.Content,
    BlockTimestamp: objectiveRecord.BlockTimestamp,
    AvgRating: objectiveRecord.AvgRating,
  }

  response.Objective = objective

  log.Printf("Objective Content is loaded for projectId %s, milestoneId %d and objectiveId %d\n",
    projectId,  milestoneId, objectiveId)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
