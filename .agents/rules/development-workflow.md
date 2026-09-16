---
trigger: model_decision
description: Defines the repository development workflow. Apply when planning, implementing, validating, or reviewing changes to determine the appropriate development process.
---

# Development Workflow Rules

## Before Changes

Before modifying code:

1. Inspect the repository structure.
2. Read relevant project files and `go.mod`.
3. Inspect existing implementations and tests.
4. Search for reusable helpers and patterns.
5. Identify relevant skills and rules.
6. Determine the smallest appropriate change.

## During Changes

- Keep changes focused on the requested task.
- Reuse existing code and patterns.
- Preserve existing behavior unless a change is required.
- Avoid unrelated refactoring.
- Avoid speculative features.
- Avoid introducing dependencies or abstractions without a concrete need.

## After Changes

Validate the implementation according to the task.

For Go changes, normally:

```bash
gofmt -w <modified-files>
go test ./...
```

Run more targeted checks when appropriate.

Inspect the final diff and remove accidental or unrelated changes.

## Agent Decision Making

When a solution appears complex:

1. Verify that the complexity is actually required.
2. Check existing project functionality.
3. Check the standard library.
4. Check existing dependencies.
5. Consider a simpler implementation.
6. Implement the smallest solution that satisfies the requirement.

When a rule or skill suggests a change, evaluate it against the repository's actual purpose and requirements.
