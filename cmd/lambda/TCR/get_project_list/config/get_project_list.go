package lambda_get_project_list_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/pkg/utils"
  "BigBang/internal/app/tcr_attributes"
)


type Request struct {
  Limit int64 `json:"limit,required"`
  Cursor string `json:"cursor,omitempty"`
}

type Response struct {
  Projects *[]tcr_attributes.Project`json:"projects,omitempty"`
  NextCursor string `json:"nextCursor,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Projects = nil
      response.NextCursor = ""
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  limit := request.Limit
  cursorStr := request.Cursor

  var cursor int64
  if cursorStr != "" {
    cursor = utils.Base64DecodeToInt64(cursorStr)
  }

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

  projectRecords := projectExecutor.GetProjectRecordsByCursor(cursor, limit + 1)

  response.NextCursor = ""

  var projects []tcr_attributes.Project
  for index, projectRecord := range *projectRecords {
    if index < int(limit) {
      project := &tcr_attributes.Project{
        ProjectId: projectRecord.ProjectId,
        Admin: projectRecord.Admin,
        Content: projectRecord.Content,
        AvgRating: projectRecord.AvgRating,
      }
      projects = append(projects, *project)
    } else {
      response.NextCursor = utils.Base64EncodeInt64(projectRecord.ID)
    }
  }

  response.Projects = &projects
  if cursorStr == "" {
    log.Printf("ProjectRecords is loaded for first query with limit %d\n", limit)
  } else {
    log.Printf("ProjectRecords is loaded for query with cursor %s and limit %d\n", cursorStr, limit)
  }
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
