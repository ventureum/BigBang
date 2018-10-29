package add_milestone

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/TCR/add_milestone/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_add_milestone_config.Request
    response lambda_add_milestone_config.Response
    err    error
  }{
    {
      request: lambda_add_milestone_config.Request {
        ProjectId: test_constants.ProjectId1,
        MilestoneId: test_constants.MilestoneId1,
        Content: test_constants.MilestoneContent1,
        BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
      },
      response: lambda_add_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_milestone_config.Request {
        ProjectId: test_constants.ProjectId1,
        MilestoneId: test_constants.MilestoneId2,
        Content: test_constants.MilestoneContent1,
        BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
      },
      response: lambda_add_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_milestone_config.Request {
        ProjectId: test_constants.ProjectId1,
        MilestoneId: test_constants.MilestoneId3,
        Content: test_constants.MilestoneContent1,
        BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
      },
      response: lambda_add_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_milestone_config.Request {
        ProjectId: test_constants.ProjectId1,
        MilestoneId: test_constants.MilestoneId4,
        Content: test_constants.MilestoneContent1,
        BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
      },
      response: lambda_add_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_milestone_config.Request {
        ProjectId: test_constants.ProjectId2,
        MilestoneId: test_constants.MilestoneId1,
        Content: test_constants.MilestoneContent1,
        BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
      },
      response: lambda_add_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_milestone_config.Request {
        ProjectId: test_constants.ProjectId2,
        MilestoneId: test_constants.MilestoneId2,
        Content: test_constants.MilestoneContent1,
        BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
      },
      response: lambda_add_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_add_milestone_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
