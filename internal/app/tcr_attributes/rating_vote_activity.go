package tcr_attributes

type RatingVoteActivity struct {
	ProjectId      string `json:"projectId,required" db:"project_id"`
	MilestoneId    int64  `json:"milestoneId,required" db:"milestone_id"`
	ObjectiveId    int64  `json:"objectiveId,required" db:"objective_id"`
	BlockTimestamp int64  `json:"blockTimestamp,required" db:"block_timestamp"`
	Rating         int64  `json:"rating,required" db:"rating"`
	Weight         int64  `json:"weight,required" db:"weight"`
}
