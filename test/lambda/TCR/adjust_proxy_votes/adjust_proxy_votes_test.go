package adjust_proxy_votes

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/TCR/adjust_proxy_votes/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_adjust_proxy_votes_config.Request
    response lambda_adjust_proxy_votes_config.Response
    err    error
  }{
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor1,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor2,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp1,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor1,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor3,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp2,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor1,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor4,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp3,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor1,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor5,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp4,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor1,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor6,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor1,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor7,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp5,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor6,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor2,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp1,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor7,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor2,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp2,
        VotesInPercent: 20,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_adjust_proxy_votes_config.Request {
        Actor:          test_constants.Actor7,
        ProjectId:      test_constants.ProjectId1,
        Proxy:          test_constants.Actor2,
        BlockTimestamp: test_constants.ProxyVotesBlockTimestamp3,
        VotesInPercent: 0,
      },
      response: lambda_adjust_proxy_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  for _, test := range tests {
    result, err := lambda_adjust_proxy_votes_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
