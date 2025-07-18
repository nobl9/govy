package rules

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	ErrorCodeRequired                      govy.ErrorCode = internal.RequiredErrorCode
	ErrorCodeForbidden                     govy.ErrorCode = "forbidden"
	ErrorCodeEqualTo                       govy.ErrorCode = "equal_to"
	ErrorCodeNotEqualTo                    govy.ErrorCode = "not_equal_to"
	ErrorCodeGreaterThan                   govy.ErrorCode = "greater_than"
	ErrorCodeGreaterThanOrEqualTo          govy.ErrorCode = "greater_than_or_equal_to"
	ErrorCodeLessThan                      govy.ErrorCode = "less_than"
	ErrorCodeLessThanOrEqualTo             govy.ErrorCode = "less_than_or_equal_to"
	ErrorCodeStringNotEmpty                govy.ErrorCode = "string_not_empty"
	ErrorCodeStringMatchRegexp             govy.ErrorCode = "string_match_regexp"
	ErrorCodeStringDenyRegexp              govy.ErrorCode = "string_deny_regexp"
	ErrorCodeStringDNSLabel                govy.ErrorCode = "string_dns_label"
	ErrorCodeStringDNSSubdomain            govy.ErrorCode = "string_dns_subdomain"
	ErrorCodeStringURL                     govy.ErrorCode = "string_url"
	ErrorCodeStringMAC                     govy.ErrorCode = "string_mac"
	ErrorCodeStringIP                      govy.ErrorCode = "string_ip"
	ErrorCodeStringIPv4                    govy.ErrorCode = "string_ipv4"
	ErrorCodeStringIPv6                    govy.ErrorCode = "string_ipv6"
	ErrorCodeStringCIDR                    govy.ErrorCode = "string_cidr"
	ErrorCodeStringCIDRv4                  govy.ErrorCode = "string_cidrv4"
	ErrorCodeStringCIDRv6                  govy.ErrorCode = "string_cidrv6"
	ErrorCodeStringASCII                   govy.ErrorCode = "string_ascii"
	ErrorCodeStringUUID                    govy.ErrorCode = "string_uuid"
	ErrorCodeStringEmail                   govy.ErrorCode = "string_email"
	ErrorCodeStringJSON                    govy.ErrorCode = "string_json"
	ErrorCodeStringContains                govy.ErrorCode = "string_contains"
	ErrorCodeStringExcludes                govy.ErrorCode = "string_excludes"
	ErrorCodeStringStartsWith              govy.ErrorCode = "string_starts_with"
	ErrorCodeStringEndsWith                govy.ErrorCode = "string_ends_with"
	ErrorCodeStringLength                  govy.ErrorCode = "string_length"
	ErrorCodeStringMinLength               govy.ErrorCode = "string_min_length"
	ErrorCodeStringMaxLength               govy.ErrorCode = "string_max_length"
	ErrorCodeStringTitle                   govy.ErrorCode = "string_title"
	ErrorCodeStringGitRef                  govy.ErrorCode = "string_git_ref"
	ErrorCodeStringFileSystemPath          govy.ErrorCode = "string_file_system_path"
	ErrorCodeStringMatchFileSystemPath     govy.ErrorCode = "string_match_file_system_path"
	ErrorCodeStringFilePath                govy.ErrorCode = "string_file_path"
	ErrorCodeStringDirPath                 govy.ErrorCode = "string_dir_path"
	ErrorCodeStringRegexp                  govy.ErrorCode = "string_regexp"
	ErrorCodeStringCrontab                 govy.ErrorCode = "string_crontab"
	ErrorCodeStringDateTime                govy.ErrorCode = "string_date_time"
	ErrorCodeStringTimeZone                govy.ErrorCode = "string_time_zone"
	ErrorCodeStringAlpha                   govy.ErrorCode = "string_alpha"
	ErrorCodeStringAlphanumeric            govy.ErrorCode = "string_alphanumeric"
	ErrorCodeStringAlphaUnicode            govy.ErrorCode = "string_alpha_unicode"
	ErrorCodeStringAlphanumericUnicode     govy.ErrorCode = "string_alphanumeric_unicode"
	ErrorCodeStringFQDN                    govy.ErrorCode = "string_fqdn"
	ErrorCodeStringKubernetesQualifiedName govy.ErrorCode = "string_kubernetes_qualified_name"
	ErrorCodeSliceLength                   govy.ErrorCode = "slice_length"
	ErrorCodeSliceMinLength                govy.ErrorCode = "slice_min_length"
	ErrorCodeSliceMaxLength                govy.ErrorCode = "slice_max_length"
	ErrorCodeMapLength                     govy.ErrorCode = "map_length"
	ErrorCodeMapMinLength                  govy.ErrorCode = "map_min_length"
	ErrorCodeMapMaxLength                  govy.ErrorCode = "map_max_length"
	ErrorCodeOneOf                         govy.ErrorCode = "one_of"
	ErrorCodeNotOneOf                      govy.ErrorCode = "not_one_of"
	ErrorCodeOneOfProperties               govy.ErrorCode = "one_of_properties"
	ErrorCodeMutuallyExclusive             govy.ErrorCode = "mutually_exclusive"
	ErrorCodeMutuallyDependent             govy.ErrorCode = "mutually_dependent"
	ErrorCodeEqualProperties               govy.ErrorCode = "equal_properties"
	ErrorCodeSliceUnique                   govy.ErrorCode = "slice_unique"
	ErrorCodeUniqueProperties              govy.ErrorCode = "unique_properties"
	ErrorCodeURL                           govy.ErrorCode = "url"
	ErrorCodeDurationPrecision             govy.ErrorCode = "duration_precision"
)
