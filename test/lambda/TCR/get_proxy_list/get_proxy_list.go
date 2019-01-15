package get_proxy_list_test

import (
	"BigBang/cmd/lambda/TCR/get_proxy_list/config"
	"BigBang/internal/pkg/utils"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

var EmptyStringLIst []string

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_proxy_list_config.Request
		response lambda_get_proxy_list_config.Response
		err      error
	}{
		{
			request: lambda_get_proxy_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_list_config.RequestContent{
					Limit: 0,
				},
			},
			response: lambda_get_proxy_list_config.Response{
				ResponseData: &lambda_get_proxy_list_config.ResponseData{
					Proxies:    &EmptyStringLIst,
					NextCursor: utils.Base64EncodeInt64(7),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_proxy_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_list_config.RequestContent{
					Limit: 2,
				},
			},
			response: lambda_get_proxy_list_config.Response{
				ResponseData: &lambda_get_proxy_list_config.ResponseData{
					Proxies: &[]string{
						test_constants.Actor7,
						test_constants.Actor6,
					},
					NextCursor: utils.Base64EncodeInt64(5),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_proxy_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_list_config.RequestContent{
					Limit:  2,
					Cursor: utils.Base64EncodeInt64(4),
				},
			},
			response: lambda_get_proxy_list_config.Response{
				ResponseData: &lambda_get_proxy_list_config.ResponseData{
					Proxies: &[]string{
						test_constants.Actor4,
						test_constants.Actor3,
					},
					NextCursor: utils.Base64EncodeInt64(2),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_proxy_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_list_config.RequestContent{
					Limit:  5,
					Cursor: utils.Base64EncodeInt64(4),
				},
			},
			response: lambda_get_proxy_list_config.Response{
				ResponseData: &lambda_get_proxy_list_config.ResponseData{
					Proxies: &[]string{
						test_constants.Actor4,
						test_constants.Actor3,
						test_constants.Actor2,
						test_constants.Actor1,
					},
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_get_proxy_list_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
