package lambda_get_milestone_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
)


type Request struct {
  ProjectId string `json:"projectId,required"`
  MilestoneId int64  `json:"milestoneId,required"`
}

type Response struct {
  Milestone *tcr_attributes.Milestone `json:"project,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
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
  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}


  milestoneExecutor.VerifyMilestoneRecordExisting(projectId, milestoneId)

  milestoneRecord := milestoneExecutor.GetMilestoneRecordByIDs(projectId, milestoneId)
  milestone := &tcr_attributes.Milestone{
    ProjectId:      milestoneRecord.ProjectId,
    MilestoneId:    milestoneRecord.MilestoneId,
    Content:        milestoneRecord.Content,
    StartTime:      milestoneRecord.StartTime,
    EndTime:        milestoneRecord.EndTime,
    BlockTimestamp: milestoneRecord.BlockTimestamp,
    NumObjectives:  milestoneRecord.NumObjectives,
    State: milestoneRecord.State,
    AvgRating:      milestoneRecord.AvgRating,
  }

  objectiveRecords := objectiveExecutor.GetObjectiveRecordsByProjectIdAndMilestoneId(projectId, milestoneId)

  var objectives []tcr_attributes.Objective
  for _, objectiveRecord := range *objectiveRecords {
    objective := tcr_attributes.Objective{
      ProjectId: objectiveRecord.ProjectId,
      MilestoneId: objectiveRecord.MilestoneId,
      ObjectiveId: objectiveRecord.ObjectiveId,
      Content: objectiveRecord.Content,
      BlockTimestamp: objectiveRecord.BlockTimestamp,
      AvgRating: objectiveRecord.AvgRating,
    }
    objectives = append(objectives, objective)
  }
  milestone.Objectives = &objectives
  response.Milestone = milestone

  log.Printf("Milestone Content is loaded for projectId %s and milestoneId %d\n",
    projectId, milestoneId)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
