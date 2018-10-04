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
        Proxy: "0xProxy001",
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
