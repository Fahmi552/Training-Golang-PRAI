package services

type Inter_WalletRepo interface {
	TopUp(userID string, amount float64)
}

// type WalletService struct {
// 	Repo *repositories.WalletRepository
// }

// func NewWalletService(repo *repositories.WalletRepository) *WalletService {
// 	return &WalletService{Repo: repo}
// }

// func (s *WalletService) TopUp(userID string, amount float64) error {
// 	return s.Repo.TopUp(userID, amount)
// }
