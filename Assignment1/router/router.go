package router

import (
	"Training/Assignment1/handler"
	"Training/Assignment1/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi dan mengatur rute untuk aplikasi
func SetupRouter(r *gin.Engine,
	userHandler handler.IUserHandler,
	submissionsHandler handler.ISubmissionHandler) {
	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.GET("getuser/:id", userHandler.GetUserWithLatestSubmission_handler)
	// usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	// usersPublicEndpoint.GET("/", userHandler.GetAllUsers)
	// usersPublicEndpoint.POST("", userHandler.CreateUser)
	usersPrivateEndpoint.POST("/create_user", userHandler.CreateUser_handler)
	usersPrivateEndpoint.PUT("/update_user/:id", userHandler.UpdateUser_handler)
	usersPrivateEndpoint.DELETE("/delete_user/:id", userHandler.DeleteUser_handler)

	submissionsPrivateEndpoint := r.Group("/submissions")
	submissionsPrivateEndpoint.Use(middleware.AuthMiddleware())
	submissionsPrivateEndpoint.GET("/submissionsGetByID/:id", submissionsHandler.GetSubmissionByID_hander)
	submissionsPrivateEndpoint.GET("/submissionsGetByUserID/:user_id", submissionsHandler.GetSubmissionByUserID_hander)
	submissionsPrivateEndpoint.GET("/submissionsGetAll/", submissionsHandler.GetAllSubmissions_handler)
	submissionsPrivateEndpoint.POST("/create_submission", submissionsHandler.CreateSubmission_hander)
	// submissionsPublicEndpoint.POST("/", submissionsHandler.CreateSubmission)
	submissionsPrivateEndpoint.DELETE("/delete_submission/:id", submissionsHandler.DeleteSubmission_handler)
}
