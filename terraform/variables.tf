variable "aws_region" {
  description = "AWS Region"
  type        = string
  default     = "us-east-1"
}

variable "table_name" {
  description = "Nome da tabela DynamoDB"
  type        = string
  default     = "Customer"
}

variable "environment" {
  description = "Environment name (staging, production, etc)"
  type        = string
  default     = "staging"
}
