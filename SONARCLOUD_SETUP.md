# SonarCloud Setup Guide - TC FIAP Customer

Este guia descreve o processo completo de integração do projeto com o SonarCloud para análise de qualidade de código.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Pré-requisitos](#pré-requisitos)
- [Setup Inicial](#setup-inicial)
- [Configuração do GitHub](#configuração-do-github)
- [Execução Local](#execução-local)
- [Análise de Resultados](#análise-de-resultados)
- [Troubleshooting](#troubleshooting)

## 🎯 Visão Geral

**SonarCloud** é uma plataforma de análise de qualidade de código que fornece:

- ✅ Detecção de bugs e vulnerabilidades
- ✅ Análise de code smells (problemas de manutenibilidade)
- ✅ Medição de cobertura de testes
- ✅ Tracking de débito técnico
- ✅ Quality Gates automáticos
- ✅ **100% gratuito para projetos públicos**

## ✅ Pré-requisitos

Antes de começar, certifique-se de ter:

- [x] Conta no GitHub
- [x] Repositório público ou privado (free tier funciona apenas para públicos)
- [x] GitHub Actions habilitado no repositório
- [x] Go 1.24+ instalado localmente

## 🚀 Setup Inicial

### Passo 1: Criar Conta no SonarCloud

1. Acesse [https://sonarcloud.io](https://sonarcloud.io)
2. Clique em **"Log in"** ou **"Start Free"**
3. Escolha **"Sign up with GitHub"**
4. Autorize o SonarCloud a acessar sua conta GitHub
5. Complete o cadastro

### Passo 2: Criar Organização

1. No SonarCloud, clique em **"+"** (canto superior direito) → **"Analyze new project"**
2. Selecione **"Create an organization"**
3. Escolha sua conta ou organização do GitHub
4. Defina um nome para a organização (ex: `viniciuscluna`)
5. Escolha o plano **Free** (para projetos open source)
6. Clique em **"Continue"**

### Passo 3: Importar Repositório

1. Na lista de repositórios, localize **`tc-fiap-customer`**
2. Clique em **"Set Up"**
3. Escolha **"With GitHub Actions"** (recomendado)
4. Copie o **Project Key** gerado (ex: `viniciuscluna_tc-fiap-customer`)
5. Copie o **Organization Key** (ex: `viniciuscluna`)

### Passo 4: Gerar Token

1. No SonarCloud, vá em **My Account** (ícone do usuário) → **Security**
2. Em **"Generate Tokens"**, digite um nome (ex: `tc-fiap-customer-token`)
3. Clique em **"Generate"**
4. **Copie o token e guarde em local seguro** (não será mostrado novamente)

## 🔧 Configuração do GitHub

### Passo 1: Adicionar Secret ao GitHub

1. Acesse seu repositório no GitHub
2. Vá em **Settings** → **Secrets and variables** → **Actions**
3. Clique em **"New repository secret"**
4. Preencha:
   - **Name:** `SONAR_TOKEN`
   - **Secret:** Cole o token copiado do SonarCloud
5. Clique em **"Add secret"**

### Passo 2: Atualizar sonar-project.properties

Edite o arquivo [`sonar-project.properties`](../sonar-project.properties):

```properties
# Substituir pelos seus valores
sonar.projectKey=SEU_USUARIO_tc-fiap-customer
sonar.organization=SEU_USUARIO

# Exemplo:
sonar.projectKey=viniciuscluna_tc-fiap-customer
sonar.organization=viniciuscluna
```

**Importante:** Use os valores copiados no Passo 3 do Setup Inicial.

### Passo 3: Verificar Workflow

O arquivo [`.github/workflows/sonarcloud.yml`](../.github/workflows/sonarcloud.yml) já está configurado e será executado automaticamente quando:

- Você fizer push para `main` ou `develop`
- Alguém abrir/atualizar um Pull Request

## 🧪 Execução Local

### Gerar Coverage Report

**Windows (PowerShell):**
```powershell
.\scripts\coverage.ps1
```

**Linux/Mac (Bash):**
```bash
chmod +x scripts/coverage.sh
./scripts/coverage.sh
```

**Manual:**
```bash
# Gerar coverage
go test ./... -coverprofile=coverage.out -covermode=atomic

# Ver resumo no terminal
go tool cover -func=coverage.out

# Gerar HTML para visualização
go tool cover -html=coverage.out -o coverage.html
```

### Executar SonarScanner Localmente (Opcional)

**1. Instalar SonarScanner CLI:**

**Windows (Chocolatey):**
```powershell
choco install sonarscanner
```

**Mac (Homebrew):**
```bash
brew install sonar-scanner
```

**Linux (Manual):**
```bash
wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-5.0.1.3006-linux.zip
unzip sonar-scanner-cli-5.0.1.3006-linux.zip
export PATH=$PATH:$(pwd)/sonar-scanner-5.0.1.3006-linux/bin
```

**2. Executar análise:**
```bash
# Gerar coverage primeiro
go test ./... -coverprofile=coverage.out

# Executar SonarScanner
sonar-scanner \
  -Dsonar.login=SEU_TOKEN_AQUI
```

## 📊 Análise de Resultados

### Dashboard do SonarCloud

Acesse: `https://sonarcloud.io/project/overview?id=SEU_PROJECT_KEY`

**Métricas Principais:**

| Métrica | Descrição | Meta |
|---------|-----------|------|
| **Quality Gate** | Status geral do projeto | ✅ Passed |
| **Coverage** | % de código coberto por testes | ≥ 80% |
| **Bugs** | Problemas funcionais detectados | 0 |
| **Vulnerabilities** | Problemas de segurança | 0 |
| **Code Smells** | Problemas de manutenibilidade | < 10 |
| **Duplications** | % de código duplicado | < 3% |
| **Security Hotspots** | Código sensível a revisar | 100% reviewed |

### Quality Gate Conditions (Padrão)

Para que o Quality Gate seja aprovado, o código deve atender:

```yaml
✅ Coverage on New Code ≥ 80%
✅ Duplicated Lines on New Code ≤ 3%
✅ Maintainability Rating on New Code = A
✅ Reliability Rating on New Code = A
✅ Security Rating on New Code = A
✅ Security Hotspots Reviewed ≥ 100%
```

### Badges

Adicione badges ao README.md (substitua `PROJECT_KEY`):

```markdown
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=PROJECT_KEY&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=PROJECT_KEY)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=PROJECT_KEY&metric=coverage)](https://sonarcloud.io/summary/new_code?id=PROJECT_KEY)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=PROJECT_KEY&metric=bugs)](https://sonarcloud.io/summary/new_code?id=PROJECT_KEY)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=PROJECT_KEY&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=PROJECT_KEY)
```

## 🔍 Interpretando os Resultados

### Bugs
**O que são:** Problemas no código que podem causar comportamento inesperado.

**Exemplos em Go:**
- Nil pointer dereference
- Resource leaks (não fechar conexões)
- Race conditions

**Como resolver:** Revisar cada bug reportado e corrigir conforme recomendação.

### Vulnerabilities
**O que são:** Problemas de segurança que podem ser explorados.

**Exemplos em Go:**
- SQL injection
- Path traversal
- Weak cryptography

**Como resolver:** Aplicar patches de segurança e seguir best practices.

### Code Smells
**O que são:** Problemas que não causam bugs, mas dificultam manutenção.

**Exemplos em Go:**
- Funções muito longas (>50 linhas)
- Complexidade ciclomática alta
- Código duplicado
- Nomes de variáveis pouco descritivos

**Como resolver:** Refatorar código seguindo princípios SOLID e Clean Code.

### Coverage
**O que é:** Percentual de código executado pelos testes.

**Como melhorar:**
- Adicionar testes para funções não cobertas
- Testar casos de erro e edge cases
- Ver relatório HTML para identificar gaps: `coverage.html`

## 🛠️ Troubleshooting

### Problema: "Could not find a default branch"

**Causa:** Repositório sem commits ou branch principal não configurada.

**Solução:**
```bash
git checkout -b main  # ou develop
git add .
git commit -m "Initial commit"
git push -u origin main
```

---

### Problema: "Invalid token" no GitHub Actions

**Causa:** Token do SonarCloud expirado ou incorreto.

**Solução:**
1. Gere um novo token no SonarCloud (My Account → Security)
2. Atualize o secret `SONAR_TOKEN` no GitHub

---

### Problema: "No coverage report found"

**Causa:** Arquivo `coverage.out` não foi gerado ou está em local errado.

**Solução:**
Verifique se o workflow está gerando coverage:
```yaml
- name: Run tests with coverage
  run: go test ./... -coverprofile=coverage.out
```

---

### Problema: "Quality Gate failed"

**Causa:** Código não atende aos critérios mínimos de qualidade.

**Solução:**
1. Acesse o dashboard do SonarCloud
2. Veja quais condições falharam
3. Corrija os problemas identificados
4. Faça novo commit/push

---

### Problema: "Analysis timed out"

**Causa:** Projeto muito grande ou muitos arquivos para analisar.

**Solução:**
Adicione mais exclusões no `sonar-project.properties`:
```properties
sonar.exclusions=**/vendor/**,**/node_modules/**,**/dist/**
```

---

## 📚 Recursos Adicionais

### Documentação Oficial
- [SonarCloud Docs](https://docs.sonarcloud.io/)
- [SonarCloud for Go](https://docs.sonarcloud.io/advanced-setup/languages/go/)
- [GitHub Actions Integration](https://docs.sonarcloud.io/advanced-setup/ci-based-analysis/github-actions-for-sonarcloud/)

### Best Practices
- [Clean Code em Go](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Testing Best Practices](https://go.dev/doc/tutorial/add-a-test)

### Comunidade
- [SonarSource Community](https://community.sonarsource.com/)
- [Stack Overflow - SonarCloud](https://stackoverflow.com/questions/tagged/sonarcloud)

---

## 🎯 Checklist de Configuração

Use este checklist para garantir que tudo está configurado:

- [ ] Conta criada no SonarCloud
- [ ] Organização criada (free tier)
- [ ] Projeto `tc-fiap-customer` importado
- [ ] Token gerado no SonarCloud
- [ ] Secret `SONAR_TOKEN` adicionado no GitHub
- [ ] Arquivo `sonar-project.properties` atualizado com project key correto
- [ ] Workflow `.github/workflows/sonarcloud.yml` commitado
- [ ] Primeiro push/PR acionou a análise
- [ ] Dashboard do SonarCloud mostrando resultados
- [ ] Badges adicionados ao README.md

---

## ✅ Próximos Passos

Após configuração completa:

1. **Monitorar Quality Gate:** Garanta que todos os PRs passem no Quality Gate
2. **Melhorar Coverage:** Meta inicial: 80% de cobertura
3. **Resolver Code Smells:** Priorize por severidade (Blocker → Critical → Major)
4. **Configurar PR Decoration:** Comentários automáticos do SonarCloud em PRs
5. **Definir Metas:** Estabeleça metas de qualidade para o time

---

**Configuração completa! 🎉**

O SonarCloud agora analisará automaticamente cada push e pull request, fornecendo feedback sobre qualidade de código, segurança e cobertura de testes.
