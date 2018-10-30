package get_actor_uuid_from_public_key

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/get_actor_uuid_from_public_key/config"
  "BigBang/internal/pkg/error_config"
  "strings"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_actor_uuid_from_public_key_config.Request
    response lambda_get_actor_uuid_from_public_key_config.Response
    err    error
  }{
    {
      request: lambda_get_actor_uuid_from_public_key_config.Request {
        PublicKey: test_constants.PublicKey1,
      },
      response: lambda_get_actor_uuid_from_public_key_config.Response {
        Actor: test_constants.Actor1,
        Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_get_actor_uuid_from_public_key_config.Request {
        PublicKey: "0xInvalidPublicKey",
      },
      response: lambda_get_actor_uuid_from_public_key_config.Response {
        Actor: "",
        Ok: false,
        Message: &error_config.ErrorInfo{
          ErrorCode: error_config.NoActorExistingForPublicKey,
          ErrorData: map[string]interface{} {
            "publicKey": strings.ToLower("0xInvalidPublicKey"),
          },
          ErrorLocation: error_config.ProfileAccountLocation,
        },
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_actor_uuid_from_public_key_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
