package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/error_config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        PostHash: "0xTestPostHash001",
        Value: -1,
      },
      response: Response {
        VoteInfo: &feed_attributes.VoteInfo{
          Actor: "0xLambdaProfileActor002",
          PostHash: "0xTestPostHash001",
          FuelCost: 20,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 80,
            Reputation: 100,
            MilestonePoints: 0,
          },
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        PostHash: "0xTestPostHash001",
        Value: 0,
      },
      response: Response {
        VoteInfo: &feed_attributes.VoteInfo{
          Actor: "0xLambdaProfileActor002",
          PostHash: "0xTestPostHash001",
          FuelCost: 10,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 80,
            Reputation: 100,
            MilestonePoints: 0,
          },
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        PostHash: "0xTestPostHash001",
        Value: 1,
      },
      response: Response {
        VoteInfo: &feed_attributes.VoteInfo{
          Actor: "0xLambdaProfileActor002",
          PostHash: "0xTestPostHash001",
          FuelCost: 10,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 70,
            Reputation: 100,
            MilestonePoints: 0,
          },
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 1,
            TotalVoteCount : 2,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 1,
            TotalVoteCount : 2,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        PostHash: "0xTestPostHash001",
        Value: 1,
      },
      response: Response {
        Message: &error_config.ErrorInfo{
          ErrorCode: "ExceedingUpvoteLimit",
          ErrorData: error_config.ErrorData {
            "actor": "0xLambdaProfileActor002",
            "postHash": "0xTestPostHash001",
          },
          ErrorLocation: "ActorVotesCountersRecordLocation",
        },
        Ok: false,
      },
      err: nil,
    },
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        PostHash: "0xTestPostHash001",
        Value: -1,
      },
      response: Response {
        Message: &error_config.ErrorInfo{
          ErrorCode: "ExceedingDownvoteLimit",
          ErrorData: error_config.ErrorData {
            "actor": "0xLambdaProfileActor002",
            "postHash": "0xTestPostHash001",
          },
          ErrorLocation: "ActorVotesCountersRecordLocation",
        },
        Ok: false,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
