package profile

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/profile/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_profile_config.Request
    response lambda_profile_config.Response
    err    error
  }{
    {
      request: lambda_profile_config.Request {
        Actor: "0xLambdaProfileActor001",
        UserType: "KOL",
        Username: "LambdaProfileActor001",
        PhotoUrl: "http://123.com",
        TelegramId: "TelegramIdLambdaProfileActor001",
        PhoneNumber: "5197290001",
      },
      response: lambda_profile_config.Response {
          Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_profile_config.Request {
        Actor: "0xLambdaProfileActor002",
        UserType: "KOL",
        Username: "LambdaProfileActor002",
        PhotoUrl: "http://567.com",
        TelegramId: "TelegramIdLambdaProfileActor002",
        PhoneNumber: "5197290002",
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
