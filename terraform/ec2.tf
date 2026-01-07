# EC2 Instance para rodar a aplicação com Docker
# Solução compatível com AWS Academy (não requer IAM roles customizadas)

# Security Group para a aplicação
resource "aws_security_group" "app" {
  name        = "tc-fiap-customer-sg"
  description = "Security group for Customer microservice"

  # HTTP para a aplicação
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Application port"
  }

  # SSH (opcional - para debug)
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "SSH access"
  }

  # Permite todo tráfego de saída
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name      = "tc-fiap-customer-sg"
    Project   = "tc-fiap-customer"
    ManagedBy = "Terraform"
  }
}

# Variável para credenciais AWS (para passar para o container)
variable "aws_access_key_id" {
  description = "AWS Access Key ID"
  type        = string
  sensitive   = true
  default     = ""
}

variable "aws_secret_access_key" {
  description = "AWS Secret Access Key"
  type        = string
  sensitive   = true
  default     = ""
}

variable "aws_session_token" {
  description = "AWS Session Token (AWS Academy)"
  type        = string
  sensitive   = true
  default     = ""
}

# User data script para configurar Docker e rodar a aplicação
locals {
  user_data = <<-EOF
    #!/bin/bash
    set -e
    
    # Update system
    yum update -y
    
    # Install Docker
    yum install -y docker
    systemctl start docker
    systemctl enable docker
    usermod -a -G docker ec2-user
    
    # Install AWS CLI v2
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    ./aws/install
    
    # Configure AWS credentials (se fornecidas)
    mkdir -p /root/.aws
    cat > /root/.aws/credentials << 'CREDS'
[default]
aws_access_key_id = ${var.aws_access_key_id}
aws_secret_access_key = ${var.aws_secret_access_key}
aws_session_token = ${var.aws_session_token}
CREDS
    
    cat > /root/.aws/config << 'CONFIG'
[default]
region = ${var.aws_region}
CONFIG
    
    # Login no ECR
    aws ecr get-login-password --region ${var.aws_region} | docker login --username AWS --password-stdin ${aws_ecr_repository.customer.repository_url}
    
    # Pull e run da imagem
    docker pull ${aws_ecr_repository.customer.repository_url}:latest
    
    # Run container com variáveis de ambiente
    docker run -d \
      --name tc-fiap-customer \
      --restart unless-stopped \
      -p 8080:8080 \
      -e AWS_REGION=${var.aws_region} \
      -e AWS_ACCESS_KEY_ID=${var.aws_access_key_id} \
      -e AWS_SECRET_ACCESS_KEY=${var.aws_secret_access_key} \
      -e AWS_SESSION_TOKEN=${var.aws_session_token} \
      ${aws_ecr_repository.customer.repository_url}:latest
    
    # Create update script
    cat > /usr/local/bin/update-app.sh << 'SCRIPT'
#!/bin/bash
aws ecr get-login-password --region ${var.aws_region} | docker login --username AWS --password-stdin ${aws_ecr_repository.customer.repository_url}
docker pull ${aws_ecr_repository.customer.repository_url}:latest
docker stop tc-fiap-customer || true
docker rm tc-fiap-customer || true
docker run -d \
  --name tc-fiap-customer \
  --restart unless-stopped \
  -p 8080:8080 \
  -e AWS_REGION=${var.aws_region} \
  -e AWS_ACCESS_KEY_ID=${var.aws_access_key_id} \
  -e AWS_SECRET_ACCESS_KEY=${var.aws_secret_access_key} \
  -e AWS_SESSION_TOKEN=${var.aws_session_token} \
  ${aws_ecr_repository.customer.repository_url}:latest
SCRIPT
    
    chmod +x /usr/local/bin/update-app.sh
    
    echo "Application deployed successfully!"
  EOF
}

# EC2 Instance
resource "aws_instance" "app" {
  ami           = "ami-0453ec754f44f9a4a" # Amazon Linux 2023 us-east-1
  instance_type = "t2.micro" # Free tier eligible
  
  vpc_security_group_ids = [aws_security_group.app.id]
  key_name               = data.aws_key_pair.existing.key_name
  
  user_data = local.user_data
  
  # Aumentar storage se necessário
  root_block_device {
    volume_size = 20
    volume_type = "gp3"
  }
  
  tags = {
    Name      = "tc-fiap-customer"
    Project   = "tc-fiap-customer"
    ManagedBy = "Terraform"
  }
}

# Outputs
output "ec2_public_ip" {
  description = "IP público da instância EC2"
  value       = aws_instance.app.public_ip
}

output "ec2_public_dns" {
  description = "DNS público da instância EC2"
  value       = aws_instance.app.public_dns
}

output "application_url" {
  description = "URL da aplicação"
  value       = "http://${aws_instance.app.public_ip}:8080"
}

output "application_health" {
  description = "Health check URL"
  value       = "http://${aws_instance.app.public_ip}:8080/health"
}

output "application_docs" {
  description = "Swagger documentation URL"
  value       = "http://${aws_instance.app.public_ip}:8080/docs/index.html"
}

output "ssh_command" {
  description = "Comando SSH para conectar (requer key pair)"
  value       = "ssh -i your-key.pem ec2-user@${aws_instance.app.public_ip}"
}
