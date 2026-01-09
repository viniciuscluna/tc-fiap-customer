package addCustomer_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/addCustomer"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/commands"
	mockRepositories "github.com/viniciuscluna/tc-fiap-customer/mocks/customer/domain/repositories"
)

type AddCustomerUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mockRepositories.MockCustomerRepository
	useCase        addCustomer.AddCustomerUseCase
}

func (suite *AddCustomerUseCaseTestSuite) SetupTest() {
	suite.mockRepository = mockRepositories.NewMockCustomerRepository(suite.T())
	suite.useCase = addCustomer.NewAddCustomerUseCaseImpl(suite.mockRepository)
}

func TestAddCustomerUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AddCustomerUseCaseTestSuite))
}

// Feature: Add Customer Use Case
// Scenario: Register a new customer in the system

func (suite *AddCustomerUseCaseTestSuite) Test_CustomerRegistration_WithValidInformation_ShouldPersistSuccessfully() {
	// GIVEN a customer with valid name, email, and CPF
	command := commands.NewAddCustomerCommand("John Doe", "john@example.com", "12345678901")

	expectedCustomer := &entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	suite.mockRepository.EXPECT().
		Add(expectedCustomer).
		Return(nil).
		Once()

	// WHEN the customer registration is executed
	err := suite.useCase.Execute(command)

	// THEN the operation should complete without errors
	assert.NoError(suite.T(), err)
	// AND the repository should have persisted the customer
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *AddCustomerUseCaseTestSuite) Test_CustomerRegistration_WithRepositoryFailure_ShouldReturnError() {
	// GIVEN a customer registration request
	command := commands.NewAddCustomerCommand("Jane Doe", "jane@example.com", "98765432109")

	expectedCustomer := &entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	// AND the repository fails with a database error
	expectedError := errors.New("database error")

	suite.mockRepository.EXPECT().
		Add(expectedCustomer).
		Return(expectedError).
		Once()

	// WHEN the customer registration is executed
	err := suite.useCase.Execute(command)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND the error should match the repository error
	assert.Equal(suite.T(), expectedError, err)
	// AND the repository should have attempted the operation
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *AddCustomerUseCaseTestSuite) Test_CustomerRegistration_WithEmptyData_ShouldAcceptAndPersist() {
	// GIVEN a customer registration request with empty data
	command := commands.NewAddCustomerCommand("", "", "0")

	expectedCustomer := &entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	// WHEN the customer registration is executed
	suite.mockRepository.EXPECT().
		Add(expectedCustomer).
		Return(nil).
		Once()

	err := suite.useCase.Execute(command)

	// THEN the operation should complete without errors
	assert.NoError(suite.T(), err)
	// AND the repository should have received the empty customer data
	suite.mockRepository.AssertExpectations(suite.T())
}
