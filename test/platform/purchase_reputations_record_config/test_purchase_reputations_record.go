package purchase_reputations_record_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/purchase_reputations_record_config"
)


func main() {
  db := client_config.ConnectPostgresClient()
  purchaseReputationsRecordExecutor := purchase_reputations_record_config.PurchaseReputationsRecordExecutor{*db}
  purchaseReputationsRecordExecutor.LoadVoteTypeEnum()
  purchaseReputationsRecordExecutor.DeletePurchaseReputationsRecordTable()
  purchaseReputationsRecordExecutor.CreatePurchaseReputationsRecordTable()

  purchaseReputationsRecord1 := &purchase_reputations_record_config.PurchaseReputationsRecord{
    Purchaser: "0x001",
    VetX: 100,
    Reputations: 2000,
  }

  purchaseReputationsRecord2 := &purchase_reputations_record_config.PurchaseReputationsRecord{
    Purchaser: "0x002",
    VetX: 100,
    Reputations: 2000,
  }

  purchaseReputationsRecordExecutor.UpsertPurchaseReputationsRecord(purchaseReputationsRecord1)
  purchaseReputationsRecordExecutor.UpsertPurchaseReputationsRecord(purchaseReputationsRecord2)
  purchaseReputationsRecordExecutor.UpsertPurchaseReputationsRecord(purchaseReputationsRecord1)
  purchaseReputationsRecordExecutor.UpsertPurchaseReputationsRecord(purchaseReputationsRecord2)
}
