package transaction_organizer_mapper

import (
	transaction_organizer_controller_model "gitlab.com/transaction-organizer/controller/model"
	transaction_organizer_domain "gitlab.com/transaction-organizer/domain"
)

type TransactionTypeMapper struct{}

func (ttm TransactionTypeMapper) DomainsToDtos(domains []transaction_organizer_domain.TransactionType) []transaction_organizer_controller_model.TransactionType {
	var dtos []transaction_organizer_controller_model.TransactionType
	for _, domain := range domains {
		dtos = append(dtos, ttm.DomainToDto(domain))
	}
	return dtos
}

func (ttm TransactionTypeMapper) DomainToDto(domain transaction_organizer_domain.TransactionType) transaction_organizer_controller_model.TransactionType {
	return transaction_organizer_controller_model.TransactionType{
		Id:    domain.Id,
		Name:  domain.Name,
		Group: domain.TypeGroup,
	}
}

func (ttm TransactionTypeMapper) DtoToDomain(dto transaction_organizer_controller_model.NewTransactionType) (transaction_organizer_domain.TransactionType, bool) {
	typeGroup, ok := transaction_organizer_domain.ParseTransactionTypeGroup(dto.Group)
	if !ok {
		return transaction_organizer_domain.TransactionType{}, true
	}

	return transaction_organizer_domain.TransactionType{
		Id:        -1,
		Name:      dto.Name,
		TypeGroup: typeGroup,
	}, false
}
