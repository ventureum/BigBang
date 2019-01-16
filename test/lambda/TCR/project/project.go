package project_test

import (
	"BigBang/cmd/lambda/TCR/project/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_project_config.Request
		response lambda_project_config.Response
		err      error
	}{
		{
			request: lambda_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_project_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					Admin:          test_constants.ProjectAdmin1,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp1,
				},
			},
			response: lambda_project_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_project_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					Admin:          test_constants.ProjectAdmin2,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp2,
				},
			},
			response: lambda_project_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_project_config.RequestContent{
					ProjectId:      test_constants.ProjectId3,
					Admin:          test_constants.ProjectAdmin3,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp3,
				},
			},
			response: lambda_project_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_project_config.RequestContent{
					ProjectId:      test_constants.ProjectId4,
					Admin:          test_constants.ProjectAdmin4,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp4,
				},
			},
			response: lambda_project_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_project_config.RequestContent{
					ProjectId:      test_constants.ProjectId5,
					Admin:          test_constants.ProjectAdmin5,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp5,
				},
			},
			response: lambda_project_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_project_config.RequestContent{
					ProjectId:      test_constants.ProjectId6,
					Admin:          test_constants.ProjectAdmin6,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp5,
				},
			},
			response: lambda_project_config.Response{
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
