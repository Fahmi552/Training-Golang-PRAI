package handler

import (
	"Training/Assignment1/entity"
	"Training/Assignment1/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	CreateUser_handler(c *gin.Context)
	// GetUser_handler(c *gin.Context)
	UpdateUser_handler(c *gin.Context)
	DeleteUser_handler(c *gin.Context)
	// GetAllUsers_handler(c *gin.Context)
	GetUserWithLatestSubmission_handler(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser_handler(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	createdUser, err := u.userService.CreateUser_service(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseSukses := gin.H{
		"message": fmt.Sprintf("User ID %d with Name %s created successfully", createdUser.ID, createdUser.Name)}
	c.JSON(http.StatusCreated, responseSukses)
}

func (u *UserHandler) DeleteUser_handler(c *gin.Context) {
	// Ambil ID dari parameter URL
	// idParam := c.Param("id")
	// log.Printf("Received ID param: %s", idParam)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada ID itu coy, yg bener ajeee"})
		return
	}

	// if err := u.userService.DeleteUser_service(c.Request.Context(), id); err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }

	// Panggil service untuk menghapus user
	err = u.userService.DeleteUser_service(c.Request.Context(), id)
	if err != nil {
		if err.Error() == fmt.Sprintf("tidak ada pengguna dengan ID %d yang ditemukan", id) {
			// Respons jika pengguna tidak ditemukan
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			// Respons jika ada kesalahan lain
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// ini jika sukses
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %d deleted successfully", id)})
}

func (u *UserHandler) UpdateUser_handler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada user ID spt ituu"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	updatedUser, err := u.userService.UpdateUser_service(c.Request.Context(), id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %d updated successfully", updatedUser.ID)})
}

// func (u *UserHandler) GetUser_handler(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	user, err := u.userService.GetUserByID_service(c.Request.Context(), id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// func (u *UserHandler) GetAllUsers_handler(c *gin.Context) {
// 	//TODO implement me
// 	panic("implement me")
// }

func convertUserMandatoryFieldErrorString(oldErrorMsg string) string {
	switch {
	case strings.Contains(oldErrorMsg, "'Name' failed on the 'required' tag"):
		return "name is mandatory"
	case strings.Contains(oldErrorMsg, "'Email' failed on the 'required' tag"):
		return "email is mandatory"
	}
	return oldErrorMsg
}

// userHandler.go
func (h *UserHandler) GetUserWithLatestSubmission_handler(c *gin.Context) {
	// Ambil ID dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Panggil service untuk mendapatkan data user dan submission terakhir
	userWithRiskProfile, err := h.userService.GetUserWithLatestSubmission_service(c.Request.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data user", "details": err.Error()})
		}
		return
	}

	// Kembalikan data sebagai respon JSON
	c.JSON(http.StatusOK, gin.H{"user": userWithRiskProfile})
}
