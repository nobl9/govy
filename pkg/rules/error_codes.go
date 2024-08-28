package rules

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	ErrorCodeRequired             govy.ErrorCode = internal.RequiredErrorCodeString
	ErrorCodeForbidden            govy.ErrorCode = "forbidden"
	ErrorCodeEqualTo              govy.ErrorCode = "equal_to"
	ErrorCodeNotEqualTo           govy.ErrorCode = "not_equal_to"
	ErrorCodeGreaterThan          govy.ErrorCode = "greater_than"
	ErrorCodeGreaterThanOrEqualTo govy.ErrorCode = "greater_than_or_equal_to"
	ErrorCodeLessThan             govy.ErrorCode = "less_than"
	ErrorCodeLessThanOrEqualTo    govy.ErrorCode = "less_than_or_equal_to"
	ErrorCodeStringNotEmpty       govy.ErrorCode = "string_not_empty"
	ErrorCodeStringMatchRegexp    govy.ErrorCode = "string_match_regexp"
	ErrorCodeStringDenyRegexp     govy.ErrorCode = "string_deny_regexp"
	ErrorCodeStringDNSLabel       govy.ErrorCode = "string_dns_label"
	ErrorCodeStringASCII          govy.ErrorCode = "string_ascii"
	ErrorCodeStringURL            govy.ErrorCode = "string_url"
	ErrorCodeStringUUID           govy.ErrorCode = "string_uuid"
	ErrorCodeStringJSON           govy.ErrorCode = "string_json"
	ErrorCodeStringContains       govy.ErrorCode = "string_contains"
	ErrorCodeStringStartsWith     govy.ErrorCode = "string_starts_with"
	ErrorCodeStringEndsWith       govy.ErrorCode = "string_ends_with"
	ErrorCodeStringLength         govy.ErrorCode = "string_length"
	ErrorCodeStringMinLength      govy.ErrorCode = "string_min_length"
	ErrorCodeStringMaxLength      govy.ErrorCode = "string_max_length"
	ErrorCodeStringTitle          govy.ErrorCode = "string_title"
	ErrorCodeSliceLength          govy.ErrorCode = "slice_length"
	ErrorCodeSliceMinLength       govy.ErrorCode = "slice_min_length"
	ErrorCodeSliceMaxLength       govy.ErrorCode = "slice_max_length"
	ErrorCodeMapLength            govy.ErrorCode = "map_length"
	ErrorCodeMapMinLength         govy.ErrorCode = "map_min_length"
	ErrorCodeMapMaxLength         govy.ErrorCode = "map_max_length"
	ErrorCodeOneOf                govy.ErrorCode = "one_of"
	ErrorCodeMutuallyExclusive    govy.ErrorCode = "mutually_exclusive"
	ErrorCodeSliceUnique          govy.ErrorCode = "slice_unique"
	ErrorCodeURL                  govy.ErrorCode = "url"
	ErrorCodeDurationPrecision    govy.ErrorCode = "duration_precision"
)
