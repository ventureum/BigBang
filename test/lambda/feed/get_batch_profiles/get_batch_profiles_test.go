package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/test/constants"
  "strings"
  "BigBang/cmd/lambda/feed/get_batch_profiles/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_batch_profiles_config.Request
    response lambda_get_batch_profiles_config.Response
    err    error
  }{
    {
      request: lambda_get_batch_profiles_config.Request{
        PrincipalId: test_constants.Actor1,
        Body: lambda_get_batch_profiles_config.RequestContent{
          Actors: []string{test_constants.Actor1},
        },
      },
      response: lambda_get_batch_profiles_config.Response{
        Profiles: &[]lambda_get_batch_profiles_config.ResponseContent{
          {
            Actor:       test_constants.Actor1,
            ActorType:   "ADMIN",
            Username:    test_constants.UserName1,
            PhotoUrl:    "http://123.com",
            TelegramId:  test_constants.TelegramId1,
            PhoneNumber: "5197290001",
            PublicKey:   strings.ToLower(test_constants.PublicKey1),
            Level:       2,
            RewardsInfo: &feed_attributes.RewardsInfo{
              Fuel:            100,
              Reputation:      100,
              MilestonePoints: 0,
            },
          },
        },
        Ok: true,
        },
        err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_batch_profiles_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
