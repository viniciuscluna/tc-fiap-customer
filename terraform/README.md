# Terraform - DynamoDB Infrastructure

Esta pasta contém a infraestrutura como código (IaC) para provisionar a tabela DynamoDB no AWS Academy.

## 📋 Pré-requisitos

- Terraform >= 1.0
- Credenciais AWS Academy configuradas
- AWS CLI configurado

## 🚀 Como usar no AWS Academy

### 1. Configurar credenciais AWS Academy

No AWS Academy, vá em **AWS Details** → **AWS CLI** e copie as credenciais temporárias:

```bash
export AWS_ACCESS_KEY_ID="..."
export AWS_SECRET_ACCESS_KEY="..."
export AWS_SESSION_TOKEN="..."
export AWS_DEFAULT_REGION="us-east-1"
```

**Importante**: As credenciais do AWS Academy expiram após algumas horas. Você precisará renová-las periodicamente.

### 2. Inicializar Terraform

```bash
cd terraform
terraform init
```

### 3. Planejar as mudanças

```bash
terraform plan
```

### 4. Aplicar a infraestrutura

```bash
terraform apply
```

Digite `yes` quando solicitado.

### 5. Verificar a tabela criada

```bash
aws dynamodb describe-table --table-name Customer --region us-east-1
```

## 🔄 Atualizar credenciais expiradas

Quando as credenciais do AWS Academy expirarem:

1. Volte ao AWS Academy e copie novas credenciais
2. Execute novamente os comandos `export` acima
3. Continue usando `terraform plan` e `terraform apply`

## 🗑️ Destruir recursos (opcional)

**⚠️ Cuidado**: Isso apagará a tabela e todos os dados!

```bash
terraform destroy
```

## 📊 Recursos criados

- **DynamoDB Table**: `Customer`
  - Billing Mode: PAY_PER_REQUEST (on-demand)
  - Hash Key: CPF (Number)
  - Encryption: Enabled

## 🔧 Customização

Edite `terraform.tfvars` para customizar:

```hcl
aws_region  = "us-east-1"
table_name  = "Customer"
environment = "staging"
```

## ⚠️ Limitações AWS Academy

- Credenciais temporárias (expiram em ~3h)
- Não permite IAM roles personalizadas
- Alguns recursos AWS podem estar limitados
- Sempre use credenciais via environment variables
