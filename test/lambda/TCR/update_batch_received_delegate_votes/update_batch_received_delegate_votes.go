package update_batch_received_delegate_votes_test

import (
	"BigBang/cmd/lambda/TCR/update_batch_received_delegate_votes/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_update_batch_received_delegate_votes_config.Request
		response lambda_update_batch_received_delegate_votes_config.Response
		err      error
	}{
		{
			request: lambda_update_batch_received_delegate_votes_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_update_batch_received_delegate_votes_config.RequestBody{
					RequestList: []lambda_update_batch_received_delegate_votes_config.RequestContent{
						{
							Actor:                      test_constants.Actor1,
							ProjectId:                  test_constants.ProjectId1,
							ReceivedDelegateVotesDelta: 60,
						},
						{
							Actor:                      test_constants.Actor2,
							ProjectId:                  test_constants.ProjectId1,
							ReceivedDelegateVotesDelta: 60,
						},
					},
				},
			},
			response: lambda_update_batch_received_delegate_votes_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_update_batch_received_delegate_votes_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
	}
}
