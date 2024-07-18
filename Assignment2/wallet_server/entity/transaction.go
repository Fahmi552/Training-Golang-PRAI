package entity

import "time"

type Transaction struct {
	ID           int32     `gorm:"primary_key"`
	UserID       int32     `gorm:"not null"`
	Type         string    `gorm:"type:varchar(50);not null"`
	Amount       float32   `gorm:"type:numeric(15,2);not null"`
	TargetUserID int32     `gorm:""`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
