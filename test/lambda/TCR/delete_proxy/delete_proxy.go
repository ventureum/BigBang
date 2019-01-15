package delete_proxy_test

import (
	"BigBang/cmd/lambda/TCR/delete_proxy/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_delete_proxy_config.Request
		response lambda_delete_proxy_config.Response
		err      error
	}{
		{
			request: lambda_delete_proxy_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_delete_proxy_config.RequestContent{
					Proxy: test_constants.Actor7,
				},
			},
			response: lambda_delete_proxy_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_delete_proxy_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
