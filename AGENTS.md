# Repository Guidelines

## Project Structure & Module Organization
- `cmd/api/main.go` is the entrypoint for the HTTP API server.
- `interface/` contains Gin router and HTTP handlers.
- `usecase/` holds application use cases (create/list/complete/delete todo).
- `domain/` defines entities and repository interfaces.
- `infrastructure/` implements persistence and database wiring (Gorm, SQLite/Postgres).
- `registry/` builds dependencies (simple DI container).
- `todo.db` is the default local SQLite database file.

## Build, Test, and Development Commands
- `go run cmd/api/main.go` starts the API with default SQLite settings.
- `go build ./...` compiles all packages (use `go build -o todo-api cmd/api/main.go` for a binary).
- `go test ./...` runs unit tests (currently none are checked in).
- `docker compose up --build` runs Postgres + API containers using `docker-compose.yml`.

## Coding Style & Naming Conventions
- Follow standard Go formatting: run `gofmt` on changed files.
- Exported identifiers use `PascalCase`; unexported use `camelCase`.
- File names are lowercase with underscores (e.g., `todo_repository.go`).
- Keep packages aligned with the DDD layers (`domain`, `usecase`, `interface`, `infrastructure`).

## Testing Guidelines
- Use Go’s `testing` package for `_test.go` files.
- Name tests `TestXxx` and prefer table-driven tests for handlers and use cases.
- When adding tests, run `go test ./...` and include any required fixtures or fakes.

## Commit & Pull Request Guidelines
- Commit messages follow a conventional prefix: `feat:`, `fix:`, `style:`, `init`.
- Keep commits scoped to a single logical change.
- Pull requests should include a short summary, test results (e.g., `go test ./...`), and any config changes (env vars, DB changes).

## Configuration Tips
- Environment variables are read via `.env` or the shell.
- Key variables: `DB_DRIVER` (`sqlite` or `postgres`), `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`, `PORT`.
- Defaults in code target SQLite with `todo.db` for local development.
