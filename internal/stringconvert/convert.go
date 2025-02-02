package stringconvert

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/nobl9/govy/internal/logging"
)

// Format converts any value to a pretty, human-readable string representation.
func Format(v any) string {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
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
		result := "["
		for i := 0; i < rv.Len(); i++ {
			if i > 0 {
				result += ", "
			}
			result += Format(rv.Index(i).Interface())
		}
		return result + "]"
	default:
		return fmt.Sprint(v)
	}
}
