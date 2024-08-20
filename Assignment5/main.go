package main

import (
	// "Assignment5/config"
	"Training/Assignment5/config"

	"Training/Assignment5/entity"
	"Training/Assignment5/handler"
	"Training/Assignment5/repositories"
	"Training/Assignment5/services"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate
	err = db.AutoMigrate(&entity.Wallet{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Setup repository, service, and handler
	walletRepo := repositories.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	// Setup repository, service, dan handler untuk TransactionCategory
	categoryRepo := repositories.NewTransactionCategoryRepository(db)
	categoryService := services.NewTransactionCategoryService(categoryRepo)
	categoryHandler := handler.NewTransactionCategoryHandler(categoryService)

	// Setup repository, service, dan handler untuk Record
	recordRepo := repositories.NewRecordRepository(db)
	recordService := services.NewRecordService(recordRepo)
	recordHandler := handler.NewRecordHandler(recordService)

	// Wallet routes
	r.POST("/wallets", walletHandler.CreateWallet)
	r.GET("/wallets/:id", walletHandler.GetWalletByID)
	r.GET("/wallets", walletHandler.GetWalletsByUserID)
	r.PUT("/wallets/:id", walletHandler.UpdateWallet)
	r.DELETE("/wallets/:id", walletHandler.DeleteWallet)

	//Transfer
	r.POST("/wallets/transfer", walletHandler.Transfer)

	//get record
	r.GET("/wallets/records", walletHandler.GetRecordsBetweenTimes)

	//get cash flow
	r.GET("/wallets/cashflow", walletHandler.GetCashFlow)

	//get expence
	r.GET("/wallets/expense_recap", walletHandler.GetExpenseRecapByCategory)

	//get 10 last record
	r.GET("/wallets/last_records", walletHandler.GetLastRecords)

	// Transaction Category routes
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories/:id", categoryHandler.GetCategoryByID)
	r.GET("/categories", categoryHandler.GetCategoriesByUserID)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Record routes
	r.POST("/records", recordHandler.CreateRecord)
	r.GET("/records/:id", recordHandler.GetRecordByID)
	//r.GET("/wallets/:wallet_id/records", recordHandler.GetRecordsByWalletID)
	r.PUT("/records/:id", recordHandler.UpdateRecord)
	r.DELETE("/records/:id", recordHandler.DeleteRecord)

	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
