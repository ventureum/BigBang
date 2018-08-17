package post_reputations_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type PostReputationsRecord struct {
  PostHash string  `db:"post_hash"`
  Actor string  `db:"actor"`
  Reputations feed_attributes.Reputation `db:"reputations"`
  LatestVoteType feed_attributes.VoteType `db:"latest_vote_type"`
  TotalVoteCount int64 `db:"total_vote_count"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
