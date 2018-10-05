package config

import (
  "BigBang/internal/app/tcr_attributes"
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/pkg/utils"
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
        Limit: 0,
      },
      response: Response {
        ObjVoteInfo: &tcr_attributes.ObjVoteInfo{
          ProjectId:   "0xProject001",
          MilestoneId: 1,
          ObjectiveId: 1,
          RatingVotes:  &[]tcr_attributes.RatingVote{},
        },
        NextCursor: utils.Base64EncodeInt64(6),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Limit: 2,
      },
      response: Response {
        ObjVoteInfo: &tcr_attributes.ObjVoteInfo{
          ProjectId:   "0xProject001",
          MilestoneId: 1,
          ObjectiveId: 1,
          RatingVotes: &[]tcr_attributes.RatingVote{
            {
              Voter:  "0xVoter006",
              Rating: 20,
              Weight: 30,
            },
            {
              Voter:  "0xVoter005",
              Rating: 20,
              Weight: 30,
            },
          },
        },
        NextCursor: utils.Base64EncodeInt64(4),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Limit: 2,
        Cursor: utils.Base64EncodeInt64(4),
      },
      response: Response {
        ObjVoteInfo: &tcr_attributes.ObjVoteInfo{
          ProjectId:   "0xProject001",
          MilestoneId: 1,
          ObjectiveId: 1,
          RatingVotes: &[]tcr_attributes.RatingVote{
            {
              Voter:  "0xVoter004",
              Rating: 20,
              Weight: 30,
            },
            {
              Voter:  "0xVoter003",
              Rating: 20,
              Weight: 30,
            },
          },
        },
        NextCursor: utils.Base64EncodeInt64(2),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject001",
        MilestoneId: 1,
        ObjectiveId: 1,
        Limit: 5,
        Cursor: utils.Base64EncodeInt64(4),
      },
      response: Response {
        ObjVoteInfo: &tcr_attributes.ObjVoteInfo{
          ProjectId:   "0xProject001",
          MilestoneId: 1,
          ObjectiveId: 1,
          RatingVotes: &[]tcr_attributes.RatingVote{
            {
              Voter:  "0xVoter004",
              Rating: 20,
              Weight: 30,
            },
            {
              Voter:  "0xVoter003",
              Rating: 20,
              Weight: 30,
            },
            {
              Voter:  "0xVoter002",
              Rating: 20,
              Weight: 30,
            },
            {
              Voter:  "0xVoter001",
              Rating: 20,
              Weight: 30,
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
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.Equal(t, test.response.NextCursor, result.NextCursor)
    objVoteInfo := *result.ObjVoteInfo
    responseObjVoteInfo := *test.response.ObjVoteInfo
    assert.Equal(t, responseObjVoteInfo.ProjectId, objVoteInfo.ProjectId)
    assert.Equal(t, responseObjVoteInfo.MilestoneId, objVoteInfo.MilestoneId)
    assert.Equal(t, responseObjVoteInfo.ObjectiveId, objVoteInfo.ObjectiveId)
    responseRatingVotes := *responseObjVoteInfo.RatingVotes
    for index, ratingVote := range *objVoteInfo.RatingVotes {
      assert.Equal(t, responseRatingVotes[index].Voter, ratingVote.Voter)
      assert.Equal(t, responseRatingVotes[index].Rating, ratingVote.Rating)
      assert.Equal(t, responseRatingVotes[index].Weight, ratingVote.Weight)
    }
  }
}
