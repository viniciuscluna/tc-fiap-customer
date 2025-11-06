package commands

type AddCustomerCommand struct {
	Name  string
	Email string
	CPF   uint
}

func NewAddCustomerCommand(name string, email string, cpf uint) *AddCustomerCommand {
	return &AddCustomerCommand{
		Name:  name,
		Email: email,
		CPF:   cpf,
	}
}
