package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gitlab.com/transaction-organizer/controller"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/transaction-organizer?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/transactions", func(c *gin.Context) {
		controller.GetTransactions(c, db)
	})
	router.POST("/transactions/kh", func(c *gin.Context) {
		controller.PostKhTransaction(c, db)
	})

	router.Run("localhost:8080")
}
