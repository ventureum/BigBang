package purchase_mps_record_config

import (
	"time"
)

type PurchaseMPsRecord struct {
	UUID      string    `db:"uuid"`
	Purchaser string    `db:"purchaser"`
	VetX      int64     `db:"vetx"`
	MPs       int64     `db:"mps"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
