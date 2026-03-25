package jsonpath

// Path is a builder for constructing valid JSONPath property paths.
// It ensures proper escaping and formatting of path segments.
type Path struct{ path string }

// NewPath creates a new empty Path.
func NewPath() Path {
	return Path{}
}

// Name appends a named segment to the path, escaping special characters as needed.
func (p Path) Name(name string) Path {
	return Path{path: Join(p.path, EscapeSegment(name))}
}

// Index appends an array index segment to the path.
func (p Path) Index(index uint) Path {
	return Path{path: JoinArray(p.path, NewArrayIndex(index))}
}

// String returns the string representation of the path.
func (p Path) String() string {
	return p.path
}
