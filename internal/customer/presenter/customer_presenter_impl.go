package presenter

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"
)

var (
	_ CustomerPresenter = (*CustomerPresenterImpl)(nil)
)

type CustomerPresenterImpl struct {
}

func NewCustomerPresenterImpl() *CustomerPresenterImpl {
	return &CustomerPresenterImpl{}
}

func (p *CustomerPresenterImpl) Present(customer *entities.Customer) *dto.GetCustomerResponseDto {
	return &dto.GetCustomerResponseDto{
		CreatedAt: customer.CreatedAt,
		Name:      customer.Name,
		CPF:       customer.CPF,
		Email:     customer.Email,
	}
}
