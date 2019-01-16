package add_tracked_wallet_addresses_test

import (
	"BigBang/cmd/lambda/feed/add_tracked_wallet_addresses/config"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_add_tracked_wallet_addresses_config.Request
		response lambda_add_tracked_wallet_addresses_config.Response
		err      error
	}{
		{
			request: lambda_add_tracked_wallet_addresses_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_add_tracked_wallet_addresses_config.RequestContent{
					Actor: test_constants.Actor1,
					WalletAddressList: []string{
						test_constants.WalletAddress1,
						test_constants.WalletAddress2,
					},
				},
			},
			response: lambda_add_tracked_wallet_addresses_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_add_tracked_wallet_addresses_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_add_tracked_wallet_addresses_config.RequestContent{
					Actor: test_constants.Actor1,
					WalletAddressList: []string{
						test_constants.WalletAddress1,
						test_constants.WalletAddress2,
					},
				},
			},
			response: lambda_add_tracked_wallet_addresses_config.Response{
				Ok: false,
				Message: &error_config.ErrorInfo{
					ErrorCode: error_config.WalletAddressAlreadyExisting,
					ErrorData: error_config.ErrorData{
						"actor":         test_constants.Actor1,
						"walletAddress": test_constants.WalletAddress1,
					},
					ErrorLocation: error_config.WalletAddressRecordLocation,
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_add_tracked_wallet_addresses_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
