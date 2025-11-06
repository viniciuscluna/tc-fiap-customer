package controller

import "github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"

type CustomerController interface {
	GetByCpf(cpf uint) (*dto.GetCustomerResponseDto, error)
	Add(customer *dto.AddCustomerRequestDto) error
}
