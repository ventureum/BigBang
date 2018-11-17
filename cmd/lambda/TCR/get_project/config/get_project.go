package lambda_get_project_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/cmd/lambda/TCR/common"
)


type Request struct {
  ProjectId string `json:"projectId,required"`
}

type Response struct {
  Project *tcr_attributes.Project `json:"project,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Project= nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

  projectExecutor.VerifyProjectRecordExisting(projectId)
  projectRecord := projectExecutor.GetProjectRecord(projectId)
  response.Project = common.ConstructProjectFromProjectRecord(projectRecord, postgresBigBangClient)

  log.Printf("Project Content is loaded for projectId %s\n", projectId)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
