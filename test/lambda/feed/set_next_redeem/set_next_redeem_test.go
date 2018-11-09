package set_next_redeem

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/set_next_redeem/config"
  "BigBang/internal/pkg/error_config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_set_next_redeem_config.Request
    response lambda_set_next_redeem_config.Response
    err    error
  }{
    {
      request: lambda_set_next_redeem_config.Request {
        Actor: test_constants.Actor1,
        MilestonePoints: test_constants.RedeemMiletonePointsRegular1,
      },
      response: lambda_set_next_redeem_config.Response {
          Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_set_next_redeem_config.Request {
        Actor: test_constants.Actor1,
        MilestonePoints: test_constants.RedeemMiletonePointsRegular2,
      },
      response: lambda_set_next_redeem_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_next_redeem_config.Request {
        Actor: test_constants.Actor2,
        MilestonePoints: test_constants.RedeemMiletonePointsMax,
      },
      response: lambda_set_next_redeem_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_next_redeem_config.Request {
        Actor: test_constants.Actor3,
        MilestonePoints: test_constants.RedeemMiletonePointsZero,
      },
      response: lambda_set_next_redeem_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_next_redeem_config.Request {
        Actor: test_constants.Actor3,
        MilestonePoints: test_constants.RedeemMiletonePointsNegative,
      },
      response: lambda_set_next_redeem_config.Response {
        Ok: false,
        Message: &error_config.ErrorInfo{
          ErrorCode: error_config.InvalidMilestonePoints,
          ErrorData: map[string]interface{} {
            "actor": test_constants.Actor3,
            "milestonePoints": float64(test_constants.RedeemMiletonePointsNegative),
          },
          ErrorLocation: error_config.MilestonePointsRedeemRequestRecordLocation,
        },
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_set_next_redeem_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}