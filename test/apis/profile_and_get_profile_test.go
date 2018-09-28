package apis

import (
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/api"
  "github.com/stretchr/testify/assert"
  "BigBang/cmd/lambda/profile/config"
  "github.com/mitchellh/mapstructure"
  config2 "BigBang/cmd/lambda/get_profile/config"
)

func TestProfileAndGetProfile(t *testing.T) {
  requestProfile := config.Request{
    Actor: Actor001,
    UserType: string(feed_attributes.USER_ACTOR_TYPE),
    Username: UserName001,
    PhotoUrl: PhotoUrl001,
    TelegramId: TelegramId001,
    PhoneNumber: PhoneNumber001,
  }

  expectedResponseProfile := config.Response{
    Ok: true,
  }

  responseMessageProfile := api.SendPost(requestProfile, api.ProfileAlphaEndingPoint)

  var responseProfile config.Response
  mapstructure.Decode(*responseMessageProfile, &responseProfile)
  assert.Equal(t, expectedResponseProfile, responseProfile)

  requestGetProfile := config2.Request{
    Actor: Actor001,
  }

  expectedResponseGetProfile :=  config2.Response{
    Ok: true,
    Profile: &config2.ResponseContent{
      Actor:       Actor001,
      ActorType:   string(feed_attributes.USER_ACTOR_TYPE),
      Username:    UserName001,
      PhotoUrl:    PhotoUrl001,
      TelegramId:  TelegramId001,
      PhoneNumber: PhoneNumber001,
      Level:       2,
      RewardsInfo: &feed_attributes.RewardsInfo{
        Fuel:            100,
        Reputation:      100,
        MilestonePoints: 0,
      },
    },
  }

  responseMessageGetProfile := api.SendPost(requestGetProfile, api.GetProfileAlphaEndingPoint)
  var responseGetProfile config2.Response
  mapstructure.Decode(*responseMessageGetProfile, &responseGetProfile)
  assert.Equal(t, responseGetProfile, expectedResponseGetProfile)
}
