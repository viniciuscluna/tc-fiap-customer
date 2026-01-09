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

// Feature: Customer Repository - Persistence Layer
// Scenario: Retrieve customer data from DynamoDB

func (suite *CustomerRepositoryTestSuite) Test_CustomerRetrieval_WithExistingCPF_ShouldReturnCustomerFromDynamoDB() {
	// GIVEN a customer exists in DynamoDB with CPF 12345678901
	cpf := "12345678901"
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

	// WHEN retrieving the customer by CPF from the repository
	result, err := suite.repository.GetByCpf(cpf)

	// THEN the customer should be found without errors
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	// AND all customer details should match
	assert.Equal(suite.T(), expectedCustomer.ID, result.ID)
	assert.Equal(suite.T(), expectedCustomer.CPF, result.CPF)
	assert.Equal(suite.T(), expectedCustomer.Name, result.Name)
	assert.Equal(suite.T(), expectedCustomer.Email, result.Email)
	// AND DynamoDB GetItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) Test_CustomerRetrieval_WithNonExistentCPF_ShouldReturnNotFoundError() {
	// GIVEN a CPF for a customer that does not exist in DynamoDB
	cpf := "99999999999"
	output := &dynamodb.GetItemOutput{
		Item: nil, // Empty result
	}

	suite.mockDB.On("GetItem", mock.Anything).Return(output, nil).Once()

	// WHEN retrieving the customer by CPF from the repository
	result, err := suite.repository.GetByCpf(cpf)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND no customer should be returned
	assert.Nil(suite.T(), result)
	// AND the error should indicate customer not found
	assert.Contains(suite.T(), err.Error(), "customer not found")
	// AND DynamoDB GetItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) Test_CustomerRetrieval_WithDynamoDBError_ShouldReturnError() {
	// GIVEN a CPF for a customer lookup
	cpf := "12345678901"
	// AND DynamoDB encounters a connection error
	expectedError := errors.New("DynamoDB connection error")

	suite.mockDB.On("GetItem", mock.Anything).Return(nil, expectedError).Once()

	// WHEN retrieving the customer by CPF from the repository
	result, err := suite.repository.GetByCpf(cpf)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND no customer should be returned
	assert.Nil(suite.T(), result)
	// AND the error should indicate the failure
	assert.Contains(suite.T(), err.Error(), "failed to get customer")
	assert.Contains(suite.T(), err.Error(), "DynamoDB connection error")
	// AND DynamoDB GetItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}

// Feature: Customer Repository - Add Customer
// Scenario: Persist a new customer to DynamoDB

func (suite *CustomerRepositoryTestSuite) Test_CustomerPersistence_WithValidCustomer_ShouldSaveToDynamoDBSuccessfully() {
	// GIVEN a valid customer entity to be persisted
	customer := &entities.Customer{
		CPF: "12345678901",
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	output := &dynamodb.PutItemOutput{}
	suite.mockDB.On("PutItem", mock.Anything).Return(output, nil).Once()

	// WHEN adding the customer to the repository
	err := suite.repository.Add(customer)

	// THEN the operation should complete without errors
	assert.NoError(suite.T(), err)
	// AND the customer should have been assigned a unique ID
	assert.NotEmpty(suite.T(), customer.ID)
	// AND DynamoDB PutItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) Test_CustomerPersistence_WithDynamoDBError_ShouldReturnError() {
	// GIVEN a valid customer entity to be persisted
	customer := &entities.Customer{
		CPF: "12345678901",
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	// AND DynamoDB encounters a write error
	expectedError := errors.New("DynamoDB write error")
	suite.mockDB.On("PutItem", mock.Anything).Return(nil, expectedError).Once()

	// WHEN adding the customer to the repository
	err := suite.repository.Add(customer)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND the error should indicate the failure
	assert.Contains(suite.T(), err.Error(), "failed to add customer")
	assert.Contains(suite.T(), err.Error(), "DynamoDB write error")
	// AND DynamoDB PutItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) Test_CustomerPersistence_WithProvidedID_ShouldGenerateNewID() {
	// GIVEN a customer entity with an existing ID
	customer := &entities.Customer{
		ID:    "existing-id",
		CPF: "12345678901",
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	output := &dynamodb.PutItemOutput{}
	suite.mockDB.On("PutItem", mock.Anything).Return(output, nil).Once()

	// WHEN adding the customer to the repository
	err := suite.repository.Add(customer)

	// THEN the operation should complete without errors
	assert.NoError(suite.T(), err)
	// AND the customer ID should be regenerated (not preserved)
	assert.NotEqual(suite.T(), "existing-id", customer.ID)
	// AND a new unique ID should have been assigned
	assert.NotEmpty(suite.T(), customer.ID)
	// AND DynamoDB PutItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *CustomerRepositoryTestSuite) Test_CustomerRetrieval_WithUnmarshalError_ShouldReturnError() {
	// GIVEN a CPF for a customer lookup
	cpf := "12345678901"

	// AND DynamoDB returns an item with invalid structure for unmarshaling
	// e.g. CPF is expected to be uint, but we return a complex map that can't be converted
	output := &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"CPF": {M: map[string]*dynamodb.AttributeValue{}}, // Invalid type for uint
		},
	}

	suite.mockDB.On("GetItem", mock.Anything).Return(output, nil).Once()

	// WHEN retrieving the customer by CPF from the repository
	result, err := suite.repository.GetByCpf(cpf)

	// THEN an error should be returned
	assert.Error(suite.T(), err)
	// AND no customer should be returned
	assert.Nil(suite.T(), result)
	// AND the error should indicate unmarshal failure
	assert.Contains(suite.T(), err.Error(), "failed to unmarshal customer")
	// AND DynamoDB GetItem should have been called
	suite.mockDB.AssertExpectations(suite.T())
}
