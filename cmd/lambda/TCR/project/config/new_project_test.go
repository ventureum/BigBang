package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/tcr_attributes"
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
        Content: "{offset: 0, length: 41, type: 'url'}",
        AvgRating: 5,
        MilestoneInfo: tcr_attributes.MilestoneInfo{
          NumMilestones: 2,
          NumMilestonesCompleted: 6,
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject002",
        Content: "{offset: 0, length: 41, type: 'url'}",
        AvgRating: 20,
        MilestoneInfo: tcr_attributes.MilestoneInfo{
          NumMilestones: 5,
          NumMilestonesCompleted: 8,
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject003",
        Content: "{offset: 0, length: 41, type: 'url'}",
        AvgRating: 20,
        MilestoneInfo: tcr_attributes.MilestoneInfo{
          NumMilestones: 5,
          NumMilestonesCompleted: 8,
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject004",
        Content: "{offset: 0, length: 41, type: 'url'}",
        AvgRating: 20,
        MilestoneInfo: tcr_attributes.MilestoneInfo{
          NumMilestones: 5,
          NumMilestonesCompleted: 8,
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject005",
        Content: "{offset: 0, length: 41, type: 'url'}",
        AvgRating: 20,
        MilestoneInfo: tcr_attributes.MilestoneInfo{
          NumMilestones: 5,
          NumMilestonesCompleted: 8,
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        ProjectId: "0xProject006",
        Content: "{offset: 0, length: 41, type: 'url'}",
        AvgRating: 20,
        MilestoneInfo: tcr_attributes.MilestoneInfo{
          NumMilestones: 5,
          NumMilestonesCompleted: 8,
        },
      },
      response: Response {
        Ok: true,
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
