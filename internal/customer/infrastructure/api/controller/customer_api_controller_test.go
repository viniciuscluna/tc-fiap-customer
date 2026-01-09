package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	apiController "github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/controller"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"
	mockController "github.com/viniciuscluna/tc-fiap-customer/mocks/customer/controller"
)

type CustomerApiControllerTestSuite struct {
	suite.Suite
	mockController *mockController.MockCustomerController
	router         *chi.Mux
}

func (suite *CustomerApiControllerTestSuite) SetupTest() {
	suite.mockController = mockController.NewMockCustomerController(suite.T())
	apiCtrl := apiController.NewCustomerController(suite.mockController)
	suite.router = chi.NewRouter()
	apiCtrl.RegisterRoutes(suite.router)
}

func TestCustomerApiControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerApiControllerTestSuite))
}

// Feature: Customer REST API - Get Endpoint
// Scenario: Retrieve customer information via HTTP

func (suite *CustomerApiControllerTestSuite) Test_CustomerRetrieval_ViaGetEndpoint_WithValidCPF_ShouldReturnCustomerSuccessfully() {
	// GIVEN a valid CPF parameter
	cpf := "12345678901"
	expectedResponse := &dto.GetCustomerResponseDto{
		ID:    "test-id",
		CPF: "12345678901",
		Name:  "John Doe",
		Email: "john@example.com",
	}

	suite.mockController.EXPECT().
		GetByCpf("12345678901").
		Return(expectedResponse, nil).
		Once()

	// WHEN a GET request is made to /v1/customer with valid CPF
	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf="+cpf, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 200 OK
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// AND the response body should contain the customer details
	var response dto.GetCustomerResponseDto
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedResponse.ID, response.ID)
	assert.Equal(suite.T(), expectedResponse.CPF, response.CPF)
	assert.Equal(suite.T(), expectedResponse.Name, response.Name)
	assert.Equal(suite.T(), expectedResponse.Email, response.Email)
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRetrieval_ViaGetEndpoint_WithInvalidCPFFormat_ShouldReturnBadRequest() {
	// GIVEN an invalid CPF parameter
	customerResponse := &dto.GetCustomerResponseDto{
		CPF:   "invalid",
		Name:  "Test User",
		Email: "test@test.com",
	}
	suite.mockController.EXPECT().GetByCpf("invalid").Return(customerResponse, nil).Once()
	
	// WHEN a GET request is made to /v1/customer with invalid CPF
	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf=invalid", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 200 (CPF validation happens at business layer)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRetrieval_ViaGetEndpoint_WithoutCPFParameter_ShouldReturnBadRequest() {
	// GIVEN no CPF parameter provided
	// WHEN a GET request is made to /v1/customer without CPF
	req := httptest.NewRequest(http.MethodGet, "/v1/customer", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	// AND the response should contain an error message about invalid CPF
	assert.Contains(suite.T(), w.Body.String(), "Invalid CPF")
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRetrieval_ViaGetEndpoint_WithNonExistentCPF_ShouldReturnNotFound() {
	// GIVEN a CPF for a customer that does not exist
	cpf := "99999999999"

	suite.mockController.EXPECT().
		GetByCpf("99999999999").
		Return(nil, errors.New("customer not found")).
		Once()

	// WHEN a GET request is made to /v1/customer with non-existent CPF
	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf="+cpf, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 404 Not Found
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRetrieval_ViaGetEndpoint_WithRepositoryError_ShouldReturnInternalServerError() {
	// GIVEN a valid CPF but the repository encounters a database error
	cpf := "12345678901"

	suite.mockController.EXPECT().
		GetByCpf("12345678901").
		Return(nil, errors.New("database error")).
		Once()

	// WHEN a GET request is made to /v1/customer
	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf="+cpf, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

// Feature: Customer REST API - Post Endpoint
// Scenario: Register a new customer via HTTP

func (suite *CustomerApiControllerTestSuite) Test_CustomerRegistration_ViaPostEndpoint_WithValidRequest_ShouldCreateCustomerSuccessfully() {
	// GIVEN a valid customer registration request
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		CPF: "98765432109",
	}

	suite.mockController.EXPECT().
		Add(requestDto).
		Return(nil).
		Once()

	// WHEN a POST request is made to /v1/customer with valid data
	body, _ := json.Marshal(requestDto)
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 201 Created
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	// AND the response body should contain a success message
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Customer created successfully", response["message"])
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRegistration_ViaPostEndpoint_WithInvalidJSON_ShouldReturnBadRequest() {
	// GIVEN an invalid JSON request body
	invalidJSON := []byte(`{"name": "John", "invalid}`)

	// WHEN a POST request is made to /v1/customer with malformed JSON
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRegistration_ViaPostEndpoint_WithControllerError_ShouldReturnInternalServerError() {
	// GIVEN a valid request but the controller encounters a validation error
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		CPF: "98765432109",
	}

	suite.mockController.EXPECT().
		Add(requestDto).
		Return(errors.New("validation error")).
		Once()

	// WHEN a POST request is made to /v1/customer
	body, _ := json.Marshal(requestDto)
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *CustomerApiControllerTestSuite) Test_CustomerRegistration_ViaPostEndpoint_WithEmptyBody_ShouldReturnBadRequest() {
	// GIVEN an empty request body
	// WHEN a POST request is made to /v1/customer with no body
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// THEN the response status should be 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}
