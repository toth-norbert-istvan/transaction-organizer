package dto

import (
	"time"
)

type Transaction struct {
	Id              int              `json:"id"`
	Partner         string           `json:"partner"`
	Amount          float64          `json:"amount"`
	Date            time.Time        `json:"date"`
	TransactionType *TransactionType `json:"type"`
}

type TransactionPatch struct {
	TransactionTypeId string `json:"id"`
}
