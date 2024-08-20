package services

import (
	"Training/Assignment5/entity"
	"Training/Assignment5/repositories"
	"errors"
	"strconv"
	"time"
)

type WalletService interface {
	CreateWallet(userID uint, name string, balance float64) (*entity.Wallet, error)
	GetWalletByID(id uint) (*entity.Wallet, error)
	GetWalletsByUserID(userID uint) ([]entity.Wallet, error)
	UpdateWallet(id uint, name string, balance float64) (*entity.Wallet, error)
	DeleteWallet(id uint) error
	IsUserExists(userID uint) (bool, error)
	Transfer(fromWalletID, toWalletID, categoryID, userID uint, amount float64) error
	GetRecordsBetweenTimes(walletID uint, startTime, endTime time.Time) ([]entity.Record, error)
	GetCashFlow(startTime, endTime time.Time) (float64, float64, error)
	GetExpenseRecapByCategory(startTime, endTime time.Time) ([]entity.CategoryExpenseRecap, error) // Perbarui tipe ini
	GetLastRecords(limit int) ([]entity.Record, error)
}

type walletService struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) WalletService {
	return &walletService{walletRepo: walletRepo}
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

func (s *walletService) Transfer(fromWalletID, toWalletID, categoryID, userID uint, amount float64) error {
	if fromWalletID == toWalletID {
		return errors.New("cannot transfer to the same wallet")
	}

	// Validasi keberadaan wallet pengirim
	fromWallet, err := s.walletRepo.FindByID(fromWalletID)
	if err != nil {
		return err
	}
	if fromWallet == nil {
		return errors.New("from_wallet not found")
	}

	// Validasi keberadaan wallet penerima
	toWallet, err := s.walletRepo.FindByID(toWalletID)
	if err != nil {
		return err
	}
	if toWallet == nil {
		return errors.New("to_wallet not found")
	}

	// Validasi pengguna
	userExists, err := s.walletRepo.UserExists(userID)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user not found")
	}

	// Validasi apakah kategori dimiliki oleh user
	categoryOwned, err := s.walletRepo.IsCategoryOwnedByUser(categoryID, userID)
	if err != nil {
		return err
	}
	if !categoryOwned {
		return errors.New("category not found or not owned by user")
	}

	// Validasi apakah saldo cukup
	if fromWallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	// Lakukan pengurangan saldo dari wallet pengirim
	fromWallet.Balance -= amount
	if err := s.walletRepo.Update(fromWallet); err != nil {
		return err
	}

	// Lakukan penambahan saldo ke wallet penerima
	toWallet.Balance += amount
	if err := s.walletRepo.Update(toWallet); err != nil {
		return err
	}

	// Buat record transaksi untuk wallet pengirim (debit)
	fromDescription := "Transfer to Wallet ID " + strconv.Itoa(int(toWalletID))
	if err := s.walletRepo.CreateTransactionRecord(fromWalletID, categoryID, -amount, "expense", fromDescription); err != nil {
		return err
	}

	// Buat record transaksi untuk wallet penerima (credit)
	toDescription := "Transfer from Wallet ID " + strconv.Itoa(int(fromWalletID))
	if err := s.walletRepo.CreateTransactionRecord(toWalletID, categoryID, amount, "income", toDescription); err != nil {
		return err
	}

	return nil
}

func (s *walletService) GetRecordsBetweenTimes(walletID uint, startTime, endTime time.Time) ([]entity.Record, error) {
	// Validasi keberadaan wallet
	wallet, err := s.walletRepo.FindByID(walletID)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	// Ambil record dari repository
	records, err := s.walletRepo.GetRecordsBetweenTimes(walletID, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *walletService) GetCashFlow(startTime, endTime time.Time) (float64, float64, error) {
	income, expense, err := s.walletRepo.GetCashFlow(startTime, endTime)
	if err != nil {
		return 0, 0, err
	}
	return income, expense, nil
}

func (s *walletService) GetExpenseRecapByCategory(startTime, endTime time.Time) ([]entity.CategoryExpenseRecap, error) {
	recaps, err := s.walletRepo.GetExpenseRecapByCategory(startTime, endTime)
	if err != nil {
		return nil, err
	}
	return recaps, nil
}
func (s *walletService) GetLastRecords(limit int) ([]entity.Record, error) {
	records, err := s.walletRepo.GetLastRecords(limit)
	if err != nil {
		return nil, err
	}
	return records, nil
}
