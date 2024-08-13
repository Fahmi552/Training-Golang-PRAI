package repositories

import (
	"gorm.io/gorm"
)

type TransferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	return &TransferRepository{db}
}

// func (r *TransferRepository) Create(transfer *entity.Transfer) error {
// 	return r.db.Create(transfer).Error
// }

// func (r *TransferRepository) FindByID(id int) (*models.Transfer, error) {
// 	var transfer models.Transfer
// 	err := r.db.First(&transfer, id).Error
// 	return &transfer, err
// }

// Tambahkan metode lain sesuai kebutuhan
