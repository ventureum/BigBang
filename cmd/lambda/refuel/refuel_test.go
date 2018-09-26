package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/pkg/error_config"
  "os"
)

func TestHandlerWithoutDebugMode(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Actor: "0xLambdaProfileActor001",
      },
      response: Response {
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
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Message.ErrorCode, result.Message.ErrorCode)
  }
}


func TestHandlerWithDebugMode(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Actor: "0xLambdaProfileActor001",
      },
      response: Response {
        Ok: true,
        RefuelAmount: 0,
      },
      err: nil,
    },
  }

  os.Setenv("DEBUG_MODE", "1")
  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
  os.Setenv("DEBUG_MODE", "0")
}
