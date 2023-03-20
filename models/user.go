package models

import (
	"time"

	"github.com/nikola43/ethfaucet/models/base"
)

type User struct {
	base.CustomGormModel
	WalletAddress string    `gorm:"index; unique; type:varchar(64) not null" json:"wallet_address"`
	IPAddress     string    `gorm:"index; unique; type:varchar(11) not null" json:"ip_address"`
	ClaimDate     time.Time `json:"claim_date"`
}
