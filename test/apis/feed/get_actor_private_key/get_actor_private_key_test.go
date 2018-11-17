package get_actor_private_key

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/get_actor_private_key/config"
  "strings"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
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
        PrivateKey: strings.ToLower(test_constants.PrivateKey1),
        Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_get_actor_private_key_config.Request {
        Actor: test_constants.Actor6,
      },
      response: lambda_get_actor_private_key_config.Response {
        Ok: false,
        Message: &error_config.ErrorInfo{
          ErrorCode: error_config.NoPrivateKeyExistingForActor,
          ErrorData: map[string]interface{} {
           "actor": test_constants.Actor6,
          },
          ErrorLocation: error_config.ProfileAccountLocation,
        },
      },
      err: nil,
    },
  }
  for _, test := range tests {
    responseMessageGetActorPrivateKey := api.SendPost(test.request, api.GetActorPrivateKeyAlphaEndingPoint)
    var responseGetActorPrivateKey lambda_get_actor_private_key_config.Response
    mapstructure.Decode(*responseMessageGetActorPrivateKey , &responseGetActorPrivateKey)
    assert.Equal(t, test.response, responseGetActorPrivateKey)
  }
}
