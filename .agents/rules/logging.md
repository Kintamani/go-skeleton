---
trigger: model_decision
description: Defines repository-specific logging conventions. Apply when modifying logging, log output, log configuration, log rotation, or observability-related behavior.
---

# Logging Rules

## Scope

These rules define logging conventions specific to this repository.

Use the `golang-observability` skill for detailed observability and logging guidance.

## Output

The default deployment model is container-based.

Prefer writing application logs to:

- stdout
- stderr

Let the container runtime or deployment platform handle log collection, retention, and rotation.

## File Logging

Do not introduce application-level file logging or log rotation unless there is a concrete deployment requirement.

Do not add file-based logging merely because it is common in traditional server deployments.

## Logging Library

Prefer the Go standard library logging facilities when they provide the required functionality.

Do not introduce a logging dependency solely for convenience.

When changing an existing logging library, evaluate the scope of the migration before making unrelated changes.

## Logging Changes

When modifying logging:

- preserve useful log context;
- avoid duplicate logging;
- avoid logging the same error at multiple unnecessary layers;
- do not log sensitive information;
- keep logging behavior consistent across the application.
