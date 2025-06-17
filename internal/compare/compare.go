package compare

import (
	"reflect"
)

// EqualExportedFields compares two structs of the same type T and checks if their exported fields are equal.
// It returns true if all exported fields are equal, false otherwise.
// The function uses generics to ensure both parameters are of the same type at compile time.
func EqualExportedFields[T any](a, b T) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// Handle nil pointers
	if !va.IsValid() && !vb.IsValid() {
		return true
	}
	if !va.IsValid() || !vb.IsValid() {
		return false
	}

	// Dereference pointers if needed
	if va.Kind() == reflect.Ptr {
		if va.IsNil() && vb.IsNil() {
			return true
		}
		if va.IsNil() || vb.IsNil() {
			return false
		}
		va = va.Elem()
		vb = vb.Elem()
	}

	// If not a struct, use reflect.DeepEqual for direct comparison
	if va.Kind() != reflect.Struct {
		return reflect.DeepEqual(va.Interface(), vb.Interface())
	}

	return compareStructFields(va, vb)
}

// compareStructFields compares the exported fields of two struct values
func compareStructFields(va, vb reflect.Value) bool {
	structType := va.Type()

	for i := 0; i < va.NumField(); i++ {
		field := structType.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		fieldA := va.Field(i)
		fieldB := vb.Field(i)

		if !compareFieldValues(fieldA, fieldB) {
			return false
		}
	}

	return true
}

// compareFieldValues recursively compares two field values
func compareFieldValues(va, vb reflect.Value) bool {
	// Handle invalid values
	if !va.IsValid() && !vb.IsValid() {
		return true
	}
	if !va.IsValid() || !vb.IsValid() {
		return false
	}

	// Handle pointers
	if va.Kind() == reflect.Ptr && vb.Kind() == reflect.Ptr {
		if va.IsNil() && vb.IsNil() {
			return true
		}
		if va.IsNil() || vb.IsNil() {
			return false
		}
		return compareFieldValues(va.Elem(), vb.Elem())
	}

	// For nested structs, recursively compare their exported fields
	if va.Kind() == reflect.Struct {
		return compareStructFields(va, vb)
	}

	// For slices, compare each element
	if va.Kind() == reflect.Slice {
		if va.Len() != vb.Len() {
			return false
		}
		for i := 0; i < va.Len(); i++ {
			if !compareFieldValues(va.Index(i), vb.Index(i)) {
				return false
			}
		}
		return true
	}

	// For maps, compare keys and values
	if va.Kind() == reflect.Map {
		if va.Len() != vb.Len() {
			return false
		}
		for _, key := range va.MapKeys() {
			valueA := va.MapIndex(key)
			valueB := vb.MapIndex(key)
			if !valueB.IsValid() || !compareFieldValues(valueA, valueB) {
				return false
			}
		}
		return true
	}

	// For all other types, use reflect.DeepEqual
	return reflect.DeepEqual(va.Interface(), vb.Interface())
}
