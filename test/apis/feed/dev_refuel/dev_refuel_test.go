package dev_refuel

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/dev_refuel/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
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
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor2,
        Fuel: 100,
        Reputation: 100,
        MilestonePoints: 100,
      },
      response: lambda_dev_refuel_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor3,
        Fuel: 100,
        Reputation: 100,
        MilestonePoints: 100,
      },
      response: lambda_dev_refuel_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor4,
        Fuel: 100,
        Reputation: 100,
        MilestonePoints: 100,
      },
      response: lambda_dev_refuel_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor5,
        Fuel: 100,
        Reputation: 100,
        MilestonePoints: 100,
      },
      response: lambda_dev_refuel_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor6,
        Fuel: 100,
        Reputation: 100,
        MilestonePoints: 100,
      },
      response: lambda_dev_refuel_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_dev_refuel_config.Request {
        Actor: test_constants.Actor7,
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
    responseMessageDevRefuel := api.SendPost(test.request, api.DevRefuelAlphaEndingPoint)
    var responseDevRefuel lambda_dev_refuel_config.Response
    mapstructure.Decode(*responseMessageDevRefuel, &responseDevRefuel)
    assert.Equal(t, test.response, responseDevRefuel)
  }
}
