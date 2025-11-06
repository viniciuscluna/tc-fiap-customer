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

func (suite *CustomerPresenterTestSuite) TestPresent_Success() {
	// Arrange
	now := time.Now()
	customer := &entities.Customer{
		ID:        "123-456",
		CPF:       12345678901,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: now,
	}

	// Act
	dto := suite.presenter.Present(customer)

	// Assert
	assert.NotNil(suite.T(), dto)
	assert.Equal(suite.T(), customer.ID, dto.ID)
	assert.Equal(suite.T(), customer.CPF, dto.CPF)
	assert.Equal(suite.T(), customer.Name, dto.Name)
	assert.Equal(suite.T(), customer.Email, dto.Email)
	assert.Equal(suite.T(), customer.CreatedAt, dto.CreatedAt)
}

func (suite *CustomerPresenterTestSuite) TestPresent_WithEmptyFields() {
	// Arrange
	customer := &entities.Customer{
		ID:        "",
		CPF:       0,
		Name:      "",
		Email:     "",
		CreatedAt: time.Time{},
	}

	// Act
	dto := suite.presenter.Present(customer)

	// Assert
	assert.NotNil(suite.T(), dto)
	assert.Empty(suite.T(), dto.ID)
	assert.Equal(suite.T(), uint(0), dto.CPF)
	assert.Empty(suite.T(), dto.Name)
	assert.Empty(suite.T(), dto.Email)
	assert.True(suite.T(), dto.CreatedAt.IsZero())
}

func (suite *CustomerPresenterTestSuite) TestPresent_PreservesAllData() {
	// Arrange
	now := time.Now()
	customer := &entities.Customer{
		ID:        "abc-123",
		CPF:       98765432100,
		Name:      "Jane Smith",
		Email:     "jane.smith@test.com",
		CreatedAt: now,
	}

	// Act
	dto := suite.presenter.Present(customer)

	// Assert
	assert.NotNil(suite.T(), dto)
	assert.Equal(suite.T(), "abc-123", dto.ID)
	assert.Equal(suite.T(), uint(98765432100), dto.CPF)
	assert.Equal(suite.T(), "Jane Smith", dto.Name)
	assert.Equal(suite.T(), "jane.smith@test.com", dto.Email)
	assert.Equal(suite.T(), now, dto.CreatedAt)
}
