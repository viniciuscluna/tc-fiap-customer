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

// Feature: Customer Controller - Get Customer by CPF
// Scenario: Retrieve and present customer data

func (suite *CustomerControllerTestSuite) Test_CustomerRetrieval_WithValidCPF_ShouldReturnPresentedCustomerSuccessfully() {
	// GIVEN a valid CPF and an existing customer
	cpf := "12345678901"
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

	// WHEN the controller retrieves and presents the customer
	result, err := suite.controller.GetByCpf(cpf)

	// THEN the customer should be returned without errors
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	// AND all customer details should be correctly presented
	assert.Equal(suite.T(), expectedDto.ID, result.ID)
	assert.Equal(suite.T(), expectedDto.CPF, result.CPF)
	assert.Equal(suite.T(), expectedDto.Name, result.Name)
	assert.Equal(suite.T(), expectedDto.Email, result.Email)
	// AND both use case and presenter should have been called
	suite.mockGetByCpfUseCase.AssertExpectations(suite.T())
	suite.mockPresenter.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) Test_CustomerRetrieval_WithNonExistentCPF_ShouldReturnError() {
	// GIVEN a CPF for a non-existent customer
	cpf := "99999999999"
	expectedError := errors.New("customer not found")

	suite.mockGetByCpfUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil, expectedError).
		Once()

	// WHEN the controller attempts to retrieve the customer
	result, err := suite.controller.GetByCpf(cpf)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND no customer data should be returned
	assert.Nil(suite.T(), result)
	// AND the error should match the use case error
	assert.Equal(suite.T(), expectedError, err)
	// AND the use case should have been called
	suite.mockGetByCpfUseCase.AssertExpectations(suite.T())
}

// Feature: Customer Controller - Add Customer
// Scenario: Register a new customer

func (suite *CustomerControllerTestSuite) Test_CustomerRegistration_WithValidRequest_ShouldAddCustomerSuccessfully() {
	// GIVEN a valid add customer request
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		CPF: "98765432109",
	}

	suite.mockAddCustomerUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil).
		Once()

	// WHEN the controller processes the customer registration
	err := suite.controller.Add(requestDto)

	// THEN the operation should complete without errors
	assert.NoError(suite.T(), err)
	// AND the use case should have been called
	suite.mockAddCustomerUseCase.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) Test_CustomerRegistration_WithUseCaseFailure_ShouldReturnError() {
	// GIVEN a customer registration request
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "John Smith",
		Email: "john.smith@example.com",
		CPF: "11111111111",
	}

	// AND the use case fails with a database error
	expectedError := errors.New("database error")

	suite.mockAddCustomerUseCase.EXPECT().
		Execute(mock.Anything).
		Return(expectedError).
		Once()

	// WHEN the controller processes the registration
	err := suite.controller.Add(requestDto)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND the error should match the use case error
	assert.Equal(suite.T(), expectedError, err)
	// AND the use case should have been called
	suite.mockAddCustomerUseCase.AssertExpectations(suite.T())
}
