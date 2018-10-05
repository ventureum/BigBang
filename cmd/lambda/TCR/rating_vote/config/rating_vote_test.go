package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
  "BigBang/internal/platform/postgres_config/client_config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Voter: "0xVoter001",
        Rating: 20,
        Weight: 30,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Voter: "0xVoter002",
        Rating: 20,
        Weight: 30,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Voter: "0xVoter003",
        Rating: 20,
        Weight: 30,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Voter: "0xVoter004",
        Rating: 20,
        Weight: 30,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Voter: "0xVoter005",
        Rating: 20,
        Weight: 30,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Voter: "0xVoter006",
        Rating: 20,
        Weight: 30,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 2,
        Voter: "0xVoter001",
        Rating: 25,
        Weight: 35,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 2,
        Voter: "0xVoter002",
        Rating: 30,
        Weight: 40,
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
  }
  postgresBigBangClient := client_config.ConnectPostgresClient()
  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
  ratingVoteExecutor.ClearRatingVoteTable()

  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
