---
trigger: model_decision
description: Enforces minimalism, YAGNI, and avoidance of unnecessary complexity, abstractions, and dependencies. Apply to all development tasks while preserving intentional skeleton architecture and extension points.
---

# Architecture Rules

## General

This repository loosely follows Clean Architecture.

Architecture is a guideline, not a requirement to introduce abstractions without a concrete need.

Prefer clear dependencies and simple package boundaries over strict adherence to patterns.

Use the `golang-architecture` skill when an architectural decision requires deeper guidance.

## Project Structure

The main structure is:

- `cmd/` — application entrypoints
- `internal/app/` — application bootstrap and dependency wiring
- `internal/delivery/` — transport and delivery layer
- `internal/usecase/` — application/business logic
- `internal/repository/` — data access
- `internal/infrastructure/` — concrete infrastructure implementations
- `internal/model/` — application/transport models
- `internal/entity/` — domain entities
- `internal/config/` — application configuration
- `db/` — migrations and seeders

Follow the existing structure before introducing new packages.

## Dependency Direction

Prefer dependencies flowing toward application/domain logic.

- Delivery handles transport concerns.
- Use cases contain application/business behavior.
- Repository handles data access.
- Infrastructure contains concrete external/system integrations.
- Entities should not depend on delivery or infrastructure details.

Do not move logic between layers solely to satisfy a pattern. The responsibility of the code should determine its location.

## Interfaces

Do not create interfaces automatically.

Prefer concrete types unless an interface provides a concrete benefit.

An interface may be justified when:

- multiple implementations are required;
- substitution is an actual requirement;
- consumer-side testing requires an abstraction;
- it represents a meaningful architectural boundary.

Do not create an interface merely because a package contains a use case or repository.

Prefer small interfaces defined by the consumer when abstraction is needed.

## Skeleton Extension Points

This repository is a reusable skeleton.

Some directories may intentionally exist as future extension points, including:

- `cmd/worker/`
- `cmd/scheduler/`
- `internal/delivery/grpc/`
- `internal/delivery/websocket/`
- `internal/delivery/worker/`
- `internal/infrastructure/queue/`
- `internal/infrastructure/redis/`
- `internal/mapper/`
- `deployments/`
- `scripts/`

Do not remove an empty directory solely because it is currently unused if it represents an intentional extension point.

Do not add implementations merely to populate these directories.

## New Packages

Before creating a package:

1. Check whether the existing package can own the responsibility.
2. Check whether an existing package already provides the required functionality.
3. Confirm that the new package represents a meaningful boundary.
4. Keep the package focused.

Avoid creating packages for trivial wrappers, single-function abstractions, or speculative future requirements.

## Changes

Preserve the existing architecture unless the task requires an architectural change.

When an architectural change is necessary:

1. identify the affected boundaries;
2. keep the change as small as practical;
3. preserve existing behavior where possible;
4. validate affected packages and tests.