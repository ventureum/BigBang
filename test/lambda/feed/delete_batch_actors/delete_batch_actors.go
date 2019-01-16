package delete_batch_actors

import (
	"BigBang/cmd/lambda/feed/delete_batch_actors/config"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_delete_batch_actors_config.Request
		response lambda_delete_batch_actors_config.Response
		err      error
	}{
		{
			request: lambda_delete_batch_actors_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_delete_batch_actors_config.RequestContent{
					ActorList: []string{
						test_constants.Actor3,
						test_constants.Actor4,
					},
				},
			},
			response: lambda_delete_batch_actors_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_delete_batch_actors_config.Request{
				PrincipalId: test_constants.Actor5,
				Body: lambda_delete_batch_actors_config.RequestContent{
					ActorList: []string{
						test_constants.Actor6,
					},
				},
			},
			response: lambda_delete_batch_actors_config.Response{
				Ok: false,
				Message: &error_config.ErrorInfo{
					ErrorCode: error_config.InvalidAuthAccess,
					ErrorData: error_config.ErrorData{
						"principalId": test_constants.Actor5,
					},
					ErrorLocation: error_config.Auth,
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_delete_batch_actors_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
