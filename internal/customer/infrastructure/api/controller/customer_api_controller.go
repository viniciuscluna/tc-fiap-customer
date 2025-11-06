package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	customerController "github.com/viniciuscluna/tc-fiap-customer/internal/customer/controller"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/dto"
)

type customerApiController struct {
	controller customerController.CustomerController
}

func NewCustomerController(controller customerController.CustomerController) *customerApiController {
	return &customerApiController{
		controller: controller,
	}
}

func (c *customerApiController) RegisterRoutes(r chi.Router) {
	prefix := "/v1/customer"
	r.Get(prefix, c.Get)
	r.Post(prefix, c.Add)
}

// @Summary     Get customer
// @Description Get customer by CPF
// @Tags        Customer
// @Accept      json
// @Produce     json
// @Param       cpf query uint true "CPF"
// @Success     200  {object} dto.GetCustomerResponseDto
// @Router      /v1/customer [get]
func (h *customerApiController) Get(w http.ResponseWriter, r *http.Request) {
	cpf := r.URL.Query().Get("cpf")

	if cpf == "" {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	cpfInt, err := strconv.ParseUint(cpf, 10, 64)
	if err != nil {
		http.Error(w, "Invalid cpf parameter", http.StatusBadRequest)
		return
	}

	customer, err := h.controller.GetByCpf(uint(cpfInt))

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// @Summary     Add customer
// @Description Add customer
// @Tags        Customer
// @Accept      json
// @Produce     json
// @Param       body body dto.AddCustomerRequestDto true "Body"
// @Success     201
// @Router      /v1/customer [post]
func (h *customerApiController) Add(w http.ResponseWriter, r *http.Request) {
	var customerRequest dto.AddCustomerRequestDto

	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	err := h.controller.Add(&customerRequest)

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
