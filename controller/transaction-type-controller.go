package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/dto"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"net/http"
)

type TransactionTypeController struct{}

func (ttc TransactionTypeController) GetTransactionTypes(c *gin.Context) {
	transactionTypes := repository.TransactionTypeRepository{}.GetTransactionTypes()
	c.IndentedJSON(http.StatusOK, mapper.TransactionTypeMapper{}.DomainsToDtos(transactionTypes))
}

func (ttc TransactionTypeController) PostTransactionType(c *gin.Context) {
	var request dto.NewTransactionType
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newTransactionType, isBadRequest := mapper.TransactionTypeMapper{}.DtoToDomain(request)
	if isBadRequest {
		c.Status(http.StatusBadRequest)
		return
	}

	savedTransactionType := repository.TransactionTypeRepository{}.SaveTransactionType(newTransactionType)
	c.IndentedJSON(http.StatusOK, mapper.TransactionTypeMapper{}.DomainToDto(savedTransactionType))
}
