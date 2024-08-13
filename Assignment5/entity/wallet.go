package entity

import (
	"time"
)

type Wallet struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Balance   float64 `gorm:"type:decimal(15,2);default:0"`
	CreatedAt time.Time
}
