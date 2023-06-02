package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"github.com/transaction-organizer/service"
	"net/http"
)

type TransactionController struct{}

func (tc TransactionController) GetTransactions(c *gin.Context, db *sql.DB) {
	transactions := repository.TransactionRepository{}.GetTransactions(db)
	c.IndentedJSON(http.StatusOK, mapper.TransactionMapper{}.DomainsToDtos(transactions, db))
}

func (tc TransactionController) PostKhTransaction(c *gin.Context, db *sql.DB) {
	var file, _ = c.FormFile("file")
	newTransactions := service.KhFileParserService{}.GetTransactionsFromExcelFile(file)
	repository.TransactionRepository{}.SaveTransactions(db, newTransactions)
	c.Status(http.StatusCreated)
}

func (tc TransactionController) PatchTransaction(c *gin.Context, db *sql.DB) {
	c.Status(http.StatusOK)
}
