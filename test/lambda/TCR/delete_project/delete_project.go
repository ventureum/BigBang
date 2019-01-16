package delete_project_test

import (
	"BigBang/cmd/lambda/TCR/delete_project/config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_delete_project_config.Request
		response lambda_delete_project_config.Response
		err      error
	}{
		{
			request: lambda_delete_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_delete_project_config.RequestContent{
					ProjectId: test_constants.ProjectId2,
				},
			},
			response: lambda_delete_project_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}

	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	postgresBigBangClient.Begin()
	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	for _, test := range tests {
		result, err := lambda_delete_project_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.False(
			t,
			projectExecutor.CheckProjectRecordExistingTx(test.request.Body.ProjectId))
	}
	postgresBigBangClient.Commit()
}
