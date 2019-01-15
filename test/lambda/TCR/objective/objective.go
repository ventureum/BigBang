package objective_test

import (
	"BigBang/cmd/lambda/TCR/objective/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_objective_config.Request
		response lambda_objective_config.Response
		err      error
	}{
		{
			request: lambda_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_objective_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId1,
					ObjectiveId:    test_constants.ObjectiveId1,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp1,
				},
			},
			response: lambda_objective_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_objective_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId1,
					ObjectiveId:    test_constants.ObjectiveId2,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp2,
				},
			},
			response: lambda_objective_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_objective_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId2,
					ObjectiveId:    test_constants.ObjectiveId1,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp3,
				},
			},
			response: lambda_objective_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_objective_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId2,
					ObjectiveId:    test_constants.ObjectiveId2,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp4,
				},
			},
			response: lambda_objective_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_objective_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId2,
					ObjectiveId:    test_constants.ObjectiveId3,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp5,
				},
			},
			response: lambda_objective_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_objective_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    test_constants.MilestoneId1,
					ObjectiveId:    test_constants.ObjectiveId1,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp2,
				},
			},
			response: lambda_objective_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_objective_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
