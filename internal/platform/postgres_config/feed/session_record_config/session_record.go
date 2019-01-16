package session_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"github.com/jmoiron/sqlx/types"
	"time"
)

type SessionRecord struct {
	Actor     string
	PostHash  string         `db:"post_hash"`
	StartTime int64          `db:"start_time"`
	EndTime   int64          `db:"end_time"`
	Content   types.JSONText `db:"content"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type SessionRecordResult struct {
	Actor     string                   `json:"actor"`
	PostHash  string                   `json:"postHash"`
	StartTime int64                    `json:"startTime"`
	EndTime   int64                    `json:"endTime"`
	Content   *feed_attributes.Content `json:"content"`
	CreatedAt time.Time                `json:"createdAt"`
	UpdatedAt time.Time                `json:"updatedAt"`
}

func (sessionRecord *SessionRecord) ToSessionRecordResult() *SessionRecordResult {
	return &SessionRecordResult{
		Actor:     sessionRecord.Actor,
		PostHash:  sessionRecord.PostHash,
		StartTime: sessionRecord.StartTime,
		EndTime:   sessionRecord.EndTime,
		Content:   feed_attributes.CreatedContentFromJsonText(sessionRecord.Content),
		CreatedAt: sessionRecord.CreatedAt,
		UpdatedAt: sessionRecord.UpdatedAt,
	}
}

func (sessionRecord *SessionRecord) EmbedSessionRecordToActivity(activity *feed_attributes.Activity) {
	activity.Extra["start_time"] = sessionRecord.StartTime
	activity.Extra["end_time"] = sessionRecord.EndTime
	activity.Extra["is_session"] = true
}
