package tcr_attributes

type MilestoneValidatorsInfoKey struct {
	ProjectId   string `json:"projectId,required"`
	MilestoneId int64  `json:"milestoneId,required"`
}

type MilestoneValidatorsInfo struct {
	MilestoneValidatorsInfoKey
	Validators *[]string `json:"validators,required"`
}
