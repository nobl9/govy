---
name: govy
description: >
  Use this skill when writing, reviewing, or explaining validation code with
  github.com/nobl9/govy. Trigger for govy validators, property rules, custom
  rules, predefined rules, govytest assertions, validation plans, error
  messages, message templates, path inference, or any task that mentions the
  govy library.
---

# Govy

This skill treats exclusively about
[github.com/nobl9/govy](https://github.com/nobl9/govy), a generic,
reflection-free validation library for Go.

The references in this skill are self-contained.
Load only the reference that matches the task.

## Reference Selection

| Task | Load |
| :--- | :--- |
| Build a validator, name it, add conditions, validate slices, or compose validators | [core-validation.md](references/core-validation.md) |
| Work with property getters, paths, pointers, transforms, required values, optional values, hidden values, or property-level conditions | [properties-and-paths.md](references/properties-and-paths.md) |
| Create or modify rules, rule sets, error details, examples, error codes, custom messages, or message templates | [rules-and-messages.md](references/rules-and-messages.md) |
| Validate nested objects, slices, maps, each element, or compose validators over collections | [collections-and-composition.md](references/collections-and-composition.md) |
| Inspect, serialize, rename, or construct govy errors | [errors.md](references/errors.md) |
| Test govy validation behavior without the govytest package | [testing.md](references/testing.md) |
| Generate or validate a validation plan | [validation-plan.md](references/validation-plan.md) |
| Configure runtime or generated path inference | [path-inference.md](references/path-inference.md) |
| Choose from existing predefined rules in `pkg/rules` | [existing-rules.md](references/existing-rules.md) |
| See examples of predefined rules in use | [predefined-rules.md](references/predefined-rules.md) |
| Use `pkg/govytest` assertion helpers | [govytest.md](references/govytest.md) |

## Usage Guidance

- Define rules and validators once and reuse them across validation calls.
  They are relatively costly to construct and may lazily initialize internal
  state during first use, so avoid rebuilding them every time validation runs.
- Keep validator declarations immutable-style: chained methods return copies,
  so derive runtime variants from existing validators instead of mutating them
  in place.
- Prefer explicit `WithName` or `WithPath` unless the user explicitly
  mentions path inference or if the inference mechanism is already used.
- Use `WithName` for one path segment and `WithPath` for multi-segment paths.
- Use `ForPointer` for optional pointer fields so rules operate on the
  pointed-to value and nil values are skipped unless `Required` is added.
- Use `Required` and `OmitEmpty` on property rules to short-circuit empty values
  before running normal rules.
- Attach plan and error metadata to the receiver that owns it:
  use `Rule.WithDescription`, `Rule.WithDetails`, `Rule.WithExamples`,
  and `Rule.WithErrorCode` for rules; use `PropertyRules.WithExamples`
  for property examples; use `govy.WhenDescription` for conditional rules.
- Use `govytest` helpers for tests that need to match structured validation
  errors rather than brittle full error strings.
