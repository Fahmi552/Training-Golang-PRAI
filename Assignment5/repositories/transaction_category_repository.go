package repositories

import (
	"Training/Assignment5/entity"

	"gorm.io/gorm"
)

type TransactionCategoryRepository interface {
	CreateCategory(category *entity.TransactionCategory) error
	GetCategoryByID(id uint) (*entity.TransactionCategory, error)
	GetCategoriesByUserID(userID uint) ([]entity.TransactionCategory, error)
	UpdateCategory(category *entity.TransactionCategory) error
	DeleteCategory(id uint) error
	CountUserByID(userID uint, count *int64) error
}

type transactionCategoryRepository struct {
	db *gorm.DB
}

func NewTransactionCategoryRepository(db *gorm.DB) TransactionCategoryRepository {
	return &transactionCategoryRepository{db}
}

func (r *transactionCategoryRepository) CreateCategory(category *entity.TransactionCategory) error {
	return r.db.Create(category).Error
}

func (r *transactionCategoryRepository) GetCategoryByID(id uint) (*entity.TransactionCategory, error) {
	var category entity.TransactionCategory
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *transactionCategoryRepository) GetCategoriesByUserID(userID uint) ([]entity.TransactionCategory, error) {
	var categories []entity.TransactionCategory
	err := r.db.Where("user_id = ?", userID).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *transactionCategoryRepository) UpdateCategory(category *entity.TransactionCategory) error {
	return r.db.Save(category).Error
}

func (r *transactionCategoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&entity.TransactionCategory{}, id).Error
}

func (r *transactionCategoryRepository) CountUserByID(userID uint, count *int64) error {
	return r.db.Table("users").Where("id = ?", userID).Count(count).Error
}
