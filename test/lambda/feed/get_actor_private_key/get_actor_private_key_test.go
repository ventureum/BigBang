package get_actor_private_key

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/get_actor_private_key/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_actor_private_key_config.Request
    response lambda_get_actor_private_key_config.Response
    err    error
  }{
    {
      request: lambda_get_actor_private_key_config.Request {
        Actor: test_constants.Actor1,
      },
      response: lambda_get_actor_private_key_config.Response {
        PrivateKey: test_constants.PrivateKey1,
        Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_get_actor_private_key_config.Request {
        Actor: test_constants.Actor6,
      },
      response: lambda_get_actor_private_key_config.Response {
        PrivateKey: "",
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_actor_private_key_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
