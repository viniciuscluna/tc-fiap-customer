package persistence

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/entities"
	"github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/repositories"
)

var (
	_ repositories.CustomerRepository = (*CustomerRepositoryImpl)(nil)
)

type CustomerRepositoryImpl struct {
	db *dynamodb.DynamoDB
}

func NewCustomerRepositoryImpl(db *dynamodb.DynamoDB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{db: db}
}

func (r *CustomerRepositoryImpl) GetByCpf(cpf uint) (*entities.Customer, error) {
	tableName := "Customer"

	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CPF": {
				N: aws.String(strconv.FormatUint(uint64(cpf), 10)),
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("customer not found")
	}

	customer := &entities.Customer{}
	err = dynamodbattribute.UnmarshalMap(result.Item, customer)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal customer: %w", err)
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) Add(customer *entities.Customer) error {
	tableName := "Customer"

	// Generate UUID for ID
	customer.ID = generateUUID()
	// Set created timestamp
	customer.CreatedAt = time.Now()

	// Marshal customer to DynamoDB attribute value map
	av, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return fmt.Errorf("failed to marshal customer: %w", err)
	}

	// Put item in DynamoDB
	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	if err != nil {
		return fmt.Errorf("failed to add customer: %w", err)
	}

	return nil
}

// generateUUID generates a simple UUID-like string
func generateUUID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), time.Now().Unix())
}
