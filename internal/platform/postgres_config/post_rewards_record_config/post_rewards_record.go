package post_rewards_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)


type PostRewardsRecord struct {
  PostHash string  `db:"post_hash"`
  Rewards feed_attributes.Reputation `db:"rewards"`
  LatestMedianReputations feed_attributes.Reputation `db:"latest_median_reputations"`
  WithdrawableRewards feed_attributes.Reputation `db:"withdrawable_rewards"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
