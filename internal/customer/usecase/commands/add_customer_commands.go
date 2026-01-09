package commands

type AddCustomerCommand struct {
	Name  string
	Email string
	CPF   string
}

func NewAddCustomerCommand(name string, email string, cpf string) *AddCustomerCommand {
	return &AddCustomerCommand{
		Name:  name,
		Email: email,
		CPF:   cpf,
	}
}
