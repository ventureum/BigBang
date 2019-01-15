package tcr_table_creator_test

import (
	"BigBang/cmd/lambda/TCR/tcr_table_creator/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_tcr_table_creator_config.Request
		response lambda_tcr_table_creator_config.Response
		err      error
	}{
		{
			request: lambda_tcr_table_creator_config.Request{
				DBInfo: nil,
			},
			response: lambda_tcr_table_creator_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_tcr_table_creator_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
	}
}
