# Ember Gate

Ember Gate is a modular Web3 infrastructure toolkit built in Go using Gin and GORM. It provides a node registry, user authentication, health checks, and an API gateway scaffold for building Web3 infrastructure services.

Key features:
- User registration and JWT authentication
- Node registry with ownership and basic health status
- Clean, modular project layout using Gin (HTTP) and GORM (ORM)
- Easy local development (Postgres or SQLite) and Docker Compose template

Architecture
------------
- cmd/server: application entrypoint
- internal/config: configuration loading
- internal/logger: structured logging (Logrus)
- internal/db: GORM database connection and migration
- internal/models: GORM models (User, Node)
- internal/services: domain logic (auth, node management)
- internal/handlers: HTTP handlers mapping to services
- internal/middleware: Gin middleware (JWT auth)
- migrations/: SQL migration examples

Installation
------------
Prerequisites:
- Go 1.20+
- PostgreSQL (optional for local; SQLite supported in tests)

Clone and build:

  git clone https://github.com/yourorg/ember-gate.git
  cd ember-gate
  go build ./cmd/server

Environment
-----------
Create a .env file or environment variables (see .env.example):
- DB_DSN: Postgres DSN, e.g. "host=localhost user=pguser password=pgpass dbname=ember sslmode=disable"
- PORT: e.g. 8080
- JWT_SECRET: a strong secret

Run (development):

  go run ./cmd/server

Or use Docker Compose (Postgres):

  docker-compose up --build

Usage / API
-----------
Base: http://localhost:8080/api/v1

Auth:
- POST /api/v1/register  { "email": "...", "password": "..." }
- POST /api/v1/login     { "email": "...", "password": "..." }

Nodes:
- POST /api/v1/nodes     (Authorization: Bearer <token>) { "name": "...", "url": "https://...", "provider": "..." }
- GET  /api/v1/nodes     list nodes (public)
- GET  /api/v1/nodes/:id detail

Example (register):

  curl -X POST http://localhost:8080/api/v1/register -H "Content-Type: application/json" -d '{"email":"alice@example.com","password":"secret"}'

Development notes
-----------------
- Tests use an in-memory SQLite database so they are fast and isolated.
- GORM AutoMigrate is used for development; for production, add proper migrations.

Contributing
------------
Contributions welcome. Please open issues or PRs. Follow the repository code style and write tests for new behavior.
