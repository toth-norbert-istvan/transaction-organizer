package repository

import (
	"github.com/transaction-organizer/db"
	"github.com/transaction-organizer/domain"
	"log"
)

type TransactionTypeRepository struct{}

func (ttr TransactionTypeRepository) GetTransactionType(transactionTypeId int) (domain.TransactionType, error) {
	var db = db.PostgreSqlDB{}.GetDb()
	var transactionType domain.TransactionType

	err := db.QueryRow("SELECT * FROM transaction_type WHERE id=$1", transactionTypeId).Scan(&transactionType.Id, &transactionType.Name, &transactionType.TypeGroup)
	if err != nil {
		return transactionType, err
	}

	return transactionType, nil
}

func (ttr TransactionTypeRepository) GetTransactionTypes() []domain.TransactionType {
	var db = db.PostgreSqlDB{}.GetDb()

	rows, err := db.Query("SELECT * FROM transaction_type")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var transactionType domain.TransactionType
	var transactionTypes []domain.TransactionType
	for rows.Next() {
		rows.Scan(&transactionType.Id, &transactionType.Name, &transactionType.TypeGroup)
		transactionTypes = append(transactionTypes, transactionType)
	}

	return transactionTypes
}

func (ttr TransactionTypeRepository) SaveTransactionType(transactionType domain.TransactionType) domain.TransactionType {
	var id int
	var db = db.PostgreSqlDB{}.GetDb()

	err := db.QueryRow("INSERT INTO transaction_type (name, type_group) VALUES ($1, $2) RETURNING id", transactionType.Name, transactionType.TypeGroup).Scan(&id)

	if err != nil {
		log.Fatalf("An error occured while executing transaction type saving: %v", err)
	}
	transactionType.Id = id

	return transactionType
}
