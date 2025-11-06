package dto

type AddCustomerRequestDto struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@doe.com"`
	CPF   uint   `json:"cpf" example:"12345678901"`
}
