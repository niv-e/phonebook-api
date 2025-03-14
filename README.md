# Phone Book API

A simple RESTful API for managing a phone book, built with Golang using Clean Architecture and TDD.

## Project Structure
- `cmd/api/` - Application entry point (main.go).
- `internal/domain/` - Business entities (e.g., Contact).
- `internal/application/` - Use cases/services.
- `internal/infrastructure/` - External systems (DB, logging).
- `internal/delivery/` - API handlers (HTTP).
- `tests/` - Test files.

## Architecture
- **Clean Architecture**: Separates domain logic from infrastructure and delivery.

## Setup Instructions
1. Ensure Docker and Docker Compose are installed.
2. Build and run the app:
   ```bash
   docker-compose up --build