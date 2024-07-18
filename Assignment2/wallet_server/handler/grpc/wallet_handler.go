package grpc

import (
	"Training/Assignment2/wallet_server/entity"
	pb "Training/Assignment2/wallet_server/proto/wallet_service/v1"
	"Training/Assignment2/wallet_server/service"
	"context"
	"fmt"
	"log"
)

// WalletHandler is used to implement UnimplementedWalletServiceServer
type WalletHandler struct {
	pb.UnimplementedWalletServiceServer
	walletService service.InterWalletService
}

// membuat instance baru dari WalletHandler
func NewWalletHandler(walletService service.InterWalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func (w *WalletHandler) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.MutationResponse, error) {
	createdWallet, err := w.walletService.CreateWallet_service(ctx, &entity.Wallet{
		UserID:  int(req.UserID),      //int32(req.GetUserID()),
		Balance: float64(req.Balance), //float32(req.GetBalance()),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success created wallet with user ID %d", createdWallet.UserID),
	}, nil
}
