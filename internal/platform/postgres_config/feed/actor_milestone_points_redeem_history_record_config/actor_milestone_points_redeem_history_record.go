package actor_milestone_points_redeem_history_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"fmt"
	"time"
)

type ActorMilestonePointsRedeemHistoryRecord struct {
	ID                           string                         `db:"id"`
	Actor                        string                         `db:"actor"`
	RedeemBlock                  feed_attributes.RedeemBlock    `db:"redeem_block"`
	TokenPool                    int64                          `db:"token_pool"`
	TotalEnrolledMilestonePoints feed_attributes.MilestonePoint `db:"total_enrolled_milestone_points"`
	TargetedMilestonePoints      feed_attributes.MilestonePoint `db:"targeted_milestone_points"`
	ActualMilestonePoints        feed_attributes.MilestonePoint `db:"actual_milestone_points"`
	ConsumedMilestonePoints      feed_attributes.MilestonePoint `db:"consumed_milestone_points"`
	RedeemedTokens               int64                          `db:"redeemed_tokens"`
	SubmittedAt                  time.Time                      `db:"submitted_at"`
	ExecutedAt                   time.Time                      `db:"executed_at"`
	CreatedAt                    time.Time                      `db:"created_at"`
	UpdatedAt                    time.Time                      `db:"updated_at"`
}

func (actorMilestonePointsRedeemHistoryRecord *ActorMilestonePointsRedeemHistoryRecord) GenerateID() {
	idStr := fmt.Sprintf("%d:%s", actorMilestonePointsRedeemHistoryRecord.RedeemBlock, actorMilestonePointsRedeemHistoryRecord.Actor)
	actorMilestonePointsRedeemHistoryRecord.ID = idStr
}
