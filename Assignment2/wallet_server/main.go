package main

import (
	grpcHandler "Training/Assignment2/wallet_server/handler/grpc"
	"Training/Assignment2/wallet_server/service"
	"log"
	"net"

	//pb2 "Training/Assignment2/wallet_server/proto/transaction_service/v1"
	pb "Training/Assignment2/wallet_server/proto/wallet_service/v1"
	"Training/Assignment2/wallet_server/repository/postgres_gorm"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// setup gorm connection
	dsn := "postgresql://postgres:admin@localhost:5432/Assignment2"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	// setup wallet
	walletRepo := postgres_gorm.NewWalletRepository(gormDB)
	walletService := service.NewWalletService(walletRepo)
	walletHandler := grpcHandler.NewWalletHandler(walletService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterWalletServiceServer(grpcServer, walletHandler)
	//pb2.RegisterTransactionServiceServer(grpcServer, transactionHandler)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50052")
	_ = grpcServer.Serve(lis)

	// // Run the grpc gateway
	// conn, err := grpc.NewClient(
	// 	"0.0.0.0:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalln("Failed to dial server:", err)
	// }
	// gwmux := runtime.NewServeMux()
	// if err = pb.RegisterWalletServiceHandler(context.Background(), gwmux, conn); err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// // dengan GIN
	// gwServer := gin.Default()
	// gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	// log.Println("Running grpc gateway server in port :8080")
	// _ = gwServer.Run()
}
