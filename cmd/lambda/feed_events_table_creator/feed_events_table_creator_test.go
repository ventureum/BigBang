package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    response Response
    err    error
  }{
    {
      response: Response {
          Ok: true,
      },
       err: nil,
    },
  }
  for _, test := range tests {
    result, err := Handler()
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
