package postgres_gorm

import (
	"Training/Assignment2/wallet_server/entity"
	"Training/Assignment2/wallet_server/service"
	"context"
	"log"

	"gorm.io/gorm"
)

// GormDBIface defines an interface for GORM DB methods used in the repository
type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type walletRepository struct {
	db GormDBIface
}

// membuat instance baru dari walletRepository
func NewWalletRepository(db GormDBIface) service.InterWalletRepo {
	return &walletRepository{db: db}
}

// CreateWallet membuat wallet baru dalam basis data
func (r *walletRepository) CreateWallet_repo(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	if err := r.db.WithContext(ctx).Create(wallet).Error; err != nil {
		log.Printf("Error creating wallet: %v\n", err)
		return entity.Wallet{}, err
	}
	return *wallet, nil
}
