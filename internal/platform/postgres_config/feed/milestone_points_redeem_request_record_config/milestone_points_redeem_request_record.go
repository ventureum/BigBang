package milestone_points_redeem_request_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type MilestonePointsRedeemRequestRecord struct {
  Actor           string                         `db:"actor"`
  NextRedeemBlock  int64     `db:"next_redeem_block"`
  TargetedMilestonePoints feed_attributes.MilestonePoint `db:"targeted_milestone_points"`
  CreatedAt       time.Time                      `db:"created_at"`
  UpdatedAt       time.Time                      `db:"updated_at"`
}
