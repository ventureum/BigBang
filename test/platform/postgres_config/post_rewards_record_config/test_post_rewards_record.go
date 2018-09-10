package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/app/feed_attributes"
)


func main() {
  db := client_config.ConnectPostgresClient()
  postRewardsRecordExecutor := post_mps_record_config.PostRewardsRecordExecutor{*db}
  postRewardsRecordExecutor.DeletePostRewardsRecordTable()
  postRewardsRecordExecutor.CreatePostRewardsRecordTable()

  postRewardsRecord1 := &post_mps_record_config.PostRewardsRecord{
   PostHash: "0xpostHash001",
    Rewards: feed_attributes.Reputation(4000000),
  }

  postRewardsRecord2 := &post_mps_record_config.PostRewardsRecord{
    PostHash: "0xpostHash002",
    Rewards: feed_attributes.Reputation(30),
  }

  postRewardsRecord3 := &post_mps_record_config.PostRewardsRecord{
    PostHash: "0xpostHash003",
    Rewards: feed_attributes.Reputation(20),
  }

  postRewardsRecordExecutor.UpsertPostRewardsRecord(postRewardsRecord1)
  postRewardsRecordExecutor.UpsertPostRewardsRecord(postRewardsRecord2)
  postRewardsRecordExecutor.UpsertPostRewardsRecord(postRewardsRecord3)

  postRewards1 := postRewardsRecordExecutor.GetPostRewards(postRewardsRecord1.PostHash)
  log.Printf("postRewards1: %+v\n", postRewards1)

  postRewards2 := postRewardsRecordExecutor.GetPostRewards(postRewardsRecord2.PostHash)
  log.Printf("postRewards2: %+v\n", postRewards2)

  postRewards3 := postRewardsRecordExecutor.GetPostRewards(postRewardsRecord3.PostHash)
  log.Printf("postRewards3: %+v\n", postRewards3)

  postRewardsRecordExecutor.AddPostRewards(
    postRewardsRecord1.PostHash,
    feed_attributes.Reputation(500000))

  postRewards1 = postRewardsRecordExecutor.GetPostRewards(postRewardsRecord1.PostHash)
  log.Printf("updated postRewards1: %+v\n", postRewards1)

  postRewardsRecordExecutor.SubPostRewards(
    postRewardsRecord2.PostHash,
    feed_attributes.Reputation(5))

  postRewards2 = postRewardsRecordExecutor.GetPostRewards(postRewardsRecord2.PostHash)
  log.Printf("updated postRewards2: %+v\n", postRewards2)

  // should fail
  postRewardsRecordExecutor.SubPostRewards(
    postRewardsRecord2.PostHash,
    feed_attributes.Reputation(5000))
}
