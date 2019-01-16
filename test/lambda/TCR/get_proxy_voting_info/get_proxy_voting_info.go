package get_proxy_voting_info_test

import (
	"BigBang/cmd/lambda/TCR/get_proxy_voting_info/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

var EmptyProxyVotesList []tcr_attributes.ProxyVoting

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_proxy_voting_info_config.Request
		response lambda_get_proxy_voting_info_config.Response
		err      error
	}{
		{
			request: lambda_get_proxy_voting_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_voting_info_config.RequestContent{
					Actor:     test_constants.Actor1,
					ProjectId: test_constants.ProjectId1,
					Limit:     0,
				},
			},
			response: lambda_get_proxy_voting_info_config.Response{
				ResponseData: &lambda_get_proxy_voting_info_config.ResponseData{
					ProxyVotingInfo: &tcr_attributes.ProxyVotingInfo{
						Actor:                  test_constants.Actor1,
						ProjectId:              test_constants.ProjectId1,
						AvailableDelegateVotes: 50,
						ReceivedDelegateVotes:  60,
						ProxyVotingList:        &EmptyProxyVotesList,
					},
					NextCursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
						test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor7, test_constants.ProxyVotesBlockTimestamp5),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_proxy_voting_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_voting_info_config.RequestContent{
					Actor:     test_constants.Actor1,
					ProjectId: test_constants.ProjectId1,
					Limit:     2,
				},
			},
			response: lambda_get_proxy_voting_info_config.Response{
				ResponseData: &lambda_get_proxy_voting_info_config.ResponseData{
					ProxyVotingInfo: &tcr_attributes.ProxyVotingInfo{
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
						},
					},
					NextCursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
						test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor5, test_constants.ProxyVotesBlockTimestamp4),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_proxy_voting_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_voting_info_config.RequestContent{
					Actor:     test_constants.Actor1,
					ProjectId: test_constants.ProjectId1,
					Limit:     2,
					Cursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
						test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor5, test_constants.ProxyVotesBlockTimestamp4),
				},
			},
			response: lambda_get_proxy_voting_info_config.Response{
				ResponseData: &lambda_get_proxy_voting_info_config.ResponseData{
					ProxyVotingInfo: &tcr_attributes.ProxyVotingInfo{
						Actor:                  test_constants.Actor1,
						ProjectId:              test_constants.ProjectId1,
						AvailableDelegateVotes: 50,
						ReceivedDelegateVotes:  60,
						ProxyVotingList: &[]tcr_attributes.ProxyVoting{
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
						},
					},
					NextCursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
						test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor3, test_constants.ProxyVotesBlockTimestamp2),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_proxy_voting_info_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_proxy_voting_info_config.RequestContent{
					Actor:     test_constants.Actor1,
					ProjectId: test_constants.ProjectId1,
					Limit:     5,
					Cursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
						test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor5, test_constants.ProxyVotesBlockTimestamp4),
				},
			},
			response: lambda_get_proxy_voting_info_config.Response{
				ResponseData: &lambda_get_proxy_voting_info_config.ResponseData{
					ProxyVotingInfo: &tcr_attributes.ProxyVotingInfo{
						Actor:                  test_constants.Actor1,
						ProjectId:              test_constants.ProjectId1,
						AvailableDelegateVotes: 50,
						ReceivedDelegateVotes:  60,
						ProxyVotingList: &[]tcr_attributes.ProxyVoting{
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
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_proxy_voting_info_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
