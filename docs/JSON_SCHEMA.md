# JSON Schema generation

Govy generates a JSON Schema Draft 2020-12 document from a validation plan.
The generated schema describes the JSON representation of the validated value.
It is not an executable translation of all Govy validation behavior.

This document records the current generation contract, known semantic gaps,
and the implementation plan for closing gaps that JSON Schema can represent.

## Current contract

`govy.JSONSchema` builds one schema tree from the property paths recorded in a
validation plan. Each rule can contribute constraints through
`Rule.WithJSONSchema`.

The current implementation has these properties:

- The root document always declares JSON Schema Draft 2020-12.
- Each generated schema node declares at most one JSON type.
- A rule without a JSON Schema builder does not contribute a constraint.
- `JSONSchemaIncludeOmittedRules()` adds optional metadata about missing
  rule and condition builders, and rules on transformed values,
  to the document root.
- A builder error stops generation and identifies the property and rule.
- Unsupported Go kinds stop generation.
- Builders receive the absolute path and JSON type of the selected value.
  They return an independent schema contribution. Govy merges non-conflicting
  keywords directly and uses `allOf` when a keyword is already set. A `nil`
  contribution adds no constraint.
- Named paths create `properties`, fixed indexes create `prefixItems`, array
  wildcards create `items`, and map wildcards create `propertyNames` or
  `additionalProperties`.
- Only paths and types present in the validation plan contribute structure.
  Govy does not discover fields without validation rules or their JSON tags.

## Proposed compatibility policy

The generated schema should be a useful, portable approximation of Govy
validation. Exact alignment with Govy is not a goal. When an exact mapping is
unavailable, generation should prefer a weaker constraint over a constraint
that rejects a value accepted by Govy.

JSON Schema consumers can also differ in their support for optional
vocabularies, formats, and regular expressions. The generated schema targets
the standard Draft 2020-12 behavior and documents known portability limits.

Each built-in rule mapping should have one documented classification:

- **Exact**: JSON Schema and Govy accept the same JSON values within the
  represented type.
- **Approximate**: the schema preserves a useful subset of the rule.
- **Annotation**: the schema documents semantics but does not enforce them.
- **Unsupported**: the rule contributes no schema constraint.

Missing builders and omitted rules on transformed values can be reported
through the optional annotation described below.
Approximate mappings are documented here
but are not classified in the generated schema.

## Omitted-rule metadata

Pass `govy.JSONSchemaIncludeOmittedRules()` to `govy.JSONSchema` to include
`x-govy-omittedRules` on the document root. Default output does not include this
annotation. The annotation is also absent when there are no omissions to report.

Each record contains:

- `path`: the absolute Govy JSON path to the validated value, including
  wildcards for collection elements.
- `rule`: the rule's error code, omitted if the rule has no code.
- `reason`: `missing JSON Schema builder` if the rule has no `WithJSONSchema`
  builder, or `missing WhenJSONSchema builder` if any guarding condition lacks
  a builder. Rules omitted because of `Transform` use
  `rule validates a transformed value`.

Govy records omissions during the existing plan traversal, before rules are
filtered or deduplicated. Transformed-value omissions take precedence over
missing builders. Otherwise, a missing rule builder takes precedence when both
the rule and a condition lack builders. The metadata does not affect validation.
Errors from executed builders still stop generation.

A rule or condition builder that runs and returns `nil` is an intentional
no-op and does not produce an omission record. Builders for transformed values
are not executed, so their omissions are reported even if they would return `nil`.
Optional-property markers are not reported.
The annotation does not classify approximate mappings.

`x-govy-omittedRules` is a Govy extension, not a standard JSON Schema keyword.
Its prefix follows the [JSON Schema custom-annotation convention] and the
`x-` naming used by [Typia custom fields]. Draft 2020-12 permits unknown
keywords and recommends treating them as annotations.
See [JSON Schema objects and keywords].

[Ajv strict mode] rejects unknown keywords unless they are registered.
Consumers that enable this metadata can register it with
`ajv.addKeyword("x-govy-omittedRules")` before compiling the schema.

## Structural and semantic gaps

### Incomplete type trees

The validation plan contains type information for recorded properties, but it
does not describe the full JSON document. Intermediate containers can therefore
have no type, and fields without validation rules do not appear in the schema.

Future plan metadata should provide the JSON field tree, including names,
container element types, map value types, and required state. Schema generation
can then merge validation constraints into that structural tree.

### Null values

