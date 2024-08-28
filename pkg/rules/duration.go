package rules

import (
	"errors"
	"fmt"
	"time"

	"github.com/nobl9/govy/pkg/govy"
)

// DurationPrecision ensures the duration is defined with the specified precision.
func DurationPrecision(precision time.Duration) govy.Rule[time.Duration] {
	msg := fmt.Sprintf("duration must be defined with %s precision", precision)
	return govy.NewRule(func(v time.Duration) error {
		if v.Nanoseconds()%int64(precision) != 0 {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeDurationPrecision).
		WithDescription(msg)
}
