package validation

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func DurationPrecision(precision time.Duration) SingleRule[time.Duration] {
	msg := fmt.Sprintf("duration must be defined with %s precision", precision)
	return NewSingleRule(func(v time.Duration) error {
		if v.Nanoseconds()%int64(precision) != 0 {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeDurationPrecision).
		WithDescription(msg)
}
