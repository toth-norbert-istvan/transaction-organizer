package repository

import (
	"database/sql"
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

	return mapRowsToTransactions(rows)
}

func (tr TransactionRepository) GetUnorganizedTransactions() []domain.Transaction {
	var db = db.PostgreSqlDB{}.GetDb()

	rows, err := db.Query("SELECT * FROM transactions t WHERE t.transaction_type_id IS NULL")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	return mapRowsToTransactions(rows)
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

func (tr TransactionRepository) UpdateTransactionTypeById(transactionId int, transactionTypeId int) error {
	var db = db.PostgreSqlDB{}.GetDb()

	result, err := db.Exec("UPDATE transactions SET transaction_type_id =$1 WHERE id = $2", transactionTypeId, transactionId)
	result.RowsAffected()
	return err
}

func mapRowsToTransactions(rows *sql.Rows) []domain.Transaction {
	var transaction domain.Transaction
	var transactions []domain.Transaction
	for rows.Next() {
		rows.Scan(&transaction.Id, &transaction.Partner, &transaction.Amount, &transaction.Date, &transaction.TransactionTypeId)
		transactions = append(transactions, transaction)
	}
	return transactions
}
