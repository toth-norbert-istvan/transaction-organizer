package domain

import (
	"time"
)

type Transaction struct {
	Id                int
	Partner           string
	Amount            float64
	Date              time.Time
	TransactionTypeId int
}
