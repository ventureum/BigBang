package delete_project

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/cmd/lambda/TCR/delete_project/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_delete_project_config.Request
    response lambda_delete_project_config.Response
    err    error
  }{
    {
      request: lambda_delete_project_config.Request {
        ProjectId: test_constants.ProjectId1,
      },
      response: lambda_delete_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  postgresBigBangClient := client_config.ConnectPostgresClient()
  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

  for _, test := range tests {
    result, err := lambda_delete_project_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.False(
      t,
      projectExecutor.CheckProjectRecordExisting(test.request.ProjectId))
  }
}
