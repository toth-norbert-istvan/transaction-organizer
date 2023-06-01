package transaction_organizer_mapper

import (
	"database/sql"
	transaction_organizer_controller_model "gitlab.com/transaction-organizer/controller/model"
	transaction_organizer_domain "gitlab.com/transaction-organizer/domain"
	transaction_organizer_domain_service "gitlab.com/transaction-organizer/service"
	"strings"
)

func DomainsToDtos(domains []transaction_organizer_domain.Transaction, db *sql.DB) []transaction_organizer_controller_model.Transaction {
	var dtos []transaction_organizer_controller_model.Transaction
	for _, domain := range domains {
		dtos = append(dtos, DomainToDto(domain, db))
	}
	return dtos
}

func DomainToDto(domain transaction_organizer_domain.Transaction, db *sql.DB) transaction_organizer_controller_model.Transaction {
	transactionTypeDomain, err := transaction_organizer_domain_service.GetTransactionType(domain.TransactionTypeId, db)

	var transactionTypeDto *transaction_organizer_controller_model.TransactionType
	if err == nil {
		transactionTypeDto = &transaction_organizer_controller_model.TransactionType{
			Id:    domain.TransactionTypeId,
			Name:  transactionTypeDomain.Name,
			Group: transactionTypeDomain.TypeGroup,
		}
	}

	return transaction_organizer_controller_model.Transaction{
		Id:              domain.Id,
		Partner:         strings.TrimSpace(domain.Partner),
		Amount:          domain.Amount,
		Date:            domain.Date,
		TransactionType: transactionTypeDto,
	}
}
