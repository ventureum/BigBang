package add_proxy

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/TCR/add_proxy/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_add_proxy_config.Request
    response lambda_add_proxy_config.Response
    err    error
  }{
    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor1,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor2,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor3,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor4,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor5,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor6,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_add_proxy_config.Request {
        Proxy: test_constants.Actor7,
      },
      response: lambda_add_proxy_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_add_proxy_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}

