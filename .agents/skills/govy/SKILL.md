---
name: govy
description: >-
  Write, review, explain, and test validation with github.com/nobl9/govy.
  Use for govy rules, validators, structured errors, plans, and path inference.
---

# Govy

Govy composes validation from four building blocks:

- `Rule`: one check on a typed value, with optional error codes,
  descriptions, and other metadata.
- `RuleSet`: a reusable group of rules for the same value type.
- `PropertyRules`: extracts a value from the input, associates an error path,
  and applies rules or nested validators.
- `Validator`: combines property validations for an input type.

Check the consumer's pinned govy version.
Load only the references needed for the task.

| Task | Read |
| :--- | :--- |
| Write, format, or review validation | [Validation patterns](references/validation-patterns.md) |
| Choose property paths, pointers, or empty-value behavior | [Properties and paths](references/properties-and-paths.md) |
| Create custom rules or messages | [Rules and messages](references/rules-and-messages.md) |
| Find a predefined rule | [Existing rules](references/existing-rules.md) |
| Inspect or construct errors | [Errors](references/errors.md) |
| Test validation | [Testing](references/testing.md) |
| Generate validation plans | [Validation plans](references/validation-plan.md) |
| Configure path inference | [Path inference](references/path-inference.md) |
