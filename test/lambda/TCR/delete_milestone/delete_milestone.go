package delete_milestone_test

import (
	"BigBang/cmd/lambda/TCR/delete_milestone/config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_delete_milestone_config.Request
		response lambda_delete_milestone_config.Response
		err      error
	}{
		{
			request: lambda_delete_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_delete_milestone_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: test_constants.MilestoneId1,
				},
			},
			response: lambda_delete_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}

	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	postgresBigBangClient.Begin()
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

	for _, test := range tests {
		result, err := lambda_delete_milestone_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.False(
			t,
			milestoneExecutor.CheckMilestoneRecordExistingTx(
				test.request.Body.ProjectId, test.request.Body.MilestoneId))
	}
	postgresBigBangClient.Commit()
}
