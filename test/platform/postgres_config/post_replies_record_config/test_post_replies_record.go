package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
)

func main() {
  db := client_config.ConnectPostgresClient()
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*db}
  postRepliesRecordExecutor.DeletePostRepliesRecordTable()
  postRepliesRecordExecutor.CreatePostRepliesRecordTable()

  postRepliesRecord1 := &post_replies_record_config.PostRepliesRecord{
   PostHash: "0xpostHash001",
   ReplyHash: "0xReplyHash001",
  }

  postRepliesRecord2 := &post_replies_record_config.PostRepliesRecord{
    PostHash: "0xpostHash001",
    ReplyHash: "0xReplyHash002",
  }

  postRepliesRecord3 := &post_replies_record_config.PostRepliesRecord{
    PostHash: "0xpostHash002",
    ReplyHash: "0xpReplyHash003",
  }

  postRepliesRecordExecutor.UpsertPostRepliesRecord(postRepliesRecord1)
  postRepliesRecordExecutor.UpsertPostRepliesRecord(postRepliesRecord2)
  postRepliesRecordExecutor.UpsertPostRepliesRecord(postRepliesRecord3)

  postReply1 := postRepliesRecordExecutor.GetPostReplies(postRepliesRecord1.PostHash)
  log.Printf("postReply1: %+v\n", postReply1)

  postReply2 := postRepliesRecordExecutor.GetPostReplies(postRepliesRecord3.PostHash)
  log.Printf("postReply2: %+v\n", postReply2)

  postReply3 := postRepliesRecordExecutor.GetPostReplies("0xpReplyHash004")
  log.Printf("postReply2: %+v\n", postReply3)

  postReplyCount1 := postRepliesRecordExecutor.GetPostRepliesRecordCount(postRepliesRecord1.PostHash)
  log.Printf("Reply Count 1: %+v\n",  postReplyCount1)

  postReplyCount2 := postRepliesRecordExecutor.GetPostRepliesRecordCount(postRepliesRecord3.PostHash)
  log.Printf("Reply Count 2: %+v\n",  postReplyCount2 )

  postRepliesRecordExecutor.DeletePostRepliesRecords(postRepliesRecord1.PostHash)
  postRepliesRecordExecutor.DeletePostRepliesRecords(postRepliesRecord3.PostHash)

  postReplyCount1 = postRepliesRecordExecutor.GetPostRepliesRecordCount(postRepliesRecord1.PostHash)
  log.Printf("Reply Count 1: %+v\n",  postReplyCount1)

  postReplyCount2 = postRepliesRecordExecutor.GetPostRepliesRecordCount(postRepliesRecord3.PostHash)
  log.Printf("Reply Count 2: %+v\n",  postReplyCount2 )
}
