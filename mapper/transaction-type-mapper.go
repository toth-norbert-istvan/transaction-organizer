package mapper

import (
	"github.com/transaction-organizer/domain"
	"github.com/transaction-organizer/dto"
)

type TransactionTypeMapper struct{}

func (ttm TransactionTypeMapper) DomainsToDtos(domains []domain.TransactionType) []dto.TransactionType {
	var dtos []dto.TransactionType
	for _, domain := range domains {
		dtos = append(dtos, ttm.DomainToDto(domain))
	}
	return dtos
}

func (ttm TransactionTypeMapper) DomainToDto(domain domain.TransactionType) dto.TransactionType {
	return dto.TransactionType{
		Id:    domain.Id,
		Name:  domain.Name,
		Group: domain.TypeGroup,
	}
}

func (ttm TransactionTypeMapper) DtoToDomain(dto dto.NewTransactionType) (domain.TransactionType, bool) {
	typeGroup, ok := domain.ParseTransactionTypeGroup(dto.Group)
	if !ok {
		return domain.TransactionType{}, true
	}

	return domain.TransactionType{
		Id:        -1,
		Name:      dto.Name,
		TypeGroup: typeGroup,
	}, false
}
