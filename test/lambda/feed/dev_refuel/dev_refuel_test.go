package dev_refuel

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/dev_refuel/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_dev_refuel_config.Request
    response lambda_dev_refuel_config.Response
    err    error
  }{
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor1,
        Fuel: 100,
        Reputation: 100,
        MilestonePoints: 100,
      },
      response: lambda_dev_refuel_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  for _, test := range tests {
    result, err := lambda_dev_refuel_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}