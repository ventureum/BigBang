package main

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
        FeedSlug: "User",
        UserId: "david3620",
      },
      response: Response {
        FeedToken: "8XDj7VcxGoOYMYigN_bIT7h9hAo",
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
