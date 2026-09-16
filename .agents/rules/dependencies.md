---
trigger: model_decision
description: Defines repository-specific dependency policies. Apply when adding, removing, upgrading, or replacing Go dependencies or evaluating whether a dependency is necessary.
---

# Dependency Rules

## General

Dependencies must solve a real problem.

Prefer, in order:

1. Go standard library
2. Existing project dependencies
3. A new dependency when it provides meaningful value

Do not add a dependency for trivial functionality that can be implemented simply with the standard library or existing code.

## Before Adding a Dependency

Before adding a new dependency:

- check whether the standard library provides the required functionality;
- check existing project dependencies;
- check whether existing project code already solves the problem;
- consider maintenance and operational cost;
- confirm that the dependency is actually required.

## Avoid

Do not introduce dependencies solely for:

- a few lines of trivial functionality;
- speculative future requirements;
- unnecessary abstractions;
- functionality already provided by an existing dependency;
- convenience when a simple standard-library solution is sufficient.

## Removing Dependencies

Remove dependencies that are no longer required when doing so is safe and relevant to the task.

Do not remove a dependency solely because it is unused in the current skeleton if it is an intentional, documented part of the template.

Verify actual usage before removing a dependency.

## Upgrading Dependencies

Do not upgrade dependencies as part of an unrelated task.

When an upgrade is required:

- review compatibility;
- inspect breaking changes;
- update affected code;
- run relevant tests;
- inspect the final dependency diff.

## Skeleton Policy

This repository is a reusable skeleton.

Dependencies should represent capabilities actually supported by the skeleton.

Do not add dependencies merely because a future project might need them.

Use the `golang-dependencies` skill for detailed dependency management and evaluation.