package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/app/feed_attributes"
)

func main() {
  db := client_config.ConnectPostgresClient()
  postExecutor := post_config.PostExecutor{*db}
  postExecutor.DeletePostTable()
  postExecutor.CreatePostTable()
  content := &feed_attributes.Content{
  Title: "titleSample1",
  Text: "hello, worl007d",
  }

  postRecord := &post_config.PostRecord{
   Actor:      "0xactor1",
   BoardId:    "0xboard01",
   ParentHash: "0xph1",
   PostHash:   "0xpostHash002",
   PostType:   feed_attributes.PostPostType.Value(),
   Content:    content.ToJsonText(),
  }


  postExecutor.UpsertPostRecord(postRecord)

  postRecordRes := postExecutor.GetPostRecord(postRecord.PostHash)
  log.Printf("%+v\n", postRecordRes.ToPostRecordResult().Content)

  postExecutor.UpsertPostRecord(postRecord)

  postUpdateCount := postExecutor.GetPostUpdateCount(postRecord.PostHash)
  log.Printf("Update Count: %+v\n",  postUpdateCount)

  postExecutor.DeletePostRecord(postRecord.PostHash)
}
