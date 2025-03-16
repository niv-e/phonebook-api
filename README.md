# Phone Book API

A simple RESTful API for managing a phone book, built with Golang using Clean Architecture and TDD.

## Project Structure
- `cmd/api/` - Application entry point (`main.go`).
- `internal/domain/` - Core domain layer containing business logic and entities.
  - `entities/` - Aggregate roots and value objects (e.g., `Contact`, `Phone`, `Address`).
  - `repositories/` - Interfaces for data access (e.g., `contact_repository.go`).
  - `errors.go` - Domain-specific error types and factory methods.
- `internal/application/` - Use cases orchestrating domain logic.
  - `commands/` - CQRS command definitions (e.g., `add_contact_command.go`).
  - `handlers/` - Handlers for commands (e.g., `add_contact_handler.go`).
- `internal/infrastructure/` - External systems (e.g., database, logging; to be implemented).
- `internal/delivery/` - API handlers and request structs.
  - `http/` - HTTP endpoints (e.g., `add_contact_request.go`).
- `tests/` - Test files (e.g., `add_contact_test.go`, `phone_test.go`).

## Architecture
- **Clean Architecture**: Separates domain logic from infrastructure and delivery, ensuring a focus on business rules.
- **CQRS Pattern**: Implements Command Query Responsibility Segregation to separate write (commands) and read (queries) operations, enhancing scalability and flexibility.

### Architecture Notes
- **UUID for Entity IDs**: The entities uses `uuid.UUID` (from `github.com/google/uuid`) as its `ID` field. This decision supports cloud-native scalability by ensuring globally unique identifiers without reliance on a centralized database sequence. UUIDs are particularly beneficial for future sharding, distributed systems, or multi-region deployments with PostgreSQL, which natively supports the `UUID` type. While this increases storage (16 bytes vs. 4-8 bytes for integers), it eliminates ID collision risks and simplifies integration in a distributed architecture.

### Domain-Driven Design (DDD) Decisions
- **Contact as Aggregate Root**: The `Contact` entity is defined as an aggregate root, encapsulating `Phone` and `Address` as value objects. All commands and queries interact with `Contact` as the entry point, ensuring consistency within its boundary. This aligns with DDD by centralizing business rules (e.g., validation of `Phone` in E.164 format) within the aggregate.
- **Nested Command Model**: Commands like `AddContactCommand` use a nested structure (e.g., `Phone` and `Address` as value objects) rather than a flat list of fields. This mirrors the `Contact` aggregate’s structure, reducing validation duplication and reinforcing domain cohesion. For example:
  ```go
  type AddContactCommand struct {
      FirstName string          `validate:"required"`
      LastName  string          `validate:"required"`
      Phone     entities.Phone  `validate:"required"`
      Address   entities.Address `validate:"required"`
  }

## Setup Instructions
1. Ensure Docker and Docker Compose are installed.
2. Build and run the app:
   ```bash
   docker-compose up --build