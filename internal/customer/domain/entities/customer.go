package entities

import "time"

type Customer struct {
	ID        string    `json:"id" dynamodbav:"id"`
	CPF       string    `json:"cpf" dynamodbav:"cpf"`
	Name      string    `json:"name" dynamodbav:"name"`
	Email     string    `json:"email" dynamodbav:"email"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
}
