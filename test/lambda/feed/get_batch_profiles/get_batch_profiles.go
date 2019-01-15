package get_batch_profiles_test

import (
	"BigBang/cmd/lambda/feed/get_batch_profiles/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_batch_profiles_config.Request
		response lambda_get_batch_profiles_config.Response
		err      error
	}{
		{
			request: lambda_get_batch_profiles_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_get_batch_profiles_config.RequestContent{
					Actors: []string{test_constants.Actor3},
				},
			},
			response: lambda_get_batch_profiles_config.Response{
				Profiles: &[]lambda_get_batch_profiles_config.ResponseContent{
					{
						Actor:          test_constants.Actor3,
						ActorType:      "KOL",
						Username:       test_constants.UserName3,
						PhotoUrl:       "http://567.com",
						PublicKey:      strings.ToLower(test_constants.PublicKey3),
						Level:          2,
						ProfileContent: test_constants.ProfileContent3,
						RewardsInfo: &feed_attributes.RewardsInfo{
							Fuel:            100,
							Reputation:      100,
							MilestonePoints: 0,
						},
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_batch_profiles_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
