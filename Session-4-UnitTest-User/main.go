package main

import (
	// "log"

	// "Training/session-2-latihan-crud-user-gin/router"
	// "Training/Session-4-UnitTest-User/entity"
	// "Training/Session-4-UnitTest-User/repository/slice"
	// "Training/Session-4-UnitTest-User/service"
	// "Training/Session-4-UnitTest-User/handler"
	// "Training/Session-4-UnitTest-User/router"
	// "github.com/gin-gonic/gin"
	// "log"

	"Training/Session-4-UnitTest-User/entity"
	"Training/Session-4-UnitTest-User/handler"
	"Training/Session-4-UnitTest-User/repository/slice"
	"Training/Session-4-UnitTest-User/router"
	"Training/Session-4-UnitTest-User/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
