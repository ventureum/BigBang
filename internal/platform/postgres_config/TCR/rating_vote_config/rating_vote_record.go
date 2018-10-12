package rating_vote_config

import (
  "time"
  "fmt"
  "BigBang/internal/pkg/utils"
)


type RatingVoteRecord struct {
  ID            string         `json:"id" db:"id"`
  ProjectId     string         `json:"projectId" db:"project_id"`
  MilestoneId   int64          `json:"milestoneId" db:"milestone_id"`
  ObjectiveId   int64          `json:"objectiveId" db:"objective_id"`
  Voter         string         `json:"voter" db:"voter"`
  BlockTimestamp int64         `json:"block_timestamp" db:"block_timestamp"`
  Rating        int64          `json:"rating" db:"rating"`
  Weight        int64          `json:"weight" db:"weight"`
  CreatedAt     time.Time      `json:"createdAt" db:"created_at"`
  UpdatedAt     time.Time      `json:"updatedAt" db:"updated_at"`
}


func (ratingVoteRecord *RatingVoteRecord) GenerateID() {
  idStr := fmt.Sprintf("%s:%09d:%09d", ratingVoteRecord.ProjectId, ratingVoteRecord.MilestoneId, ratingVoteRecord.ObjectiveId)
  ratingVoteRecord.ID = idStr
}

func (ratingVoteRecord *RatingVoteRecord) EncodeID() string {
  idStr := fmt.Sprintf("%s:%09d:%09d", ratingVoteRecord.ProjectId, ratingVoteRecord.MilestoneId, ratingVoteRecord.ObjectiveId)
  return utils.Base64EncodeIdByInt64AndStr(ratingVoteRecord.BlockTimestamp, idStr)
}