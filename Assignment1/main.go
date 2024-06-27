package main

import (
	"Training/Assignment1/handler"
	"Training/Assignment1/middleware"
	postgresgorm "Training/Assignment1/repository/postgres_gorm"
	"Training/Assignment1/router"
	"Training/Assignment1/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Inisialisai router gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// integrasi middleware
	r.Use(middleware.AuthMiddleware())

	// setup gorm connection
	dsn := "postgresql://postgres:admin@localhost:5432/Assignment1"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	// setup repository
	userRepo := postgresgorm.NewUserRepository(gormDB)
	submissionRepo := postgresgorm.NewSubmissionRepository(gormDB)

	// service and handler declaration
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)

	// Routes
	router.SetupRouter(r, userHandler, submissionHandler) //

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
