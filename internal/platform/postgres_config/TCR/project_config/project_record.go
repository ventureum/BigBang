package project_config

import (
	"time"
)

type ProjectRecord struct {
	ID                     string    `json:"id" db:"id"`
	ProjectId              string    `json:"projectId" db:"project_id"`
	Admin                  string    `json:"admin" db:"admin"`
	Content                string    `json:"content" db:"content"`
	BlockTimestamp         int64     `json:"block_timestamp" db:"block_timestamp"`
	AvgRating              int64     `json:"avgRating" db:"avg_rating"`
	TotalRating            int64     `json:"totalRating" db:"total_rating"`
	TotalWeight            int64     `json:"totalWeight" db:"total_weight"`
	CurrentMilestone       int64     `json:"currentMilestone" db:"current_milestone"`
	NumMilestones          int64     `json:"numMilestones" db:"num_milestones"`
	NumMilestonesCompleted int64     `json:"numMilestonesCompleted" db:"num_milestones_completed"`
	CreatedAt              time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time `json:"updatedAt" db:"updated_at"`
}
