package refuel

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/pkg/error_config"
  "BigBang/cmd/lambda/feed/refuel/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandlerWithoutDebugMode(t *testing.T) {
  tests := []struct{
    request lambda_refuel_config.Request
    response lambda_refuel_config.Response
    err    error
  }{
    {
      request: lambda_refuel_config.Request {
        Actor: test_constants.Actor1,
      },
      response: lambda_refuel_config.Response {
          Message: &error_config.ErrorInfo{
            ErrorCode: "InsufficientWaitingTimeToRefuel",
            ErrorData: error_config.ErrorData{
              "lastRefuelTimestamp": 100,
            },
          },
          Ok: false,
      },
       err: nil,
    },
  }
  for _, test := range tests {
    responseMessageRefuel := api.SendPost(test.request, api.RefuelAlphaEndingPoint)
    var responseRefuel lambda_refuel_config.Response
    mapstructure.Decode(*responseMessageRefuel, &responseRefuel)
    assert.Equal(t, test.response.Ok, responseRefuel.Ok)
  }
}
