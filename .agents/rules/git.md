---
trigger: model_decision
description: Defines repository-specific Git conventions. Apply when inspecting Git state, creating commits, managing branches, or performing Git operations.
---

# Git Rules

## Commits

Use Conventional Commits.

Common prefixes:

- `feat:`
- `fix:`
- `refactor:`
- `test:`
- `docs:`
- `style:`
- `chore:`

Keep commits focused and descriptive.

## Scope

Do not mix unrelated changes into the same commit.

Do not include generated files, temporary files, secrets, or local environment configuration unless explicitly required.

## Existing Changes

Before modifying files, be aware of the existing working tree state.

Do not overwrite or discard unrelated user changes.

## History

Do not rewrite Git history unless explicitly requested.

Do not use destructive Git commands on user changes without explicit approval.

## Skill

Use the relevant Git/tooling skill when performing non-trivial Git operations.
