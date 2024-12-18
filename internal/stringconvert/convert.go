package stringconvert

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"
)

// Convert converts any value to a pretty, human-readable string representation.
func Convert(v any) string {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		return Convert(rv.Elem())
	}
	switch rv.Kind() {
	case reflect.Struct, reflect.Map:
		data, err := json.Marshal(v)
		if err != nil {
			slog.Error("unexpected error", slog.String("err", err.Error()))
		}
		return string(data)
	case reflect.Slice, reflect.Array:
		result := "["
		for i := 0; i < rv.Len(); i++ {
			if i > 0 {
				result += ", "
			}
			result += Convert(rv.Index(i).Interface())
		}
		return result + "]"
	default:
		return fmt.Sprint(v)
	}
}
