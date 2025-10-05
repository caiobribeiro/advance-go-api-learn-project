# Arquitetura do Projeto

Este documento descreve a arquitetura atual do projeto `advance-go-api-learn-project`, uma API desenvolvida em Go para fins de aprendizado e experimentação.

## Visão Geral

O projeto é estruturado em módulos internos que separam responsabilidades de configuração, acesso a dados, manipulação de requisições, rotas, utilitários e modelos. Utiliza o padrão MVC simplificado, com camadas bem definidas para facilitar manutenção e evolução.

## Estrutura de Pastas

- **main.go**: Ponto de entrada da aplicação. Realiza o carregamento das configurações, inicializa o banco de dados, handlers e rotas, e inicia o servidor HTTP.
- **internal/**: Contém a lógica principal da aplicação, dividida em submódulos:
  - **db_config/**: Configuração e conexão com o banco de dados PostgreSQL.
  - **dtos/**: Estruturas para requisições (Data Transfer Objects).
  - **handlers/**: Implementação dos handlers HTTP (core, health, user).
  - **migrations/**: Scripts SQL para schema e queries do banco de dados.
  - **routes/**: Definição das rotas HTTP e agrupamento por contexto (health, user, setup).
  - **store/**: Camada de acesso ao banco de dados, modelos e queries geradas pelo sqlc.
  - **utils/**: Funções utilitárias para JWT, senhas e respostas HTTP.
- **models/**: Modelos de domínio (ex: usuário, blog).
- **server_config/**: Carregamento e estrutura das configurações do servidor.
- **.env**: Variáveis de ambiente para configuração local.
- **sqlc.yaml**: Configuração do sqlc para geração de código Go a partir de queries SQL.
- **podman-compose.yaml**: Orquestração de containers para banco de dados PostgreSQL e pgAdmin.

## Fluxo Principal

1. **Configuração**: Carregamento das variáveis de ambiente e inicialização das configurações do servidor.
2. **Banco de Dados**: Conexão com PostgreSQL usando parâmetros do arquivo `.env`.
3. **Store**: Instanciação da camada de acesso a dados via sqlc.
4. **Handlers**: Criação dos handlers HTTP, recebendo dependências do banco e store.
5. **Rotas**: Definição das rotas e vinculação dos handlers.
6. **Servidor**: Inicialização do servidor HTTP na porta definida.

## Dependências
- Go 1.25.1
- PostgreSQL
- sqlc
- jwt-go
- godotenv
- lib/pq
- x/crypto

## Sugestões de Melhoria

- Implementar testes automatizados (unitários e de integração).
- Adicionar documentação das rotas e exemplos de uso (ex: Swagger/OpenAPI).
- Melhorar tratamento de erros e respostas padronizadas.
- Separar camadas de serviço para lógica de negócio.
- Adicionar autenticação/autorização nas rotas sensíveis.
- Gerenciar migrações de banco com ferramenta dedicada (ex: golang-migrate).
- Adicionar logs estruturados e monitoramento.
- Criar scripts Makefile para facilitar comandos de build/test/run.

---

Esta arquitetura reflete o estado atual do projeto, sem adição de novos componentes ou funcionalidades além do que já existe.
