package actor_profile_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type ActorProfileRecord struct {
	Actor              string                             `db:"actor"`
	ActorType          feed_attributes.ActorType          `db:"actor_type"`
	Username           string                             `db:"username"`
	PhotoUrl           string                             `db:"photo_url"`
	TelegramId         string                             `db:"telegram_id"`
	PhoneNumber        string                             `db:"phone_number"`
	PrivateKey         string                             `db:"private_key"`
	PublicKey          string                             `db:"public_key"`
	ActorProfileStatus feed_attributes.ActorProfileStatus `db:"actor_profile_status"`
	ProfileContent     string                             `db:"profile_content"`
	CreatedAt          time.Time                          `db:"created_at"`
	UpdatedAt          time.Time                          `db:"updated_at"`
}
