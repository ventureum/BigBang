package finalize_validators_test

import (
	"BigBang/cmd/lambda/TCR/finalize_validators/config"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_finalize_validators_config.Request
		response lambda_finalize_validators_config.Response
		err      error
	}{
		{
			request: lambda_finalize_validators_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_finalize_validators_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: test_constants.MilestoneId1,
					Validators: []string{
						test_constants.Actor1,
						test_constants.Actor2,
					},
				},
			},
			response: lambda_finalize_validators_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_finalize_validators_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_finalize_validators_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: test_constants.MilestoneId1,
					Validators: []string{
						test_constants.Actor3,
						test_constants.Actor2,
					},
				},
			},
			response: lambda_finalize_validators_config.Response{
				Ok: false,
				Message: &error_config.ErrorInfo{
					ErrorData: error_config.ErrorData{
						"projectId":   test_constants.ProjectId1,
						"milestoneId": float64(test_constants.MilestoneId1),
						"validator":   test_constants.Actor2,
					},
					ErrorCode:     error_config.MilestoneValidatorAlreadyExisting,
					ErrorLocation: error_config.MilestoneValidatorRecordLocation,
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_finalize_validators_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
