package repository

import (
	"Training/Assignment1/service"
	"Training/e-wallet/entity"
	"Training/e-wallet/wallet_server"
	"context"
	"log"

	"gorm.io/gorm"
)


type WalletRepository struct {
	DB GormDBIface
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{DB: db}
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db GormDBIface) service service.IUserRepository {
	return &userRepository{db: db}
}

func (r *WalletRepository) TopUp(userID string, amount float64) error {
	var wallet entity.Wallet
	result := r.DB.Where("user_id = ?", userID).First(&wallet)
	if result.Error != nil {
		return result.Error
	}

	// Lakukan penambahan saldo
	Wallet.Balance += amount
	result = r.DB.Save(&wallet)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Top up sejumlah %f berhasil untuk user %s", amount, userID)
	return nil
}
