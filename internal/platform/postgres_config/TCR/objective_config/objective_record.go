package objective_config

import (
	"time"
)

type ObjectiveRecord struct {
	ProjectId      string    `json:"projectId" db:"project_id"`
	MilestoneId    int64     `json:"milestoneId" db:"milestone_id"`
	ObjectiveId    int64     `json:"objectiveId" db:"objective_id"`
	Content        string    `json:"content" db:"content"`
	BlockTimestamp int64     `json:"block_timestamp" db:"block_timestamp"`
	AvgRating      int64     `json:"avgRating" db:"avg_rating"`
	TotalRating    int64     `json:"totalRating" db:"total_rating"`
	TotalWeight    int64     `json:"totalWeight" db:"total_weight"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}
