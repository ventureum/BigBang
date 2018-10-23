package profile

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/profile/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_profile_config.Request
    response lambda_profile_config.Response
    err    error
  }{
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor1,
        UserType: "KOL",
        Username: test_constants.UserName1,
        PhotoUrl: "http://123.com",
        TelegramId: test_constants.TelegramId1,
        PhoneNumber: "5197290001",
        PrivateKey: test_constants.PrivateKey1,
      },
      response: lambda_profile_config.Response {
          Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor2,
        UserType: "KOL",
        Username: test_constants.UserName2,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId2,
        PhoneNumber: "5197290002",
        PrivateKey: test_constants.PrivateKey2,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor3,
        UserType: "KOL",
        Username: test_constants.UserName3,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId3,
        PhoneNumber: "5197290002",
        PrivateKey: test_constants.PrivateKey3,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor4,
        UserType: "KOL",
        Username: test_constants.UserName4,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId4,
        PhoneNumber: "5197290002",
        PrivateKey: test_constants.PrivateKey4,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor5,
        UserType: "KOL",
        Username: test_constants.UserName5,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId5,
        PhoneNumber: "5197290002",
        PrivateKey: test_constants.PrivateKey5,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor6,
        UserType: "KOL",
        Username: test_constants.UserName6,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId6,
        PhoneNumber: "5197290002",
        PrivateKey: test_constants.PrivateKey6,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor7,
        UserType: "KOL",
        Username: test_constants.UserName7,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId7,
        PhoneNumber: "5197290002",
        PrivateKey: test_constants.PrivateKey7,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_profile_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}