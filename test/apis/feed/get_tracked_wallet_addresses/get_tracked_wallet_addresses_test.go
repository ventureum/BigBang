package get_tracked_wallet_addresses

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/get_tracked_wallet_addresses/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_tracked_wallet_addresses_config.Request
    response lambda_get_tracked_wallet_addresses_config.Response
    err    error
  }{
    {
      request: lambda_get_tracked_wallet_addresses_config.Request {
        Actor:          test_constants.Actor1,
      },
      response: lambda_get_tracked_wallet_addresses_config.Response {
        Ok: true,
        WalletAddressList: &[]string {
          test_constants.WalletAddress1,
          test_constants.WalletAddress2,
        },
      },
      err: nil,
    },
  }

  for _, test := range tests {
    result, err := lambda_get_tracked_wallet_addresses_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
