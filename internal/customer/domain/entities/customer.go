package entities

import "time"

type Customer struct {
	CPF       uint      `json:"cpf" dynamodbav:"CPF"`
	Name      string    `json:"name" dynamodbav:"Name"`
	Email     string    `json:"email" dynamodbav:"Email"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"CreatedAt"`
}

func (Customer) TableName() string {
	return "Customer"
}
