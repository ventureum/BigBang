package get_batch_finalized_validators_test

import (
	"BigBang/cmd/lambda/TCR/get_batch_finalized_validators/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_batch_finalized_validators_config.Request
		response lambda_get_batch_finalized_validators_config.Response
		err      error
	}{
		{
			request: lambda_get_batch_finalized_validators_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_batch_finalized_validators_config.RequestContent{
					MilestoneValidatorsInfoKeyList: []tcr_attributes.MilestoneValidatorsInfoKey{
						{
							ProjectId:   test_constants.ProjectId1,
							MilestoneId: test_constants.MilestoneId1,
						},
					},
				},
			},
			response: lambda_get_batch_finalized_validators_config.Response{
				MilestoneValidatorsInfoList: &[]tcr_attributes.MilestoneValidatorsInfo{
					{
						MilestoneValidatorsInfoKey: tcr_attributes.MilestoneValidatorsInfoKey{
							ProjectId:   test_constants.ProjectId1,
							MilestoneId: test_constants.MilestoneId1,
						},
						Validators: &[]string{
							test_constants.Actor1,
							test_constants.Actor2,
						},
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_get_batch_finalized_validators_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
