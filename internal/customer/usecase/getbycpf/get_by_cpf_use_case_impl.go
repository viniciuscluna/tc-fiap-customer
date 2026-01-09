package getbycpf

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/repositories"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
)

var (
	_ GetByCpfUseCase = (*GetByCpfUseCaseImpl)(nil)
)

type GetByCpfUseCaseImpl struct {
	customerRepository repositories.CustomerRepository
}

func NewGetByCpfUseCaseImpl(customerRepository repositories.CustomerRepository) *GetByCpfUseCaseImpl {
	return &GetByCpfUseCaseImpl{customerRepository: customerRepository}
}

func (u *GetByCpfUseCaseImpl) Execute(command *commands.GetCustomerByCpfCommand) (*entities.Customer, error) {
	entity, err := u.customerRepository.GetByCpf(command.CPF)
	if err != nil {
		return nil, err
	}

	return entity, err
}
