package actor_profile_record_config

import (
  "time"
  "BigBang/internal/app/feed_attributes"
)

type ActorProfileRecord struct {
  Actor string  `db:"actor"`
  ActorType feed_attributes.ActorType  `db:"actor_type"`
  Username string `db:"username"`
  PhotoUrl string `db:"photo_url"`
  TelegramId string `db:"telegram_id"`
  PhoneNumber string `db:"phone_number"`
  ActorProfileStatus feed_attributes.ActorProfileStatus  `db:"actor_profile_status"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
