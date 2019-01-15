package post_votes_counters_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type PostVotesCountersRecord struct {
	PostHash                   string                     `db:"post_hash"`
	LatestVoteType             feed_attributes.VoteType   `db:"latest_vote_type"`
	LatestActorReputation      feed_attributes.Reputation `db:"latest_actor_reputation"`
	DownVoteCount              int64                      `db:"downvote_count"`
	UpVoteCount                int64                      `db:"upvote_count"`
	TotalVoteCount             int64                      `db:"total_vote_count"`
	TotalReputationForUpvote   feed_attributes.Reputation `db:"total_reputation_for_upvote"`
	TotalReputationForDownvote feed_attributes.Reputation `db:"total_reputation_for_downvote"`
	TotalReputationForVote     feed_attributes.Reputation `db:"total_reputation_for_vote"`
	CreatedAt                  time.Time                  `db:"created_at"`
	UpdatedAt                  time.Time                  `db:"updated_at"`
}
