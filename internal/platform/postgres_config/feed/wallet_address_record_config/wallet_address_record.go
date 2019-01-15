package wallet_address_record_config

import (
	"time"
)

type WalletAddressRecord struct {
	UUID          string    `json:"uuid,required" db:"uuid"`
	Actor         string    `json:"actor,required" db:"actor"`
	WalletAddress string    `json:"walletAddress,required" db:"wallet_address"`
	CreatedAt     time.Time `json:"createdAt,required" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt,required" db:"updated_at"`
}
