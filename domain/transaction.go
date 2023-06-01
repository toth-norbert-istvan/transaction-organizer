package transaction_organizer_domain

import "time"

type Transaction struct {
	Id                int       `json:"id"`
	Partner           string    `json:"partner"`
	Amount            float64   `json:"amount"`
	Date              time.Time `json:"date"`
	TransactionTypeId int       `json:"transactionTypeId"`
}

type TransactionType struct {
	Id        int                  `json:"id"`
	Name      string               `json:"name"`
	TypeGroup TransactionTypeGroup `json:"typeGroup"`
}

// TransactionTypeGroup - Custom type to hold value for weekday ranging from 1-7
type TransactionTypeGroup int

// Declare related constants for each weekday starting with index 1
const (
	Overhead TransactionTypeGroup = iota + 1
	OccasionalExpense
	ExtraExpense
)

// String - Creating common behavior - give the type a String function
func (w TransactionTypeGroup) String() string {
	return [...]string{"Overhead", "OccasionalExpense", "ExtraExpense"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w TransactionTypeGroup) EnumIndex() int {
	return int(w)
}
