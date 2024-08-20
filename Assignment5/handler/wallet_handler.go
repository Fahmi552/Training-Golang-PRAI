package handler

import (
	"Training/Assignment5/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletService services.WalletService
}

func NewWalletHandler(walletService services.WalletService) *WalletHandler {
	return &WalletHandler{walletService}
}

func (h *WalletHandler) CreateWallet(c *gin.Context) {
	var input struct {
		UserID  uint    `json:"user_id"`
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}
	// Bind JSON input ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		// Jika JSON tidak valid, kembalikan status 400 dengan pesan error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi apakah user_id ada di database
	userExists, err := h.walletService.IsUserExists(input.UserID)
	if err != nil {
		// Jika ada error saat validasi user, kembalikan status 500 dengan pesan error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !userExists {
		// Jika user_id tidak ditemukan, kembalikan status 404 dengan pesan error
		c.JSON(http.StatusNotFound, gin.H{"message": "Tidak ada user id itu coy"})
		return
	}

	// Jika user_id valid, buat wallet baru
	_, err = h.walletService.CreateWallet(input.UserID, input.Name, input.Balance)
	if err != nil {
		// Jika ada error saat membuat wallet, kembalikan status 500 dengan pesan error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika berhasil, kembalikan status 201 dengan pesan sukses
	c.JSON(http.StatusCreated, gin.H{"message": "Wallet created successfully"})
}

func (h *WalletHandler) GetWalletByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	wallet, err := h.walletService.GetWalletByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if wallet == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}
	c.JSON(http.StatusOK, wallet)
}

func (h *WalletHandler) GetWalletsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Validasi apakah user_id ada
	userExists, err := h.walletService.IsUserExists(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User ID not found"})
		return
	}

	// Ambil semua wallet milik user dengan user_id tertentu
	wallets, err := h.walletService.GetWalletsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wallets)
}

func (h *WalletHandler) UpdateWallet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	// Ambil userID dari query parameter
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Validasi apakah user_id ada
	userExists, err := h.walletService.IsUserExists(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User ID not found"})
		return
	}

	var input struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update wallet jika user_id valid
	wallet, err := h.walletService.UpdateWallet(uint(id), input.Name, input.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Wallet updated successfully",
		"wallet_name": wallet.Name,
		"balance":     wallet.Balance,
	})
}

func (h *WalletHandler) DeleteWallet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	// Ambil userID dari query parameter
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Validasi apakah user_id ada
	userExists, err := h.walletService.IsUserExists(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User ID not found"})
		return
	}

	// Hapus wallet jika user_id valid
	err = h.walletService.DeleteWallet(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Wallet deleted successfully"})
}

func (h *WalletHandler) Transfer(c *gin.Context) {
	fromWalletID, err := strconv.Atoi(c.Query("from_wallet_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid from_wallet_id"})
		return
	}

	toWalletID, err := strconv.Atoi(c.Query("to_wallet_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid to_wallet_id"})
		return
	}

	categoryID, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id"})
		return
	}

	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	amount, err := strconv.ParseFloat(c.Query("amount"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	// Validasi keberadaan wallet dan user
	err = h.walletService.Transfer(uint(fromWalletID), uint(toWalletID), uint(categoryID), uint(userID), amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}

func (h *WalletHandler) GetRecordsBetweenTimes(c *gin.Context) {
	walletID, err := strconv.Atoi(c.Query("wallet_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet_id"})
		return
	}

	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	// Parsing tanggal dengan format YYYY-MM-DD
	startTime, err := time.Parse("2006-01-02", startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTime, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	// Sesuaikan waktu akhir agar mencakup seluruh hari
	endTime = endTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	records, err := h.walletService.GetRecordsBetweenTimes(uint(walletID), startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func (h *WalletHandler) GetCashFlow(c *gin.Context) {
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	// Parsing tanggal dengan format YYYY-MM-DD
	startTime, err := time.Parse("2006-01-02", startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTime, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	// Sesuaikan waktu akhir agar mencakup seluruh hari
	endTime = endTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	income, expense, err := h.walletService.GetCashFlow(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_income":  income,
		"total_expense": expense,
	})
}

func (h *WalletHandler) GetExpenseRecapByCategory(c *gin.Context) {
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	// Parsing tanggal dengan format YYYY-MM-DD
	startTime, err := time.Parse("2006-01-02", startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTime, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	// Sesuaikan waktu akhir agar mencakup seluruh hari
	endTime = endTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	recaps, err := h.walletService.GetExpenseRecapByCategory(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recaps)
}

func (h *WalletHandler) GetLastRecords(c *gin.Context) {
	limit := 10 // Set limit untuk mengambil 10 record terakhir

	records, err := h.walletService.GetLastRecords(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}
