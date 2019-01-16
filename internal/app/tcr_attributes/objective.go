package tcr_attributes

type Objective struct {
	ProjectId      string `json:"projectId,required"`
	MilestoneId    int64  `json:"milestoneId,required"`
	ObjectiveId    int64  `json:"objectiveId,required"`
	Content        string `json:"content,required"`
	BlockTimestamp int64  `json:"blockTimestamp,required"`
	AvgRating      int64  `json:"avgRating,required"`
}
