package set_token_pool_test

import (
	"BigBang/cmd/lambda/feed/set_token_pool/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

var NextRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(1)
var NextNRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(10000)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_set_token_pool_config.Request
		response lambda_set_token_pool_config.Response
		err      error
	}{
		{
			request: lambda_set_token_pool_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_set_token_pool_config.RequestContent{
					RedeemBlockTimestamp: NextRedeemBlock.ConvertToTime().Unix(),
					TokenPool:            test_constants.TokenPoolSize1,
				},
			},
			response: lambda_set_token_pool_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_set_token_pool_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_set_token_pool_config.RequestContent{
					RedeemBlockTimestamp: NextNRedeemBlock.ConvertToTime().Unix(),
					TokenPool:            test_constants.TokenPoolSize1,
				},
			},
			response: lambda_set_token_pool_config.Response{
				Ok: false,
				Message: &error_config.ErrorInfo{
					ErrorCode: error_config.NoReDeemBlockInfoRecordExisting,
					ErrorData: map[string]interface{}{
						"redeemBlock": float64(NextNRedeemBlock),
					},
					ErrorLocation: error_config.RedeemBlockInfoRecordLocation,
				},
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_set_token_pool_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
