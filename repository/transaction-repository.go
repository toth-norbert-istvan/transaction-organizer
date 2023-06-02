package repository

import (
	"github.com/transaction-organizer/db"
	"github.com/transaction-organizer/domain"
	"log"
)

type TransactionRepository struct{}

func (tr TransactionRepository) GetTransactions() []domain.Transaction {
	var db = db.PostgreSqlDB{}.GetDb()

	rows, err := db.Query("SELECT * FROM transactions")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var transaction domain.Transaction
	var transactions []domain.Transaction
	for rows.Next() {
		rows.Scan(&transaction.Id, &transaction.Partner, &transaction.Amount, &transaction.Date, &transaction.TransactionTypeId)
		transactions = append(transactions, transaction)
	}
	return transactions
}

func (tr TransactionRepository) SaveTransactions(transactions []domain.Transaction) {
	var db = db.PostgreSqlDB{}.GetDb()

	for _, transaction := range transactions {
		_, err := db.Exec("INSERT INTO transactions (partner, amount, date) VALUES ($1, $2, $3)", transaction.Partner, transaction.Amount, transaction.Date)

		if err != nil {
			log.Fatalf("An error occured while executing transaction saving: %v", err)
		}
	}
}
