# advance-go-api-learn-project

Go REST API project for learning purposes, using PostgreSQL as the database and sqlc for query generation. The project is structured in internal modules for configuration, data access, request handling, routing, utilities, and domain models.

## Main Features
- Modular structure (config, handlers, routes, store, utils, models)
- PostgreSQL connection via lib/pq
- SQL code generation with sqlc
- Handlers for health, users, and core
- Utilities for JWT, password, and HTTP responses
- Configuration via `.env` file
- Container orchestration with Podman Compose

## How to run
1. Configure the `.env` file with the required variables.
2. Start the database with `podman-compose up`.
3. Run the Go application with `go run main.go`.

## Architecture
See the [ARCHITECTURE.md](ARCHITECTURE.md) file for details about the project architecture.