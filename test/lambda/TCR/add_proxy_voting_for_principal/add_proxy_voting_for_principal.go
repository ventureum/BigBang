package add_proxy_voting_for_principal_test

import (
	"BigBang/cmd/lambda/TCR/add_proxy_voting_for_principal/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_add_proxy_voting_for_principal_config.Request
		response lambda_add_proxy_voting_for_principal_config.Response
		err      error
	}{
		{
			request: lambda_add_proxy_voting_for_principal_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_add_proxy_voting_for_principal_config.RequestContent{
					Actor:     test_constants.Actor1,
					ProjectId: test_constants.ProjectId1,
					ProxyVotingList: []tcr_attributes.ProxyVoting{
						{
							Proxy:          test_constants.Actor2,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp1,
							VotesInPercent: 20,
						},
						{
							Proxy:          test_constants.Actor3,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp2,
							VotesInPercent: 20,
						},
						{
							Proxy:          test_constants.Actor4,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp3,
							VotesInPercent: 20,
						},
						{
							Proxy:          test_constants.Actor5,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp4,
							VotesInPercent: 20,
						},
						{
							Proxy:          test_constants.Actor6,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
							VotesInPercent: 10,
						},
						{
							Proxy:          test_constants.Actor7,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
							VotesInPercent: 10,
						},
					},
				},
			},
			response: lambda_add_proxy_voting_for_principal_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_add_proxy_voting_for_principal_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_add_proxy_voting_for_principal_config.RequestContent{
					Actor:     test_constants.Actor6,
					ProjectId: test_constants.ProjectId1,
					ProxyVotingList: []tcr_attributes.ProxyVoting{
						{
							Proxy:          test_constants.Actor2,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp1,
							VotesInPercent: 20,
						},
					},
				},
			},
			response: lambda_add_proxy_voting_for_principal_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_add_proxy_voting_for_principal_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_add_proxy_voting_for_principal_config.RequestContent{
					Actor:     test_constants.Actor7,
					ProjectId: test_constants.ProjectId1,
					ProxyVotingList: []tcr_attributes.ProxyVoting{
						{
							Proxy:          test_constants.Actor2,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp2,
							VotesInPercent: 20,
						},
					},
				},
			},
			response: lambda_add_proxy_voting_for_principal_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_add_proxy_voting_for_principal_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_add_proxy_voting_for_principal_config.RequestContent{
					Actor:     test_constants.Actor7,
					ProjectId: test_constants.ProjectId1,
					ProxyVotingList: []tcr_attributes.ProxyVoting{
						{
							Proxy:          test_constants.Actor2,
							BlockTimestamp: test_constants.ProxyVotesBlockTimestamp3,
							VotesInPercent: 0,
						},
					},
				},
			},
			response: lambda_add_proxy_voting_for_principal_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_add_proxy_voting_for_principal_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
