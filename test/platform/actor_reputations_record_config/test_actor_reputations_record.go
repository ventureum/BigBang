package actor_reputations_record_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/actor_reputations_record_config"
  "BigBang/internal/app/feed_attributes"
)


func main() {
  db := client_config.ConnectPostgresClient()
  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{*db}
  actorReputationsRecordExecutor.DeleteActorReputationsRecordTable()
  actorReputationsRecordExecutor.CreateActorReputationsRecordTable()

  actor1 := "0xactor001"
  actor2 := "0xactor002"
  actor3 := "0xactor003"
  actor4 := "0xactor004"
  actorReputationsRecord1 := &actor_reputations_record_config.ActorReputationsRecord{
    Actor: actor1,
    Reputations: feed_attributes.Reputation(4000000),
  }

  actorReputationsRecord2 := &actor_reputations_record_config.ActorReputationsRecord{
    Actor: actor2,
    Reputations: feed_attributes.Reputation(30),
  }

  actorReputationsRecord3 := &actor_reputations_record_config.ActorReputationsRecord{
    Actor: actor3,
    Reputations: feed_attributes.Reputation(20),
  }

  actorReputationsRecordExecutor.UpsertActorReputationsRecord(actorReputationsRecord1)
  actorReputationsRecordExecutor.UpsertActorReputationsRecord(actorReputationsRecord2)
  actorReputationsRecordExecutor.UpsertActorReputationsRecord(actorReputationsRecord3)

  actorReputations1 := actorReputationsRecordExecutor.GetActorReputations(actorReputationsRecord1.Actor)
  log.Printf("actorReputations1: %+v\n", actorReputations1)

  actorReputations2 := actorReputationsRecordExecutor.GetActorReputations(actorReputationsRecord3.Actor)
  log.Printf("actorReputations2: %+v\n", actorReputations2)

  actorReputations3 := actorReputationsRecordExecutor.GetActorReputations(actorReputationsRecord3.Actor)
  log.Printf("actorReputations3: %+v\n", actorReputations3)

  actorReputationsRecordExecutor.AddActorReputations(
    actorReputationsRecord1.Actor,
    feed_attributes.Reputation(500000))

  actorReputations1 = actorReputationsRecordExecutor.GetActorReputations(actorReputationsRecord1.Actor)
  log.Printf("updated actorReputations1: %+v\n", actorReputations1)

  actorReputationsRecordExecutor.SubActorReputations(
    actorReputationsRecord2.Actor,
    feed_attributes.Reputation(5))

  actorReputations2 = actorReputationsRecordExecutor.GetActorReputations(actorReputationsRecord2.Actor)
  log.Printf("updated actorReputations2: %+v\n", actorReputations2)

  actorReputations3 = actorReputationsRecordExecutor.GetActorReputations(actor4)
  log.Printf("updated actorReputations3: %+v\n", actorReputations3)

  // should fail
  actorReputationsRecordExecutor.SubActorReputations(
    actorReputationsRecord2.Actor,
    feed_attributes.Reputation(5000))
}
