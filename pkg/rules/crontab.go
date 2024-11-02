package rules

import (
	"errors"
	"strconv"
	"strings"
)

var errInvalidCrontab = errors.New("invalid crontab expression")

var crontabMonthsMap = map[string]int{
	"JAN": 1, "FEB": 2, "MAR": 3, "APR": 4,
	"MAY": 5, "JUN": 6, "JUL": 7, "AUG": 8,
	"SEP": 9, "OCT": 10, "NOV": 11, "DEC": 12,
}

var crontabDaysMap = map[string]int{
	"SUN": 0, "MON": 1, "TUE": 2, "WED": 3,
	"THU": 4, "FRI": 5, "SAT": 6,
}

func parseCrontab(c string) error {
	if strings.HasPrefix(c, "@") {
		switch c {
		case "@reboot", "@yearly", "@annually", "@monthly", "@weekly", "@daily", "@hourly":
			return nil
		default:
			return errInvalidCrontab
		}
	}
	fields := strings.Fields(c)
	if len(fields) != 5 {
		return errors.New("crontab expression must have exactly 5 fields")
	}
	for i, field := range fields {
		if field == "*" {
			continue
		}
		if strings.HasPrefix(field, "*/") {
			if len(field) < 3 {
				return errInvalidCrontab
			}
			if _, err := strconv.Atoi(field[2:]); err != nil {
				return errInvalidCrontab
			}
			continue
		}
		switch i {
		case 0:
			if !validateCrontabField(field, 0, 59, crontabParseStandardField) {
				return errInvalidCrontab
			}
		case 1:
			if !validateCrontabField(field, 0, 23, crontabParseStandardField) {
				return errInvalidCrontab
			}
		case 2:
			if !validateCrontabField(field, 1, 31, crontabParseStandardField) {
				return errInvalidCrontab
			}
		case 3:
			if !validateCrontabField(field, 1, 12, crontabParseMonthField) {
				return errInvalidCrontab
			}
		case 4:
			if !validateCrontabField(field, 0, 7, crontabParseDayField) {
				return errInvalidCrontab
			}
		}
	}
	return nil
}

type crontabFieldParseFunc func(string) (int, bool)

func crontabParseStandardField(v string) (int, bool) {
	if v == "" {
		return -1, false
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return -1, false
	}
	return i, true
}

func crontabParseMonthField(v string) (int, bool) {
	if v == "" {
		return -1, false
	}
	if i, ok := crontabMonthsMap[strings.ToUpper(v)]; ok {
		return i, true
	}
	return crontabParseStandardField(v)
}

func crontabParseDayField(v string) (int, bool) {
	if v == "" {
		return -1, false
	}
	if i, ok := crontabDaysMap[strings.ToUpper(v)]; ok {
		return i, true
	}
	return crontabParseStandardField(v)
}

func validateCrontabField(field string, lowerLimit, upperLimit int, parse crontabFieldParseFunc) bool {
	for _, el := range strings.Split(field, ",") {
		rangeIdx := strings.Index(el, "-")
		if rangeIdx != -1 {
			if rangeIdx == 0 || rangeIdx == len(el)-1 {
				return false
			}
			// Check lower range bound.
			l, ok := parse(el[:rangeIdx])
			if !ok {
				return false
			}
			if l < lowerLimit || l > upperLimit {
				return false
			}
			// Take step value into account.
			stepIdx := strings.Index(el, "/")
			if stepIdx == -1 {
				stepIdx = len(el)
			} else if stepIdx == len(el)-1 {
				return false
			} else {
				if v, err := strconv.Atoi(el[stepIdx+1:]); err != nil || v < 0 {
					return false
				}
			}
			// Check upper range bound.
			u, ok := parse(el[rangeIdx+1 : stepIdx])
			if !ok {
				return false
			}
			if u < lowerLimit || u > upperLimit {
				return false
			}
			// Compare lower and upper bounds.
			if l > u {
				return false
			}
			break
		} else {
			v, ok := parse(el)
			if !ok {
				return false
			}
			if v < lowerLimit || v > upperLimit {
				return false
			}
		}
	}
	return true
}
