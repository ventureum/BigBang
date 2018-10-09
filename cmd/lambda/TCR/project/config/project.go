package lambda_project_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
)


type Request struct {
  ProjectId   string  `json:"projectId,required"`
  Admin       string  `json:"admin,required"`
  Content     string  `json:"content,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToProjectRecord() (record *project_config.ProjectRecord) {
  return &project_config.ProjectRecord{
    ProjectId:     request.ProjectId,
    Admin:         request.Admin,
    Content:       request.Content,
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
