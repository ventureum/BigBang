package adjust_proxy_votes

import (
  "BigBang/test/constants"
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
  "BigBang/cmd/lambda/TCR/get_proxy_votes_list_by_principal/config"
)

var EmptyProxyVotesList []tcr_attributes.ProxyVotes

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_proxy_votes_list_by_principal_config.Request
    response lambda_get_proxy_votes_list_by_principal_config.Response
    err    error
  }{
    {
      request: lambda_get_proxy_votes_list_by_principal_config.Request {
        Actor: test_constants.Actor1,
        ProjectId: test_constants.ProjectId1,
        Limit: 0,
      },
      response: lambda_get_proxy_votes_list_by_principal_config.Response {
        ProxyVotesInfo: &tcr_attributes.ProxyVotesInfo{
          Actor: test_constants.Actor1,
          ProjectId: test_constants.ProjectId1,
          ProxyVotesList:  &EmptyProxyVotesList,
        },
        NextCursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
          test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor7, test_constants.ProxyVotesBlockTimestamp5),
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_proxy_votes_list_by_principal_config.Request {
        Actor: test_constants.Actor1,
        ProjectId: test_constants.ProjectId1,
        Limit: 2,
      },
      response: lambda_get_proxy_votes_list_by_principal_config.Response {
        ProxyVotesInfo: &tcr_attributes.ProxyVotesInfo{
          Actor: test_constants.Actor1,
          ProjectId:   test_constants.ProjectId1,
          ProxyVotesList: &[]tcr_attributes.ProxyVotes{
            {
              Proxy:          test_constants.Actor7,
              BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
              VotesInPercent: 20,
            },
            {
              Proxy:          test_constants.Actor6,
              BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
              VotesInPercent: 20,
            },
          },
        },
        NextCursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
          test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor5, test_constants.ProxyVotesBlockTimestamp4),
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_proxy_votes_list_by_principal_config.Request {
        Actor: test_constants.Actor1,
        ProjectId: test_constants.ProjectId1,
        Limit: 2,
        Cursor:  principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
          test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor5, test_constants.ProxyVotesBlockTimestamp4),
      },
      response: lambda_get_proxy_votes_list_by_principal_config.Response {
        ProxyVotesInfo: &tcr_attributes.ProxyVotesInfo{
          Actor: test_constants.Actor1,
          ProjectId:  test_constants.ProjectId1,
          ProxyVotesList: &[]tcr_attributes.ProxyVotes{
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
        NextCursor:  principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
          test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor3, test_constants.ProxyVotesBlockTimestamp2),
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_proxy_votes_list_by_principal_config.Request {
        Actor: test_constants.Actor1,
        ProjectId: test_constants.ProjectId1,
        Limit: 5,
        Cursor: principal_proxy_votes_config.GenerateEncodedPrincipalProxyVotesRecordID(
          test_constants.Actor1, test_constants.ProjectId1, test_constants.Actor5, test_constants.ProxyVotesBlockTimestamp4),
      },
      response: lambda_get_proxy_votes_list_by_principal_config.Response {
        ProxyVotesInfo: &tcr_attributes.ProxyVotesInfo{
          Actor: test_constants.Actor1,
          ProjectId:   test_constants.ProjectId1,
          ProxyVotesList: &[]tcr_attributes.ProxyVotes{
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
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_proxy_votes_list_by_principal_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
