package stringconvert

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"github.com/nobl9/govy/internal/logging"
)

// Format converts any value to a pretty, human-readable string representation.
func Format(v any) string {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		return Format(rv.Elem())
	}
	switch rv.Kind() {
	case reflect.Struct, reflect.Map:
		data, err := json.Marshal(v)
		if err != nil {
			logging.Logger().Error("unexpected error", slog.String("err", err.Error()))
		}
		return string(data)
	case reflect.Slice, reflect.Array:
		var result strings.Builder
		result.WriteString("[")
		for i := 0; i < rv.Len(); i++ {
			if i > 0 {
				result.WriteString(", ")
			}
			result.WriteString(Format(rv.Index(i).Interface()))
		}
		return result.String() + "]"
	default:
		return fmt.Sprint(v)
	}
}
