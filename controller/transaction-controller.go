package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	transaction_organizer_domain_service "gitlab.com/transaction-organizer/service"
	"net/http"
)

func GetTransactions(c *gin.Context, db *sql.DB) {
	c.IndentedJSON(http.StatusOK, transaction_organizer_domain_service.GetTransactions(db))
}

func PostKhTransaction(c *gin.Context, db *sql.DB) {
	var file, _ = c.FormFile("file")
	newTransactions := transaction_organizer_domain_service.GetTransactionsFromExcelFile(file)
	transaction_organizer_domain_service.SaveTransactions(db, newTransactions)
	c.Status(http.StatusCreated)
}
