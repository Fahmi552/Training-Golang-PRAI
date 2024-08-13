package handler

import (
	"Training/Assignment5/services"
	"net/http"
	"strconv"

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
