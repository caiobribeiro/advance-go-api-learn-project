# Project Architecture

This document describes the current architecture of the `advance-go-api-learn-project`, a Go API developed for learning and experimentation purposes.

## Overview

The project is structured in internal modules that separate responsibilities for configuration, data access, request handling, routing, utilities, and models. It uses a simplified MVC pattern, with well-defined layers to facilitate maintenance and evolution.

## Folder Structure

- **main.go**: Application entry point. Loads configuration, initializes the database, handlers, and routes, and starts the HTTP server.
- **internal/**: Contains the main application logic, divided into submodules:
  - **db_config/**: PostgreSQL database configuration and connection.
  - **dtos/**: Data Transfer Object structures for requests.
  - **handlers/**: HTTP handler implementations (core, health, user).
  - **migrations/**: SQL scripts for schema and database queries.
  - **routes/**: HTTP route definitions and grouping by context (health, user, setup).
  - **store/**: Data access layer, models, and sqlc-generated queries.
  - **utils/**: Utility functions for JWT, password, and HTTP responses.
- **models/**: Domain models (e.g., user, blog).
- **server_config/**: Server configuration loading and structure.
- **.env**: Environment variables for local configuration.
- **sqlc.yaml**: sqlc configuration for Go code generation from SQL queries.
- **podman-compose.yaml**: Container orchestration for PostgreSQL and pgAdmin.

## Main Flow

1. **Configuration**: Loads environment variables and initializes server configuration.
2. **Database**: Connects to PostgreSQL using parameters from the `.env` file.
3. **Store**: Instantiates the data access layer via sqlc.
4. **Handlers**: Creates HTTP handlers, receiving database and store dependencies.
5. **Routes**: Defines routes and binds handlers.
6. **Server**: Starts the HTTP server on the defined port.

## Dependencies
- Go 1.25.1
- PostgreSQL
- sqlc
- jwt-go
- godotenv
- lib/pq
- x/crypto


