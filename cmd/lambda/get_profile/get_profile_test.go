package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Actor: "0xLambdaProfileActor001",
      },
      response: Response {
        Profile: &ResponseContent{
          Actor: "0xLambdaProfileActor001",
          ActorType: "KOL",
          Username: "LambdaProfileActor001",
          PhotoUrl: "http://123.com",
          TelegramId: "TelegramIdLambdaProfileActor001",
          PhoneNumber: "5197290001",
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
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
