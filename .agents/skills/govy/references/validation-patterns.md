# Validation Patterns

- Define shared rules and validators once.
Use factories when configuration varies.
  Builder methods return copies, so retain the returned value.
- Choose an existing `pkg/rules` rule before writing a custom rule.
  Keep domain-specific rules in the consumer.
- Match the data structure with `Include` and `IncludeForEach`.
  Govy adds parent paths and collection indexes to nested errors.
- Use `WithName` for one JSON path segment and `WithPath` for multiple segments.
  Preserve existing path inference when the project uses it.
  Otherwise, use explicit paths unless the task requests inference.
- Separate presence from content. `ForPointer(...).Required()` rejects `nil`,
  but accepts a pointer to a zero value unless a content rule rejects it.
  Direct-value `Required()` rejects zero values.
Use `OmitEmpty()` only when zero means absent.
- Put conditions at the narrowest applicable level.
  Use cascade stops when later checks depend on an earlier check passing.
  Keep independent property failures visible.
- Give custom rules stable error codes and useful descriptions.
  When generating plans, describe conditions with `govy.WhenDescription`.
  Add examples or details when they clarify accepted input or a failure.
- Test through the consumer's public validation entrypoint and existing test helpers.
  Check paths, codes, and error counts.
Assert rendered messages when their wording is part of the contract.

## Readable rule layout

Make the validator, property, and rule levels easy to distinguish.
Use the following layout for new validation code.
When editing existing code, preserve its local style where it remains clear.

- Give each property its own block in `govy.New(...)`.
  Keep a simple getter inline, such as `func(u User) string { return u.Name }`.
  Put chained property methods on separate lines,
  with the dot on the preceding line.
  Place `WithName` or `WithPath` near the getter so the property path is visible.
- Keep a short single-rule call compact, such as `Rules(rules.StringNotEmpty())`.
  For multiple rules, expand `Rules(...)` or `govy.NewRuleSet(...)`
  with one rule per line and a trailing comma.
  Also expand a single-rule call when its configuration makes the nesting
  hard to follow.
- Indent rule metadata, such as `WithDetails` and `WithErrorCode`,
  beneath the rule constructor.
  Keep property methods aligned at the property level
  and validator methods at the validator level.
  This makes the scope of `When` and `Cascade` visible.
- Keep getters focused on value access.
  Extract a custom rule or predicate when its callback body obscures
  the property list.
  Name the rule or predicate after its purpose.
  Keep short, clear expressions inline.
- Preserve rule order and the scope of conditions, transforms,
  and cascade settings during formatting.
  Moving a setting between a rule, property, and validator can change
  validation behavior.
- Run the consumer's Go formatter after choosing the layout.

The [rule set example](rules-and-messages.md#group-multiple-rules-into-a-reusable-set)
shows a multiline rule list, rule metadata, and a compact property rule call.
For a full validator with nested properties, see
[core validation](core-validation.md#build-a-full-teacher-validator).

## Examples

The examples in these references are excerpts from tested Go source.
They can use supporting types or helpers from their original package.
Adapt those definitions and imports to the consumer.

- For construction, naming, and validator conditions, read [core validation](core-validation.md).
- For property paths, pointers, and transforms, read [properties and paths](properties-and-paths.md).
- For nested objects and collections, read [collections and composition](collections-and-composition.md).
- For relationships between fields, read [predefined rule examples](predefined-rules.md).
- For custom rules and metadata, read [rules and messages](rules-and-messages.md).
- For tests, read [testing](testing.md).
