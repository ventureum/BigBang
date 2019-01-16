package get_actor_private_key_test

import (
	"BigBang/cmd/lambda/feed/get_actor_private_key/config"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_actor_private_key_config.Request
		response lambda_get_actor_private_key_config.Response
		err      error
	}{
		{
			request: lambda_get_actor_private_key_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_actor_private_key_config.RequestContent{
					Actor: test_constants.Actor1,
				},
			},
			response: lambda_get_actor_private_key_config.Response{
				PrivateKey: strings.ToLower(test_constants.PrivateKey1),
				Ok:         true,
			},
			err: nil,
		},
		{
			request: lambda_get_actor_private_key_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_actor_private_key_config.RequestContent{
					Actor: test_constants.Actor6,
				},
			},
			response: lambda_get_actor_private_key_config.Response{
				Ok: false,
				Message: &error_config.ErrorInfo{
					ErrorCode: error_config.NoPrivateKeyExistingForActor,
					ErrorData: map[string]interface{}{
						"actor": test_constants.Actor6,
					},
					ErrorLocation: error_config.ProfileAccountLocation,
				},
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
