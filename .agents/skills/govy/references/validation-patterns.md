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
