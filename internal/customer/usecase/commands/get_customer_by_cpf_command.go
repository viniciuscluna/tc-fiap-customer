package commands

type GetCustomerByCpfCommand struct {
	CPF string
}

func NewGetCustomerByCpfCommand(cpf string) *GetCustomerByCpfCommand {
	return &GetCustomerByCpfCommand{
		CPF: cpf,
	}
}
