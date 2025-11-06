package repositories

import "github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"

type CustomerRepository interface {
	GetByCpf(cpf uint) (*entities.Customer, error)
	Add(customer *entities.Customer) error
}
