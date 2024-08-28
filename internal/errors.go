package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// JoinErrors joins multiple errors into a single pretty-formatted string.
func JoinErrors[T error](b *strings.Builder, errs []T, indent string) {
	for i, err := range errs {
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
// Structs, interfaces, maps and slices are converted to compacted JSON strings.
// It tries to improve readability by:
// - limiting the string to 100 characters
// - removing leading and trailing whitespaces
// - escaping newlines
// If value is a struct implementing [fmt.Stringer] String method will be used
// only if the struct does not contain any JSON tags.
func PropertyValueString(v interface{}) string {
	if v == nil {
		return ""
	}
	rv := reflect.ValueOf(v)
	ft := reflect.Indirect(rv)
	var s string
	switch ft.Kind() {
	case reflect.Interface, reflect.Map, reflect.Slice:
		if reflect.ValueOf(v).IsZero() {
			break
		}
		raw, _ := json.Marshal(v)
		s = string(raw)
	case reflect.Struct:
		if reflect.ValueOf(v).IsZero() {
			break
		}
		if stringer, ok := v.(fmt.Stringer); ok && !hasJSONTags(v, rv.Kind() == reflect.Pointer) {
			s = stringer.String()
			break
		}
		raw, _ := json.Marshal(v)
		s = string(raw)
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
