package set_token_pool

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/cmd/lambda/feed/get_redeem_block_info/config"
  "BigBang/internal/pkg/error_config"
)


var NextRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(1)
var ExecutedAt = NextRedeemBlock.ConvertToTime()
var NextNRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(10000)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_redeem_block_info_config.Request
    response lambda_get_redeem_block_info_config.Response
    err    error
  }{
    {
      request: lambda_get_redeem_block_info_config.Request {
        RedeemBlockTimestamp: NextRedeemBlock.ConvertToTime().Unix(),
      },
      response: lambda_get_redeem_block_info_config.Response {
        RedeemBlockInfo: &feed_attributes.RedeemBlockInfo {
          RedeemBlock:                  NextRedeemBlock,
          TotalEnrolledMilestonePoints: 400,
          TokenPool:                    10000,
          ExecutedAt:                   ExecutedAt,
        },
        Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_get_redeem_block_info_config.Request {
        RedeemBlockTimestamp: NextNRedeemBlock.ConvertToTime().Unix(),
      },
      response: lambda_get_redeem_block_info_config.Response {
        Ok: false,
        Message: &error_config.ErrorInfo{
          ErrorCode: error_config.NoReDeemBlockInfoRecordExisting,
          ErrorData: map[string]interface{} {
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
