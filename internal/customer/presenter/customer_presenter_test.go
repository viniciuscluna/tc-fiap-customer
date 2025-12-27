package presenter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/presenter"
)

type CustomerPresenterTestSuite struct {
	suite.Suite
	presenter presenter.CustomerPresenter
}

func (suite *CustomerPresenterTestSuite) SetupTest() {
	suite.presenter = presenter.NewCustomerPresenterImpl()
}

func TestCustomerPresenterTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerPresenterTestSuite))
}

// Feature: Customer Presentation
// Scenario: Transform customer entity to response DTO successfully

func (suite *CustomerPresenterTestSuite) Test_CustomerPresentation_WithValidCustomerData_ShouldTransformToDTOSuccessfully() {
	// GIVEN a customer entity with complete data
	now := time.Now()
	customer := &entities.Customer{
		ID:        "123-456",
		CPF:       12345678901,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: now,
	}

	// WHEN the presenter transforms the customer to DTO
	dto := suite.presenter.Present(customer)

	// THEN the DTO should not be nil
	assert.NotNil(suite.T(), dto)
	// AND all customer fields should be preserved in the DTO
	assert.Equal(suite.T(), customer.ID, dto.ID)
	assert.Equal(suite.T(), customer.CPF, dto.CPF)
	assert.Equal(suite.T(), customer.Name, dto.Name)
	assert.Equal(suite.T(), customer.Email, dto.Email)
	assert.Equal(suite.T(), customer.CreatedAt, dto.CreatedAt)
}

func (suite *CustomerPresenterTestSuite) Test_CustomerPresentation_WithEmptyFields_ShouldPreserveEmptyValues() {
	// GIVEN a customer entity with empty/zero fields
	customer := &entities.Customer{
		ID:        "",
		CPF:       0,
		Name:      "",
		Email:     "",
		CreatedAt: time.Time{},
	}

	// WHEN the presenter transforms the customer to DTO
	dto := suite.presenter.Present(customer)

	// THEN the DTO should not be nil
	assert.NotNil(suite.T(), dto)
	// AND empty fields should remain empty in the DTO
	assert.Empty(suite.T(), dto.ID)
	assert.Equal(suite.T(), uint(0), dto.CPF)
	assert.Empty(suite.T(), dto.Name)
	assert.Empty(suite.T(), dto.Email)
	assert.True(suite.T(), dto.CreatedAt.IsZero())
}

func (suite *CustomerPresenterTestSuite) Test_CustomerPresentation_WithAlternativeData_ShouldPreserveAllDataAccurately() {
	// GIVEN a customer entity with alternative data
	now := time.Now()
	customer := &entities.Customer{
		ID:        "abc-123",
		CPF:       98765432100,
		Name:      "Jane Smith",
		Email:     "jane.smith@test.com",
		CreatedAt: now,
	}

	// WHEN the presenter transforms the customer to DTO
	dto := suite.presenter.Present(customer)

	// THEN the DTO should not be nil
	assert.NotNil(suite.T(), dto)
	// AND all fields should be accurately transformed
	assert.Equal(suite.T(), "abc-123", dto.ID)
	assert.Equal(suite.T(), uint(98765432100), dto.CPF)
	assert.Equal(suite.T(), "Jane Smith", dto.Name)
	assert.Equal(suite.T(), "jane.smith@test.com", dto.Email)
	assert.Equal(suite.T(), now, dto.CreatedAt)
}
