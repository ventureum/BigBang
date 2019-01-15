package feed_attributes

import "time"

type RedeemBlockInfo struct {
	RedeemBlock                  RedeemBlock    `json:"redeemBlock" db:"redeem_block"`
	TotalEnrolledMilestonePoints MilestonePoint `json:"totalEnrolledMilestonePoints" db:"total_enrolled_milestone_points"`
	TokenPool                    int64          `json:"tokenPool" db:"token_pool"`
	ExecutedAt                   time.Time      `json:"executedAt" db:"executed_at"`
}
