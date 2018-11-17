package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/cmd/lambda/feed/get_profile/config"
  "BigBang/test/constants"
  "strings"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_profile_config.Request
    response lambda_get_profile_config.Response
    err    error
  }{
    {
      request: lambda_get_profile_config.Request {
        Actor: test_constants.Actor1,
      },
      response: lambda_get_profile_config.Response {
        Profile: &lambda_get_profile_config.ResponseContent{
          Actor: test_constants.Actor1,
          ActorType: "KOL",
          Username:test_constants.UserName1,
          PhotoUrl: "http://123.com",
          TelegramId: test_constants.TelegramId1,
          PhoneNumber: "5197290001",
          PublicKey: strings.ToLower(test_constants.PublicKey1),
          Level: 2,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 100,
            Reputation: 100,
            MilestonePoints: 0,
          },
        },
        Ok: true,
      },
       err: nil,
    },
  }
  for _, test := range tests {
    responseMessageGetProfile := api.SendPost(test.request, api.GetProfileAlphaEndingPoint)
    var responseGetProfile lambda_get_profile_config.Response
    mapstructure.Decode(*responseMessageGetProfile, &responseGetProfile)
    assert.Equal(t, test.response, responseGetProfile)
  }
}
