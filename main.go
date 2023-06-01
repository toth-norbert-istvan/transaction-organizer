package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/controller"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/transaction-organizer?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	initTransactionMethods(router, db)
	initTransactionTypeMethods(router, db)
	router.Run("localhost:8080")
}

func initTransactionTypeMethods(router *gin.Engine, db *sql.DB) {
	router.GET("/transaction-types", func(c *gin.Context) {
		controller.GetTransactionTypes(c, db)
	})
	router.POST("/transaction-types", func(c *gin.Context) {
		controller.PostTransactionType(c, db)
	})
}

func initTransactionMethods(router *gin.Engine, db *sql.DB) {
	router.GET("/transactions", func(c *gin.Context) {
		controller.GetTransactions(c, db)
	})
	router.POST("/transactions/kh", func(c *gin.Context) {
		controller.PostKhTransaction(c, db)
	})
	router.PATCH("/transactions/:transactionId", func(c *gin.Context) {
		controller.PatchTransaction(c, db)
	})
}
