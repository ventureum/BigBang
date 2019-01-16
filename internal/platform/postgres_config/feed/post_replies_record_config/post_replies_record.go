package post_replies_record_config

import (
	"time"
)

type PostRepliesRecord struct {
	PostHash  string    `db:"post_hash"`
	ReplyHash string    `db:"reply_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
