package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
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
        Limit: 0,
      },
      response: Response {
        Projects: &[]project_config.ProjectRecordResult{},
        NextCursor: utils.Base64EncodeInt64(6),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Limit: 2,
      },
      response: Response {
        Projects: &[]project_config.ProjectRecordResult{
           {
             ProjectId: "0xProject006",
             Content:   "{offset: 0, length: 41, type: 'url'}",
             AvgRating: 20,
             MilestoneInfo: &tcr_attributes.MilestoneInfo{
               NumMilestones: 5,
               NumMilestonesCompleted: 8,
             },
           },
           {
            ProjectId: "0xProject005",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 20,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 5,
              NumMilestonesCompleted: 8,
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
        Limit: 2,
        Cursor: utils.Base64EncodeInt64(4),
      },
      response: Response {
        Projects: &[]project_config.ProjectRecordResult{
          {
            ProjectId: "0xProject004",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 20,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 5,
              NumMilestonesCompleted: 8,
            },
          },
          {
            ProjectId: "0xProject003",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 20,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 5,
              NumMilestonesCompleted: 8,
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
        Limit: 5,
        Cursor: utils.Base64EncodeInt64(4),
      },
      response: Response {
        Projects: &[]project_config.ProjectRecordResult{
          {
            ProjectId: "0xProject004",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 20,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 5,
              NumMilestonesCompleted: 8,
            },
          },
          {
            ProjectId: "0xProject003",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 20,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 5,
              NumMilestonesCompleted: 8,
            },
          },
          {
            ProjectId: "0xProject002",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 20,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 5,
              NumMilestonesCompleted: 8,
            },
          },
          {
            ProjectId: "0xProject001",
            Content:   "{offset: 0, length: 41, type: 'url'}",
            AvgRating: 5,
            MilestoneInfo: &tcr_attributes.MilestoneInfo{
              NumMilestones: 2,
              NumMilestonesCompleted: 6,
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
    resultProjects := *result.Projects
    responseProjects := *test.response.Projects
    assert.Equal(t, len(resultProjects), len(responseProjects))
    for index, responseProject := range responseProjects {
      assert.Equal(t, resultProjects[index].ProjectId, responseProject.ProjectId)
      assert.Equal(t, resultProjects[index].Content, responseProject.Content)
      assert.Equal(t, resultProjects[index].AvgRating, responseProject.AvgRating)
      assert.Equal(t, resultProjects[index].MilestoneInfo, responseProject.MilestoneInfo)
    }
  }
}
