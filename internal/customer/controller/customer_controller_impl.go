package controller

import (
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"
	customerPresenter "github.com/viniciuscluna/tc-fiap-customer/internal/customer/presenter"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/addCustomer"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
	getbycpf "github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/getbycpf"
)

var (
	_ CustomerController = (*CustomerControllerImpl)(nil)
)

type CustomerControllerImpl struct {
	presenter          customerPresenter.CustomerPresenter
	addCustomerUseCase addCustomer.AddCustomerUseCase
	getByCpfUseCase    getbycpf.GetByCpfUseCase
}

func NewCustomerControllerImpl(
	presenter customerPresenter.CustomerPresenter,
	addCustomerUseCase addCustomer.AddCustomerUseCase,
	getByCpfUseCase getbycpf.GetByCpfUseCase) *CustomerControllerImpl {
	return &CustomerControllerImpl{
		presenter:          presenter,
		addCustomerUseCase: addCustomerUseCase,
		getByCpfUseCase:    getByCpfUseCase,
	}
}

func (c *CustomerControllerImpl) GetByCpf(cpf string) (*dto.GetCustomerResponseDto, error) {
	customer, err := c.getByCpfUseCase.Execute(commands.NewGetCustomerByCpfCommand(cpf))
	if err != nil {
		return nil, err
	}

	return c.presenter.Present(customer), nil
}

func (c *CustomerControllerImpl) Add(customer *dto.AddCustomerRequestDto) error {
	command := commands.NewAddCustomerCommand(customer.Name, customer.Email, customer.CPF)
	err := c.addCustomerUseCase.Execute(command)
	if err != nil {
		return err
	}
	return nil
}
