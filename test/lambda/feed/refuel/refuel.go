package refuel_test

import (
	"BigBang/cmd/lambda/feed/refuel/config"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandlerWithoutDebugMode(t *testing.T) {
	tests := []struct {
		request  lambda_refuel_config.Request
		response lambda_refuel_config.Response
		err      error
	}{
		{
			request: lambda_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_refuel_config.RequestContent{
					Actor: test_constants.Actor1,
				},
			},
			response: lambda_refuel_config.Response{
				Message: &error_config.ErrorInfo{
					ErrorCode: "InsufficientWaitingTimeToRefuel",
					ErrorData: error_config.ErrorData{
						"lastRefuelTimestamp": 100,
					},
				},
				Ok: false,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_refuel_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Message.ErrorCode, result.Message.ErrorCode)
	}
}

func TestHandlerWithDebugMode(t *testing.T) {
	tests := []struct {
		request  lambda_refuel_config.Request
		response lambda_refuel_config.Response
		err      error
	}{
		{
			request: lambda_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_refuel_config.RequestContent{
					Actor: test_constants.Actor1,
				},
			},
			response: lambda_refuel_config.Response{
				Ok:           true,
				RefuelAmount: 0,
			},
			err: nil,
		},
	}

	os.Setenv("DEBUG_MODE", "1")
	for _, test := range tests {
		result, err := lambda_refuel_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
	os.Setenv("DEBUG_MODE", "0")
}
