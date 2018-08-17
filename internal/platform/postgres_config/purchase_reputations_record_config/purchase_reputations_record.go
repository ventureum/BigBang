package purchase_reputations_record_config

import (
  "time"
)

type PurchaseReputationsRecord struct {
  UUID string `db:"uuid"`
  Purchaser string  `db:"purchaser"`
  VetX int64 `db:"vetx"`
  Reputations int64 `db:"reputations"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
