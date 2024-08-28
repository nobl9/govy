# Roadmap - what's coming

## Move name inference out of experimental stage

It is up for debate whether name inference should be ever enabled by default.
If so, it would have to be with runtime mode, and that would come with
a significant performance cost.
We should still remove it from experimental stage once the feature gets
battle tested enough.

## Move plan generation out of experimental stage

The plan generation should be moved out of experimental stage once it is
stable enough.

## Placeholders for error messages and descriptions

Provide replaceable placeholders in a form of `{{ <placeholder_name> }}`
which would allow an easy access to common variables used when constructing
error messages and descriptions like `WithDescription`.

Unlike [Fluent Validation](https://docs.fluentvalidation.net/en/latest/configuring.html#),
we would opt for double curly braces as to meet the Go templates standard.
It is up to debate whether to use Go templates or simply `strings.Replace`.
Go templates would definitely provide more flexibility but might come at
a higher performance cost.

## Localization

Allow defining different locales for error messages.

## Add custom linter

The linter could parse generated validation plan and enforce customizable
rules.
For instance, if a property has no name, it would raise an error.
