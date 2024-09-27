package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// JoinErrors joins multiple errors into a single pretty-formatted string.
// JoinErrors assumes the errors are not nil, if this presumption is broken the formatting might not be correct.
func JoinErrors[T error](b *strings.Builder, errs []T, indent string) {
	for i, err := range errs {
		if error(err) == nil {
			continue
		}
		buildErrorMessage(b, err.Error(), indent)
		if i < len(errs)-1 {
			b.WriteString("\n")
		}
	}
}

const listPoint = "- "

func buildErrorMessage(b *strings.Builder, errMsg, indent string) {
	b.WriteString(indent)
	if !strings.HasPrefix(errMsg, listPoint) {
		b.WriteString(listPoint)
	}
	// Indent the whole error message.
	errMsg = strings.ReplaceAll(errMsg, "\n", "\n"+indent)
	b.WriteString(errMsg)
}

var newLineReplacer = strings.NewReplacer("\n", "\\n", "\r", "\\r")

// PropertyValueString returns the string representation of the given value.
// Structs, interfaces, maps and slices are converted to compacted JSON strings (see struct exceptions below).
// It tries to improve readability by:
//   - limiting the string to 100 characters
//   - removing leading and trailing whitespaces
//   - escaping newlines
//
// If value is a struct implementing [fmt.Stringer] [fmt.Stringer.String] method will be used only if:
//   - the struct does not contain any JSON tags
//   - the struct is not empty or it is empty but does not have any fields
//
// If a value is a struct of type [time.Time] it will be formatted using [time.RFC3339] layout.
func PropertyValueString(v interface{}) string {
	if v == nil {
		return ""
	}
	rv := reflect.ValueOf(v)
	ft := reflect.Indirect(rv)
	var s string
	switch ft.Kind() {
	case reflect.Interface, reflect.Map, reflect.Slice:
		if rv.IsZero() {
			break
		}
		raw, _ := json.Marshal(v)
		s = string(raw)
	case reflect.Struct:
		// If the struct is empty and it has.
		if rv.IsZero() && rv.NumField() != 0 {
			break
		}
		if timeDate, ok := v.(time.Time); ok {
			s = timeDate.Format(time.RFC3339)
			break
		}
		if stringer, ok := v.(fmt.Stringer); ok && !hasJSONTags(v, rv.Kind() == reflect.Pointer) {
			s = stringer.String()
			break
		}
		raw, _ := json.Marshal(v)
		s = string(raw)
	case reflect.Ptr:
		if rv.IsNil() {
			return ""
		}
		deref := rv.Elem().Interface()
		return PropertyValueString(deref)
	case reflect.Func:
		return "func"
	case reflect.Invalid:
		return ""
	default:
		s = fmt.Sprint(ft.Interface())
	}
	s = limitString(s, 100)
	s = strings.TrimSpace(s)
	s = newLineReplacer.Replace(s)
	return s
}

func hasJSONTags(v interface{}, isPointer bool) bool {
	t := reflect.TypeOf(v)
	if isPointer {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if _, hasTag := field.Tag.Lookup("json"); hasTag {
			return true
		}
	}
	return false
}

func limitString(s string, limit int) string {
	if len(s) > limit {
		return s[:limit] + "..."
	}
	return s
}
