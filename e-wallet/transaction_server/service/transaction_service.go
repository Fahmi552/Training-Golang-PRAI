package service

import (
	"Training/e-wallet/entity"
	"context"
	"fmt"
)

type InterTransService interface {
	CreateTransaction_service(ctx context.Context, trans *entity.Transaction) (entity.Transaction, error)
}

type InterTransRepository interface {
	CreateTransaction_repo(ctx context.Context, trans *entity.Transaction) (entity.Transaction, error)
}

type transService struct {
	transRepo InterTransRepository
}

// NewTransService membuat instance baru dari TransService
func NewTransService(transRepo InterTransRepository) InterTransService {
	return &transService{transRepo: transRepo}
}

func (t *transService) CreateTransaction_service(ctx context.Context, trans *entity.Transaction) (entity.Transaction, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	createdTransaction, err := t.transRepo.CreateTransaction_repo(ctx, trans)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("gagal melakukan transaksi: %v", err)
	}
	return createdTransaction, nil
}
