package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/mapper"
	"github.com/transaction-organizer/repository"
	"github.com/transaction-organizer/service"
	"net/http"
)

func GetTransactions(c *gin.Context, db *sql.DB) {
	transactions := repository.GetTransactions(db)
	c.IndentedJSON(http.StatusOK, mapper.DomainsToDtos(transactions, db))
}

func PostKhTransaction(c *gin.Context, db *sql.DB) {
	var file, _ = c.FormFile("file")
	newTransactions := service.GetTransactionsFromExcelFile(file)
	repository.SaveTransactions(db, newTransactions)
	c.Status(http.StatusCreated)
}

func PatchTransaction(c *gin.Context, db *sql.DB) {
	c.Status(http.StatusOK)
}
