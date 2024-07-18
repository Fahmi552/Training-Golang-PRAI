package entity

import "time"

type User struct {
	ID        int32     `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Wallet struct {
	UserID  int32   `gorm:"primary_key"`
	Balance float32 `gorm:"type:numeric(15,2);default:0.00"`
}

type Transaction struct {
	ID           int32     `gorm:"primary_key"`
	UserID       int32     `gorm:"not null"`
	Type         string    `gorm:"type:varchar(50);not null"`
	Amount       float32   `gorm:"type:numeric(15,2);not null"`
	TargetUserID int32     `gorm:""`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
