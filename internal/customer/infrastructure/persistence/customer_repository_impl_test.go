package persistence_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/persistence"
)

// Mock DynamoDB Client
type MockDynamoDBClient struct {
	mock.Mock
	dynamodbiface.DynamoDBAPI
}

func (m *MockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func (m *MockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

type CustomerRepositoryTestSuite struct {
	suite.Suite
	mockDB     *MockDynamoDBClient
	repository *persistence.CustomerRepositoryImpl
}

func (suite *CustomerRepositoryTestSuite) SetupTest() {
	suite.mockDB = new(MockDynamoDBClient)
	suite.repository = persistence.NewCustomerRepositoryImpl(suite.mockDB)
}

func TestCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}

func (suite *CustomerRepositoryTestSuite) TestGetByCpf_Success() {
	// Arrange
	cpf := uint(12345678901)
	expectedCustomer := &entities.Customer{
		ID:    "test-id-123",
		CPF:   cpf,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	output := &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"ID":    {S: aws.String("test-id-123")},
			"CPF":   {N: aws.String("12345678901")},
			"Name":  {S: aws.String("John Doe")},
			"Email": {S: aws.String("john@example.com")},
		},
	}

	suite.mockDB.On("GetItem", mock.Anything).Return(output, nil).Once()

	// Act
	result, err := suite.repository.GetByCpf(cpf)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedCustomer.ID, result.ID)
	assert.Equal(suite.T(), expectedCustomer.CPF, result.CPF)
	assert.Equal(suite.T(), expectedCustomer.Name, result.Name)
	assert.Equal(suite.T(), expectedCustomer.Email, result.Email)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) TestGetByCpf_NotFound() {
	// Arrange
	cpf := uint(99999999999)
	output := &dynamodb.GetItemOutput{
		Item: nil, // Empty result
	}

	suite.mockDB.On("GetItem", mock.Anything).Return(output, nil).Once()

	// Act
	result, err := suite.repository.GetByCpf(cpf)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "customer not found")
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) TestGetByCpf_DynamoDBError() {
	// Arrange
	cpf := uint(12345678901)
	expectedError := errors.New("DynamoDB connection error")

	suite.mockDB.On("GetItem", mock.Anything).Return(nil, expectedError).Once()

	// Act
	result, err := suite.repository.GetByCpf(cpf)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "failed to get customer")
	assert.Contains(suite.T(), err.Error(), "DynamoDB connection error")
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) TestAdd_Success() {
	// Arrange
	customer := &entities.Customer{
		CPF:   12345678901,
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	output := &dynamodb.PutItemOutput{}
	suite.mockDB.On("PutItem", mock.Anything).Return(output, nil).Once()

	// Act
	err := suite.repository.Add(customer)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), customer.ID) // ID should be generated
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) TestAdd_DynamoDBError() {
	// Arrange
	customer := &entities.Customer{
		CPF:   12345678901,
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	expectedError := errors.New("DynamoDB write error")
	suite.mockDB.On("PutItem", mock.Anything).Return(nil, expectedError).Once()

	// Act
	err := suite.repository.Add(customer)

	// Assert
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "failed to add customer")
	assert.Contains(suite.T(), err.Error(), "DynamoDB write error")
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) TestAdd_WithExistingID() {
	// Arrange
	customer := &entities.Customer{
		ID:    "existing-id",
		CPF:   12345678901,
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	output := &dynamodb.PutItemOutput{}
	suite.mockDB.On("PutItem", mock.Anything).Return(output, nil).Once()

	// Act
	err := suite.repository.Add(customer)

	// Assert
	assert.NoError(suite.T(), err)
	// Note: O repository sempre gera um novo ID, ignorando o existente
	assert.NotEqual(suite.T(), "existing-id", customer.ID)
	assert.NotEmpty(suite.T(), customer.ID) // ID should be generated
	suite.mockDB.AssertExpectations(suite.T())
}
