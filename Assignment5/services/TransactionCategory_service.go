package services

import (
	"Training/Assignment5/entity"
	"Training/Assignment5/repositories"
)

type TransactionCategoryService interface {
	CreateCategory(userID uint, name string) (*entity.TransactionCategory, error)
	GetCategoryByID(id uint) (*entity.TransactionCategory, error)
	GetCategoriesByUserID(userID uint) ([]entity.TransactionCategory, error)
	UpdateCategory(id uint, name string) (*entity.TransactionCategory, error)
	DeleteCategory(id uint) error
	IsUserExists(userID uint) (bool, error)
}

type transactionCategoryService struct {
	repo repositories.TransactionCategoryRepository
}

func NewTransactionCategoryService(repo repositories.TransactionCategoryRepository) TransactionCategoryService {
	return &transactionCategoryService{repo}
}

func (s *transactionCategoryService) IsUserExists(userID uint) (bool, error) {
	var count int64
	err := s.repo.CountUserByID(userID, &count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *transactionCategoryService) CreateCategory(userID uint, name string) (*entity.TransactionCategory, error) {
	category := &entity.TransactionCategory{
		UserID: int(userID),
		Name:   name,
	}
	err := s.repo.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *transactionCategoryService) GetCategoryByID(id uint) (*entity.TransactionCategory, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *transactionCategoryService) GetCategoriesByUserID(userID uint) ([]entity.TransactionCategory, error) {
	return s.repo.GetCategoriesByUserID(userID)
}

func (s *transactionCategoryService) UpdateCategory(id uint, name string) (*entity.TransactionCategory, error) {
	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	category.Name = name
	err = s.repo.UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *transactionCategoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}
