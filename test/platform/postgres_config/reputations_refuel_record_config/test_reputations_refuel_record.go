package main

import (
  "log"
  "BigBang/internal/platform/postgres_config/reputations_refuel_record_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
)


func main() {
  db := client_config.ConnectPostgresClient()
  reputationsRefuelRecordExecutor := reputations_refuel_record_config.ReputationsRefuelRecordExecutor{*db}
  reputationsRefuelRecordExecutor.LoadUuidExtension()
  reputationsRefuelRecordExecutor.DeleteReputationsRefuelRecordTable()
  reputationsRefuelRecordExecutor.CreateReputationsRefuelRecordTable()

  reputationsRefuelRecord1 := &reputations_refuel_record_config.ReputationsRefuelRecord{
    Actor: "0xpostHash001",
    Reputations: feed_attributes.Reputation(4000000),
  }

  reputationsRefuelRecord2 := &reputations_refuel_record_config.ReputationsRefuelRecord{
    Actor: "0xpostHash002",
    Reputations: feed_attributes.Reputation(30),
  }

  reputationsRefuelRecord3 := &reputations_refuel_record_config.ReputationsRefuelRecord{
    Actor: "0xpostHash003",
    Reputations: feed_attributes.Reputation(20),
  }

  reputationsRefuelRecord4 := &reputations_refuel_record_config.ReputationsRefuelRecord{
    Actor: "0xpostHash003",
    Reputations: feed_attributes.Reputation(10),
  }

  reputationsRefuelRecordExecutor.UpsertReputationsRefuelRecord(reputationsRefuelRecord1)
  reputationsRefuelRecordExecutor.UpsertReputationsRefuelRecord(reputationsRefuelRecord2)
  reputationsRefuelRecordExecutor.UpsertReputationsRefuelRecord(reputationsRefuelRecord3)
  reputationsRefuelRecordExecutor.UpsertReputationsRefuelRecord(reputationsRefuelRecord4)

  reputationsRefuelRecords1 := reputationsRefuelRecordExecutor.GetReputationsRefuelRecord(reputationsRefuelRecord1.Actor)
  log.Printf("reputationsRefuelRecords1: %+v\n", reputationsRefuelRecords1)

  reputationsRefuelRecords2 := reputationsRefuelRecordExecutor.GetReputationsRefuelRecord(reputationsRefuelRecord2.Actor)
  log.Printf("reputationsRefuelRecords2: %+v\n", reputationsRefuelRecords2)

  reputationsRefuelRecords3 := reputationsRefuelRecordExecutor.GetReputationsRefuelRecord(reputationsRefuelRecord3.Actor)
  log.Printf("reputationsRefuelRecords3: %+v\n", reputationsRefuelRecords3)
}
