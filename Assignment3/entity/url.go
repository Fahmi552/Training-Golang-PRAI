package entity

import "time"

type URL struct {
	ID          uint      `gorm:"primary_key"`
	OriginalURL string    `gorm:"type:text;not null"`
	ShortURL    string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
