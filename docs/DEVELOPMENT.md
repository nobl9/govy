# Development

This document describes the intricacies of govy development workflow.
If you see anything missing, feel free to contribute to this document :)

## Pull requests

[Pull request template](../.github/pull_request_template.md)
is provided when you create new PR.
Section worth noting and getting familiar with is located under
`## Release Notes` header.

## Makefile

Govy ships with a Makefile which is well documented and should cover most if
not all development cycle needs.
Run `make help` to display short description for each target.

## CI

Continuous integration pipelines utilize the same Makefile commands which
you run locally within reproducible `devbox` environment.
This ensures consistent behavior of the executed checks
and makes local debugging easier.

## Testing

You can run all unit tests with `make test`.
We also encourage inspecting test coverage during development, you can verify
if the paths you're interested in are covered with `make test/coverage`.

## Releases

Govy adheres to the Go's official release workflow recommendations and
requirements. Refer to the official
[Go docs](https://go.dev/doc/modules/release-workflow) for more details.

### Release automation

We're using [Release Drafter](https://github.com/release-drafter/release-drafter)
to automate release notes creation. Drafter also does its best to propose
the next release version based on commit messages from `main` branch.

Release Drafter is also responsible for auto-labeling pull requests.
It checks both title and body of the pull request and adds appropriate labels. \
**NOTE:** The auto-labeling mechanism will not remove labels once they're
created. For example, If you end up changing PR title from `sec:` to `fix:`
you'll have to manually remove `security` label.

On each commit to `main` branch, Release Drafter will update the next release
draft.

In addition to Release Drafter, we're also running a script which extracts
explicitly listed release notes and breaking changes which are optionally
defined in `## Release Notes` and `## Breaking Changes` headers.
It also performs a cleanup of the PR draft mitigating Release Drafter
shortcomings.

## Code generation

Some parts of the codebase are automatically generated.
We use the following tools to do that:

- [embed-examples-in-markdown.bash](../scripts/embed-examples-in-markdown.bash)
  for embedding tested examples in Markdown documents.

## Validation

We're using our own validation library to write validation for all objects.
Refer to this [README.md](../internal/validation/README.md) for more information.

## Predefined rules

Predefined rules live in [`pkg/rules`](../pkg/rules/).
Add a new rule there only when the validation is generally useful to govy users.
If the rule exists only for one consumer, prefer a consumer-owned
[`govy.NewRule`](../pkg/govy/rule.go) instead.

Most rule constructors follow the same shape:

1. Define an exported constructor in the matching `pkg/rules/*.go` file.
   Constructors return `govy.Rule[T]` for one validation decision
   or `govy.RuleSet[T]` when the public rule is composed from existing rules.
2. Add a stable `ErrorCode...` constant in
   [`pkg/rules/error_codes.go`](../pkg/rules/error_codes.go).
   Tests and integrations use these codes, so do not change existing values
   unless the change is intentionally breaking.
3. Add a message template in
   [`internal/messagetemplates/templates.go`](../internal/messagetemplates/templates.go)
   when no existing template describes the failure.
   Return `govy.NewRuleErrorTemplate(govy.TemplateVars{...})`
   from the rule body so message rendering receives the failed value,
   comparison value, parse error, or custom fields it needs.
4. Attach `WithErrorCode`, `WithMessageTemplate`, and a description
   to the returned rule.
   Use `WithDescription` for fixed text or `WithDescriptionTemplate`
   for descriptions derived from a template.
5. Add details, examples, cascade mode, and plan modifiers only when they are
   part of the rule contract.
   For example, composed DNS rules use `Cascade(govy.CascadeModeStop)`,
   and finite-value rules such as `OneOf` add
   `govy.RulePlanModifierValidValues`.

Rule names should describe the validated value and the predicate.
Existing examples include `StringEmail`, `StringKubernetesQualifiedName`,
`SliceUnique`, `MapMinLength`, and `GTProperties`.
Use generics when the rule applies to a family of types;
for example, length rules support strings, slices, and maps,
while comparison rules use `comparable` or `cmp.Ordered` constraints.

### Rule tests

Keep each rule's tests together in the same `*_test.go` file, in this order:

1. `Test<Rule>`: the standard table-driven validation test.
2. `Test<Rule>_JSONSchema`: the JSON Schema test, when the rule has schema support.
3. `Benchmark<Rule>`: the validation benchmark.

Place shared inputs near this group.
Use the same shared input tables for standard validation, JSON Schema, and benchmarks.
Also include applicable published corpora in the JSON Schema test.
Record expected JSON Schema differences on the input cases,
with a reason in `JSONSchemaDifference` or the equivalent case field.

Generate schemas through the public `govy.JSONSchema` function.
Use `internal/jsonschematest` to compare the complete schema with a fixture
in `pkg/rules/testdata/jsonschema`.
The helper also validates the inputs with Ajv.
Keep shared helpers and cross-rule schema cases in the matching rule test file.
Do not add separate `*_jsonschema_test.go` files.

A nonempty `JSONSchemaDifference` requires the schema result to differ from Govy.
It does not skip the case.
Both unexpected agreement and unexpected disagreement fail the test.
Devbox supplies Ajv and `ajv-formats` for Draft 2020-12 validation
with format assertions enabled in full mode.
Content keywords remain annotations.
See the [runner configuration](../internal/jsonschematest/testdata/validate.cjs)
for its strictness settings.

Every exported predefined rule must have a standard test and benchmark.
The presence check in [`pkg/rules/rules_test.go`](../pkg/rules/rules_test.go)
requires both functions for each exported rule constructor.
Keep test cases table-driven where that matches the surrounding file.
Assert the stable error code with `govy.HasErrorCode`.
Assert the exact message when the rule has a specific rendered message.

When adding a new rule, run at least:

```sh
make test
make check/markdown
```

For changes to templates, generated references, spelling-sensitive text,
or public API,
run the full `make check`.
Do this before opening a pull request.

## Dependencies

Renovate is configured to automatically merge minor and patch updates.
For major versions, which sadly includes GitHub Actions, manual approval
is required.

## Tests coverage

Tests coverage reporting is automated using the following actions:

- [go-coverage-report](https://github.com/ncruces/go-coverage-report) which
  is responsible for updating the coverage badge in main README.md.
  It stores the coverage results in GitHub wiki and it can be easily inspected
  [here](https://raw.githack.com/wiki/nobl9/govy/coverage.html).
  This action is run only on `push` events to _main_ branch.
- [coverdiff](https://github.com/kskitek/coverdiff) which is executed on each
  PR runs the tests coverage and posts a summary report as a comment.
  It highlights positive and negative changes.

## Benchmarks

[github-action-benchmark](https://github.com/benchmark-action/github-action-benchmark)
is used to collect and store benchmarks' results.
It inspects PRs and if a configured threshold difference between previous
and current results is breached it will leave a comment on the affected PR.

On top of that it publishes benchmarks' history charts onto
[GitHub Pages](https://nobl9.github.io/govy/dev/bench).
