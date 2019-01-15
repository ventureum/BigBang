package deactivate_actor_test

import (
	"BigBang/cmd/lambda/feed/deactivate_actor/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_deactivate_actor_config.Request
		response lambda_deactivate_actor_config.Response
		err      error
	}{
		{
			request: lambda_deactivate_actor_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_deactivate_actor_config.RequestContent{
					Actor: test_constants.Actor2,
				},
			},
			response: lambda_deactivate_actor_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_deactivate_actor_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