When `encoding/json` decodes into a zero-valued Go struct, an absent pointer
property remains `nil` and an explicit JSON `null` sets it to `nil`. Govy
validates the resulting Go value and cannot recover which JSON representation
was used.

Generated schemas resolve this ambiguity by modeling optionality through
`required`. A pointer property uses its dereferenced, non-null JSON type. The
property can be absent unless required, but it cannot be an explicit `null`.
Slices and maps likewise use `array` and `object` without accepting `null`.

### Required properties and zero values

`PropertyRules.Required` and `rules.Required` have the same schema behavior.
Both add the property name to the parent `required` list. Neither rejects a
present zero value in the generated schema.

This behavior is an intentional approximation. JSON Schema `required` checks
property presence only, while Govy also checks whether the Go value is its
type's zero value.

The multi-property presence rules have the same difference. Govy considers a
property set when its value is non-zero. Their JSON Schema mappings use property
presence.

### Conditional rules

Govy conditions are arbitrary Go predicates. The plan records their
descriptions, but schema generation cannot infer a condition from executable
Go code.

`WhenJSONSchema` lets the caller define the equivalent JSON Schema condition.
Schema generation emits the condition through `if` and applies guarded rule
builders through `then`. It preserves the validator or parent-property scope
where Govy evaluates each predicate. Multiple predicates retain Govy's AND
semantics, including predicates inherited from different nested scopes.

If any predicate guarding a rule has no JSON Schema builder, schema generation
omits that rule's builders. `WhenDescription` remains independent and is not
required for schema generation.

### Transformed properties

`govy.Transform` preserves the original input type in the generated schema.
Rules on the transformed value do not contribute value constraints.
For example, parsing a string into an integer and applying `EQ(12)` emits
`type: string`, without the incompatible numeric `const: 12`.
The same policy applies when a transform preserves the type,
such as lowercasing a string.

Directly attached `PropertyRules.Required` and `rules.Required` retain their
property-presence effect, subject to conditions on the original parent.
Their value constraints are not carried over.
Included validators describe the transformed result, so their rules and
property structure are omitted, including their required properties.
Govy does not execute the transformer or schema builders for transformed values.
Runtime validation and regular `govy.Plan` output remain unchanged.

With `JSONSchemaIncludeOmittedRules()`, each omitted rule is reported with
`rule validates a transformed value`.
Included rule paths identify values in the transformed result.
Those paths do not add properties to the input schema.

To supply an explicit input mapping, add a separate `govy.For` with the original
getter and the same property path, then attach a rule with `WithJSONSchema`.
Its schema contribution combines with the preserved input type.
The rule can use a no-op validation function if it only supplies schema metadata.
A custom builder attached to a rule under `Transform` is still omitted.
This approach reuses the existing API without changing runtime validation.

### Go equality and JSON equality

`EQ`, `NEQ`, `OneOf`, and `NotOneOf` use Go equality. Their schema mappings use
JSON value equality through `const` and `enum`.

The mapping works for primitive values and many comparable structs and arrays.
It is not exact for pointer identity, values with custom JSON marshaling, or
distinct Go values that produce the same JSON value.

### Go numeric domains

The generated schema distinguishes JSON integers and numbers, but it does not
emit the bounds of the source Go type. A schema for `int8` can therefore accept
an integer that cannot be decoded into `int8`.

Floating-point validation also uses Go's finite precision, while JSON Schema
defines numbers without Go-specific precision limits.
Ajv uses JavaScript numbers, so integers outside the safe integer range can
lose precision. The rule tests record this consumer-specific difference
separately from the generated constraints.

### Regular expressions

JSON Schema uses ECMA-262 regular expressions. Govy uses Go's RE2-based
`regexp` package.

The internal translator preserves supported Go expressions and reports an
error for multiline anchors. JSON Schema pattern behavior can still vary across
consumers because their ECMA-262 support differs.

## Built-in rule gaps

### No portable exact mapping

The following rules compare multiple instance values or depend on arbitrary Go
functions:

- `EqualProperties`
- `GTProperties`, `GTEProperties`, `LTProperties`, and `LTEProperties`
- all four `*ComparableProperties` rules
- `UniqueProperties`
- `SliceUnique` with an arbitrary hash function
- `GT`, `GTE`, `LT`, and `LTE` for strings

Standard JSON Schema cannot reference one instance value as another keyword's
comparison value. Ajv `$data` references can express some comparisons, but they
are nonportable. A Govy-specific vocabulary is another option, but every
consumer would need to implement it.

