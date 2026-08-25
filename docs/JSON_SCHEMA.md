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

Generation should make unsupported and approximate mappings observable.
The exact reporting API is still an open decision.

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

`govy.Transform` records the original property type, while its rules validate
the transformed Go value. A numeric rule applied after a string-to-number
transform can therefore add a numeric keyword to a string schema. That keyword
does not validate the original JSON representation.

Transformed property rules need an explicit schema mapping. Without one,
generation should omit their builders.

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

### Content annotation candidates

Adding `contentEncoding`, `contentMediaType`, and `contentSchema` to the schema
model would support useful annotations for these rules:

- `StringBase64`
- `StringJSON`
- `StringJWT`

Content keywords are annotations and do not require a validator to decode or
validate embedded content.

### Pattern mappings and candidates

`StringEIN` and `StringSSN` emit exact ECMA-262 patterns.

The following rules can also use patterns or finite enums without adding a
custom vocabulary:

- `StringMAC`
- `StringCIDR`, `StringCIDRv4`, and `StringCIDRv6`
- `StringBIC` and `StringBICISO93622014`
- the three Base64 rules
- `StringLatitude` and `StringLongitude`
- the non-length part of `StringKubernetesQualifiedName`
- the ISO 3166 and ISO 4217 rules

The ISO rules can produce exact enums, but the resulting schemas can be large.
The other rules need tests that define whether their patterns are exact or
approximate.

### Algorithmic string rules

The following rules validate checksums, decoded content, or complex parsers:

- `StringCreditCard` and `StringLuhnChecksum`
- `StringISBN`, `StringISBN10`, `StringISBN13`, and `StringISSN`
- both BCP 47 rules
- `StringGitRef`
- `StringCrontab`
- `StringTitle`
- `StringJWT`

Patterns can document the outer shape of some values, but they do not preserve
the complete Govy validation. These mappings should be approximate or
unsupported.

## Implementation plan

### 1. Define observable generation behavior

Decide how callers learn that a rule was omitted or approximated. The main
options are generation diagnostics, a strict mode that returns an error, or
schema annotations.

### 2. Complete structural generation

Merge a complete JSON field tree into the validation plan when that metadata is
available. Add explicit handling for transformed properties.

### 3. Add missing standard keywords

Extend `jsonschema.Schema` only when a built-in rule needs the keyword. The
first candidates are:

- `uniqueItems`
- `contentEncoding`
- `contentMediaType`
- `contentSchema`

### 4. Implement rule mappings by confidence

Implement exact mappings first. Then add approximate mappings with documented
differences. Keep runtime-dependent and cross-property rules unsupported unless
the project explicitly accepts a nonportable extension.

### 5. Verify complete schema documents

Tests should read expected JSON documents from `pkg/govy/test_data` and compare
the complete marshaled schema. Each semantic gap that is fixed should have a
focused fixture before it is removed from this document.

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
- Should unsupported mappings be silent, reported, annotated, or errors?
- Should nonportable Ajv or Govy vocabulary output be supported?
- Should large finite enums be emitted by default?

## References

- [JSON Schema Draft 2020-12]
- [JSON Schema validation vocabulary]
- [Ajv `$data` references]

[Ajv `$data` references]: https://ajv.js.org/guide/combining-schemas.html#data-reference
[JSON Schema Draft 2020-12]: https://json-schema.org/draft/2020-12/schema
[JSON Schema validation vocabulary]: https://json-schema.org/draft/2020-12/json-schema-validation
