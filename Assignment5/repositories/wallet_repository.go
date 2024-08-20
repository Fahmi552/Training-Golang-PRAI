package repositories

import (
	"Training/Assignment5/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(wallet *entity.Wallet) error
	FindByID(id uint) (*entity.Wallet, error)
	FindAllByUserID(userID uint) ([]entity.Wallet, error)
	Update(wallet *entity.Wallet) error
	Delete(id uint) error
	CountUserByID(userID uint, count *int64) error
	TransferBetweenWallets(fromWalletID, toWalletID, categoryID, userID int, amount float64) error
	UserExists(userID uint) (bool, error)
	CreateTransactionRecord(walletID, categoryID uint, amount float64, transactionType, description string) error
	IsCategoryOwnedByUser(categoryID, userID uint) (bool, error)
	GetRecordsBetweenTimes(walletID uint, startTime, endTime time.Time) ([]entity.Record, error)
	GetCashFlow(startTime, endTime time.Time) (float64, float64, error)
	GetExpenseRecapByCategory(startTime, endTime time.Time) ([]entity.CategoryExpenseRecap, error)
	GetLastRecords(limit int) ([]entity.Record, error)
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
	if err := r.db.First(&wallet, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) UserExists(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *walletRepository) Update(wallet *entity.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *walletRepository) CreateTransactionRecord(walletID, categoryID uint, amount float64, transactionType, description string) error {
	record := entity.Record{
		WalletID:    int(walletID),
		CategoryID:  int(categoryID),
		Amount:      amount,
		Type:        transactionType,
		Description: description,
	}
	return r.db.Create(&record).Error
}

func (r *walletRepository) FindAllByUserID(userID uint) ([]entity.Wallet, error) {
	var wallets []entity.Wallet
	err := r.db.Where("user_id = ?", userID).Find(&wallets).Error
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (r *walletRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Wallet{}, id).Error
}

// Tambahkan metode lain sesuai kebutuhan

func (r *walletRepository) TransferBetweenWallets(fromWalletID, toWalletID, categoryID, userID int, amount float64) error {
	// Mulai transaksi
	tx := r.db.Begin()

	// Cari wallet pengirim
	var fromWallet entity.Wallet
	if err := tx.First(&fromWallet, fromWalletID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("from_wallet not found: %v", err)
	}

	// Cari wallet penerima
	var toWallet entity.Wallet
	if err := tx.First(&toWallet, toWalletID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("to_wallet not found: %v", err)
	}

	// Periksa apakah ada cukup saldo di wallet pengirim
	if fromWallet.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance")
	}

	// Kurangi saldo dari wallet pengirim
	fromWallet.Balance -= amount
	if err := tx.Save(&fromWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update from_wallet: %v", err)
	}

	// Tambah saldo ke wallet penerima
	toWallet.Balance += amount
	if err := tx.Save(&toWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update to_wallet: %v", err)
	}

	// Catat record pengeluaran untuk wallet pengirim
	recordOut := entity.Record{
		WalletID:    fromWalletID,
		CategoryID:  categoryID,
		Amount:      amount,
		Type:        "expense",
		Description: fmt.Sprintf("Transferred to wallet %d", toWalletID),
	}
	if err := tx.Create(&recordOut).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create expense record: %v", err)
	}

	// Catat record pemasukan untuk wallet penerima
	recordIn := entity.Record{
		WalletID:    toWalletID,
		CategoryID:  categoryID,
		Amount:      amount,
		Type:        "income",
		Description: fmt.Sprintf("Received from wallet %d", fromWalletID),
	}
	if err := tx.Create(&recordIn).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create income record: %v", err)
	}

	// Commit transaksi
	return tx.Commit().Error
}

func (r *walletRepository) IsCategoryOwnedByUser(categoryID, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entity.TransactionCategory{}).Where("id = ? AND user_id = ?", categoryID, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *walletRepository) GetRecordsBetweenTimes(walletID uint, startTime, endTime time.Time) ([]entity.Record, error) {
	var records []entity.Record
	err := r.db.Where("wallet_id = ? AND created_at BETWEEN ? AND ?", walletID, startTime, endTime).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *walletRepository) GetCashFlow(startTime, endTime time.Time) (float64, float64, error) {
	var totalIncome, totalExpense float64

	// Hitung total income
	err := r.db.Model(&entity.Record{}).
		Where("type = ? AND created_at BETWEEN ? AND ?", "income", startTime, endTime).
		Select("SUM(amount)").
		Scan(&totalIncome).Error
	if err != nil {
		return 0, 0, err
	}

	// Hitung total expense
	err = r.db.Model(&entity.Record{}).
		Where("type = ? AND created_at BETWEEN ? AND ?", "expense", startTime, endTime).
		Select("SUM(amount)").
		Scan(&totalExpense).Error
	if err != nil {
		return 0, 0, err
	}

	return totalIncome, totalExpense, nil
}

func (r *walletRepository) GetExpenseRecapByCategory(startTime, endTime time.Time) ([]entity.CategoryExpenseRecap, error) {
	var recaps []entity.CategoryExpenseRecap

	err := r.db.Table("records").
		Select("transaction_categories.id as category_id, transaction_categories.name as category, SUM(records.amount) as total").
		Joins("left join transaction_categories on records.category_id = transaction_categories.id").
		Where("records.type = ? AND records.created_at BETWEEN ? AND ?", "expense", startTime, endTime).
		Group("transaction_categories.id, transaction_categories.name").
		Scan(&recaps).Error

	if err != nil {
		return nil, err
	}

	return recaps, nil
}

func (r *walletRepository) GetLastRecords(limit int) ([]entity.Record, error) {
	var records []entity.Record

	err := r.db.Order("created_at DESC").Limit(limit).Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}
