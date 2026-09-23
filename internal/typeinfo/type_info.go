package typeinfo

import (
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// TypeInfo stores the Go type information.
type TypeInfo struct {
	Name        string
	Kind        string
	Package     string
	ReflectKind reflect.Kind
	// JSONKind overrides ReflectKind for special encodings.
	// Invalid means no override; Interface means an unknown custom JSON representation.
	JSONKind reflect.Kind
}

// Get returns the information for the type T.
// It returns the type name without package path or name.
// It strips the pointer '*' from the type name.
// Package is only available if the type is not a built-in type.
//
// It has a special treatment for slices of type definitions.
// Instead of having:
//
//	TypeInfo{Name: "[]mypkg.Bar"}
//
// It will produce:
//
//	TypeInfo{Name: "[]Bar", Package: ".../mypkg"}.
func Get[T any]() TypeInfo {
	typ := reflect.TypeOf(*new(T))
	if typ == nil {
		return TypeInfo{}
	}
	jsonType := jsonKind(typ)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	info := TypeInfo{
		Kind:        getKindString(typ),
		ReflectKind: typ.Kind(),
		JSONKind:    jsonType,
	}

	if typ.PkgPath() == "" && typ.Kind() == reflect.Slice {
		info.Name = "[]"
		typ = typ.Elem()
	}
	switch {
	case typ.PkgPath() == "":
		info.Name += typ.String()
	default:
		info.Name += typ.Name()
		info.Package = typ.PkgPath()
	}
	return info
}

func jsonKind(typ reflect.Type) reflect.Kind {
	base := typ
	if base.Kind() == reflect.Pointer {
		base = base.Elem()
	}
	switch base {
	case reflect.TypeFor[json.Number]():
		return reflect.Float64
	case reflect.TypeFor[time.Time]():
		return reflect.String
	}
	if typ.Implements(reflect.TypeFor[json.Marshaler]()) {
		return reflect.Interface
	}
	if typ.Implements(reflect.TypeFor[encoding.TextMarshaler]()) {
		return reflect.String
	}
	// Pointer methods can change the encoding of addressable struct fields.
	// The selected property's type alone does not establish addressability.
	if reflect.PointerTo(typ).Implements(reflect.TypeFor[json.Marshaler]()) ||
		reflect.PointerTo(typ).Implements(reflect.TypeFor[encoding.TextMarshaler]()) {
		return reflect.Interface
	}
	if base.Kind() == reflect.Slice && base.Elem().Kind() == reflect.Uint8 {
		element := reflect.PointerTo(base.Elem())
		if !element.Implements(reflect.TypeFor[json.Marshaler]()) &&
			!element.Implements(reflect.TypeFor[encoding.TextMarshaler]()) {
			return reflect.String
		}
	}
	return reflect.Invalid
}

func getKindString(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", getKindString(typ.Key()), getKindString(typ.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", getKindString(typ.Elem()))
	default:
		return typ.Kind().String()
	}
}
