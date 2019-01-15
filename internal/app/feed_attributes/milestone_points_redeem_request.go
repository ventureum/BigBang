package feed_attributes

import "time"

type MilestonePointsRedeemRequest struct {
	Actor                   string         `json:"actor" db:"actor"`
	NextRedeemBlock         int64          `json:"nextRedeemBlock" db:"next_redeem_block"`
	TargetedMilestonePoints MilestonePoint `json:"targetedMilestonePoints" db:"targeted_milestone_points"`
	SubmittedAt             time.Time      `json:"submittedAt" db:"updated_at"`
}
