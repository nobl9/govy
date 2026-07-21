# AGENTS.md

## Project Overview

Govy is a Go validation library that provides a generic,
reflection-free, functional API for building strongly typed validators
and readable validation errors.

CI runs project commands inside the `devbox` environment, so prefer the
Makefile targets over direct tool invocations.

## Repository Map

- `pkg/govy`: core public validation API.
- `pkg/rules`: built-in rule constructors and rule error codes.
- `pkg/govytest`: public test assertion helpers for govy users.
- `pkg/jsonpath`: JSON path model used in validation errors and path inference.
- `cmd/govy`: sidecar CLI, currently focused on `inferpath` generation.
- `internal/*`: implementation helpers shared by the public packages.
- `internal/examples`: tested examples embedded into `README.md`.
- `docs/validator-comparison`: separate workspace module used by tests.
- `tests/examplemodule`: separate workspace module used by path inference tests.
- `scripts`: project checks and generation helpers used by the Makefile.

The workspace in `go.work` includes the root module,
`docs/validator-comparison`, and `tests/examplemodule`.

## Development

Run `make help` first when you are unsure which target to use.
The Makefile is the source of truth for local development commands,
and CI calls the same targets through `devbox`.

Use `make activate` to enter the reproducible devbox shell.
If devbox is missing, `make install/devbox` installs it.

For code changes, run at least `make test` and the relevant `make check/*`
target.
Before handing off broad changes, run `make check` unless a required tool or
environment dependency is unavailable.
Report the exact failing command and error if verification cannot be completed.

Do not substitute raw `go test ./...` or `golangci-lint run ./...` for final
verification when a Makefile target exists.
The targets include extra packages, build tags, formatting, generated-file,
spelling, Markdown, and vulnerability checks.

### Code Generation

Do not hand-edit generated files.
Generated files include files with `Code generated ... DO NOT EDIT`,
for example `internal/messagetemplates/templatekey_string.go`.

Generation entry points:

- `make generate/code` runs `go generate ./...`.
- `make generate/readme` refreshes embedded README examples.
- `make check/generate` verifies generated output is committed.

`README.md` contains embed directives such as:

```md
[//]: # (embed: internal/examples/readme_intro_example_test.go)
```

Edit the source example under `internal/examples`, then run
`make generate/readme`.
Do not edit the embedded README code block directly unless the source example
is intentionally unchanged.

### Go Style

Follow the existing immutable builder style.
Methods that customize validators, rules, and rule sets should return modified
copies rather than mutating shared state.

Use `make format/go` rather than running separate formatting tools.

Public API additions need Go doc comments.
Use package links such as `[Rule]`, `[Validator]`, and `[jsonpath.Path]`
where they improve documentation.
Keep comments focused on API contracts, invariants, side effects,
and non-obvious behavior.

When adding a rule in `pkg/rules`:

- Return `govy.Rule[T]` or `govy.RuleSet[T]` consistently with nearby rules.
- Assign an exported error code from `pkg/rules/error_codes.go`.
- Always use message templates from `internal/messagetemplates`.
- Set a useful description with `WithDescription` so validation plans remain informative.
- Add examples or details when they materially improve user-facing errors.
- Add tests for success, failure message, and error-code behavior.

When changing core validation behavior in `pkg/govy`, preserve the public error
shape unless the task explicitly calls for a breaking change.
Tests often assert exact `ValidatorError`, `PropertyError`, and message output.

### Testing

Tests use the standard `testing` package with local helpers from
`internal/assert` and public helpers from `pkg/govytest`.
Use the helper style already present in the package you are editing.

When an authoritative source publishes a finite valid/invalid corpus,
record its URL and immutable version or revision alongside the test table,
then copy every input applicable to the documented contract verbatim.
Do not replace source vectors with representative or equivalent inputs.
Keep derived boundary cases in addition to, not instead of, the source corpus.
If a source vector is intentionally excluded,
enumerate the literal and explain why it falls outside the documented contract.

For rules, keep table data close to the rule test and cover both passing and
failing inputs.
When expected output includes validation messages, assert the exact message
unless the existing package uses a looser helper for that case.

For user-facing examples, prefer Go testable examples.
Examples in `internal/examples` are embedded into `README.md`, so changes there
must still pass `make test` and `make generate/readme`.

Benchmarks live next to the related tests.
If a change affects validation hot paths, run `make test/benchmark` or explain
why it was not run.

## Documentation

Markdown is linted by `make check/markdown`.
Use semantic line breaks and fenced code blocks with language identifiers.

Documentation claims must match the current code.
If a referenced file or workflow is missing, do not repeat the stale reference
in new docs.

When editing `README.md`, prefer changing tested source examples and
regenerating embedded blocks.

## Pull Requests

PR titles must match the rules defined in `.github/workflows/pr-title.yml`

The PR template requires motivation, summary, related changes, testing,
and release-note handling.
Each code change should be covered by unit tests.

For release notes, either replace the `## Release Notes` placeholder with
1-3 sentences or remove the section entirely.
Add `## Breaking Changes` only when the change is actually breaking.
