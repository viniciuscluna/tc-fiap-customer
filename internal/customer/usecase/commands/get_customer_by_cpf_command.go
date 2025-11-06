package commands

type GetCustomerByCpfCommand struct {
	Cpf uint
}

func NewGetCustomerByCpfCommand(cpf uint) *GetCustomerByCpfCommand {
	return &GetCustomerByCpfCommand{
		Cpf: cpf,
	}
}
