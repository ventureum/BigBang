package get_batch_proxy_voting_info_test

import (
	"BigBang/cmd/lambda/TCR/get_batch_proxy_voting_info/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_batch_proxy_voting_info_config.Request
		response lambda_get_batch_proxy_voting_info_config.Response
		err      error
	}{
		{
			request: lambda_get_batch_proxy_voting_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_batch_proxy_voting_info_config.RequestContent{
					ProxyVotingInfoKeyList: []tcr_attributes.ProxyVotingInfoKey{
						{
							Actor:     test_constants.Actor1,
							ProjectId: test_constants.ProjectId1,
						},
					},
				},
			},
			response: lambda_get_batch_proxy_voting_info_config.Response{
				ProxyVotingInfoList: &[]tcr_attributes.ProxyVotingInfo{
					{
						Actor:                  test_constants.Actor1,
						ProjectId:              test_constants.ProjectId1,
						AvailableDelegateVotes: 50,
						ReceivedDelegateVotes:  60,
						ProxyVotingList: &[]tcr_attributes.ProxyVoting{
							{
								Proxy:          test_constants.Actor7,
								BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
								VotesInPercent: 10,
							},
							{
								Proxy:          test_constants.Actor6,
								BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
								VotesInPercent: 10,
							},
							{
								Proxy:          test_constants.Actor5,
								BlockTimestamp: test_constants.ProxyVotesBlockTimestamp4,
								VotesInPercent: 20,
							},
							{
								Proxy:          test_constants.Actor4,
								BlockTimestamp: test_constants.ProxyVotesBlockTimestamp3,
								VotesInPercent: 20,
							},
							{
								Proxy:          test_constants.Actor3,
								BlockTimestamp: test_constants.ProxyVotesBlockTimestamp2,
								VotesInPercent: 20,
							},
							{
								Proxy:          test_constants.Actor2,
								BlockTimestamp: test_constants.ProxyVotesBlockTimestamp1,
								VotesInPercent: 20,
							},
						},
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_batch_proxy_voting_info_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
