package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/app/feed_attributes"
)


func main() {
  db := client_config.ConnectPostgresClient()
  db.LoadActorTypeEnum()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*db}
  actorProfileRecordExecutor.DeleteActorProfileRecordTable()
  actorProfileRecordExecutor.CreateActorProfileRecordTable()

  actor1 := "0xactor001"
  actor2 := "0xactor002"
  actor3 := "0xactor003"

  actorProfileRecord1 := &actor_profile_record_config.ActorProfileRecord{
    Actor: actor1,
    ActorType: feed_attributes.USER_ACTOR_TYPE,
  }

  actorProfileRecord2 := &actor_profile_record_config.ActorProfileRecord{
    Actor: actor2,
    ActorType: feed_attributes.KOL_ACTOR_TYPE,
  }

  actorProfileRecord3 := &actor_profile_record_config.ActorProfileRecord{
    Actor: actor3,
    ActorType: feed_attributes.ADMIN_ACTOR_TYPE,
  }

  actorProfileRecordExecutor.UpsertActorProfileRecord(actorProfileRecord1)
  actorProfileRecordExecutor.UpsertActorProfileRecord(actorProfileRecord2)
  actorProfileRecordExecutor.UpsertActorProfileRecord(actorProfileRecord3)

  actorProfile1 := actorProfileRecordExecutor.GetActorProfileRecord(actorProfileRecord1.Actor)
  log.Printf("actorProfile1: %+v\n", actorProfile1)

  actorProfile2 := actorProfileRecordExecutor.GetActorProfileRecord(actorProfileRecord2.Actor)
  log.Printf("actorProfile2: %+v\n", actorProfile2)

  actorProfile3 := actorProfileRecordExecutor.GetActorProfileRecord(actorProfileRecord3.Actor)
  log.Printf("actorProfile3: %+v\n", actorProfile3)
}
