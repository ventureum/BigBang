package get_project

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/cmd/lambda/TCR/get_project/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_project_config.Request
    response lambda_get_project_config.Response
    err    error
  }{
    {
      request: lambda_get_project_config.Request {
        ProjectId: test_constants.ProjectId1,
      },
      response: lambda_get_project_config.Response {
        Project: &tcr_attributes.Project{
          ProjectId: test_constants.ProjectId1,
          Admin: test_constants.ProjectAdmin1,
          Content:   test_constants.ProjectContent1,
          BlockTimestamp: test_constants.ProjectBlockTimestamp1,
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
    assert.Equal(t, test.response.Project.BlockTimestamp, result.Project.BlockTimestamp)
  }
}
