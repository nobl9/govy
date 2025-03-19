package govy

// ErrorCode is a unique string that represents a specific [RuleError].
// It can be used to precisely identify the error without inspecting its message.
type ErrorCode = string

const (
	ErrorCodeTransform ErrorCode = "transform"
)

func addErrorCode(c1, c2 ErrorCode) ErrorCode {
	return concatStrings(c1, c2, ErrorCodeSeparator)
}
