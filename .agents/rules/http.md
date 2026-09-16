---
trigger: model_decision
description: Defines repository-specific HTTP conventions. Apply when creating or modifying HTTP endpoints, handlers, middleware, routing, or HTTP responses.
---

# HTTP Rules

## Scope

These rules define HTTP conventions specific to this repository.

Use the `golang-api-documentation` and relevant Go API skills when implementing or documenting HTTP APIs.

## Framework

The skeleton uses the existing Echo setup.

Reuse the existing Echo configuration and conventions.

Do not introduce another HTTP framework unless explicitly required.

## Layering

Keep HTTP-specific concerns in the delivery layer.

HTTP handlers should:

- receive transport input;
- validate or translate transport data;
- call the appropriate use case;
- translate the result into an HTTP response.

Do not place business logic directly in handlers when it belongs in a use case.

## Routing

Keep route registration separate from business logic.

Reuse existing routing patterns before introducing new abstractions or wrappers.

## Wrappers

Do not create wrappers around Echo or standard HTTP functionality unless they provide meaningful project-specific behavior.

## Changes

When adding or modifying an endpoint:

- follow existing API conventions;
- reuse existing response/error patterns;
- keep the change focused;
- add or update tests when appropriate.
