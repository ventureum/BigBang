package main

import (
  "log"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/fuels_refuel_record_config"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
)


func main() {
  db := client_config.ConnectPostgresClient()
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*db}
  fuelsRefuelRecordExecutor := refuel_record_config.FuelsRefuelRecordExecutor{*db}
  fuelsRefuelRecordExecutor.LoadUuidExtension()
  fuelsRefuelRecordExecutor.DeleteFuelsRefuelRecordTable()
  fuelsRefuelRecordExecutor.CreateFuelsRefuelRecordTable()
  actor1 := "0xActor001"
  actor2 := "0xActor002"
  actor3 := "0xActor003"
  actorProfileRecordExecutor.UpsertActorProfileRecord(&actor_profile_record_config.ActorProfileRecord{
    Actor: actor1,
    ActorType: feed_attributes.USER_ACTOR_TYPE,
  })

  actorProfileRecordExecutor.UpsertActorProfileRecord(&actor_profile_record_config.ActorProfileRecord{
    Actor: actor2,
    ActorType: feed_attributes.USER_ACTOR_TYPE,
  })

  actorProfileRecordExecutor.UpsertActorProfileRecord(&actor_profile_record_config.ActorProfileRecord{
    Actor: actor3,
    ActorType: feed_attributes.USER_ACTOR_TYPE,
  })

  fuelsRefuelRecord1 := &refuel_record_config.FuelsRefuelRecord{
    Actor: actor1,
    Fuels: feed_attributes.Fuel(4000000),
  }

  fuelsRefuelRecord2 := &refuel_record_config.FuelsRefuelRecord{
    Actor: actor2,
    Fuels: feed_attributes.Fuel(30),
  }

  fuelsRefuelRecord3 := &refuel_record_config.FuelsRefuelRecord{
    Actor: actor3,
    Fuels: feed_attributes.Fuel(20),
  }

  fuelsRefuelRecord4 := &refuel_record_config.FuelsRefuelRecord{
    Actor: actor3,
    Fuels: feed_attributes.Fuel(10),
  }

  fuelsRefuelRecordExecutor.UpsertFuelsRefuelRecord(fuelsRefuelRecord1)
  fuelsRefuelRecordExecutor.UpsertFuelsRefuelRecord(fuelsRefuelRecord2)
  fuelsRefuelRecordExecutor.UpsertFuelsRefuelRecord(fuelsRefuelRecord3)
  fuelsRefuelRecordExecutor.UpsertFuelsRefuelRecord(fuelsRefuelRecord4)

  fuelsRefuelRecords1 := fuelsRefuelRecordExecutor.GetFuelsRefuelRecord(fuelsRefuelRecord1.Actor)
  log.Printf("fuelsRefuelRecords1: %+v\n", fuelsRefuelRecords1)

  fuelsRefuelRecords2 := fuelsRefuelRecordExecutor.GetFuelsRefuelRecord(fuelsRefuelRecord2.Actor)
  log.Printf("fuelsRefuelRecords2: %+v\n", fuelsRefuelRecords2)

  fuelsRefuelRecords3 := fuelsRefuelRecordExecutor.GetFuelsRefuelRecord(fuelsRefuelRecord3.Actor)
  log.Printf("fuelsRefuelRecords3: %+v\n", fuelsRefuelRecords3)
}
