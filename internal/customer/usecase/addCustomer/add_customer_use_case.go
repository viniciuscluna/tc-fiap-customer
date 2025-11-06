package addCustomer

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
)

type AddCustomerUseCase interface {
	Execute(command *commands.AddCustomerCommand) error
}
