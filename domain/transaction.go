package transaction_organizer_domain

import (
	"strings"
	"time"
)

type Transaction struct {
	Id                int
	Partner           string
	Amount            float64
	Date              time.Time
	TransactionTypeId int
}

type TransactionType struct {
	Id        int
	Name      string
	TypeGroup TransactionTypeGroup
}

// TransactionTypeGroup - Custom type to hold value for TransactionTypeGroup
type TransactionTypeGroup string

const (
	Overhead          TransactionTypeGroup = "OVERHEAD"
	OccasionalExpense TransactionTypeGroup = "OCCASIONAL_EXPENSE"
	ExtraExpense      TransactionTypeGroup = "EXTRA_EXPENSE"
)

var (
	capabilitiesMap = map[string]TransactionTypeGroup{
		"overhead":           Overhead,
		"occasional_expense": OccasionalExpense,
		"extra_expense":      ExtraExpense,
	}
)

func ParseTransactionTypeGroup(str string) (TransactionTypeGroup, bool) {
	c, ok := capabilitiesMap[strings.ToLower(str)]
	return c, ok
}
