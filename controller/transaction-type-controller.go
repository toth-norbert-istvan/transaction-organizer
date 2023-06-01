package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/dto"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"net/http"
)

func GetTransactionTypes(c *gin.Context, db *sql.DB) {
	transactionTypes := repository.GetTransactionTypes(db)
	c.IndentedJSON(http.StatusOK, mapper.TransactionTypeMapper{}.DomainsToDtos(transactionTypes))
}

func PostTransactionType(c *gin.Context, db *sql.DB) {
	var request dto.NewTransactionType
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newTransactionType, isBadRequest := mapper.TransactionTypeMapper{}.DtoToDomain(request)
	if isBadRequest {
		c.Status(http.StatusBadRequest)
		return
	}

	savedTransactionType := repository.SaveTransactionType(db, newTransactionType)
	c.IndentedJSON(http.StatusOK, mapper.TransactionTypeMapper{}.DomainToDto(savedTransactionType))
}
