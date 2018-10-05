package config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
)


type Request struct {
  ProjectId   string  `json:"projectId,required"`
  Content     string  `json:"content,required"`
  AvgRating   int64   `json:"avgRating,required"`
  MilestoneInfo tcr_attributes.MilestoneInfo `json:"milestoneInfo,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToProjectRecord() (record *project_config.ProjectRecord) {
  return &project_config.ProjectRecord{
    ProjectId:     request.ProjectId,
    Content:       request.Content,
    AvgRating:     request.AvgRating,
    MilestoneInfo: request.MilestoneInfo.ToJsonText(),
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()
  postgresBigBangClient.Begin()

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  projectExecutor.UpsertProjectRecordTx(request.ToProjectRecord())

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
