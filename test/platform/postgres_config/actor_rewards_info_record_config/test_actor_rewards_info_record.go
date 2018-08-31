package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
)


func main() {
  db := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      res := error_config.CreatedErrorInfoFromString(errPanic)
      log.Printf("%+v", res)
      db.Close()
    }
  }()

  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*db}
  actorRewardsInfoRecordExecutor.DeleteActorRewardsInfoRecordTable()
  actorRewardsInfoRecordExecutor.CreateActorRewardsInfoRecordTable()

  actor1 := "0xactor001"
  actor2 := "0xactor002"
  actor3 := "0xactor003"
  actor4 := "0xactor004"
  actorRewardsInfoRecord1 := &actor_rewards_info_record_config.ActorRewardsInfoRecord{
   Actor:      actor1,
   Reputation: feed_attributes.Reputation(4000000),
  }

  actorRewardsInfoRecord2 := &actor_rewards_info_record_config.ActorRewardsInfoRecord{
   Actor:      actor2,
   Reputation: feed_attributes.Reputation(30),
  }

  actorRewardsInfoRecord3 := &actor_rewards_info_record_config.ActorRewardsInfoRecord{
   Actor:      actor3,
   Reputation: feed_attributes.Reputation(20),
  }

  actorRewardsInfoRecordExecutor.UpsertActorRewardsInfoRecord(actorRewardsInfoRecord1)
  actorRewardsInfoRecordExecutor.UpsertActorRewardsInfoRecord(actorRewardsInfoRecord2)
  actorRewardsInfoRecordExecutor.UpsertActorRewardsInfoRecord(actorRewardsInfoRecord3)

  actorRewardsInfo1 := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actorRewardsInfoRecord1.Actor)
  log.Printf("actorRewardsInfo1: %+v\n", actorRewardsInfo1)

  actorRewardsInfo2 := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actorRewardsInfoRecord3.Actor)
  log.Printf("actorRewardsInfo2: %+v\n", actorRewardsInfo2)

  actorRewardsInfo3 := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actorRewardsInfoRecord3.Actor)
  log.Printf("actorRewardsInfo3: %+v\n", actorRewardsInfo3)

  actorRewardsInfoRecordExecutor.AddActorReputation(
   actorRewardsInfoRecord1.Actor,
   feed_attributes.Reputation(500000))

  actorRewardsInfo1 = actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actorRewardsInfoRecord1.Actor)
  log.Printf("updated actorRewardsInfo1: %+v\n", actorRewardsInfo1)

  actorRewardsInfoRecordExecutor.SubActorReputation(
   actorRewardsInfoRecord2.Actor,
   feed_attributes.Reputation(5))

  actorRewardsInfo2 = actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actorRewardsInfoRecord2.Actor)
  log.Printf("updated actorRewardsInfo2: %+v\n", actorRewardsInfo2)

  actorRewardsInfo3 = actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actor4)
  log.Printf("updated actorRewardsInfo3: %+v\n", actorRewardsInfo3)



  // should fail
  actorRewardsInfoRecordExecutor.SubActorReputation(
   actorRewardsInfoRecord2.Actor,
   feed_attributes.Reputation(5000))

  actorRewardsInfoRecordExecutor.SubActorReputation(
   "0x020130",
   feed_attributes.Reputation(4000))
}
