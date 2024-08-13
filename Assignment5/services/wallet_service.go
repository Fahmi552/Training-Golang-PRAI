package services

import (
	"Training/Assignment5/entity"
	"Training/Assignment5/repositories"
	"errors"
)

type WalletService interface {
	CreateWallet(userID uint, name string, balance float64) (*entity.Wallet, error)
	GetWalletByID(id uint) (*entity.Wallet, error)
	GetWalletsByUserID(userID uint) ([]entity.Wallet, error)
	UpdateWallet(id uint, name string, balance float64) (*entity.Wallet, error)
	DeleteWallet(id uint) error
	IsUserExists(userID uint) (bool, error)
}

type walletService struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) WalletService {
	return &walletService{walletRepo}
}

func (s *walletService) CreateWallet(userID uint, name string, balance float64) (*entity.Wallet, error) {
	if name == "" {
		return nil, errors.New("wallet name cannot be empty")
	}

	wallet := &entity.Wallet{
		UserID:  userID,
		Name:    name,
		Balance: balance,
	}

	err := s.walletRepo.Create(wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *walletService) GetWalletByID(id uint) (*entity.Wallet, error) {
	return s.walletRepo.FindByID(id)
}

func (s *walletService) GetWalletsByUserID(userID uint) ([]entity.Wallet, error) {
	return s.walletRepo.FindAllByUserID(userID)
}

func (s *walletService) UpdateWallet(id uint, name string, balance float64) (*entity.Wallet, error) {
	wallet, err := s.walletRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	wallet.Name = name
	wallet.Balance = balance

	err = s.walletRepo.Update(wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *walletService) DeleteWallet(id uint) error {
	return s.walletRepo.Delete(id)
}

func (s *walletService) IsUserExists(userID uint) (bool, error) {
	var count int64
	err := s.walletRepo.CountUserByID(userID, &count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
