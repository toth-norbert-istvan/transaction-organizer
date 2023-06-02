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
	router.POST("/transactions/kh", func(c *gin.Context) {
		controller.TransactionController{}.PostKhTransaction(c)
	})
	router.PATCH("/transactions/:transactionId", func(c *gin.Context) {
		controller.TransactionController{}.PatchTransaction(c)
	})
}
