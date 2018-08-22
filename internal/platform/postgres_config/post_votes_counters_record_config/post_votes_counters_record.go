package post_votes_counters_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type PostVotesCountersRecord struct {
  PostHash string  `db:"post_hash"`
  LatestVoteType feed_attributes.VoteType `db:"latest_vote_type"`
  DownVoteCount int64 `db:"downvote_count"`
  UpVoteCount int64 `db:"upvote_count"`
  TotalVoteCount int64 `db:"total_vote_count"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
