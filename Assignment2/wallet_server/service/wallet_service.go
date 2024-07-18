package service

import (
	"Training/Assignment2/wallet_server/entity"
	"context"
	"fmt"
)

// mendefinisikan interface untuk layanan Wallet
type InterWalletService interface {
	CreateWallet_service(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error)
}

// mendefinisikan interface untuk repository wallet
type InterWalletRepo interface {
	CreateWallet_repo(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error)
}

// walletService adalah implementasi dari InterWalletService yang menggunakan InterWalletRepo
type walletService struct {
	walletRepo InterWalletRepo
}

// NewWalletService membuat instance baru dari walletService
func NewWalletService(walletRepo InterWalletRepo) InterWalletService {
	return &walletService{walletRepo: walletRepo}
}

// Create Wallet
func (s *walletService) CreateWallet_service(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	// Memanggil CreateWallet dari repository untuk membuat wallet baru
	createdWallet, err := s.walletRepo.CreateWallet_repo(ctx, wallet)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("gagal membuat wallet: %v", err)
	}
	return createdWallet, nil
}
