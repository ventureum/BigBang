package rating_vote_config

import (
	"BigBang/internal/pkg/utils"
	"fmt"
	"time"
)

type RatingVoteRecord struct {
	ID             string    `json:"id" db:"id"`
	ProjectId      string    `json:"projectId" db:"project_id"`
	MilestoneId    int64     `json:"milestoneId" db:"milestone_id"`
	ObjectiveId    int64     `json:"objectiveId" db:"objective_id"`
	Voter          string    `json:"voter" db:"voter"`
	BlockTimestamp int64     `json:"block_timestamp" db:"block_timestamp"`
	Rating         int64     `json:"rating" db:"rating"`
	Weight         int64     `json:"weight" db:"weight"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

func (ratingVoteRecord *RatingVoteRecord) GenerateID() {
	idStr := fmt.Sprintf("%d:%s:%09d:%09d:%s", ratingVoteRecord.BlockTimestamp, ratingVoteRecord.ProjectId, ratingVoteRecord.MilestoneId, ratingVoteRecord.ObjectiveId, ratingVoteRecord.Voter)
	ratingVoteRecord.ID = idStr
}

func (ratingVoteRecord *RatingVoteRecord) EncodeID() string {
	idStr := fmt.Sprintf("%s:%09d:%09d:%s", ratingVoteRecord.ProjectId, ratingVoteRecord.MilestoneId, ratingVoteRecord.ObjectiveId, ratingVoteRecord.Voter)
	return utils.Base64EncodeIdByInt64AndStr(ratingVoteRecord.BlockTimestamp, idStr)
}

func GenerateEncodedRatingVoteRecordID(projectId string, milestoneId int64, objectiveId int64, voter string, blockTimestamp int64) string {
	idStr := fmt.Sprintf("%s:%09d:%09d:%s", projectId, milestoneId, objectiveId, voter)
	return utils.Base64EncodeIdByInt64AndStr(blockTimestamp, idStr)
}

func GenerateRatingVoteRecordID(projectId string, milestoneId int64, objectiveId int64, voter string, blockTimestamp int64) string {
	idStr := fmt.Sprintf("%d:%s:%09d:%09d:%s", blockTimestamp, projectId, milestoneId, objectiveId, voter)
	return idStr
}
