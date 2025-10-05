# advance-go-api-learn-project

Projeto de API REST em Go para fins de aprendizado, utilizando PostgreSQL como banco de dados e sqlc para geração de queries. Estruturado em módulos internos para configuração, acesso a dados, manipulação de requisições, rotas, utilitários e modelos de domínio.

## Principais Características
- Estrutura modular (config, handlers, rotas, store, utils, models)
- Conexão com PostgreSQL via lib/pq
- Geração de código SQL com sqlc
- Handlers para saúde, usuários e núcleo
- Utilitários para JWT, senhas e respostas HTTP
- Configuração via arquivo `.env`
- Orquestração de containers com Podman Compose

## Como executar
1. Configure o arquivo `.env` com as variáveis necessárias.
2. Suba o banco de dados com `podman-compose up`.
3. Execute a aplicação Go com `go run main.go`.

## Arquitetura
Veja o arquivo [ARCHITECTURE.md](ARCHITECTURE.md) para detalhes sobre a arquitetura do projeto.