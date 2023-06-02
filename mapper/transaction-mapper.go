package mapper

import (
	"github.com/transaction-organizer/domain"
	"github.com/transaction-organizer/dto"
	"github.com/transaction-organizer/repository"
)

type TransactionMapper struct{}

func (tm TransactionMapper) DomainsToDtos(domains []domain.Transaction) []dto.Transaction {
	var dtos []dto.Transaction
	for _, domain := range domains {
		dtos = append(dtos, tm.DomainToDto(domain))
	}
	return dtos
}

func (tm TransactionMapper) DomainToDto(domain domain.Transaction) dto.Transaction {
	transactionTypeDomain, err := repository.TransactionTypeRepository{}.GetTransactionType(domain.TransactionTypeId)

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
		Partner:         domain.Partner,
		Amount:          domain.Amount,
		Date:            domain.Date,
		TransactionType: transactionTypeDto,
	}
}
