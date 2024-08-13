package entity

import (
	"time"
)

type Record struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	WalletID    int       `gorm:"not null" json:"wallet_id"`
	CategoryID  int       `json:"category_id"`
	Amount      float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	Type        string    `gorm:"type:varchar(10);not null" json:"type"` // income or expense
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
