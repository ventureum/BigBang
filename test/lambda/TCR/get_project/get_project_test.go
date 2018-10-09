package get_project

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/cmd/lambda/TCR/get_project/config"
  "BigBang/test/contants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_project_config.Request
    response lambda_get_project_config.Response
    err    error
  }{
    {
      request: lambda_get_project_config.Request {
        ProjectId: test_contants.ProjectId1,
      },
      response: lambda_get_project_config.Response {
        Project: &tcr_attributes.Project{
          ProjectId: test_contants.ProjectId1,
          Admin: test_contants.Admin1,
          Content:   test_contants.ProjectContent1,
        },
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_project_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.Equal(t, test.response.Project.ProjectId, result.Project.ProjectId)
    assert.Equal(t, test.response.Project.Admin, result.Project.Admin)
    assert.Equal(t, test.response.Project.Content, result.Project.Content)
  }
}
