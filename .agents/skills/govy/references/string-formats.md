# String Formats

Match the rule to the input contract.
For the full constructor list, read [Existing rules](existing-rules.md).

## Base64

- Use `StringBase64` for the standard alphabet with required padding.
- Use `StringBase64URL` for the URL-safe alphabet with required padding.
- Use `StringBase64RawURL` for the URL-safe alphabet without padding.

Padding is required only when the payload needs it.
For example, `Zg==` is padded and `Zg` is unpadded.
These rules reject whitespace and nonzero trailing padding bits.
They accept empty strings, so add `Required()` when the property must be present.

## UUIDs

`StringUUID` checks the hexadecimal and hyphen layout.
It accepts the nil UUID and does not restrict version or variant.
Use `StringUUIDRFC4122` for versions 1–5 with the RFC 4122 variant.
Use `StringUUIDv3`, `StringUUIDv4`, or `StringUUIDv5` when one version is required.

## Language tags

`StringBCP47LanguageTag` accepts case variants and deprecated aliases.
`StringBCP47StrictLanguageTag` requires canonical case and preferred forms.
For example, `en-US` is canonical.
`en-us` and `iw` require normalization.
Both rules reject underscores and duplicate variants or extensions.

## JWTs

`StringJWT` checks JWS compact syntax, including unsigned `alg: "none"` tokens.
It does not verify signatures, trusted algorithms, expiration, or other claims.
Use a JWT verification library when authenticity or claims matter.

## File paths

`StringAbsoluteFilePath` checks whether a path is absolute on the current platform.
It does not check whether the path exists.
Use `StringFilePath`, `StringDirPath`, or `StringFileSystemPath` for existence checks.
