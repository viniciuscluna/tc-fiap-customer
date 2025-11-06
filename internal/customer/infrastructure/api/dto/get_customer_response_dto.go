package dto

import "time"

type GetCustomerResponseDto struct {
	ID        string    `json:"id"`
	CPF       uint      `json:"cpf"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