### Runtime-dependent rules

The following rules depend on the host environment or on a Go-specific
representation:

- `StringFileSystemPath`, `StringFilePath`, and `StringDirPath`
- `StringMatchFileSystemPath`
- `StringTimeZone`
- `StringRegexp`
- `URL` and its options

These rules should remain unsupported unless generation receives an explicit,
portable representation. A fixed IANA time-zone snapshot could support
`StringTimeZone`, but it would make the schema depend on that snapshot.

### Format annotations

Schema generation emits these Draft 2020-12 format annotations:

- `StringEmail`: `email`
- `StringURL`: `uri`
- `StringIP`: `anyOf` with `ipv4` and `ipv6`
- `StringIPv4`: `ipv4`
- `StringIPv6`: `ipv6`
- `StringDateTime`: `date-time` for `time.RFC3339` and `time.RFC3339Nano`

Other `StringDateTime` layouts remain unsupported because Go layouts can
describe formats other than RFC 3339.

These mappings are approximate. `StringEmail` uses Go's RFC 5322 address
parser, while the JSON Schema `email` format refers to an RFC 5321 mailbox.
`StringURL` uses Go's URL parser, while `uri` refers to RFC 3986. JSON Schema
implementations can also differ in format validation. Draft 2020-12 treats
`format` as an annotation unless the consumer enables format assertions.

### Content annotations

Schema generation emits these content annotations:

- `StringBase64`: `contentEncoding: base64`
- `StringJSON`: `contentMediaType: application/json`
- `StringJWT`: `contentMediaType: application/jwt`

Content keywords are annotations and do not require a validator to decode or
validate embedded content.
The URL-safe Base64 rules have no `contentEncoding` annotation because there is
no portable `base64url` content encoding value. The schema model does not yet
include `contentSchema` because no built-in rule supplies a decoded-content
schema.

### Pattern mappings and candidates

