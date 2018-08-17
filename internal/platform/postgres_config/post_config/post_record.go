package post_config

import (
  "github.com/jmoiron/sqlx/types"
  "time"
  "BigBang/internal/app/feed_attributes"
)


type PostRecord struct {
  Actor       string
  BoardId     string         `db:"board_id"`
  ParentHash  string         `db:"parent_hash"`
  PostHash    string         `db:"post_hash"`
  PostType    string         `db:"post_type"`
  Content     types.JSONText `db:"content"`
  UpdateCount int64          `db:"update_count"`
  CreatedAt   time.Time      `db:"created_at"`
  UpdatedAt   time.Time      `db:"updated_at"`
}

type PostRecordResult struct {
  Actor       string
  BoardId     string
  ParentHash  string
  PostHash    string
  PostType    string
  Content     *feed_attributes.Content
  UpdateCount int64
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

func (postRecord *PostRecord) ToPostRecordResult() *PostRecordResult{
  return &PostRecordResult{
    Actor:       postRecord.Actor,
    BoardId:     postRecord.BoardId,
    ParentHash:  postRecord.ParentHash,
    PostHash:    postRecord.PostHash,
    PostType:    postRecord.PostType,
    Content:     feed_attributes.CreatedContentFromToJsonText(postRecord.Content),
    UpdateCount: postRecord.UpdateCount,
    CreatedAt:   postRecord.CreatedAt,
    UpdatedAt:   postRecord.UpdatedAt,
  }
}
