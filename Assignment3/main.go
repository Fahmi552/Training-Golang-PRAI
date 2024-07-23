package main

import (
	grpcHandler "Training/Assignment3/handler"
	pb "Training/Assignment3/proto"
	"Training/Assignment3/repository"
	"Training/Assignment3/service"
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// set redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// Koneksi database
	dsn := "postgresql://postgres:admin@localhost:5432/Assignment3"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalf("gagal terhubung ke database: %v", err)
	}

	// Inisialisasi repository dan service
	urlRepo := repository.NewURLRepository(gormDB)
	urlService := service.NewURLService(urlRepo, rdb)
	urlHandler := grpcHandler.NewURLHandlerServer(urlService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterURLShortenerServer(grpcServer, urlHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Println("Running grpc server in port :50051")
		_ = grpcServer.Serve(lis)
	}()
	time.Sleep(1 * time.Second)

	// Run the grpc gateway
	// conn, err := grpc.NewClient(
	// 	"0.0.0.0:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterURLShortenerHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// dengan GIN
	// gwServer := gin.Default()
	// //gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	// log.Println("Running grpc gateway server in port :8080")
	// _ = gwServer.Run()

	// Run HTTP server with Gin
	gwServer := gin.Default()
	//gwServer.Any("/*any", gin.WrapH(gwmux))
	gwServer.Any("/v1/*any", gin.WrapH(gwmux))

	// Route untuk redirect
	gwServer.GET("/:short_url", func(c *gin.Context) {
		shortURL := c.Param("short_url")
		originalURL, err := urlService.GetOriginalURL(shortURL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(http.StatusMovedPermanently, originalURL)
	})

	log.Println("Running gRPC gateway server on port :8080")
	if err := gwServer.Run(":8080"); err != nil {
		log.Fatalf("failed to run gin server: %v", err)
	}
}
