package project

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/TCR/project/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_project_config.Request
    response lambda_project_config.Response
    err    error
  }{
    {
      request: lambda_project_config.Request {
        ProjectId: test_constants.ProjectId1,
        Admin: test_constants.Admin1,
        Content: test_constants.ProjectContent1,
        BlockTimestamp: test_constants.BlockTimestamp1,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_constants.ProjectId2,
        Admin: test_constants.Admin1,
        Content: test_constants.ProjectContent1,
        BlockTimestamp: test_constants.BlockTimestamp2,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_constants.ProjectId3,
        Admin: test_constants.Admin1,
        Content: test_constants.ProjectContent1,
        BlockTimestamp: test_constants.BlockTimestamp3,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_constants.ProjectId4,
        Admin: test_constants.Admin1,
        Content: test_constants.ProjectContent1,
        BlockTimestamp: test_constants.BlockTimestamp4,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_constants.ProjectId5,
        Admin: test_constants.Admin1,
        Content: test_constants.ProjectContent1,
        BlockTimestamp: test_constants.BlockTimestamp5,
      },
      response: lambda_project_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_project_config.Request {
        ProjectId: test_constants.ProjectId6,
        Admin: test_constants.Admin1,
        Content: test_constants.ProjectContent1,
        BlockTimestamp: test_constants.BlockTimestamp5,
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
