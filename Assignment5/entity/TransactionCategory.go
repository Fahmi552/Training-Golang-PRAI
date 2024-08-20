package entity

import (
	"time"
)

type TransactionCategory struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Struktur untuk menampung rekapitulasi pengeluaran per kategori
type CategoryExpenseRecap struct {
	CategoryID uint    `json:"category_id"`
	Category   string  `json:"category"`
	Total      float64 `json:"total"`
}
