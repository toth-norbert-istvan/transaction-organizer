package transaction_organizer_domain_service

import (
	"database/sql"
	transaction_organizer_db "gitlab.com/transaction-organizer/db"
	"gitlab.com/transaction-organizer/domain"
	"log"
)

func GetTransactions(db *sql.DB) []transaction_organizer_domain.Transaction {
	rows, err := db.Query("SELECT * FROM transactions")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var transaction transaction_organizer_domain.Transaction
	var transactions []transaction_organizer_domain.Transaction
	for rows.Next() {
		rows.Scan(&transaction.Id, &transaction.Partner, &transaction.Amount, &transaction.Date, &transaction.TransactionTypeId)
		transactions = append(transactions, transaction)
	}
	return transactions
}

func GetTransactionsFromDB() []transaction_organizer_domain.Transaction {
	transaction_organizer_db.Connect()
	return nil
}

func SaveTransactions(db *sql.DB, transactions []transaction_organizer_domain.Transaction) {
	for _, transaction := range transactions {
		_, err := db.Exec("INSERT INTO transactions (partner, amount, date) VALUES ($1, $2, $3)", transaction.Partner, transaction.Amount, transaction.Date)

		if err != nil {
			log.Fatalf("An error occured while executing transaction saving: %v", err)
		}
	}
}
