# ECR Repository for Customer microservice
resource "aws_ecr_repository" "customer" {
  name                 = "tc-fiap-customer"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  # Encryption
  encryption_configuration {
    encryption_type = "AES256"
  }

  tags = {
    Name        = "Customer Microservice Repository"
    Environment = var.environment
    Project     = "tc-fiap-customer"
    ManagedBy   = "Terraform"
  }
}

# Lifecycle policy to keep only recent images (economizar espaço)
resource "aws_ecr_lifecycle_policy" "customer" {
  repository = aws_ecr_repository.customer.name

  policy = jsonencode({
    rules = [{
      rulePriority = 1
      description  = "Keep last 10 images"
      selection = {
        tagStatus     = "any"
        countType     = "imageCountMoreThan"
        countNumber   = 10
      }
      action = {
        type = "expire"
      }
    }]
  })
}

# Output para usar no pipeline
output "ecr_repository_url" {
  description = "URL do repositório ECR"
  value       = aws_ecr_repository.customer.repository_url
}

output "ecr_repository_name" {
  description = "Nome do repositório ECR"
  value       = aws_ecr_repository.customer.name
}

output "ecr_repository_arn" {
  description = "ARN do repositório ECR"
  value       = aws_ecr_repository.customer.arn
}
