package commands_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
)

func TestNewAddCustomerCommand(t *testing.T) {
	// GIVEN valid customer data
	name := "John Doe"
	email := "john@example.com"
	cpf := "12345678901"

	// WHEN creating a new AddCustomerCommand
	command := commands.NewAddCustomerCommand(name, email, cpf)

	// THEN the command should be created with the correct values
	assert.NotNil(t, command)
	assert.Equal(t, name, command.Name)
	assert.Equal(t, email, command.Email)
	assert.Equal(t, cpf, command.CPF)
}
