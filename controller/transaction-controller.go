package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"github.com/transaction-organizer/service"
	"log"
	"net/http"
	"strconv"
)

type TransactionController struct{}

func (tc TransactionController) GetTransactions(c *gin.Context) {
	transactions := repository.TransactionRepository{}.GetTransactions()
	c.IndentedJSON(http.StatusOK, mapper.TransactionMapper{}.DomainsToDtos(transactions))
}

func (tc TransactionController) GetUnorganizedTransactions(c *gin.Context) {
	transactions := repository.TransactionRepository{}.GetUnorganizedTransactions()
	c.IndentedJSON(http.StatusOK, mapper.TransactionMapper{}.DomainsToDtos(transactions))
}

func (tc TransactionController) PostKhTransaction(c *gin.Context) {
	var file, _ = c.FormFile("file")
	newTransactions := service.KhFileParserService{}.GetTransactionsFromExcelFile(file)
	repository.TransactionRepository{}.SaveTransactions(newTransactions)
	c.Status(http.StatusCreated)
}

func (tc TransactionController) PatchTransaction(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		log.Println("Invalid transaction id: ", c.Param("transactionId"))
		c.Status(http.StatusBadRequest)
		return
	}

	transactionTypeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		log.Println("Invalid transaction type id: ", c.Param("typeId"))
		c.Status(http.StatusBadRequest)
		return
	}

	_, err = repository.TransactionTypeRepository{}.GetTransactionType(transactionTypeId)
	if err != nil {
		log.Println("Transaction type does not exist with id: ", transactionTypeId)
		c.Status(http.StatusBadRequest)
		return
	}

	err = repository.TransactionRepository{}.UpdateTransactionTypeById(transactionId, transactionTypeId)
	if err != nil {
		log.Println("Transaction update failed: ", err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
