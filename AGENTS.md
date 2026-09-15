# Repository Instructions

## Repository Purpose

This repository is a reusable Go project skeleton.

It is intentionally designed to serve as the starting point for new projects.

Architectural directories may exist as placeholders even when they are currently empty.

Do not remove an empty directory solely because it has no current implementation if it represents a documented extension point of the skeleton.

However, do not add speculative implementations, abstractions, interfaces, dependencies, or configuration merely to populate placeholder directories.

## Core Engineering Principles

This repository follows the Ponytail minimalism principles defined in:

`.agents/rules/ponytail/ponytail.md`

Ponytail is the default engineering philosophy for this repository.

Before implementing changes:

1. Understand the actual requirement.
2. Inspect the existing codebase before creating new code.
3. Reuse existing implementations and patterns whenever possible.
4. Prefer the Go standard library when it is sufficient.
5. Prefer existing dependencies over introducing new dependencies.
6. Avoid abstractions without a concrete use case.
7. Avoid interfaces when there is only one implementation and no meaningful substitution requirement.
8. Keep implementations small and focused.
9. Delete unnecessary code before adding new code when appropriate.
10. Do not introduce architectural complexity merely for future possibilities.

Do not blindly apply these principles if they conflict with an explicit project requirement.

---

## Architecture & Code Organization

This repository is a Go REST API skeleton loosely structured around Clean Architecture principles.

### `cmd/`

Application entry points.

- `cmd/api/main.go` — HTTP API server.
- `cmd/seed/main.go` — database seeding.

### `internal/app/`

Application bootstrap and dependency wiring.

- `web.go` — HTTP application initialization.

### `internal/delivery/http/`

HTTP delivery layer.

- `route/` — HTTP routing.
- `handler/` — HTTP handlers/controllers.
- `response/` — HTTP response helpers.

Handlers should focus on HTTP delivery concerns and delegate business logic to use cases.

### `internal/usecase/`

Business logic and orchestration.

Use cases coordinate application behavior and interact with repository contracts when persistence is required.

Do not create a use-case interface automatically. Prefer concrete implementations unless abstraction is required by the project.

### `internal/repository/`

Repository contracts and data-access implementations.

Prefer the simplest repository structure that satisfies the current requirements.

Do not create repository interfaces solely because they are commonly used in Clean Architecture.

### `internal/infrastructure/`

Concrete infrastructure adapters.

Examples include:

- database connections
- Echo setup
- logging
- external infrastructure integrations

Infrastructure-specific concerns should remain in this layer.

### `internal/model/`

DTOs and request/response models.

### `internal/entity/`

Core domain entities.

### `internal/config/`

Configuration loading and environment variable mapping.

Prefer simple configuration loading over unnecessary configuration abstractions.

### `db/migrations/`

Database SQL migrations.

### `db/seeders/`

Database seeders.

---

## Go Development Conventions

- Write idiomatic Go.
- Prefer simple code over unnecessary abstraction.
- Keep functions small and focused.
- Prefer composition over unnecessary inheritance-like abstractions.
- Avoid unnecessary interfaces.
- Avoid unnecessary wrapper types and wrapper packages.
- Avoid adding dependencies unless there is a concrete reason.
- Prefer standard library functionality when sufficient.
- Follow existing project conventions before introducing a new pattern.
- Do not refactor unrelated code while implementing a feature.

---

## Development Workflow

Before modifying code:

1. Inspect the relevant repository structure.
2. Read `go.mod`.
3. Locate existing implementations related to the requested change.
4. Check existing utilities, services, handlers, repositories, and tests.
5. Determine whether the requested functionality already exists.
6. Identify the smallest appropriate change.

After modifying code:

1. Run `gofmt` on modified Go files.
2. Run relevant tests.
3. Run `go test ./...` when appropriate.
4. Review the diff.
5. Remove unnecessary changes before finishing.

Do not modify unrelated files.

---

## Testing

Tests should verify behavior rather than implementation details.

Prefer:

- table-driven tests when they improve clarity
- focused unit tests
- integration tests where integration behavior is important

Do not create mocks, interfaces, or test abstractions unless they provide a concrete benefit.

---

## Database Migrations

Follow the existing migration conventions in:

`db/migrations/`

Use the commands defined in `taskfile.yml`.

Examples:

```bash
task migrate
task create-migration-sequence name=<name>
```

Do not manually modify existing migrations unless explicitly required.

## Dependencies

Before adding a dependency:

- Check whether the Go standard library provides the required functionality.
- Check whether an existing project dependency already provides it.
- Consider whether a small local implementation is simpler.
- Add a dependency only when it provides a meaningful benefit.

Avoid dependencies for trivial functionality.

## Git & Commits

Follow Conventional Commits:

``` bash
feat:
fix:
refactor:
test:
docs:
style:
chore:
```

Keep commits focused and avoid unrelated changes.

## Agent Behavior

When a task appears larger than necessary:

- question unnecessary requirements
- inspect existing code before creating new abstractions
- prefer deletion or simplification where appropriate
- explain important architectural trade-offs briefly
- do not implement speculative functionality

When Ponytail identifies an opportunity to simplify code, treat it as a recommendation to evaluate, not an automatic instruction to delete code.

Preserve behavior and existing contracts unless the requested change explicitly requires otherwise.