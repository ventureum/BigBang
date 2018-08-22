package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_votes_counters_record_config"
  "BigBang/internal/app/feed_attributes"
)

func main() {
  db := client_config.ConnectPostgresClient()
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*db}
  postVotesCountersRecordExecutor.LoadVoteTypeEnum()
  postVotesCountersRecordExecutor.DeletePostVotesCountersRecordTable()
  postVotesCountersRecordExecutor.CreatePostVotesCountersRecordTable()
  
  postHashOne := "0xpostHash001"
  postHashTwo := "0xpostHash002"

  postVotesCountersRecord1 := &post_votes_counters_record_config.PostVotesCountersRecord{
    PostHash: postHashOne,
    LatestVoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  postVotesCountersRecord2 := &post_votes_counters_record_config.PostVotesCountersRecord{
    PostHash: postHashTwo,
    LatestVoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  postVotesCountersRecord3 := &post_votes_counters_record_config.PostVotesCountersRecord{
    PostHash: postHashOne,
    LatestVoteType: feed_attributes.UP_VOTE_TYPE,
  }

  postVotesCountersRecord4 := &post_votes_counters_record_config.PostVotesCountersRecord{
    PostHash: postHashTwo,
    LatestVoteType: feed_attributes.UP_VOTE_TYPE,
  }

  postVotesCountersRecord5 := &post_votes_counters_record_config.PostVotesCountersRecord{
    PostHash: postHashTwo,
    LatestVoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  updatedRes := postVotesCountersRecordExecutor.UpsertPostVotesCountersRecord(postVotesCountersRecord1)
  log.Printf("Upserted PostVotesCountersRecord for postHash %s: %+v\n",
    updatedRes.PostHash, updatedRes)

  updatedRes = postVotesCountersRecordExecutor.UpsertPostVotesCountersRecord(postVotesCountersRecord2)
  log.Printf("Upserted PostVotesCountersRecord for postHash %s: %+v\n",
    updatedRes.PostHash, updatedRes)


  updatedRes = postVotesCountersRecordExecutor.UpsertPostVotesCountersRecord(postVotesCountersRecord3)
  log.Printf("Upserted PostVotesCountersRecord for postHash %s: %+v\n",
    updatedRes.PostHash, updatedRes)

  updatedRes = postVotesCountersRecordExecutor.UpsertPostVotesCountersRecord(postVotesCountersRecord4)
  log.Printf("Upserted PostVotesCountersRecord for postHash %s: %+v\n",
    updatedRes.PostHash, updatedRes)

  updatedRes = postVotesCountersRecordExecutor.UpsertPostVotesCountersRecord(postVotesCountersRecord5)
  log.Printf("Upserted PostVotesCountersRecord for postHash %s: %+v\n",
    updatedRes.PostHash, updatedRes)
}
