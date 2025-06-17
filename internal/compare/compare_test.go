package compare

import (
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
)

func TestEqualExportedFields(t *testing.T) {
	type Person struct {
		Name    string
		Age     int
		private string //nolint:unused
	}
	type PersonWithPointers struct {
		Name    *string
		Age     *int
		private *string //nolint:unused
	}
	type Company struct {
		Name      string
		CEO       Person
		Employees []Person
	}
	type ComplexStruct struct {
		ID       int
		Tags     []string
		Metadata map[string]string
		People   []Person
	}
	type Event struct {
		Name      string
		Timestamp time.Time
		EndTime   *time.Time
	}

	// Helper function to create string pointer
	ptr := func(s string) *string { return &s }
	intPtr := func(i int) *int { return &i }

	tests := []struct {
		name     string
		testFunc func() bool
		expected bool
	}{
		// Basic struct tests
		{
			name: "simple structs equal ignoring private fields",
			testFunc: func() bool {
				a := Person{Name: "John", Age: 30, private: "secret1"}
				b := Person{Name: "John", Age: 30, private: "secret2"}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "simple structs different names",
			testFunc: func() bool {
				a := Person{Name: "John", Age: 30}
				b := Person{Name: "Jane", Age: 30}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},
		{
			name: "simple structs different ages",
			testFunc: func() bool {
				a := Person{Name: "John", Age: 30}
				b := Person{Name: "John", Age: 25}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},

		// Pointer tests
		{
			name: "pointer structs both nil",
			testFunc: func() bool {
				a := PersonWithPointers{Name: nil, Age: nil}
				b := PersonWithPointers{Name: nil, Age: nil}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "pointer structs equal values",
			testFunc: func() bool {
				name1, name2 := "John", "John"
				age1, age2 := 30, 30
				a := PersonWithPointers{Name: &name1, Age: &age1}
				b := PersonWithPointers{Name: &name2, Age: &age2}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "pointer structs different values",
			testFunc: func() bool {
				a := PersonWithPointers{Name: ptr("John"), Age: intPtr(30)}
				b := PersonWithPointers{Name: ptr("Jane"), Age: intPtr(30)}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},
		{
			name: "pointer structs one nil one not",
			testFunc: func() bool {
				a := PersonWithPointers{Name: ptr("John"), Age: nil}
				b := PersonWithPointers{Name: ptr("John"), Age: intPtr(30)}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},

		// Struct pointer tests
		{
			name: "struct pointers equal",
			testFunc: func() bool {
				a := &Person{Name: "John", Age: 30}
				b := &Person{Name: "John", Age: 30}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "struct pointers different",
			testFunc: func() bool {
				a := &Person{Name: "John", Age: 30}
				b := &Person{Name: "Jane", Age: 30}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},
		{
			name: "struct pointers both nil",
			testFunc: func() bool {
				var a, b *Person
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "struct pointers one nil",
			testFunc: func() bool {
				var a *Person
				b := &Person{Name: "John", Age: 30}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},
		{
			name: "nested structs equal",
			testFunc: func() bool {
				a := Company{
					Name: "ACME Corp",
					CEO:  Person{Name: "John", Age: 45},
					Employees: []Person{
						{Name: "Alice", Age: 30},
						{Name: "Bob", Age: 25},
					},
				}
				b := Company{
					Name: "ACME Corp",
					CEO:  Person{Name: "John", Age: 45},
					Employees: []Person{
						{Name: "Alice", Age: 30},
						{Name: "Bob", Age: 25},
					},
				}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "nested structs different CEO",
			testFunc: func() bool {
				a := Company{
					Name: "ACME Corp",
					CEO:  Person{Name: "John", Age: 45},
				}
				b := Company{
					Name: "ACME Corp",
					CEO:  Person{Name: "Jane", Age: 45},
				}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},

		// Complex struct tests with maps and slices
		{
			name: "complex structs equal",
			testFunc: func() bool {
				a := ComplexStruct{
					ID:   1,
					Tags: []string{"important", "urgent"},
					Metadata: map[string]string{
						"department": "engineering",
						"project":    "alpha",
					},
					People: []Person{
						{Name: "John", Age: 30},
					},
				}
				b := ComplexStruct{
					ID:   1,
					Tags: []string{"important", "urgent"},
					Metadata: map[string]string{
						"department": "engineering",
						"project":    "alpha",
					},
					People: []Person{
						{Name: "John", Age: 30},
					},
				}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "complex structs different slice length",
			testFunc: func() bool {
				a := ComplexStruct{
					Tags: []string{"important", "urgent"},
				}
				b := ComplexStruct{
					Tags: []string{"important"},
				}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},
		{
			name: "complex structs different map values",
			testFunc: func() bool {
				a := ComplexStruct{
					Metadata: map[string]string{"key": "value1"},
				}
				b := ComplexStruct{
					Metadata: map[string]string{"key": "value2"},
				}
				return EqualExportedFields(a, b)
			},
			expected: false,
		},
		{
			name: "events with same time",
			testFunc: func() bool {
				timestamp := time.Date(2023, 6, 15, 10, 0, 0, 0, time.UTC)
				endTime := time.Date(2023, 6, 15, 11, 0, 0, 0, time.UTC)
				a := Event{
					Name:      "Meeting",
					Timestamp: timestamp,
					EndTime:   &endTime,
				}
				b := Event{
					Name:      "Meeting",
					Timestamp: timestamp,
					EndTime:   &endTime,
				}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "events with different time",
			testFunc: func() bool {
				a := Event{
					Name:      "Meeting",
					Timestamp: time.Date(2023, 6, 15, 10, 0, 0, 0, time.UTC),
				}
				b := Event{
					Name:      "Meeting",
					Timestamp: time.Date(2023, 6, 15, 11, 0, 0, 0, time.UTC), // Different hour
				}
				return EqualExportedFields(a, b)
			},
			expected: true,
		},
		{
			name: "string comparison equal",
			testFunc: func() bool {
				return EqualExportedFields("hello", "hello")
			},
			expected: true,
		},
		{
			name: "string comparison different",
			testFunc: func() bool {
				return EqualExportedFields("hello", "world")
			},
			expected: false,
		},
		{
			name: "int comparison equal",
			testFunc: func() bool {
				return EqualExportedFields(42, 42)
			},
			expected: true,
		},
		{
			name: "slice comparison equal",
			testFunc: func() bool {
				return EqualExportedFields([]int{1, 2, 3}, []int{1, 2, 3})
			},
			expected: true,
		},
		{
			name: "slice comparison different",
			testFunc: func() bool {
				return EqualExportedFields([]int{1, 2, 3}, []int{1, 2, 4})
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.testFunc()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkEqualExportedFields(b *testing.B) {
	type BenchStruct struct {
		Name    string
		Age     int
		Email   string
		Active  bool
		private string //nolint:unused
	}

	type NestedBenchStruct struct {
		ID     int
		User   BenchStruct
		Tags   []string
		Meta   map[string]string
		Nested *BenchStruct
	}

	user := BenchStruct{
		Name:    "John Doe",
		Age:     30,
		Email:   "john@example.com",
		Active:  true,
		private: "secret",
	}

	nested1 := NestedBenchStruct{
		ID:   1,
		User: user,
		Tags: []string{"admin", "user", "active"},
		Meta: map[string]string{
			"department": "engineering",
			"role":       "senior",
		},
		Nested: &user,
	}

	nested2 := NestedBenchStruct{
		ID:   1,
		User: user,
		Tags: []string{"admin", "user", "active"},
		Meta: map[string]string{
			"department": "engineering",
			"role":       "senior",
		},
		Nested: &user,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualExportedFields(nested1, nested2)
	}
}
