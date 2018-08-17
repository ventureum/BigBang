package actor_reputations_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type ActorReputationsRecord struct {
  Actor string  `db:"actor"`
  Reputations feed_attributes.Reputation `db:"reputations"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
