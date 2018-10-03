package project_config

import (
  "github.com/jmoiron/sqlx/types"
  "time"
  "BigBang/internal/app/tcr_attributes"
)


type ProjectRecord struct {
  ID            int64          `json:"id" db:"id"`
  ProjectId     string         `json:"projectId" db:"project_id"`
  Content       string         `json:"content" db:"content"`
  AvgRating     int64          `json:"avgRating" db:"avg_rating"`
  MilestoneInfo types.JSONText `json:"milestoneInfo" db:"milestone_info"`
  CreatedAt     time.Time      `json:"createdAt" db:"created_at"`
  UpdatedAt     time.Time      `json:"updatedAt" db:"updated_at"`
}

type ProjectRecordResult struct {
  ProjectId     string                        `json:"projectId,required"`
  Content       string                        `json:"content,required"`
  AvgRating     int64                         `json:"avgRating,required"`
  MilestoneInfo *tcr_attributes.MilestoneInfo  `json:"milestoneInfo,required"`
  CreatedAt     time.Time                     `json:"createdAt,omitempty"`
  UpdatedAt     time.Time                     `json:"updatedAt,omitempty"`
}

func (projectRecord *ProjectRecord) ToProjectRecordResult() *ProjectRecordResult{
  return &ProjectRecordResult{
    ProjectId:     projectRecord.ProjectId,
    Content:       projectRecord.Content,
    AvgRating:     projectRecord.AvgRating,
    MilestoneInfo: tcr_attributes.CreatedMilestoneInfoFromJsonText(projectRecord.MilestoneInfo),
    CreatedAt:     projectRecord.CreatedAt,
    UpdatedAt:     projectRecord.UpdatedAt,
  }
}
