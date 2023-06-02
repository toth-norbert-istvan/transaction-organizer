package repository

import (
	"database/sql"
	"github.com/transaction-organizer/db"
	"github.com/transaction-organizer/domain"
	"log"
)

type TransactionRepository struct{}

func (tr TransactionRepository) GetTransactions(db *sql.DB) []domain.Transaction {
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

func (tr TransactionRepository) GetTransactionsFromDB() []domain.Transaction {
	db.Connect()
	return nil
}

func (tr TransactionRepository) SaveTransactions(db *sql.DB, transactions []domain.Transaction) {
	for _, transaction := range transactions {
		_, err := db.Exec("INSERT INTO transactions (partner, amount, date) VALUES ($1, $2, $3)", transaction.Partner, transaction.Amount, transaction.Date)

		if err != nil {
			log.Fatalf("An error occured while executing transaction saving: %v", err)
		}
	}
}
