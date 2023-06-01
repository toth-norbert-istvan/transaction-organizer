package transaction_organizer_domain_service

import (
	"database/sql"
	"gitlab.com/transaction-organizer/domain"
	"log"
)

func GetTransactionType(transactionTypeId int, db *sql.DB) (transaction_organizer_domain.TransactionType, error) {
	var transactionType transaction_organizer_domain.TransactionType

	err := db.QueryRow("SELECT * FROM transaction_type WHERE id=$1", transactionTypeId).Scan(&transactionType.Id, &transactionType.Name, &transactionType.TypeGroup)
	if err != nil {
		return transactionType, err
	}
	return transactionType, nil
}

func GetTransactionTypes(db *sql.DB) []transaction_organizer_domain.TransactionType {
	rows, err := db.Query("SELECT * FROM transaction_type")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var transactionType transaction_organizer_domain.TransactionType
	var transactionTypes []transaction_organizer_domain.TransactionType
	for rows.Next() {
		rows.Scan(&transactionType.Id, &transactionType.Name, &transactionType.TypeGroup)
		transactionTypes = append(transactionTypes, transactionType)
	}

	return transactionTypes
}

func SaveTransactionType(db *sql.DB, transactionType transaction_organizer_domain.TransactionType) transaction_organizer_domain.TransactionType {
	var id int
	err := db.QueryRow("INSERT INTO transaction_type (name, type_group) VALUES ($1, $2) RETURNING id", transactionType.Name, transactionType.TypeGroup).Scan(&id)

	if err != nil {
		log.Fatalf("An error occured while executing transaction type saving: %v", err)
	}
	transactionType.Id = id
	return transactionType
}
