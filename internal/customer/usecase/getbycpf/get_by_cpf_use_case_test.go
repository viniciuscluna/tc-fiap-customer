package getbycpf_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/getbycpf"
	mockRepositories "github.com/viniciuscluna/tc-fiap-customer/mocks/customer/domain/repositories"
)

type GetByCpfUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mockRepositories.MockCustomerRepository
	useCase        getbycpf.GetByCpfUseCase
}

func (suite *GetByCpfUseCaseTestSuite) SetupTest() {
	suite.mockRepository = mockRepositories.NewMockCustomerRepository(suite.T())
	suite.useCase = getbycpf.NewGetByCpfUseCaseImpl(suite.mockRepository)
}

func TestGetByCpfUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetByCpfUseCaseTestSuite))
}

// Feature: Get Customer by CPF Use Case
// Scenario: Retrieve customer information from the system

func (suite *GetByCpfUseCaseTestSuite) Test_CustomerRetrieval_WithExistingCPF_ShouldReturnCustomerSuccessfully() {
	// GIVEN an existing customer with CPF 12345678901
	cpf := "12345678901"
	command := commands.NewGetCustomerByCpfCommand(cpf)

	expectedCustomer := &entities.Customer{
		ID:        "123",
		CPF:       cpf,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	suite.mockRepository.EXPECT().
		GetByCpf(cpf).
		Return(expectedCustomer, nil).
		Once()

	// WHEN searching for the customer by CPF
	customer, err := suite.useCase.Execute(command)

	// THEN the customer should be found without errors
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), customer)
	// AND all customer details should match
	assert.Equal(suite.T(), expectedCustomer.ID, customer.ID)
	assert.Equal(suite.T(), expectedCustomer.CPF, customer.CPF)
	assert.Equal(suite.T(), expectedCustomer.Name, customer.Name)
	assert.Equal(suite.T(), expectedCustomer.Email, customer.Email)
	// AND the repository should have been called
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *GetByCpfUseCaseTestSuite) Test_CustomerRetrieval_WithNonExistentCPF_ShouldReturnError() {
	// GIVEN a CPF that does not exist in the system
	cpf := "99999999999"
	command := commands.NewGetCustomerByCpfCommand(cpf)

	expectedError := errors.New("customer not found")

	suite.mockRepository.EXPECT().
		GetByCpf(cpf).
		Return(nil, expectedError).
		Once()

	// WHEN searching for the customer by CPF
	customer, err := suite.useCase.Execute(command)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND no customer should be returned
	assert.Nil(suite.T(), customer)
	// AND the error message should indicate customer not found
	assert.Equal(suite.T(), expectedError, err)
	// AND the repository should have been called
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *GetByCpfUseCaseTestSuite) Test_CustomerRetrieval_WithRepositoryFailure_ShouldReturnError() {
	// GIVEN a customer retrieval request
	cpf := "12345678901"
	command := commands.NewGetCustomerByCpfCommand(cpf)

	// AND the repository fails with a database connection error
	expectedError := errors.New("database connection error")

	suite.mockRepository.EXPECT().
		GetByCpf(cpf).
		Return(nil, expectedError).
		Once()

	// WHEN searching for the customer by CPF
	customer, err := suite.useCase.Execute(command)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND no customer should be returned
	assert.Nil(suite.T(), customer)
	// AND the error should match the repository error
	assert.Equal(suite.T(), expectedError, err)
	// AND the repository should have been called
	suite.mockRepository.AssertExpectations(suite.T())
}
