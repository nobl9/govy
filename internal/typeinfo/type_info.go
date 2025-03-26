package typeinfo

import (
	"fmt"
	"reflect"
)

// TypeInfo stores the type information.
type TypeInfo struct {
	Name    string
	Package string
	Kind    string
}

// Get returns the information for the type T.
// It returns the type name without package path or name.
// It strips the pointer '*' from the type name.
// Package is only available if the type is not a built-in type.
//
// It has a special treatment for slices of type definitions.
// Istead of having:
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
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	result := TypeInfo{
		Kind: getKindString(typ),
	}

	if typ.PkgPath() == "" && typ.Kind() == reflect.Slice {
		result.Name = "[]"
		typ = typ.Elem()
	}
	switch {
	case typ.PkgPath() == "":
		result.Name += typ.String()
	default:
		result.Name += typ.Name()
		result.Package = typ.PkgPath()
	}
	return result
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
