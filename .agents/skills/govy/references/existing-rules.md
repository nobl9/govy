# Existing Rules

Generated reference of exported predefined rule constructors from `pkg/rules`.
Use this when choosing an existing rule before writing a custom one.
For UUID, base64, language-tag, JWT, and path variants, read
[String formats](string-formats.md).

<!-- Generated summaries retain the source documentation line layout. -->
<!-- markdownlint-disable MD013 -->

[//]: # (docs: pkg/rules?kind=func&returns=govy.Rule,govy.RuleSet)

- `DurationPrecision` - DurationPrecision ensures the duration is defined with the specified precision.
- `EQ` - EQ ensures the property's value is equal to the compared value.
- `EqualProperties` - EqualProperties checks if all the specified properties are equal.
- `Forbidden` - Forbidden ensures the property's value is its type's zero value, i.e. it's empty.
- `GT` - GT ensures the property's value is greater than the compared value.
- `GTComparableProperties` - GTComparableProperties ensures the first property's value is greater than the second property's value.
- `GTE` - GTE ensures the property's value is greater than or equal to the compared value.
- `GTEComparableProperties` - GTEComparableProperties ensures the first property's value is greater than or equal to the second property's value.
- `GTEProperties` - GTEProperties ensures the first property's value is greater than or equal to the second property's value.
- `GTProperties` - GTProperties ensures the first property's value is greater than the second property's value.
- `LT` - LT ensures the property's value is less than the compared value.
- `LTComparableProperties` - LTComparableProperties ensures the first property's value is less than the second property's value.
- `LTE` - LTE ensures the property's value is less than or equal to the compared value.
- `LTEComparableProperties` - LTEComparableProperties ensures the first property's value is less than or equal to the second property's value.
- `LTEProperties` - LTEProperties ensures the first property's value is less than or equal to the second property's value.
- `LTProperties` - LTProperties ensures the first property's value is less than the second property's value.
- `MapLength` - MapLength ensures the map's length is between min and max (closed interval).
- `MapMaxLength` - MapMaxLength ensures the map's length is less than or equal to the limit.
- `MapMinLength` - MapMinLength ensures the map's length is greater than or equal to the limit.
- `MutuallyDependent` - MutuallyDependent checks if properties are mutually dependent.
- `MutuallyExclusive` - MutuallyExclusive checks if properties are mutually exclusive.
- `NEQ` - NEQ ensures the property's value is not equal to the compared value.
- `NotOneOf` - NotOneOf checks if the property's value does not match any of the provided values.
- `OneOf` - OneOf checks if the property's value matches one of the provided values.
- `OneOfProperties` - OneOfProperties checks if at least one of the properties is set.
- `Required` - Required ensures the property's value is not empty (i.e. it's not its type's zero value).
- `SliceLength` - SliceLength ensures the slice's length is between min and max (closed interval).
- `SliceMaxLength` - SliceMaxLength ensures the slice's length is less than or equal to the limit.
- `SliceMinLength` - SliceMinLength ensures the slice's length is greater than or equal to the limit.
- `SliceUnique` - SliceUnique ensures that a slice contains unique elements based on a provided [HashFunction].
- `StringASCII` - StringASCII ensures property's value contains only ASCII characters.
- `StringAbsoluteFilePath` - StringAbsoluteFilePath ensures the property's value is an absolute file system path.
- `StringAlpha` - StringAlpha ensures the property's value consists only of ASCII letters.
- `StringAlphaUnicode` - StringAlphaUnicode ensures the property's value consists only of Unicode letters.
- `StringAlphanumeric` - StringAlphanumeric ensures the property's value consists only of ASCII letters and numbers.
- `StringAlphanumericUnicode` - StringAlphanumericUnicode ensures the property's value consists only of Unicode letters and numbers.
- `StringBCP47LanguageTag` - StringBCP47LanguageTag ensures the property's value is a valid BCP 47 language tag.
- `StringBCP47StrictLanguageTag` - StringBCP47StrictLanguageTag ensures the property's value is a valid canonical BCP 47 language tag.
- `StringBIC` - StringBIC ensures the property's value matches the current Business Identifier Code (BIC) syntax.
- `StringBICISO93622014` - StringBICISO93622014 ensures the property's value matches the ISO 9362:2014 Business Identifier Code (BIC) syntax.
- `StringBase64` - StringBase64 ensures the property's value is a standard padded base64 string.
- `StringBase64RawURL` - StringBase64RawURL ensures the property's value is a URL-safe base64 string without padding.
- `StringBase64URL` - StringBase64URL ensures the property's value is a URL-safe padded base64 string.
- `StringCIDR` - StringCIDR ensures property's value is a valid CIDR notation IP address.
- `StringCIDRv4` - StringCIDRv4 ensures property's value is a valid CIDR notation IPv4 address.
- `StringCIDRv6` - StringCIDRv6 ensures property's value is a valid CIDR notation IPv6 address.
- `StringCVE` - StringCVE ensures the property's value is a valid CVE ID.
- `StringContains` - StringContains ensures the property's value contains all the provided substrings.
- `StringCreditCard` - StringCreditCard ensures the property's value is a plausible digit-only payment card number.
- `StringCrontab` - StringCrontab ensures the property's value is a valid crontab schedule expression.
- `StringDNSLabel` - StringDNSLabel ensures the property's value is a valid DNS label as defined by [RFC 1123].
- `StringDNSSubdomain` - StringDNSSubdomain ensures the property's value is a valid DNS subdomain as defined by [RFC 1123].
- `StringDateTime` - StringDateTime ensures the property's value is a valid date and time in the specified layout.
- `StringDenyRegexp` - StringDenyRegexp ensures the property's value does not match the regular expression.
- `StringDirPath` - StringDirPath ensures the property's value is a file system path pointing to an existing directory.
- `StringE164` - StringE164 ensures the property's value is a valid E.164 phone number.
- `StringEIN` - StringEIN ensures the property's value is a United States Employer Identification Number (EIN) in NN-NNNNNNN format with a recognized prefix.
- `StringEmail` - StringEmail ensures the property's value is a valid email address.
- `StringEndsWith` - StringEndsWith ensures the property's value ends with one of the provided suffixes.
- `StringExcludes` - StringExcludes ensures the property's value does not contain any of the provided substrings.
- `StringFQDN` - StringFQDN ensures the property's value is a fully qualified domain name (FQDN).
- `StringFilePath` - StringFilePath ensures the property's value is a file system path pointing to an existing file.
- `StringFileSystemPath` - StringFileSystemPath ensures the property's value is an existing file system path.
- `StringGitRef` - StringGitRef ensures a git reference name follows the [git-check-ref-format] rules.
- `StringHexadecimal` - StringHexadecimal ensures the property's value is a hexadecimal string.
- `StringIP` - StringIP ensures property's value is a valid IP address.
- `StringIPv4` - StringIPv4 ensures property's value is a valid IPv4 address.
- `StringIPv6` - StringIPv6 ensures property's value is a valid IPv6 address.
- `StringISBN` - StringISBN ensures the property's value is a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format.
- `StringISBN10` - StringISBN10 ensures the property's value is a valid International Standard Book Number (ISBN) in ISBN-10 format.
- `StringISBN13` - StringISBN13 ensures the property's value is a valid International Standard Book Number (ISBN) in ISBN-13 format.
- `StringISO31662` - StringISO31662 ensures the property's value is a valid ISO 3166-2 country subdivision code.
- `StringISO3166Alpha2` - StringISO3166Alpha2 ensures the property's value is a valid ISO 3166-1 alpha-2 country code.
- `StringISO3166Alpha3` - StringISO3166Alpha3 ensures the property's value is a valid ISO 3166-1 alpha-3 country code.
- `StringISO3166Numeric` - StringISO3166Numeric ensures the property's value is a valid ISO 3166-1 numeric-3 country code.
- `StringISO4217` - StringISO4217 ensures the property's value is a valid ISO 4217 three-letter alphabetic currency code.
- `StringISSN` - StringISSN ensures the property's value is a valid International Standard Serial Number (ISSN).
- `StringJSON` - StringJSON ensures property's value is a valid JSON literal.
- `StringJWT` - StringJWT ensures the property's value is a JSON Web Token (JWT) represented using [JWS Compact Serialization].
- `StringKubernetesQualifiedName` - StringKubernetesQualifiedName ensures the property's value is a valid "qualified name" as defined by [Kubernetes validation].
- `StringLatitude` - StringLatitude ensures the property's value is a decimal latitude coordinate between -90 and 90 degrees.
- `StringLength` - StringLength ensures the string's length is between min and max (closed interval).
- `StringLongitude` - StringLongitude ensures the property's value is a decimal longitude coordinate between -180 and 180 degrees.
- `StringLuhnChecksum` - StringLuhnChecksum ensures the property's value is a digit-only string that passes the Luhn checksum.
- `StringMAC` - StringMAC ensures property's value is a valid MAC address.
- `StringMD5` - StringMD5 ensures the property's value is a lowercase hexadecimal MD5 digest.
- `StringMatchFileSystemPath` - StringMatchFileSystemPath ensures the property's value matches the provided file path pattern.
- `StringMatchRegexp` - StringMatchRegexp ensures the property's value matches the regular expression.
- `StringMaxLength` - StringMaxLength ensures the string's length is less than or equal to the limit.
- `StringMinLength` - StringMinLength ensures the string's length is greater than or equal to the limit.
- `StringMongoDBObjectID` - StringMongoDBObjectID ensures the property's value is a 24-character hexadecimal MongoDB ObjectID.
- `StringNotEmpty` - StringNotEmpty ensures the property's value is not empty.
- `StringRegexp` - StringRegexp ensures the property's value is a valid regular expression.
- `StringSHA256` - StringSHA256 ensures the property's value is a lowercase hexadecimal SHA-256 digest.
- `StringSHA384` - StringSHA384 ensures the property's value is a lowercase hexadecimal SHA-384 digest.
- `StringSHA512` - StringSHA512 ensures the property's value is a lowercase hexadecimal SHA-512 digest.
- `StringSSN` - StringSSN ensures the property's value is a United States Social Security Number (SSN) in NNN-NN-NNNN format.
- `StringSemver` - StringSemver ensures the property's value is a valid Semantic Versioning 2.0.0 version.
- `StringStartsWith` - StringStartsWith ensures the property's value starts with one of the provided prefixes.
- `StringTimeZone` - StringTimeZone ensures the property's value is a valid time zone name which uniquely identifies a time zone in the IANA Time Zone database.
- `StringTitle` - StringTitle ensures each word in a string starts with a capital letter.
- `StringULID` - StringULID ensures the property's value is a 26-character Crockford Base32 Universally Unique Lexicographically Sortable Identifier (ULID) string.
- `StringURL` - StringURL ensures property's value is a valid URL as defined by [url.Parse] function.
- `StringUUID` - StringUUID ensures property's value is a valid UUID string as defined by [RFC 4122].
- `StringUUIDRFC4122` - StringUUIDRFC4122 ensures the property's value is a Universally Unique Identifier (UUID) string as defined by RFC 4122.
- `StringUUIDv3` - StringUUIDv3 ensures the property's value is a version 3 Universally Unique Identifier (UUID) string as defined by RFC 4122.
- `StringUUIDv4` - StringUUIDv4 ensures the property's value is a version 4 Universally Unique Identifier (UUID) string as defined by RFC 4122.
- `StringUUIDv5` - StringUUIDv5 ensures the property's value is a version 5 Universally Unique Identifier (UUID) string as defined by RFC 4122.
- `URL` - URL ensures the URL is valid.
- `UniqueProperties` - UniqueProperties ensures each property is unique based on a provided [HashFunction].

[//]: # (end-docs)
