package project

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/TCR/project/config"
  "BigBang/test/contants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_project_config.Request
    response lambda_project_config.Response
    err    error
  }{
    {
      request: lambda_project_config.Request {
        ProjectId: test_contants.ProjectId1,
        Admin: test_contants.Admin1,
        Content: test_contants.ProjectContent1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_contants.ProjectId2,
        Admin: test_contants.Admin1,
        Content: test_contants.ProjectContent1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_contants.ProjectId3,
        Admin: test_contants.Admin1,
        Content: test_contants.ProjectContent1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_contants.ProjectId4,
        Admin: test_contants.Admin1,
        Content: test_contants.ProjectContent1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_contants.ProjectId5,
        Admin: test_contants.Admin1,
        Content: test_contants.ProjectContent1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_contants.ProjectId6,
        Admin: test_contants.Admin1,
        Content: test_contants.ProjectContent1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_project_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
