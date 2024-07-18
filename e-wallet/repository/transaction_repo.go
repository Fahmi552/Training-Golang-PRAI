package repository

import (
	"Training/e-wallet/entity"
	"Training/e-wallet/transaction_server/service"
	"context"
	"log"
)

type transaction_Repository struct {
	db GormDBIface
}

func NewTransactionRepository(db GormDBIface) service.InterTransRepository {
	return &transaction_Repository{db: db}
}

func (r *transaction_Repository) CreateTransaction_repo(ctx context.Context, trans *entity.Transaction) (entity.Transaction, error) {
	if err := r.db.WithContext(ctx).Create(trans).Error; err != nil {
		log.Printf("Error creating trans: %v\n", err)
		return entity.Transaction{}, err
	}
	return *trans, nil
}

// -- DISINI HARUS ADA GET TRANSACTION BY USER ID
// func (r *transaction_Repository) GetTransactionsByUserID(userID string) ([]models.Transaction, error) {
// 	var transactions []models.Transaction
// 	result := r.DB.Where("user_id = ?", userID).Find(&transactions)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return transactions, nil
// }
