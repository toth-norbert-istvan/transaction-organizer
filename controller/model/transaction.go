package transaction_organizer_controller_model

import transaction_organizer_domain "gitlab.com/transaction-organizer/domain"

type transaction struct {
	ID              string          `json:"id"`
	Partner         string          `json:"title"`
	Amount          string          `json:"artist"`
	Date            float64         `json:"price"`
	transactionType transactionType `json:type`
}

type transactionType struct {
	ID    string                                            `json:"id"`
	Name  string                                            `json:"name"`
	Group transaction_organizer_domain.TransactionTypeGroup `json:"group"`
}
