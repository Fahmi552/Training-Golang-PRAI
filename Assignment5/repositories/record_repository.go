package repositories

import (
	"Training/Assignment5/entity"

	"gorm.io/gorm"
)

type RecordRepository interface {
	Create(record *entity.Record) error
	FindByID(id int) (*entity.Record, error)
	FindAllByWalletID(walletID int) ([]entity.Record, error)
	Update(record *entity.Record) error
	Delete(id int) error
}

type recordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db}
}

func (r *recordRepository) Create(record *entity.Record) error {
	return r.db.Create(record).Error
}

func (r *recordRepository) FindByID(id int) (*entity.Record, error) {
	var record entity.Record
	err := r.db.First(&record, id).Error
	return &record, err
}

func (r *recordRepository) FindAllByWalletID(walletID int) ([]entity.Record, error) {
	var records []entity.Record
	err := r.db.Where("wallet_id = ?", walletID).Find(&records).Error
	return records, err
}

func (r *recordRepository) Update(record *entity.Record) error {
	return r.db.Save(record).Error
}

func (r *recordRepository) Delete(id int) error {
	return r.db.Delete(&entity.Record{}, id).Error
}
