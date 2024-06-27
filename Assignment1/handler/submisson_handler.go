package handler

import (
	"Training/Assignment1/entity"
	"Training/Assignment1/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ISubmissionHandler mendefinisikan interface untuk handler submission
type ISubmissionHandler interface {
	CreateSubmission_hander(c *gin.Context)
	GetSubmissionByID_hander(c *gin.Context)
	GetSubmissionByUserID_hander(c *gin.Context)
	GetSubmission(c *gin.Context)
	GetAllSubmissions_handler(c *gin.Context)
	DeleteSubmission_handler(c *gin.Context)
}

// NewSubmissionHandler membuat instance baru dari SubmissionHandler
func NewSubmissionHandler(submissionService service.ISubmissionService) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

type SubmissionHandler struct {
	submissionService service.ISubmissionService
}

// get submission by id
func (s *SubmissionHandler) GetSubmissionByID_hander(c *gin.Context) {
	// Ambil ID dari URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		// Jika ID tidak valid, kembalikan respon 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Panggil service untuk mendapatkan submission berdasarkan ID
	submission, err := s.submissionService.GetSubmissionByID_service(c.Request.Context(), id)
	if err != nil {
		// Tangani kasus ketika submission tidak ditemukan
		if err.Error() == fmt.Sprintf("submission dengan ID %d tidak ditemukan", id) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			// Tangani error lainnya dengan status 500 Internal Server Error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan submission", "details": err.Error()})
		}
		return
	}

	// Kembalikan data submission dengan status 200 OK
	c.JSON(http.StatusOK, gin.H{"submission": submission})
}

// get submission by user id
func (s *SubmissionHandler) GetSubmissionByUserID_hander(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	log.Print(userID)
	log.Print(err)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid euy"})
		return
	}

	submissions, err := s.submissionService.GetSubmissionByUserID_service(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan submissions", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"submissions": submissions})
}

// CreateSubmission adalah handler untuk membuat submission baru
func (s *SubmissionHandler) CreateSubmission_hander(c *gin.Context) {
	var submission entity.SubmissionRequest

	// Bind JSON request body ke struct Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Panggil service untuk membuat submission
	if err := s.submissionService.CreateSubmission_service(c.Request.Context(), &submission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create submission", "details": err.Error()})
		return
	}

	// Berhasil membuat submission, kirim respons berhasil
	c.JSON(http.StatusCreated, gin.H{"message": "Submission created successfully"})

}

func (s *SubmissionHandler) GetSubmission(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *SubmissionHandler) DeleteSubmission_handler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada ID itu coy, yg bener ajeee"})
		return
	}

	// Panggil service untuk menghapus
	err = s.submissionService.DeleteSubmission_service(c.Request.Context(), id)
	if err != nil {
		if err.Error() == fmt.Sprintf("tidak ada submission dengan ID %d yang ditemukan", id) {
			// Respons jika pengguna tidak ditemukan
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			// Respons jika ada kesalahan lain
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// ini jika sukses
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Submission with ID %d deleted successfully", id)})
}

func (s *SubmissionHandler) GetAllSubmissions_handler(c *gin.Context) {
	submissions, err := s.submissionService.GetAllSubmissions_service(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan submissions", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"submissions": submissions})
}
