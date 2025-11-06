package dynamodb

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Table names constants
const (
	CustomerTableName = "Customer"
)

// NewDynamoDBClient creates and returns a new DynamoDB client
func NewDynamoDBClient() *dynamodb.DynamoDB {
	// Get AWS configuration from environment variables
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1" // Default region
	}

	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// Configure AWS session
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	// If running locally or with custom endpoint (e.g., DynamoDB Local)
	if endpoint != "" {
		awsConfig.Endpoint = aws.String(endpoint)

		// For local development, use dummy credentials
		if awsAccessKey == "" || awsSecretKey == "" {
			awsConfig.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "")
		}
	}

	// If credentials are provided, use them
	if awsAccessKey != "" && awsSecretKey != "" {
		awsConfig.Credentials = credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")
	}

	// Create AWS session
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	log.Println("DynamoDB client initialized successfully")

	// Ensure table exists
	ensureTableExists(svc)

	return svc
}

// ensureTableExists creates the Customer table if it doesn't exist
func ensureTableExists(svc *dynamodb.DynamoDB) {
	// Check if table exists
	_, err := svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(CustomerTableName),
	})

	if err == nil {
		log.Printf("Table %s already exists\n", CustomerTableName)
		return
	}

	// Create table
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(CustomerTableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CPF"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CPF"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"),
	}

	_, err = svc.CreateTable(input)
	if err != nil {
		log.Printf("Warning: Failed to create table %s: %v\n", CustomerTableName, err)
		return
	}

	log.Printf("Table %s created successfully\n", CustomerTableName)
}