`StringEIN` and `StringSSN` emit exact ECMA-262 patterns.
`StringASCII`, `StringAlpha`, and `StringAlphanumeric` emit exact translated
patterns. `StringAlphaUnicode` and `StringAlphanumericUnicode` also emit exact
patterns, but they expand their Unicode classes according to the Go toolchain's
Unicode tables at generation time.
`StringDNSLabel` and `StringDNSSubdomain` emit their length bounds and exact
translated patterns. `StringFQDN` emits its exact translated pattern.
The UUID family emits exact structural patterns for each supported version and
variant. `StringMD5`, `StringSHA256`, `StringSHA384`, and `StringSHA512` emit
exact lowercase hexadecimal length patterns.
The three Base64 rules emit approximate patterns that constrain the alphabet
and padding shape. They do not validate trailing bits or invalid unpadded
lengths as strictly as the Go decoders.
`StringBIC` and `StringBICISO93622014` emit the same approximate structural
pattern. It does not enforce Govy's country-code allowlist.
`StringLatitude` and `StringLongitude` emit the same approximate decimal
coordinate pattern. It does not enforce the latitude range of -90 to 90 or the
longitude range of -180 to 180.
`StringMAC` emits an exact pattern for the forms accepted by `net.ParseMAC`.
It covers 6-, 8-, and 20-octet hexadecimal addresses with colons, hyphens,
dots, or no separators.
`StringCIDR`, `StringCIDRv4`, and `StringCIDRv6` emit approximate structural
patterns. They do not validate address components, prefix ranges, or the
network-alignment check in `StringCIDRv4`.
`StringGitRef` emits an approximate pattern. It accepts `HEAD` or two or more
nonempty slash-separated components. It rejects ASCII control characters,
spaces, and the unconditional forbidden characters `\`, `?`, `*`, `[`, `~`,
`^`, and `:`. It does not reject `..`, `@{`, leading dots, `.lock` suffixes,
trailing dots, components equal to `@`, or leading hyphens in branch and tag
components.
`StringKubernetesQualifiedName` emits its global length bounds and a pattern
for the prefix and name structure. It does not enforce the separate
253-character prefix and 63-character name limits.

`StringISO3166Alpha2`, `StringISO3166Alpha3`, and `StringISO3166Numeric` emit
exact enums of accepted country codes. `StringISO4217` emits an exact enum of
accepted currency codes. `StringISO31662` emits an exact 5,213-value enum of
accepted subdivision codes, which makes schemas that use it large.

### Algorithmic string rules

The checksum rules emit approximate patterns for their outer syntax:

- `StringCreditCard` and `StringLuhnChecksum`
- `StringISBN`, `StringISBN10`, `StringISBN13`, and `StringISSN`

These patterns do not validate checksums. The payment-card pattern also does
not reject values in which every digit is the same. The ISBN patterns enforce
the supported digit counts, separators, check-character positions, and ISBN-13
prefixes. The ISSN pattern enforces its fixed hyphenated form.

The following rules validate decoded content or use complex parsers:

- both BCP 47 rules
- `StringCrontab`
- `StringTitle`
- `StringJWT`

Patterns can document the outer shape of some values, but they do not preserve
the complete Govy validation. These mappings should be approximate or
unsupported.

## Implementation plan

### 1. Extend observable generation behavior

Optional root metadata now reports missing builders and rules on transformed values.
Decide whether approximate mappings also need machine-readable classifications
and whether callers need a strict mode that rejects omitted rules.

### 2. Complete structural generation

Merge a complete JSON field tree into the validation plan when that metadata is
available. Transformed properties already preserve their input types and omit
constraints on transformed values.

### 3. Add missing standard keywords

Extend `jsonschema.Schema` only when a built-in rule needs the keyword. The
first candidates are:

- `uniqueItems`
- `contentSchema`

### 4. Implement rule mappings by confidence

Implement exact mappings first. Then add approximate mappings with documented
differences. Keep runtime-dependent and cross-property rules unsupported unless
the project explicitly accepts a nonportable extension.

### 5. Verify complete schema documents

Rule-specific tests live in `pkg/rules` and use the public `govy.JSONSchema`
function. They compare the complete schema with a JSON fixture in
`pkg/rules/testdata/jsonschema`, then validate the existing rule inputs with Ajv.
Published input corpora remain part of those tests.

The shared input tables record expected differences through
`jsonschematest.Case.JSONSchemaDifference` or an equivalent field on the rule's
case type. A nonempty reason requires the schema result to differ from Govy's
expected result. It does not skip the case. Both unexpected agreement and an
unexpected disagreement fail the test.

Core generation tests stay in `pkg/govy`, with fixtures in
`pkg/govy/test_data`. They cover composition, conditions, paths, transforms,
omitted-rule metadata, and generation errors through the public API.

`internal/jsonschematest` owns the shared fixture assertion and Node runner.
Devbox installs the pinned Ajv and `ajv-formats` dependencies.
The runner compiles each schema once and validates its input batch in one
process. It uses Draft 2020-12 and [Ajv format assertions] in full mode.
It does not coerce values, insert defaults, or remove properties.
Content keywords remain annotations.

Ajv's `strictTypes`, `strictRequired`, and `strictTuples` checks are disabled
because Govy permits partial type trees, presence-only conditions, and open
tuples. Schema validation and unknown-keyword checks remain enabled.
The runner registers `x-govy-omittedRules` as an annotation.
Invalid schemas, unknown formats, and missing dependencies fail the test.

Each semantic gap that is fixed should have a focused fixture and input case
before it is removed from this document.

Run at least these repository targets after each implementation batch:

```shell
make test
make check/markdown
make check/spell
make check/vet
make check/lint
```

Run `make check` before the completed feature is handed off.

## Open decisions

- Must the default schema be a conservative superset of Govy-valid JSON?
- Should generation offer a strict mode that rejects omitted rules?
- Should nonportable Ajv or Govy vocabulary output be supported?
- Should large finite enums be emitted by default?

## References

- [JSON Schema Draft 2020-12]
- [JSON Schema validation vocabulary]
- [JSON Schema objects and keywords]
- [JSON Schema custom-annotation convention]
- [Typia custom fields]
- [Ajv `$data` references]
- [Ajv strict mode]

[Ajv `$data` references]: https://ajv.js.org/guide/combining-schemas.html#data-reference
[Ajv format assertions]: https://ajv.js.org/packages/ajv-formats.html
[Ajv strict mode]: https://ajv.js.org/strict-mode.html#unknown-keywords
[JSON Schema custom-annotation convention]: https://json-schema.org/blog/posts/custom-annotations-will-continue
[JSON Schema Draft 2020-12]: https://json-schema.org/draft/2020-12/schema
[JSON Schema objects and keywords]: https://json-schema.org/draft/2020-12/json-schema-core#section-4.3.1
[JSON Schema validation vocabulary]: https://json-schema.org/draft/2020-12/json-schema-validation
[Typia custom fields]: https://typia.io/docs/json/schema/#custom-fields
