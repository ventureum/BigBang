package get_objective

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/cmd/lambda/TCR/delete_milestone/config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_delete_milestone_config.Request
    response lambda_delete_milestone_config.Response
    err    error
  }{
    {
      request: lambda_delete_milestone_config.Request {
        ProjectId: test_constants.ProjectId1,
        MilestoneId: test_constants.MilestoneId1,
      },
      response: lambda_delete_milestone_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  postgresBigBangClient := client_config.ConnectPostgresClient()
  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

  for _, test := range tests {
    result, err := lambda_delete_milestone_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.False(
      t,
      milestoneExecutor.CheckMilestoneRecordExisting(
        test.request.ProjectId, test.request.MilestoneId))
  }
}