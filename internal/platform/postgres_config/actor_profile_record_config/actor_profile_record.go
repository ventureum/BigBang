package actor_profile_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type ActorProfileRecord struct {
  Actor string  `db:"actor"`
  ActorType feed_attributes.ActorType  `db:"actor_type"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
