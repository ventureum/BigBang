package batch_objectives_test

import (
	"BigBang/cmd/lambda/TCR/batch_objectives/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_batch_objectives_config.Request
		response lambda_batch_objectives_config.Response
		err      error
	}{
		{
			request: lambda_batch_objectives_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_batch_objectives_config.RequestBody{
					RequestList: []lambda_batch_objectives_config.RequestContent{
						{
							ProjectId:      test_constants.ProjectId1,
							MilestoneId:    test_constants.MilestoneId2,
							ObjectiveId:    test_constants.ObjectiveId3,
							Content:        test_constants.ObjectiveContent1,
							BlockTimestamp: test_constants.ObjectiveBlockTimestamp5,
						},
						{
							ProjectId:      test_constants.ProjectId2,
							MilestoneId:    test_constants.MilestoneId1,
							ObjectiveId:    test_constants.ObjectiveId1,
							Content:        test_constants.ObjectiveContent1,
							BlockTimestamp: test_constants.ObjectiveBlockTimestamp2,
						},
						{
							ProjectId:      test_constants.ProjectId2,
							MilestoneId:    test_constants.MilestoneId1,
							ObjectiveId:    test_constants.ObjectiveId2,
							Content:        test_constants.ObjectiveContent1,
							BlockTimestamp: test_constants.ObjectiveBlockTimestamp2,
						},
					},
				},
			},
			response: lambda_batch_objectives_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_batch_objectives_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
