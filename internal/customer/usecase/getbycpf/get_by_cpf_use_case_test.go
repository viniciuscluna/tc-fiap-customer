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

func (suite *GetByCpfUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	cpf := uint(12345678901)
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

	// Act
	customer, err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), customer)
	assert.Equal(suite.T(), expectedCustomer.ID, customer.ID)
	assert.Equal(suite.T(), expectedCustomer.CPF, customer.CPF)
	assert.Equal(suite.T(), expectedCustomer.Name, customer.Name)
	assert.Equal(suite.T(), expectedCustomer.Email, customer.Email)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *GetByCpfUseCaseTestSuite) TestExecute_CustomerNotFound() {
	// Arrange
	cpf := uint(99999999999)
	command := commands.NewGetCustomerByCpfCommand(cpf)

	expectedError := errors.New("customer not found")

	suite.mockRepository.EXPECT().
		GetByCpf(cpf).
		Return(nil, expectedError).
		Once()

	// Act
	customer, err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), customer)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *GetByCpfUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	cpf := uint(12345678901)
	command := commands.NewGetCustomerByCpfCommand(cpf)

	expectedError := errors.New("database connection error")

	suite.mockRepository.EXPECT().
		GetByCpf(cpf).
		Return(nil, expectedError).
		Once()

	// Act
	customer, err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), customer)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}
