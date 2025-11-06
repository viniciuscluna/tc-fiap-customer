package addCustomer

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/repositories"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
)

var (
	_ AddCustomerUseCase = (*AddCustomerUseCaseImpl)(nil)
)

type AddCustomerUseCaseImpl struct {
	customerRepository repositories.CustomerRepository
}

func NewAddCustomerUseCaseImpl(customerRepository repositories.CustomerRepository) *AddCustomerUseCaseImpl {
	return &AddCustomerUseCaseImpl{customerRepository: customerRepository}
}

func (u *AddCustomerUseCaseImpl) Execute(command *commands.AddCustomerCommand) error {
	entity := entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	return u.customerRepository.Add(&entity)
}
