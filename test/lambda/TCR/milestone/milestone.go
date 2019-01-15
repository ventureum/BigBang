package milestone_test

import (
	"BigBang/cmd/lambda/TCR/milestone/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_milestone_config.Request
		response lambda_milestone_config.Response
		err      error
	}{
		{
			request: lambda_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId1,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
					StartTime:      test_constants.MilestoneStartTime1,
					EndTime:        test_constants.MilestoneEndTime1,
					State:          tcr_attributes.CompleteMilestoneState,
				},
			},
			response: lambda_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId2,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
					StartTime:      test_constants.MilestoneStartTime2,
					EndTime:        test_constants.MilestoneEndTime2,
					State:          tcr_attributes.InProgressMilestoneState,
				},
			},
			response: lambda_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId3,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
					StartTime:      test_constants.MilestoneStartTime3,
					EndTime:        test_constants.MilestoneEndTime3,
					State:          tcr_attributes.PendingMilestoneState,
				},
			},
			response: lambda_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId4,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
					StartTime:      test_constants.MilestoneStartTime4,
					EndTime:        test_constants.MilestoneEndTime4,
					State:          tcr_attributes.PendingMilestoneState,
				},
			},
			response: lambda_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    test_constants.MilestoneId1,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
					StartTime:      test_constants.MilestoneStartTime1,
					EndTime:        test_constants.MilestoneEndTime1,
					State:          tcr_attributes.PendingMilestoneState,
				},
			},
			response: lambda_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_milestone_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    test_constants.MilestoneId2,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
					StartTime:      test_constants.MilestoneStartTime2,
					EndTime:        test_constants.MilestoneEndTime2,
					State:          tcr_attributes.PendingMilestoneState,
				},
			},
			response: lambda_milestone_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_milestone_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
