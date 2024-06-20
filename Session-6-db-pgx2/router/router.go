package router

import (
	"Training/Session-6-db-pgx2/handler"
	"Training/Session-6-db-pgx2/middleware"

	"github.com/gin-gonic/gin"
)

// func SetupRouter(r *gin.Engine) {
// 	usersPublicEndpoint := r.Group("/users")
// 	usersPublicEndpoint.GET("/:id", handler.GetUser)
// 	usersPublicEndpoint.GET("/", handler.GetAllUsers)
// 	usersPublicEndpoint.GET("/ALL/:name", handler.GetUserByName)
// 	usersPrivateEndpoint := r.Group("/users")
// 	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
// 	usersPrivateEndpoint.POST("/", handler.CreateUser)
// 	usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
// 	usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
// }

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	// Mengatur endpoint publik untuk pengguna
	usersPublicEndpoint := r.Group("/users")
	// Rute untuk mendapatkan pengguna berdasarkan ID
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	// Rute untuk mendapatkan semua pengguna
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)

	// Mengatur endpoint privat untuk pengguna dengan middleware autentikasi
	usersPrivateEndpoint := r.Group("/users")
	// Menambahkan middleware autentikasi untuk endpoint privat
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	// Rute untuk membuat pengguna baru
	usersPrivateEndpoint.POST("", userHandler.CreateUser)
	usersPrivateEndpoint.POST("/", userHandler.CreateUser)
	// Rute untuk memperbarui pengguna berdasarkan ID
	usersPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	// Rute untuk menghapus pengguna berdasarkan ID
	usersPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
}
