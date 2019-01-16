package tcr_attributes

type Milestone struct {
	ProjectId      string         `json:"projectId,required"`
	MilestoneId    int64          `json:"milestoneId,required"`
	Content        string         `json:"content,required"`
	StartTime      int64          `json:"startTime,required"`
	EndTime        int64          `json:"endTime,required"`
	BlockTimestamp int64          `json:"blockTimestamp,required"`
	State          MilestoneState `json:"state,required"`
	NumObjectives  int64          `json:"numObjectives,required"`
	AvgRating      int64          `json:"avgRating,required"`
	Objectives     *[]Objective   `json:"objectives,required"`
}
