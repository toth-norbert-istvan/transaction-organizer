package transaction_organizer_controller_model

import (
	transaction_organizer_domain "gitlab.com/transaction-organizer/domain"
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

type NewTransactionType struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

type TransactionType struct {
	Id    int                                               `json:"id"`
	Name  string                                            `json:"name"`
	Group transaction_organizer_domain.TransactionTypeGroup `json:"group"`
}
