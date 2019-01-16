package actor_votes_counters_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type ActorVotesCountersRecord struct {
	PostHash                    string                     `db:"post_hash"`
	Actor                       string                     `db:"actor"`
	LatestReputation            feed_attributes.Reputation `db:"latest_reputation"`
	LatestVoteType              feed_attributes.VoteType   `db:"latest_vote_type"`
	LatestReputationForUpvote   feed_attributes.Reputation `db:"latest_reputation_for_upvote"`
	LatestReputationForDownvote feed_attributes.Reputation `db:"latest_reputation_for_downvote"`
	DownVoteCount               int64                      `db:"downvote_count"`
	UpVoteCount                 int64                      `db:"upvote_count"`
	TotalVoteCount              int64                      `db:"total_vote_count"`
	CreatedAt                   time.Time                  `db:"created_at"`
	UpdatedAt                   time.Time                  `db:"updated_at"`
}
