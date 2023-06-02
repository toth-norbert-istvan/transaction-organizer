package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"github.com/transaction-organizer/service"
	"net/http"
)

type TransactionController struct{}

func (tc TransactionController) GetTransactions(c *gin.Context) {
	transactions := repository.TransactionRepository{}.GetTransactions()
	c.IndentedJSON(http.StatusOK, mapper.TransactionMapper{}.DomainsToDtos(transactions))
}

func (tc TransactionController) PostKhTransaction(c *gin.Context) {
	var file, _ = c.FormFile("file")
	newTransactions := service.KhFileParserService{}.GetTransactionsFromExcelFile(file)
	repository.TransactionRepository{}.SaveTransactions(newTransactions)
	c.Status(http.StatusCreated)
}

func (tc TransactionController) PatchTransaction(c *gin.Context) {
	c.Status(http.StatusOK)
}
