terraform {
  required_version = ">= 1.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  
  # AWS Academy usa credenciais temporárias via environment variables
  # Não configurar access_key/secret_key aqui
}

# DynamoDB Table - Customer
resource "aws_dynamodb_table" "customer" {
  name           = var.table_name
  billing_mode   = "PAY_PER_REQUEST"  # On-demand pricing (melhor para Academy)
  hash_key       = "CPF"

  attribute {
    name = "CPF"
    type = "N"  # Number type para CPF
  }

  # Optional: Enable point-in-time recovery (pode não estar disponível no Academy)
  # point_in_time_recovery {
  #   enabled = true
  # }

  # Optional: Server-side encryption (geralmente disponível)
  server_side_encryption {
    enabled = true
  }

  tags = {
    Name        = "Customer Table"
    Environment = var.environment
    Project     = "tc-fiap-customer"
    ManagedBy   = "Terraform"
  }
}

# Output útil para o pipeline
output "dynamodb_table_name" {
  description = "Nome da tabela DynamoDB"
  value       = aws_dynamodb_table.customer.name
}

output "dynamodb_table_arn" {
  description = "ARN da tabela DynamoDB"
  value       = aws_dynamodb_table.customer.arn
}
