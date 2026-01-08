# Tech Challenge - Customer Microservice

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=coverage)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=bugs)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer)

Bem-vindo ao projeto Tech Challenge - Customer Microservice! Este microserviço é responsável pelo gerenciamento de clientes do sistema de restaurante.

## Índice

- [Sobre](#sobre)
- [Funcionalidades](#funcionalidades)
- [Tecnologias](#tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Configuração](#configuracao)
- [Uso](#uso)
- [Testes com BDD](#testes-com-bdd)
- [Qualidade de Código](#qualidade-de-código)
- [Migração PostgreSQL → DynamoDB](#migracao)

## Sobre

Este microserviço faz parte de uma arquitetura de microserviços para gestão de restaurantes. Ele foi desenvolvido em Golang com **DynamoDB** como banco de dados NoSQL, implementando **Clean Architecture** para separação clara entre regras de negócio e infraestrutura.

### Arquitetura Clean

O projeto segue os princípios da Clean Architecture, organizando o código em camadas bem definidas:

- **Domain (Domínio)**: Contém as entidades e regras de negócio centrais
- **Use Cases (Casos de Uso)**: Implementa a lógica de aplicação e orquestra as operações
- **Controllers**: Gerenciam o fluxo de dados entre a camada de apresentação e casos de uso
- **Infrastructure (Infraestrutura)**: Implementa detalhes técnicos como persistência DynamoDB e APIs REST
- **Presenters**: Formatam os dados para apresentação

## Funcionalidades

### API REST de Gerenciamento de Clientes

- ✅ **Cadastro de Clientes**: Registre novos clientes com CPF, nome e email
- ✅ **Consulta por CPF**: Busque informações de clientes pelo CPF
- ✅ **Validação de Dados**: Validação automática de CPF e campos obrigatórios
- ✅ **API RESTful**: Interface padronizada seguindo boas práticas REST
- ✅ **Documentação Swagger**: API totalmente documentada com OpenAPI 3.0

### Infraestrutura e DevOps

- ✅ **AWS DynamoDB**: Banco de dados NoSQL serverless com auto-scaling
- ✅ **DynamoDB Local**: Suporte para desenvolvimento local sem custos
- ✅ **Docker & Docker Compose**: Containerização completa da aplicação
- ✅ **AWS ECS Fargate**: Deploy serverless em containers na AWS
- ✅ **CI/CD GitHub Actions**: Pipeline automático de build, test e deploy
- ✅ **Infraestrutura como Código**: Terraform para provisionamento AWS
- ✅ **Kubernetes Ready**: Manifestos K8s para orquestração em clusters

### Qualidade e Arquitetura

- ✅ **Clean Architecture**: Separação clara de responsabilidades e camadas
- ✅ **Dependency Injection**: Gerenciamento com Uber FX
- ✅ **Testes Unitários**: Cobertura de código com testes automatizados
- ✅ **SonarCloud**: Análise contínua de qualidade de código
- ✅ **Mocks Automatizados**: Geração de mocks para testes isolados
- ✅ **Health Checks**: Endpoints de monitoramento de saúde da aplicação

### Banco de Dados

- **Tabela DynamoDB**: `tc-fiap-staging-customer`
- **Chave de Partição**: CPF (número único do cliente)
- **Modo de Cobrança**: Pay-per-request (ideal para cargas variáveis)
- **Criação Automática**: A tabela é criada automaticamente na primeira execução

## Tecnologias

- **Go (Golang)** - Linguagem de programação principal
- **AWS DynamoDB** - Banco de dados NoSQL
- **DynamoDB Local** - Para desenvolvimento local
- **Docker & Docker Compose** - Containerização
- **Kubernetes** - Orquestração de containers
- **Swagger (OpenAPI)** - Documentação da API
- **[go-chi](https://github.com/go-chi/chi)** - Router HTTP leve e performático
- **[GORM](https://gorm.io/)** - ORM para Go
- **[Uber FX](https://uber-go.github.io/fx/)** - Framework de injeção de dependências
- **Mercado Pago API** - Gateway de pagamento
- **Terraform** - Infraestrutura como código
- **AWS EKS** - Kubernetes gerenciado na AWS
- **Chi Router** - Framework HTTP minimalista e performático
- **Uber FX** - Framework de injeção de dependências
- **Swagger** - Documentação de API
- **Docker & Docker Compose** - Containerização
- **AWS SDK for Go** - Integração com DynamoDB
- **GitHub Actions** - CI/CD

## Estrutura do Projeto

O projeto segue os princípios da **Clean Architecture**, organizando o código em camadas bem definidas:

```
cmd/api/                    # Entrada da aplicação (main.go)
docs/                       # Documentação da API gerada pelo Swagger
http/                       # Arquivos para testar endpoints
internal/
  app/                      # Inicialização e injeção de dependências 
  customer/                 # Domínio de Clientes
    controller/             # Controllers (orquestração)
    domain/
      entities/             # Entidades do domínio
      repositories/         # Interfaces dos repositórios
    infrastructure/
      api/                  # Controllers HTTP e DTOs
      persistence/          # Implementação dos repositórios (DynamoDB)
    presenter/              # Formatação de dados para apresentação
    usecase/                # Casos de uso (regras de negócio)
      addCustomer/
      getbycpf/
      commands/             # Command objects (padrão Command)
pkg/                        # Pacotes compartilhados
  rest/                     # Interfaces HTTP comuns
  storage/dynamodb/         # Cliente e configuração DynamoDB
k8s/                        # Manifestos Kubernetes
```

## Configuração

### Rodando com Docker Compose

Para rodar o projeto localmente com Docker Compose:

1. Clone o repositório:
    ```bash
    git clone https://github.com/viniciuscluna/tc-fiap-customer.git
    cd tc-fiap-customer
    ```

2. Configure as variáveis de ambiente:
   - Copie `.env.example` para `.env`
   - Para desenvolvimento local, as configurações padrão já funcionam com DynamoDB Local
   - Para produção na AWS, configure as credenciais AWS apropriadas

3. Suba a aplicação e o DynamoDB Local com Docker Compose:
    ```bash
    docker compose up
    ```
   O servidor ficará disponível em `localhost:8080` e o DynamoDB Local em `localhost:8000`.

4. Acesse a documentação da API:
   - Abra [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) no navegador.

### Rodando Localmente (sem Docker)

1. Certifique-se de ter Go 1.24+ instalado
2. Instale as dependências:
   ```bash
   go mod download
   ```
3. Execute o DynamoDB Local (ou configure credenciais AWS):
   ```bash
   docker run -p 8000:8000 amazon/dynamodb-local
   ```
4. Configure o arquivo `.env` com `DYNAMODB_ENDPOINT=http://localhost:8000`
5. Execute a aplicação:
   ```bash
   go run cmd/api/main.go
   ```

## Uso

### Endpoints Disponíveis

#### Adicionar Cliente
```bash
POST /v1/customer
Content-Type: application/json

{
  "name": "João Silva",
  "email": "joao@example.com",
  "cpf": 12345678901
}
```

#### Consultar Cliente por CPF
```bash
GET /v1/customer?cpf=12345678901
```

### Swagger UI

Utilize a Swagger UI em `http://localhost:8080/swagger/index.html` para:
- Visualizar todos os endpoints disponíveis
- Testar requisições interativamente
- Ver exemplos de requisições e respostas

## Migração

### PostgreSQL → DynamoDB

Este projeto foi migrado de PostgreSQL para DynamoDB. Principais mudanças:

**Antes (PostgreSQL/GORM):**
- Banco relacional com tabelas SQL
- ORM GORM para mapeamento
- IDs auto-incrementais
- Constraints e relacionamentos SQL

**Depois (DynamoDB):**
- Banco NoSQL com tabelas DynamoDB
- AWS SDK para Go
- IDs gerados pela aplicação (timestamp-based)
- CPF como chave primária (partition key)
- Sem relacionamentos, design para consultas diretas

**Estrutura da Tabela Customer:**
- `CPF` (Partition Key): Chave primária
- `ID`: Identificador único gerado
- `Name`: Nome do cliente
- `Email`: Email do cliente
- `CreatedAt`: Timestamp de criação

## Arquivos HTTP

No diretório [`http/`](http/) está o arquivo `customer.http` com exemplos prontos para testar a API usando a extensão [REST Client para VS Code](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

**Como usar:**
1. Abra o arquivo `customer.http` no VS Code
2. Ajuste a variável `baseUrl` se necessário (ex: `@baseUrl = http://localhost:8080/`)
3. Clique em "Send Request" para executar e ver a resposta

Os exemplos cobrem:
- **Clientes**: Cadastro e consulta por CPF
- **Pagamentos**: Processamento e consulta de status
- **Webhooks**: Notificações do Mercado Pago

## Testes com BDD

Este projeto implementa **Behavior-Driven Development (BDD)** usando [testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite).

- **25 testes** em 6 camadas (presenter, use case, controller, API, repositório)
- Padrão **Given/When/Then** para clareza e legibilidade
- Sem dependências extras - apenas testify
- Testes descritivos que funcionam como documentação

### Executar Testes

```bash
go test ./internal/customer/... -v
```


## Qualidade de Código

Este projeto utiliza **SonarCloud** para análise contínua de qualidade de código, segurança e cobertura de testes.

### Métricas Monitoradas

| Métrica | Status | Objetivo |
|---------|--------|----------|
| **Quality Gate** | [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer) | ✅ Passed |
| **Coverage** | [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=coverage)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer) | ≥ 80% |
| **Bugs** | [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=bugs)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer) | 0 |
| **Code Smells** | [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer) | < 10 |
| **Security** | [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=viniciuscluna_tc-fiap-customer&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=viniciuscluna_tc-fiap-customer) | A |

### Análise Automática

A análise de código é executada automaticamente via **GitHub Actions** em:
- ✅ Cada push para branches `main` e `develop`
- ✅ Todos os Pull Requests

### Gerar Coverage Localmente

**Windows (PowerShell):**
```powershell
.\scripts\coverage.ps1
```

**Linux/Mac (Bash):**
```bash
chmod +x scripts/coverage.sh
./scripts/coverage.sh
```

Isso gera:
- `coverage.out` - Formato para SonarCloud
- `coverage.html` - Visualização no browser

## 🚀 Deploy na AWS

Este projeto usa **EC2 t2.micro** com Docker para deploy na AWS Academy.

### Setup Inicial

1. **Configure credenciais AWS Academy**:
   ```powershell
   # Copie do AWS Academy → AWS Details → Show
   $env:AWS_ACCESS_KEY_ID="ASIA..."
   $env:AWS_SECRET_ACCESS_KEY="..."
   $env:AWS_SESSION_TOKEN="..."
   $env:AWS_DEFAULT_REGION="us-east-1"
   ```

2. **Crie arquivo com credenciais** (terraform/terraform.tfvars):
   ```hcl
   aws_access_key_id     = "ASIA..."
   aws_secret_access_key = "..."
   aws_session_token     = "..."
   ```

3. **Crie a infraestrutura**:
   ```bash
   cd terraform
   terraform init
   terraform apply
   ```
   
   Isso cria:
   - ✅ Tabela DynamoDB `Customer`
   - ✅ Repositório ECR `tc-fiap-customer`
   - ✅ EC2 t2.micro com Docker
   - ✅ Security Group (portas 8080 e 22)

4. **Configure GitHub Secrets**:
   
   Em **Settings → Secrets and variables → Actions**, adicione:
   - `AWS_ACCESS_KEY_ID`
   - `AWS_SECRET_ACCESS_KEY`
   - `AWS_SESSION_TOKEN`

### Acessar a aplicação

Após o `terraform apply`, copie o IP:

```bash
terraform output application_url
# http://XX.XXX.XXX.XX:8080
```

**Endpoints**:
- 🌐 App: `http://IP:8080`
- ❤️ Health: `http://IP:8080/health`
- 📚 Swagger: `http://IP:8080/docs/index.html`

### Atualizar a aplicação

Após push de nova imagem no ECR:

```bash
# SSH na instância
ssh ec2-user@SEU_IP

# Atualizar (script já criado pelo Terraform)
sudo /usr/local/bin/update-app.sh
```

### Monitorar

```bash
# Via SSH
ssh ec2-user@SEU_IP

# Ver logs do container
sudo docker logs -f tc-fiap-customer

# Status do container
sudo docker ps
```

### 💰 Custo estimado
- **EC2 t2.micro**: Grátis (Free Tier) ou ~$8/mês
- **DynamoDB**: Pay-per-request (~$1-5/mês)
- **ECR**: ~$0.10/GB/mês

**Total**: ~$0-15/mês

### ⚠️ Renovar credenciais AWS Academy

As credenciais expiram a cada ~3 horas:

1. AWS Academy → AWS Details → Show (novas credenciais)
2. Atualize `terraform/terraform.tfvars`
3. Execute:
   ```bash
   terraform apply -var="aws_access_key_id=NOVA_KEY" \
                   -var="aws_secret_access_key=NOVA_SECRET" \
                   -var="aws_session_token=NOVO_TOKEN"
   ```
4. SSH na EC2 e rode: `sudo /usr/local/bin/update-app.sh`

### 🗑️ Destruir recursos

```bash
cd terraform
terraform destroy
```


⚠️ **Importante**: Credenciais AWS Academy expiram em ~3h e precisam ser renovadas periodicamente.