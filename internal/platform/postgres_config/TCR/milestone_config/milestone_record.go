package milestone_config

import (
  "time"
)


type MilestoneRecord struct {
  ProjectId     string         `json:"projectId" db:"project_id"`
  MilestoneId   int64          `json:"milestoneId" db:"milestone_id"`
  Content       string         `json:"content" db:"content"`
  StartTime     int64          `json:"startTime" db:"start_time"`
  EndTime       int64          `json:"endTime" db:"end_time"`
  NumObjs       int64          `json:"numObjs" db:"num_objs"`
  AvgRating     int64          `json:"avgRating" db:"avg_rating"`
  CreatedAt     time.Time      `json:"createdAt" db:"created_at"`
  UpdatedAt     time.Time      `json:"updatedAt" db:"updated_at"`
}
