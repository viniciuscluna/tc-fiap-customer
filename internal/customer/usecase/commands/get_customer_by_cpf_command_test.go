package commands_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
)

func TestNewGetCustomerByCpfCommand(t *testing.T) {
	// GIVEN a valid CPF
	cpf := "12345678901"

	// WHEN creating a new GetCustomerByCpfCommand
	command := commands.NewGetCustomerByCpfCommand(cpf)

	// THEN the command should be created with the correct values
	assert.NotNil(t, command)
	assert.Equal(t, cpf, command.CPF)
}
