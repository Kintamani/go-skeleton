# Agent Instructions

## Repository

This repository is a reusable Go project skeleton for backend/API projects.

It is a starting point for new projects, not a finished application.

## Core Rules

- Inspect the existing repository before making changes.
- Understand the requirement before implementing it.
- Reuse existing code, patterns, and dependencies.
- Prefer simple and idiomatic Go solutions.
- Prefer the standard library when it is sufficient.
- Avoid unnecessary abstractions, interfaces, wrappers, and dependencies.
- Do not implement speculative features.
- Do not modify unrelated files.
- Preserve existing behavior unless the task requires changing it.

## Rules

Read only the rules relevant to the current task.

### Architecture

`.agents/rules/architecture.md`

Use when creating, moving, or refactoring packages, changing dependencies between layers, or making architectural decisions.

Related skill:

`golang-architecture`
`golang-code-quality`

### Dependencies

`.agents/rules/dependencies.md`

Use when adding, removing, upgrading, or replacing dependencies.

Related skill:

`golang-dependencies`

### Logging

`.agents/rules/logging.md`

Use when modifying logging, log output, log configuration, or observability-related behavior.

Related skill:

`golang-observability`

### HTTP

`.agents/rules/http.md`

Use when creating or modifying HTTP endpoints, handlers, middleware, routing, or HTTP responses.

Related skills:

`golang-api-documentation`

### Configuration

`.agents/rules/configuration.md`

Use when modifying application configuration, environment variables, configuration loading, or configuration defaults.

### Database

`.agents/rules/database.md`

Use when modifying database access, repositories, migrations, seeders, queries, or database configuration.

Related skill:

`golang-database`

### Testing

`.agents/rules/testing.md`

Use when creating or modifying tests, test infrastructure, benchmarks, or test strategy.

Related skills:

`golang-testing`

### Development Workflow

`.agents/rules/development-workflow.md`

Use when planning or executing development tasks, especially when deciding how to inspect, modify, and validate the repository.

### Git

`.agents/rules/git.md`

Use when creating commits, modifying branches, inspecting Git state, or performing Git operations.

## Skeleton

This repository is a reusable project skeleton.

Some directories may intentionally exist as extension points even when they are currently empty.

Do not:

- remove intentional extension points merely because they are unused;
- add fake implementations to empty directories;
- add speculative infrastructure;
- add unused interfaces or dependencies.

## Skills

Use relevant skills from `.agents/skills/*` when specialized knowledge is required.

Prefer existing skills instead of duplicating their instructions in repository rules.

Use only the skills relevant to the current task.

## Ponytail

Follow:

`.agents/rules/ponytail.md`

Ponytail identifies unnecessary complexity and encourages minimal implementations.

Evaluate Ponytail recommendations against the purpose of this reusable skeleton.

Do not blindly remove intentional architecture or extension points.

## Change Scope

Keep changes focused on the requested task.

Do not perform unrelated refactoring.

If a separate improvement is discovered, mention it instead of implementing it automatically.

## Validation

Before completing a task:

- format modified Go code;
- run relevant tests;
- inspect the final diff;
- confirm no unnecessary changes were introduced.