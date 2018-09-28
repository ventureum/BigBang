package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
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
        UserType: "KOL",
        Username: "LambdaProfileActor001",
        PhotoUrl: "http://123.com",
        TelegramId: "TelegramIdLambdaProfileActor001",
        PhoneNumber: "5197290001",
      },
      response: Response {
          Ok: true,
      },
       err: nil,
    },
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        UserType: "KOL",
        Username: "LambdaProfileActor002",
        PhotoUrl: "http://567.com",
        TelegramId: "TelegramIdLambdaProfileActor002",
        PhoneNumber: "5197290002",
      },
      response: Response {
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
