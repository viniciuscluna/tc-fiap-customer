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

func (suite *CustomerApiControllerTestSuite) TestGet_Success() {
	// Arrange
	cpf := "12345678901"
	expectedResponse := &dto.GetCustomerResponseDto{
		ID:    "test-id",
		CPF:   12345678901,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	suite.mockController.EXPECT().
		GetByCpf(uint(12345678901)).
		Return(expectedResponse, nil).
		Once()

	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf="+cpf, nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response dto.GetCustomerResponseDto
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedResponse.ID, response.ID)
	assert.Equal(suite.T(), expectedResponse.CPF, response.CPF)
	assert.Equal(suite.T(), expectedResponse.Name, response.Name)
	assert.Equal(suite.T(), expectedResponse.Email, response.Email)
}

func (suite *CustomerApiControllerTestSuite) TestGet_InvalidCPF() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf=invalid", nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid CPF")
}

func (suite *CustomerApiControllerTestSuite) TestGet_MissingCPF() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/v1/customer", nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid CPF")
}

func (suite *CustomerApiControllerTestSuite) TestGet_CustomerNotFound() {
	// Arrange
	cpf := "99999999999"

	suite.mockController.EXPECT().
		GetByCpf(uint(99999999999)).
		Return(nil, errors.New("customer not found")).
		Once()

	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf="+cpf, nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}

func (suite *CustomerApiControllerTestSuite) TestGet_InternalError() {
	// Arrange
	cpf := "12345678901"

	suite.mockController.EXPECT().
		GetByCpf(uint(12345678901)).
		Return(nil, errors.New("database error")).
		Once()

	req := httptest.NewRequest(http.MethodGet, "/v1/customer?cpf="+cpf, nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *CustomerApiControllerTestSuite) TestAdd_Success() {
	// Arrange
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		CPF:   98765432109,
	}

	suite.mockController.EXPECT().
		Add(requestDto).
		Return(nil).
		Once()

	body, _ := json.Marshal(requestDto)
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Customer created successfully", response["message"])
}

func (suite *CustomerApiControllerTestSuite) TestAdd_InvalidJSON() {
	// Arrange
	invalidJSON := []byte(`{"name": "John", "invalid}`)
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *CustomerApiControllerTestSuite) TestAdd_ControllerError() {
	// Arrange
	requestDto := &dto.AddCustomerRequestDto{
		Name:  "Jane Doe",
		Email: "jane@example.com",
		CPF:   98765432109,
	}

	suite.mockController.EXPECT().
		Add(requestDto).
		Return(errors.New("validation error")).
		Once()

	body, _ := json.Marshal(requestDto)
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *CustomerApiControllerTestSuite) TestAdd_EmptyBody() {
	// Arrange
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}
