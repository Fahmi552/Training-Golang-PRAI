package repositories

import (
	"Training/Assignment5/entity"

	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(wallet *entity.Wallet) error
	FindByID(id uint) (*entity.Wallet, error)
	FindAllByUserID(userID uint) ([]entity.Wallet, error)
	Update(wallet *entity.Wallet) error
	Delete(id uint) error
	CountUserByID(userID uint, count *int64) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db}
}

func (r *walletRepository) CountUserByID(userID uint, count *int64) error {
	return r.db.Table("users").Where("id = ?", userID).Count(count).Error
}

func (r *walletRepository) Create(wallet *entity.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *walletRepository) FindByID(id uint) (*entity.Wallet, error) {
	var wallet entity.Wallet
	err := r.db.First(&wallet, id).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) FindAllByUserID(userID uint) ([]entity.Wallet, error) {
	var wallets []entity.Wallet
	err := r.db.Where("user_id = ?", userID).Find(&wallets).Error
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (r *walletRepository) Update(wallet *entity.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *walletRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Wallet{}, id).Error
}

// Tambahkan metode lain sesuai kebutuhan
