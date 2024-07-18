package handler

import (
	"Training/e-wallet/entity"
	"Training/e-wallet/transaction_server/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InterTransHandler interface {
	CreatedTransaction_handler(c *gin.Context)
}

type TransHandler struct {
	transService service.InterTransService
}

// func NewUserHandler(userService service.IUserService) IUserHandler {
// 	return &UserHandler{
// 		userService: userService,
// 	}
// }

func NewTransHandler(TransService service.InterTransService) InterTransHandler {
	return &TransHandler{
		transService: TransService,
	}
}

func (t *TransHandler) CreatedTransaction_handler(c *gin.Context) {
	var trans entity.Transaction
	if err := c.ShouldBindJSON(&trans); err != nil {
		errMsg := err.Error()
		//custom error disini
		//errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	createdTransaction, err := t.transService.CreateTransaction_service(c.Request.Context(), &trans)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseSukses := gin.H{
		"message": fmt.Sprintf("Transaction ID %d with Type %s created successfully", createdTransaction.ID, createdTransaction.Type)}
	c.JSON(http.StatusCreated, responseSukses)
}
