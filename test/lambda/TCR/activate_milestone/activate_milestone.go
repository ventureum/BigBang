package activate_milestone_test

import (
	"BigBang/cmd/lambda/TCR/activate_milestone/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_activate_milestone_config.Request
		response lambda_activate_milestone_config.Response
		err      error
	}{
		{
			request: lambda_activate_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_activate_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
					StartTime:      test_constants.MilestoneStartTime1,
				},
			},
			response: lambda_activate_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_activate_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_activate_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId2,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
					StartTime:      test_constants.MilestoneStartTime2,
				},
			},
			response: lambda_activate_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_activate_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_activate_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId3,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
					StartTime:      test_constants.MilestoneStartTime3,
				},
			},
			response: lambda_activate_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_activate_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_activate_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId4,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
					StartTime:      test_constants.MilestoneStartTime4,
				},
			},
			response: lambda_activate_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_activate_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_activate_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    test_constants.MilestoneId1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
					StartTime:      test_constants.MilestoneStartTime1,
				},
			},
			response: lambda_activate_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_activate_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_activate_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    test_constants.MilestoneId2,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
					StartTime:      test_constants.MilestoneStartTime2,
				},
			},
			response: lambda_activate_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_activate_milestone_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
