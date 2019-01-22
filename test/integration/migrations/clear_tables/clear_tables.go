package clear_tables_test

import (
	"BigBang/cmd/lambda/migrations/clear_tables/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  clear_tables_config.Request
		response clear_tables_config.Response
		err      error
	}{
		{
			request: clear_tables_config.Request{
				DBInfo: nil,
			},
			response: clear_tables_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := clear_tables_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
	}
}
