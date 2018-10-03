package TCR

import (
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/api"
  "github.com/stretchr/testify/assert"
  profileConfig "BigBang/cmd/lambda/feed/profile/config"
  "github.com/mitchellh/mapstructure"
  getProfileConfig "BigBang/cmd/lambda/feed/get_profile/config"
)

func TestProfileAndGetProfile(t *testing.T) {
  requestProfile := profileConfig.Request{
    Actor: Actor001,
    UserType: string(feed_attributes.USER_ACTOR_TYPE),
    Username: UserName001,
    PhotoUrl: PhotoUrl001,
    TelegramId: TelegramId001,
    PhoneNumber: PhoneNumber001,
  }

  expectedResponseProfile := profileConfig.Response{
    Ok: true,
  }

  responseMessageProfile := api.SendPost(requestProfile, api.ProfileAlphaEndingPoint)

  var responseProfile profileConfig.Response
  mapstructure.Decode(*responseMessageProfile, &responseProfile)
  assert.Equal(t, expectedResponseProfile, responseProfile)

  requestGetProfile := getProfileConfig.Request{
    Actor: Actor001,
  }

  expectedResponseGetProfile :=  getProfileConfig.Response{
    Ok: true,
    Profile: &getProfileConfig.ResponseContent{
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
  var responseGetProfile getProfileConfig.Response
  mapstructure.Decode(*responseMessageGetProfile, &responseGetProfile)
  assert.Equal(t, responseGetProfile, expectedResponseGetProfile)
}
