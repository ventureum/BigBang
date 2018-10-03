package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
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
      },
      response: Response {
        Project: &project_config.ProjectRecordResult{
          ProjectId: "0xProject001",
          Content:   "{offset: 0, length: 41, type: 'url'}",
          AvgRating: 5,
          MilestoneInfo: &tcr_attributes.MilestoneInfo{
            NumMilestones:          2,
            NumMilestonesCompleted: 6,
          },
        },
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.Equal(t, test.response.Project.ProjectId, result.Project.ProjectId)
    assert.Equal(t, test.response.Project.MilestoneInfo, result.Project.MilestoneInfo)
    assert.Equal(t, test.response.Project.AvgRating, result.Project.AvgRating)
    assert.Equal(t, test.response.Project.Content, result.Project.Content)
  }
}
