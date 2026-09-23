package govy_test

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_NumericMapKeyRules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		run  func(*testing.T) error
		kind string
	}{
		{
			name: "signed inequality", kind: "int",
			run: func(t *testing.T) error {
				v := govy.New(govy.ForMap(govy.GetSelf[map[int]string]()).RulesForKeys(rules.GTE(1)))
				assert.NoError(t, v.Validate(map[int]string{1: "ok"}))
				_, err := govy.JSONSchema(v)
				return err
			},
		},
		{
			name: "unsigned constant", kind: "uint",
			run: func(t *testing.T) error {
				v := govy.New(govy.ForMap(govy.GetSelf[map[uint]string]()).RulesForKeys(rules.EQ(uint(1))))
				assert.NoError(t, v.Validate(map[uint]string{1: "ok"}))
				_, err := govy.JSONSchema(v)
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(
				t,
				tt.run(t),
				`cannot generate JSON Schema for "$.*~" property: map key rules require string keys, got Go kind "`+tt.kind+`"`,
			)
		})
	}
	t.Run("map without key rules", func(t *testing.T) {
		t.Parallel()
		_, err := govy.JSONSchema(govy.New(govy.ForMap(govy.GetSelf[map[int]string]()).
			RulesForValues(rules.StringMinLength(1))))
		assert.NoError(t, err)
	})
}
