# Tech Challenge

Bem-vindo ao projeto Tech Challenge Fase 2! Este documento apresenta uma visão geral do sistema, instruções para rodar localmente, detalhes de uso e informações para contribuir.

## Índice

- [Sobre](#sobre)
- [Funcionalidades](#funcionalidades)
- [Tecnologias](#tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Configuração](#configuracao)
- [Uso](#uso)
- [Arquivos HTTP](#arquivos-http)

## Sobre

Esse projeto apresenta uma solução para gestão de restaurantes, com controle de pedidos, produtos, clientes e pagamentos. O backend foi desenvolvido em Golang com PostgreSQL, implementando **Clean Architecture**, que promove a separação clara entre regras de negócio e infraestrutura.

### Arquitetura Clean

O projeto segue os princípios da Clean Architecture, organizando o código em camadas bem definidas:

- **Domain (Domínio)**: Contém as entidades e regras de negócio centrais
- **Use Cases (Casos de Uso)**: Implementa a lógica de aplicação e orquestra as operações
- **Controllers**: Gerenciam o fluxo de dados entre a camada de apresentação e casos de uso
- **Infrastructure (Infraestrutura)**: Implementa detalhes técnicos como persistência, APIs externas e interfaces web
- **Presenters**: Formatam os dados para apresentação

### Diagrama de Arquitetura

Para visualizar a arquitetura completa incluindo a infraestrutura Kubernetes, acesse o [Diagrama no Miro](https://miro.com/app/board/uXjVIDe9VAo=/) que contém:
- Estrutura da Clean Architecture
- Fluxo de dados entre as camadas
- Infraestrutura Kubernetes (pods, services, deployments)
- Integração com banco de dados PostgreSQL
- Gateway de pagamento Mercado Pago 

## Funcionalidades

- Cadastro, edição, remoção e listagem de produtos
- Cadastro e gerenciamento de clientes
- Criação de pedidos e acompanhamento de status
- Processamento de pagamentos com integração ao Mercado Pago
- Webhooks para notificações de pagamento
- API documentada com Swagger (OpenAPI)
- Arquitetura preparada para  Kubernetes

## Tecnologias

- **Go (Golang)** - Linguagem de programação principal
- **PostgreSQL** - Banco de dados relacional
- **Docker & Docker Compose** - Containerização
- **Kubernetes** - Orquestração de containers
- **Swagger (OpenAPI)** - Documentação da API
- **[go-chi](https://github.com/go-chi/chi)** - Router HTTP leve e performático
- **[GORM](https://gorm.io/)** - ORM para Go
- **[Uber FX](https://uber-go.github.io/fx/)** - Framework de injeção de dependências
- **Mercado Pago API** - Gateway de pagamento
- **Terraform** - Infraestrutura como código
- **AWS EKS** - Kubernetes gerenciado na AWS
- **GitHub Actions** - CI/CD

## Estrutura do Projeto

O projeto segue os princípios da **Clean Architecture**, organizando o código em camadas bem definidas:

```
cmd/api/                    # Entrada da aplicação (main.go)
docs/                       # Documentação da API, Swagger e Banco de Dados
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
      persistence/          # Implementação dos repositórios
    presenter/              # Formatação de dados para apresentação
    usecase/                # Casos de uso (regras de negócio)
      addCustomer/
      getByCpf/
      commands/             # Command objects (padrão Command)
  order/                    # Domínio de Pedidos
  payment/                  # Domínio de Pagamentos
  product/                  # Domínio de Produtos
pkg/                        # Pacotes compartilhados
  rest/                     # Interfaces HTTP comuns
  storage/postgres/         # Configuração do banco de dados
k8s/                        # Manifestos Kubernetes
terraform/                  # Infraestrutura como código (AWS EKS)
```

## Configuração

### Rodando com Docker Compose

Para rodar o projeto localmente com Docker Compose:

1. Clone o repositório:
    ```bash
    git clone https://github.com/viniciuscluna/tc-fiap-50.git
    cd tc-fiap-50
    ```

2. Configure as variáveis de ambiente:
   - Copie `.env.example` para `.env` e preencha com os dados do banco e Mercado Pago.
   - Para produção na AWS, configure os secrets no GitHub Actions (veja [documentação de secrets](docs/md/github-secrets.md)).

3. Suba a aplicação e o banco com Docker Compose:
    ```bash
    docker compose up
    ```
   O servidor ficará disponível em `localhost:8080` e o banco Postgres será iniciado automaticamente.

4. Acesse a documentação da API:
   - Abra [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) no navegador.

### Rodando com Kubernetes (Kind)

Como alternativa, é possível rodar a aplicação em um ambiente Kubernetes local utilizando o Kind. Esta abordagem é ideal para testar os manifestos de implantação e simular um ambiente de produção.

- **[Guia de Configuração do Kubernetes com Kind](docs/md/k8s.md)**: Siga este guia detalhado para criar o cluster local, implantar a aplicação e interagir com o ambiente.

### Rodando na AWS com EKS

Para ambiente de produção, recomendamos o deploy na AWS utilizando EKS (Elastic Kubernetes Service) com infraestrutura gerenciada pelo Terraform.

- **[Guia de Deploy na AWS](docs/md/aws-deploy.md)**: Documentação completa para deploy em produção na AWS com EKS, Terraform e GitHub Actions.

#### Deploy Separado de Banco de Dados

O projeto agora utiliza uma arquitetura separada onde o banco de dados PostgreSQL é gerenciado pelo projeto `tc-fiap-database`:

1. **Deploy do Banco de Dados** (pré-requisito):
   ```bash
   cd ../tc-fiap-database/k8s
   ./deploy-database.sh
   ```

2. **Deploy da Aplicação**:
   ```bash
   cd k8s
   ./deploy-app.sh
   ```

Esta separação permite:
- Gerenciamento independente do banco de dados
- Reutilização do banco entre diferentes aplicações
- Melhor organização da infraestrutura

### CI/CD

- **[Documentação do Pipeline de CI/CD](docs/md/cicd.md)**: Entenda como nosso pipeline unificado funciona, desde os testes até o deploy automático na AWS ou ambiente local, dependendo da branch.

### Banco de Dados

- **[Documentação do Banco de Dados](docs/md/database-documentation.md)**: Documentação completa da estrutura do banco de dados, incluindo ERD, descrição das tabelas, relacionamentos e otimizações de performance com índices.

## Uso

- Utilize a Swagger UI para testar os endpoints e ver exemplos de requisições e respostas.
- Os logs da aplicação aparecem no terminal.

## Arquivos HTTP

No diretório [`http/`](http/) estão arquivos `.http` (ex: `product.http`, `order.http`). Eles trazem exemplos prontos para testar a API usando a extensão [REST Client para VS Code](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) ou ferramentas similares.

**Como usar:**
1. Abra um arquivo `.http` no VS Code.
2. Ajuste a variável `baseUrl` se necessário (ex: `@baseUrl = http://localhost:8080/`).
3. Clique em "Send Request" para executar e ver a resposta.

Esses arquivos cobrem:
- **Produtos**: Cadastro, consulta, atualização e remoção
- **Pedidos**: Criação, consulta e atualização de status
- **Clientes**: Cadastro e consulta por CPF
- **Pagamentos**: Processamento e consulta de status
- **Webhooks**: Notificações do Mercado Pago