package actor_rewards_info_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type ActorRewardsInfoRecord struct {
	Actor                     string                         `db:"actor"`
	Fuel                      feed_attributes.Fuel           `db:"fuel"`
	Reputation                feed_attributes.Reputation     `db:"reputation"`
	MilestonePointsFromVotes  feed_attributes.MilestonePoint `db:"milestone_points_from_votes"`
	MilestonePointsFromPosts  feed_attributes.MilestonePoint `db:"milestone_points_from_posts"`
	MilestonePointsFromOthers feed_attributes.MilestonePoint `db:"milestone_points_from_others"`
	ConsumedMilestonePoints   feed_attributes.MilestonePoint `db:"consumed_milestone_points"`
	MilestonePoints           feed_attributes.MilestonePoint `db:"milestone_points"`
	CreatedAt                 time.Time                      `db:"created_at"`
	UpdatedAt                 time.Time                      `db:"updated_at"`
}
