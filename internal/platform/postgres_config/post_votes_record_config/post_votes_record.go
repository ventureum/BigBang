package post_votes_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type PostVotesRecord struct {
  UUID string `db:"uuid"`
  Actor string  `db:"actor"`
  PostHash string  `db:"post_hash"`
  VoteType feed_attributes.VoteType `db:"vote_type"`
  SignedReputations int64 `db:"signed_reputations"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}


const STACK_FRACTION float64 = 0.01
