package generateddocs

// Zulu validates a value. Additional detail.
func Zulu[T any]() govy.Rule[T] {
	panic("fixture")
}

// StringValue returns a string.
func StringValue() string {
	return "fixture"
}

// hidden is not exported.
func hidden[T any]() govy.Rule[T] {
	panic("fixture")
}
