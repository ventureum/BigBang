package feed_attributes

type RewardsInfo struct {
	Fuel            Fuel           `json:"fuel"`
	Reputation      Reputation     `json:"reputation"`
	MilestonePoints MilestonePoint `json:"milestonePoints" db:"milestone_points"`
}
