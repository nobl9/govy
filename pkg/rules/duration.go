package rules

import (
	"errors"
	"fmt"
	"time"

	"github.com/nobl9/govy/pkg/govy"
)

func DurationPrecision(precision time.Duration) govy.SingleRule[time.Duration] {
	msg := fmt.Sprintf("duration must be defined with %s precision", precision)
	return govy.NewSingleRule(func(v time.Duration) error {
		if v.Nanoseconds()%int64(precision) != 0 {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeDurationPrecision).
		WithDescription(msg)
}
