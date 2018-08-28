package session_record_config

import (
  "github.com/jmoiron/sqlx/types"
  "time"
  "BigBang/internal/app/feed_attributes"
)


type SessionRecord struct {
  Actor       string
  PostHash    string         `db:"post_hash"`
  StartTime   int64    `db:"start_time"`
  EndTime     int64    `db:"end_time"`
  Content     types.JSONText `db:"content"`
  CreatedAt   time.Time      `db:"created_at"`
  UpdatedAt   time.Time      `db:"updated_at"`
}

type SessionRecordResult struct {
  Actor       string
  PostHash    string
  StartTime   int64
  EndTime     int64
  Content     *feed_attributes.Content
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

func (sessionRecord *SessionRecord) ToSessionRecordResult() *SessionRecordResult{
  return &SessionRecordResult{
    Actor:       sessionRecord.Actor,
    PostHash:    sessionRecord.PostHash,
    StartTime:   sessionRecord.StartTime,
    EndTime:     sessionRecord.EndTime,
    Content:     feed_attributes.CreatedContentFromJsonText(sessionRecord.Content),
    CreatedAt:   sessionRecord.CreatedAt,
    UpdatedAt:   sessionRecord.UpdatedAt,
  }
}


func (sessionRecord *SessionRecord) EmbedSessionRecordToActivity(activity *feed_attributes.Activity) {
  activity.Extra["start_time"] = sessionRecord.StartTime
  activity.Extra["end_time"] = sessionRecord.EndTime
  activity.Extra["is_session"] = true
}
