package tcr_attributes

type ObjectiveVotesInfoKey struct {
	ProjectId   string `json:"projectId,required"`
	MilestoneId int64  `json:"milestoneId,required"`
	ObjectiveId int64  `json:"objectiveId,required"`
}

type ObjectiveVotesInfo struct {
	ProjectId   string        `json:"projectId,required"`
	MilestoneId int64         `json:"milestoneId,required"`
	ObjectiveId int64         `json:"objectiveId,required"`
	RatingVotes *[]RatingVote `json:"ratingVotes,required"`
}
