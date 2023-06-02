package dto

import (
	"github.com/transaction-organizer/domain"
)

type GroupSummaryReport struct {
	Group domain.TransactionTypeGroup `json:"group"`
	Sum   float64                     `json:"sum"`
}
