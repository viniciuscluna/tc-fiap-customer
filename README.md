# Tech Challenge - Customer Microservice

Bem-vindo ao projeto Tech Challenge - Customer Microservice! Este microserviço é responsável pelo gerenciamento de clientes do sistema de restaurante.

## Índice

- [Sobre](#sobre)
- [Funcionalidades](#funcionalidades)
- [Tecnologias](#tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Configuração](#configuracao)
- [Uso](#uso)
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

- Cadastro de clientes (nome, email, CPF)
- Consulta de cliente por CPF
- API documentada com Swagger (OpenAPI)
- Suporte a DynamoDB (AWS e Local)
- Arquitetura preparada para Kubernetes

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

### Documentação Completa

Veja [BDD_TESTING_GUIDE.md](BDD_TESTING_GUIDE.md) para:
- Padrão completo de Given/When/Then
- Convenção de nomenclatura
- Melhores práticas
- Exemplos por camada
- Debugging de testes
