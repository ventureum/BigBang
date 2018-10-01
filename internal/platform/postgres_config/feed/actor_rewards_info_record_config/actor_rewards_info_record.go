package actor_rewards_info_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type ActorRewardsInfoRecord struct {
  Actor           string                         `db:"actor"`
  Fuel            feed_attributes.Fuel           `db:"fuel"`
  Reputation      feed_attributes.Reputation     `db:"reputation"`
  MilestonePoints feed_attributes.MilestonePoint `db:"milestone_points"`
  CreatedAt       time.Time                      `db:"created_at"`
  UpdatedAt       time.Time                      `db:"updated_at"`
}
