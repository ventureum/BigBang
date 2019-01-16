package get_redeem_block_info_test

import (
	"BigBang/cmd/lambda/feed/get_redeem_block_info/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

var NextRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(1)
var ExecutedAt = NextRedeemBlock.ConvertToTime()
var NextNRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(10000)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_redeem_block_info_config.Request
		response lambda_get_redeem_block_info_config.Response
		err      error
	}{
		{
			request: lambda_get_redeem_block_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_redeem_block_info_config.RequestContent{
					RedeemBlockTimestamp: NextRedeemBlock.ConvertToTime().Unix(),
				},
			},
			response: lambda_get_redeem_block_info_config.Response{
				RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
					RedeemBlock:                  NextRedeemBlock,
					TotalEnrolledMilestonePoints: 400,
					TokenPool:                    test_constants.TokenPoolSize1,
					ExecutedAt:                   ExecutedAt,
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_redeem_block_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_redeem_block_info_config.RequestContent{
					RedeemBlockTimestamp: NextNRedeemBlock.ConvertToTime().Unix(),
				},
			},
			response: lambda_get_redeem_block_info_config.Response{
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
		result, err := lambda_get_redeem_block_info_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
