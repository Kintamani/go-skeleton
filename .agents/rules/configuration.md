---
trigger: model_decision
description: Defines repository-specific configuration conventions. Apply when modifying environment variables, configuration loading, defaults, validation, or application configuration.
---

# Configuration Rules

## Scope

These rules define configuration conventions specific to this repository.

Use the relevant Go project/configuration skill when implementation details require deeper guidance.

## General

Configuration should be:

- explicit;
- simple;
- easy to understand;
- environment-friendly;
- validated when required.

Prefer straightforward configuration loading over elaborate configuration frameworks.

## Environment

Use environment variables for deployment-specific configuration and secrets.

Do not hardcode credentials, tokens, passwords, or environment-specific values.

## Initialization

Required configuration should fail fast when the application cannot operate without it.

Avoid silently using invalid or unexpected defaults for required values.

## Abstractions

Do not introduce:

- unnecessary functional-option patterns;
- configuration factories;
- multiple configuration loaders;
- speculative configuration sources.

Use the simplest configuration structure that satisfies the current requirements.

## Changes

When modifying configuration:

- preserve existing variable names unless a change is required;
- consider compatibility with existing environments;
- update relevant documentation or examples when necessary;
- do not introduce unrelated configuration changes.
