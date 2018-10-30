package profile

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/profile/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/error_config"
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
        PublicKey: test_constants.PublicKey1,
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
        PublicKey: test_constants.PublicKey2,
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
        PublicKey: test_constants.PublicKey3,
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
        PublicKey: test_constants.PublicKey4,
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
        PublicKey: test_constants.PublicKey5,
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
        PublicKey: test_constants.PublicKey6,
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
        PublicKey: test_constants.PublicKey7,
      },
      response: lambda_profile_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: test_constants.Actor8,
        UserType: "InvalidUserType",
        Username: test_constants.UserName8,
        PhotoUrl: "http://567.com",
        TelegramId: test_constants.TelegramId8,
        PhoneNumber: "5197290002",
        PublicKey: test_constants.PublicKey8,
      },
      response: lambda_profile_config.Response {
        Ok: false,
        Message: &error_config.ErrorInfo{
           ErrorCode: error_config.InvalidActorType,
           ErrorData: error_config.ErrorData {
              "actorType": "InvalidUserType",
           },
           ErrorLocation: error_config.ActorTypeLocation,
        },
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
