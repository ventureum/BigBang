package lambda_project_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/pkg/utils"
)


type Request struct {
  ProjectId   string  `json:"projectId,required"`
  Admin       string  `json:"admin,required"`
  Content     string  `json:"content,required"`
  BlockTimestamp  int64 `json:"blockTimestamp,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToProjectRecord() (record *project_config.ProjectRecord) {
  projectRecord := &project_config.ProjectRecord{
    ID: utils.GenerateIdByInt64AndStr(request.BlockTimestamp, request.ProjectId),
    ProjectId:     request.ProjectId,
    Admin:         request.Admin,
    Content:       request.Content,
    BlockTimestamp: request.BlockTimestamp,
  }
  return projectRecord
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
  projectExecutor.UpsertProjectRecordTx(request.ToProjectRecord())

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
