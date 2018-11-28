package delete_batch_objectives

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
  "BigBang/cmd/lambda/TCR/delete_batch_objectives/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_delete_batch_objectives_config.Request
    response lambda_delete_batch_objectives_config.Response
    err    error
  }{
    {
      request: lambda_delete_batch_objectives_config.Request{
        PrincipalId: test_constants.Actor1,
        Body: lambda_delete_batch_objectives_config.RequestBody{
          RequestList: []lambda_delete_batch_objectives_config.RequestContent{
            {
              ProjectId:   test_constants.ProjectId1,
              MilestoneId: test_constants.MilestoneId1,
              ObjectiveId: test_constants.ObjectiveId2,
            },
            {
              ProjectId:   test_constants.ProjectId1,
              MilestoneId: test_constants.MilestoneId2,
              ObjectiveId: test_constants.ObjectiveId1,
            },
          },
        },
      },
      response: lambda_delete_batch_objectives_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}

  for _, test := range tests {
    result, err := lambda_delete_batch_objectives_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    for _, singleRequest := range test.request.Body.RequestList {
      assert.False(
        t,
        objectiveExecutor.CheckObjectiveRecordExisting(
          singleRequest.ProjectId, singleRequest.MilestoneId, singleRequest.ObjectiveId))
    }
  }
}
