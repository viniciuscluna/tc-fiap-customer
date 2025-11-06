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

func (suite *AddCustomerUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	command := commands.NewAddCustomerCommand("John Doe", "john@example.com", 12345678901)

	expectedCustomer := &entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	suite.mockRepository.EXPECT().
		Add(expectedCustomer).
		Return(nil).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *AddCustomerUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	command := commands.NewAddCustomerCommand("Jane Doe", "jane@example.com", 98765432109)

	expectedCustomer := &entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	expectedError := errors.New("database error")

	suite.mockRepository.EXPECT().
		Add(expectedCustomer).
		Return(expectedError).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *AddCustomerUseCaseTestSuite) TestExecute_ValidatesCustomerData() {
	// Arrange
	command := commands.NewAddCustomerCommand("", "", 0)

	expectedCustomer := &entities.Customer{
		Name:  command.Name,
		Email: command.Email,
		CPF:   command.CPF,
	}

	suite.mockRepository.EXPECT().
		Add(expectedCustomer).
		Return(nil).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepository.AssertExpectations(suite.T())
}
