package rating_vote_config

import (
  "time"
)


type RatingVoteRecord struct {
  ProjectId     string         `json:"projectId" db:"project_id"`
  MilestoneId   int64          `json:"milestoneId" db:"milestone_id"`
  ObjectiveId   int64          `json:"objId" db:"objective_id"`
  Voter         string         `json:"voter" db:"voter"`
  Rating        int64          `json:"rating" db:"rating"`
  Weight        int64          `json:"weight" db:"weight"`
  CreatedAt     time.Time      `json:"createdAt" db:"created_at"`
  UpdatedAt     time.Time      `json:"updatedAt" db:"updated_at"`
}
