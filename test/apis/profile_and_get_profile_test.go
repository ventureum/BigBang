package apis

import (
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/api"
  "github.com/stretchr/testify/assert"
)

func TestProfileAndGetProfile(t *testing.T) {
  messageProfile := api.Message(map[string]interface{}{
    "actor": Actor001,
    "userType": feed_attributes.USER_ACTOR_TYPE,
    "username": UserName001,
    "photoUrl": PhotoUrl001,
    "telegramId": TelegramId001,
    "phoneNumber": PhoneNumber001,
  })

  expectedResponseProfile :=  api.Message(map[string]interface{}{
    "ok": true,
  })

  responseProfile := api.SendPost(messageProfile, api.ProfileAlphaEndingPoint)
  assert.Equal(t, responseProfile, &expectedResponseProfile)

  messageGetProfile := api.Message(map[string]interface{}{
    "actor": Actor001,
  })

  expectedResponseGetProfile :=  api.Message(map[string]interface{}{
    "ok": true,
    "profile": map[string]interface{}{
      "actor": Actor001,
      "actorType": string(feed_attributes.USER_ACTOR_TYPE),
      "username": UserName001,
      "photoUrl": PhotoUrl001,
      "telegramId": TelegramId001,
      "phoneNumber": PhoneNumber001,
      "level": float64(2),
      "rewardsInfo": map[string]interface{}{
        "fuel": float64(100),
        "reputation": float64(100),
        "milestonePoints": float64(0),
      },
    },
  })

  responseGetProfile := api.SendPost(messageGetProfile, api.GetProfileAlphaEndingPoint)
  assert.Equal(t, responseGetProfile, &expectedResponseGetProfile)
}
