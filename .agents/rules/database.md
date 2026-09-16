---
trigger: model_decision
description: Defines repository-specific database conventions. Apply when modifying database access, repositories, queries, migrations, seeders, or database configuration.
---

# Database Rules

## Scope

These rules define database conventions specific to this repository.

Use the `golang-database` skill for detailed database implementation guidance.

## Structure

Keep database-specific concerns within the repository/infrastructure boundaries defined by the project architecture.

Do not expose database implementation details unnecessarily to delivery or domain code.

## Migrations

Use the existing migration system and conventions.

Database schema changes must be represented by migrations.

Do not modify existing migration history unless explicitly required.

Prefer creating a new migration for a schema change rather than rewriting an already-applied migration.

## Seeders

Seeders should remain simple and deterministic where practical.

Do not introduce large data-generation dependencies for trivial seed data.

Keep development/example data separate from production business logic.

## Queries

Reuse existing database patterns before introducing new abstractions.

Avoid unnecessary query wrappers, repositories, or generic data-access layers.

## Changes

When modifying database code:

- inspect existing schema and migrations first;
- preserve existing data behavior unless the task requires a change;
- consider indexes and constraints when changing schema;
- test affected database behavior when practical.
