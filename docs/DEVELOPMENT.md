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

- [gomarkdoc](https://github.com/princjef/gomarkdoc)
  for generating Markdown docs from Go code docs.
- [embed-example-in-readme.bash](../scripts/embed-example-in-readme.bash)
  for embedding tested examples in [README.md](../README.md).

## Validation

We're using our own validation library to write validation for all objects.
Refer to this [README.md](../internal/validation/README.md) for more information.

## Dependencies

Renovate is configured to automatically merge minor and patch updates.
For major versions, which sadly includes GitHub Actions, manual approval
is required.
