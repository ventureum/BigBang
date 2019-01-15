package post_votes_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type PostVotesRecord struct {
	UUID                 string                   `json:"uuid,required" db:"uuid"`
	Actor                string                   `json:"actor,required" db:"actor"`
	PostHash             string                   `json:"postHash,required" db:"post_hash"`
	PostType             string                   `json:"postType,required" db:"post_type"`
	VoteType             feed_attributes.VoteType `json:"voteType,required" db:"vote_type"`
	DeltaFuel            int64                    `json:"deltaFuel,required" db:"delta_fuel"`
	DeltaReputation      int64                    `json:"deltaReputation,required" db:"delta_reputation"`
	DeltaMilestonePoints int64                    `json:"deltaMilestonePoints,required" db:"delta_milestone_points"`
	SignedReputation     int64                    `json:"signedReputation,required" db:"signed_reputation"`
	CreatedAt            time.Time                `json:"createdAt,required" db:"created_at"`
	UpdatedAt            time.Time                `json:"updatedAt,required" db:"updated_at"`
}

const STACK_FRACTION float64 = 0.001
