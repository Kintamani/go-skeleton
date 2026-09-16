---
trigger: model_decision
description: Defines repository-specific testing conventions. Apply when creating or modifying tests, test infrastructure, benchmarks, mocks, or test strategy.
---

# Testing Rules

## Scope

These rules define testing conventions specific to this repository.

Use the relevant `golang-testing`, `golang-unit-testing`, and `golang-benchmark-testing` skills when appropriate.

## General

Tests should verify behavior and meaningful requirements.

Prefer simple tests that clearly communicate expected behavior.

Do not test implementation details merely to increase coverage.

## Unit Tests

Use unit tests for isolated business logic and components where they provide value.

Prefer table-driven tests when they improve readability or make multiple cases clearer.

## Integration Tests

Use integration tests when behavior depends on real infrastructure or integration boundaries.

Do not replace meaningful integration behavior with mocks solely for convenience.

## Interfaces and Mocks

Do not introduce interfaces only to make mocking possible.

When an abstraction is genuinely required for testing, keep the interface small and define it at the consumer boundary where practical.

## Test Changes

When changing behavior:

- update existing tests when necessary;
- add tests for new behavior;
- preserve relevant regression coverage;
- run targeted tests first when useful;
- run `go test ./...` when appropriate.

## Test Quality

Avoid:

- brittle tests;
- excessive setup;
- unnecessary mocks;
- duplicated test logic;
- tests coupled tightly to internal implementation details.
