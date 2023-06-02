package main

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/controller"
	"github.com/transaction-organizer/db"
)

func main() {
	// Init DB connection
	db.PostgreSqlDB{}.Connect()

	// Init REST API endpoints
	router := gin.Default()
	initTransactionMethods(router)
	initTransactionTypeMethods(router)
	initReportMethods(router)
	router.Run("localhost:8080")
}

func initTransactionTypeMethods(router *gin.Engine) {
	router.GET("/transaction-types", func(c *gin.Context) {
		controller.TransactionTypeController{}.GetTransactionTypes(c)
	})
	router.POST("/transaction-types", func(c *gin.Context) {
		controller.TransactionTypeController{}.PostTransactionType(c)
	})
}

func initTransactionMethods(router *gin.Engine) {
	router.GET("/transactions", func(c *gin.Context) {
		controller.TransactionController{}.GetTransactions(c)
	})
	router.GET("/transactions/unorganized", func(c *gin.Context) {
		controller.TransactionController{}.GetUnorganizedTransactions(c)
	})
	router.POST("/transactions/kh", func(c *gin.Context) {
		controller.TransactionController{}.PostKhTransaction(c)
	})
	router.PATCH("/transactions/:transactionId/type/:typeId", func(c *gin.Context) {
		controller.TransactionController{}.PatchTransaction(c)
	})
}

func initReportMethods(router *gin.Engine) {
	router.GET("/report/group-summary", func(c *gin.Context) {
		controller.ReportController{}.GetGroupSummaryReport(c)
	})
	router.GET("/report/transaction-type-summary", func(c *gin.Context) {
		controller.ReportController{}.GetSummaryReportByTransactionType(c)
	})
}
