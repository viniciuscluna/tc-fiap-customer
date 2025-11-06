package controller_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/controller"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"
	mockPresenter "github.com/viniciuscluna/tc-fiap-customer/mocks/customer/presenter"
	mockAddCustomer "github.com/viniciuscluna/tc-fiap-customer/mocks/customer/usecase/addCustomer"
	mockGetByCpf "github.com/viniciuscluna/tc-fiap-customer/mocks/customer/usecase/getbycpf"
)

type CustomerControllerTestSuite struct {
	suite.Suite
	mockPresenter          *mockPresenter.MockCustomerPresenter
	mockAddCustomerUseCase *mockAddCustomer.MockAddCustomerUseCase
	mockGetByCpfUseCase    *mockGetByCpf.MockGetByCpfUseCase
	controller             controller.CustomerController
}

func (suite *CustomerControllerTestSuite) SetupTest() {
	suite.mockPresenter = mockPresenter.NewMockCustomerPresenter(suite.T())
	suite.mockAddCustomerUseCase = mockAddCustomer.NewMockAddCustomerUseCase(suite.T())
	suite.mockGetByCpfUseCase = mockGetByCpf.NewMockGetByCpfUseCase(suite.T())

	suite.controller = controller.NewCustomerControllerImpl(
		suite.mockPresenter,
		suite.mockAddCustomerUseCase,
		suite.mockGetByCpfUseCase,
	)
}

func TestCustomerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerControllerTestSuite))
}

func (suite *CustomerControllerTestSuite) TestGetByCpf_Success() {
	// Arrange
	cpf := uint(12345678901)
	now := time.Now()

	customerEntity := &entities.Customer{
		ID:        "123",
		CPF:       cpf,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: now,
	}

	expectedDto := &dto.GetCustomerResponseDto{
		ID:        "123",
		CPF:       cpf,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: now,
	}

	suite.mockGetByCpfUseCase.EXPECT().
		Execute(mock.Anything).
		Return(customerEntity, nil).
		Once()

	suite.mockPresenter.EXPECT().
		Present(customerEntity).
		Return(expectedDto).
		Once()

	// Act
	result, err := suite.controller.GetByCpf(cpf)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedDto.ID, result.ID)
	assert.Equal(suite.T(), expectedDto.CPF, result.CPF)
	assert.Equal(suite.T(), expectedDto.Name, result.Name)
	assert.Equal(suite.T(), expectedDto.Email, result.Email)
	suite.mockGetByCpfUseCase.AssertExpectations(suite.T())
	suite.mockPresenter.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) TestGetByCpf_UseCaseError() {
	// Arrange
	cpf := uint(99999999999)
	expectedError := errors.New("customer not found")

	suite.mockGetByCpfUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil, expectedError).
		Once()

	// Act
	result, err := suite.controller.GetByCpf(cpf)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockGetByCpfUseCase.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) TestAdd_Success() {
	// Arrange
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		CPF:   98765432109,
	}

	suite.mockAddCustomerUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil).
		Once()

	// Act
	err := suite.controller.Add(requestDto)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockAddCustomerUseCase.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) TestAdd_UseCaseError() {
	// Arrange
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "John Smith",
		Email: "john.smith@example.com",
		CPF:   11111111111,
	}

	expectedError := errors.New("database error")

	suite.mockAddCustomerUseCase.EXPECT().
		Execute(mock.Anything).
		Return(expectedError).
		Once()

	// Act
	err := suite.controller.Add(requestDto)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockAddCustomerUseCase.AssertExpectations(suite.T())
}
