package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/dto"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"net/http"
)

type TransactionTypeController struct{}

func (ttc TransactionTypeController) GetTransactionTypes(c *gin.Context, db *sql.DB) {
	transactionTypes := repository.TransactionTypeRepository{}.GetTransactionTypes(db)
	c.IndentedJSON(http.StatusOK, mapper.TransactionTypeMapper{}.DomainsToDtos(transactionTypes))
}

func (ttc TransactionTypeController) PostTransactionType(c *gin.Context, db *sql.DB) {
	var request dto.NewTransactionType
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newTransactionType, isBadRequest := mapper.TransactionTypeMapper{}.DtoToDomain(request)
	if isBadRequest {
		c.Status(http.StatusBadRequest)
		return
	}

	savedTransactionType := repository.TransactionTypeRepository{}.SaveTransactionType(db, newTransactionType)
	c.IndentedJSON(http.StatusOK, mapper.TransactionTypeMapper{}.DomainToDto(savedTransactionType))
}
