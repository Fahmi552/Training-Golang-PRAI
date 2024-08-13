package services

import (
	"Training/Assignment5/entity"
	"Training/Assignment5/repositories"
)

type RecordService interface {
	CreateRecord(record *entity.Record) error
	GetRecordByID(id int) (*entity.Record, error)
	GetRecordsByWalletID(walletID int) ([]entity.Record, error)
	UpdateRecord(record *entity.Record) error
	DeleteRecord(id int) error
}

type recordService struct {
	repo repositories.RecordRepository
}

func NewRecordService(repo repositories.RecordRepository) RecordService {
	return &recordService{repo}
}

func (s *recordService) CreateRecord(record *entity.Record) error {
	return s.repo.Create(record)
}

func (s *recordService) GetRecordByID(id int) (*entity.Record, error) {
	return s.repo.FindByID(id)
}

func (s *recordService) GetRecordsByWalletID(walletID int) ([]entity.Record, error) {
	return s.repo.FindAllByWalletID(walletID)
}

func (s *recordService) UpdateRecord(record *entity.Record) error {
	return s.repo.Update(record)
}

func (s *recordService) DeleteRecord(id int) error {
	return s.repo.Delete(id)
}
