package feed_attributes

import (
	"BigBang/internal/pkg/utils"
	"fmt"
	"time"
)

type MilestonePointsRedeemHistory struct {
	Actor                        string         `json:"actor" db:"actor"`
	RedeemBlock                  RedeemBlock    `json:"redeemBlock" db:"redeem_block"`
	TokenPool                    int64          `json:"tokenPool" db:"token_pool"`
	TotalEnrolledMilestonePoints MilestonePoint `json:"totalEnrolledMilestonePoints" db:"total_enrolled_milestone_points"`
	TargetedMilestonePoints      MilestonePoint `json:"targetedMilestonePoints" db:"targeted_milestone_points"`
	ActualMilestonePoints        MilestonePoint `json:"actualMilestonePoints" db:"actual_milestone_points"`
	ConsumedMilestonePoints      MilestonePoint `json:"consumedMilestonePoints" db:"consumed_milestone_points"`
	RedeemedTokens               int64          `json:"redeemedTokens" db:"redeemed_tokens"`
	SubmittedAt                  time.Time      `json:"submittedAt" db:"submitted_at"`
	ExecutedAt                   time.Time      `json:"executedAt" db:"executed_at"`
}

func (milestonePointsRedeemHistory MilestonePointsRedeemHistory) GenerateRecordID() string {
	idStr := fmt.Sprintf("%d:%s", milestonePointsRedeemHistory.RedeemBlock, milestonePointsRedeemHistory.Actor)
	return utils.Base64EncodeStr(idStr)
}
