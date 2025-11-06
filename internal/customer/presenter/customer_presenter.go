package presenter

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"
)

type CustomerPresenter interface {
	Present(customer *entities.Customer) *dto.GetCustomerResponseDto
}
