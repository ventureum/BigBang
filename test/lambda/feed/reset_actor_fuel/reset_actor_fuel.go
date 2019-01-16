package reset_actor_fuel_test

import (
	"BigBang/cmd/lambda/feed/reset_actor_fuel/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		response lambda_reset_actor_fuel_config.Response
		err      error
	}{
		{
			response: lambda_reset_actor_fuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_reset_actor_fuel_config.Handler()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
