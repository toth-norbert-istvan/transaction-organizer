package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	transaction_organizer_controller_model "gitlab.com/transaction-organizer/controller/model"
	transaction_organizer_mapper "gitlab.com/transaction-organizer/mapper"
	transaction_organizer_domain_service "gitlab.com/transaction-organizer/service"
	"net/http"
)

func GetTransactionTypes(c *gin.Context, db *sql.DB) {
	transactionTypes := transaction_organizer_domain_service.GetTransactionTypes(db)
	c.IndentedJSON(http.StatusOK, transaction_organizer_mapper.TransactionTypeMapper{}.DomainsToDtos(transactionTypes))
}

func PostTransactionType(c *gin.Context, db *sql.DB) {
	var request transaction_organizer_controller_model.NewTransactionType
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newTransactionType, isBadRequest := transaction_organizer_mapper.TransactionTypeMapper{}.DtoToDomain(request)
	if isBadRequest {
		c.Status(http.StatusBadRequest)
		return
	}

	savedTransactionType := transaction_organizer_domain_service.SaveTransactionType(db, newTransactionType)
	c.IndentedJSON(http.StatusOK, transaction_organizer_mapper.TransactionTypeMapper{}.DomainToDto(savedTransactionType))
}
