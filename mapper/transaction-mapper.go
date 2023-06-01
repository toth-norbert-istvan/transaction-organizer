package mapper

import (
	"database/sql"
	"github.com/transaction-organizer/domain"
	"github.com/transaction-organizer/dto"
	"github.com/transaction-organizer/repository"
	"strings"
)

func DomainsToDtos(domains []domain.Transaction, db *sql.DB) []dto.Transaction {
	var dtos []dto.Transaction
	for _, domain := range domains {
		dtos = append(dtos, DomainToDto(domain, db))
	}
	return dtos
}

func DomainToDto(domain domain.Transaction, db *sql.DB) dto.Transaction {
	transactionTypeDomain, err := repository.GetTransactionType(domain.TransactionTypeId, db)

	var transactionTypeDto *dto.TransactionType
	if err == nil {
		transactionTypeDto = &dto.TransactionType{
			Id:    domain.TransactionTypeId,
			Name:  transactionTypeDomain.Name,
			Group: transactionTypeDomain.TypeGroup,
		}
	}

	return dto.Transaction{
		Id:              domain.Id,
		Partner:         strings.TrimSpace(domain.Partner),
		Amount:          domain.Amount,
		Date:            domain.Date,
		TransactionType: transactionTypeDto,
	}
}
