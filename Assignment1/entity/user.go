package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar;not null" json:"name" binding:"required"`
	Email     string    `gorm:"type:varchar;uniqueIndex;not null" json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// entity/user.go
type UserWithRiskProfile struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	RiskScore      int       `json:"risk_score"`
	RiskCategory   string    `json:"risk_category"`
	RiskDefinition string    `json:"risk_definition"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
