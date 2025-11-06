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
		http.Error(w, `{"error":"Invalid CPF parameter"}`, http.StatusBadRequest)
		return
	}

	cpfInt, err := strconv.ParseUint(cpf, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid CPF format"}`, http.StatusBadRequest)
		return
	}

	customer, err := h.controller.GetByCpf(uint(cpfInt))

	if err != nil {
		if err.Error() == "customer not found" {
			http.Error(w, `{"error":"Customer not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"Error processing request"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// @Summary     Add customer
// @Description Add customer
// @Tags        Customer
// @Accept      json
// @Produce     json
// @Param       body body dto.AddCustomerRequestDto true "Body"
// @Success     201  {object} map[string]string
// @Router      /v1/customer [post]
func (h *customerApiController) Add(w http.ResponseWriter, r *http.Request) {
	var customerRequest dto.AddCustomerRequestDto

	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		http.Error(w, `{"error":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	err := h.controller.Add(&customerRequest)

	if err != nil {
		http.Error(w, `{"error":"Error processing request"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer created successfully"})
}
