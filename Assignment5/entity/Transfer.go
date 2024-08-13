package entity

import (
	"time"
)

type Transfer struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FromWalletID int       `gorm:"not null" json:"from_wallet_id"`
	ToWalletID   int       `gorm:"not null" json:"to_wallet_id"`
	Amount       float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
