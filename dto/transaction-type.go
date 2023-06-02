package dto

import (
	"github.com/transaction-organizer/domain"
)

type NewTransactionType struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

type TransactionType struct {
	Id    int                         `json:"id"`
	Name  string                      `json:"name"`
	Group domain.TransactionTypeGroup `json:"group"`
}
