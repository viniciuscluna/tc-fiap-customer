package getbycpf

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
)

type GetByCpfUseCase interface {
	Execute(command *commands.GetCustomerByCpfCommand) (*entities.Customer, error)
}
