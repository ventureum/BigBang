package get_project_id_by_admin_test

import (
	"BigBang/cmd/lambda/TCR/get_project_id_by_admin/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

var EmptyObjectives []tcr_attributes.Objective

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_project_id_by_admin_config.Request
		response lambda_get_project_id_by_admin_config.Response
		err      error
	}{
		{
			request: lambda_get_project_id_by_admin_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_id_by_admin_config.RequestContent{
					Admin: test_constants.Actor1,
				},
			},
			response: lambda_get_project_id_by_admin_config.Response{
				ProjectId: "",
				Ok:        true,
			},
			err: nil,
		},
		{
			request: lambda_get_project_id_by_admin_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_id_by_admin_config.RequestContent{
					Admin: test_constants.ProjectAdmin1,
				},
			},
			response: lambda_get_project_id_by_admin_config.Response{
				ProjectId: test_constants.ProjectId1,
				Ok:        true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_get_project_id_by_admin_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.ProjectId, result.ProjectId)
	}
}
