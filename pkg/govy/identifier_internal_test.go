package govy

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestCollectionPropertyRulesInternalRulesHaveNoIDs(t *testing.T) {
	slice := ForSlice(func(struct{}) []string { return nil })
	sliceID := slice.GetID()
	assert.True(t, sliceID != "")
	slice = slice.
		RulesForEach(NewRule(func(string) error { return nil })).
		IncludeForEach(New(For(GetSelf[string]()))).
		Cascade(CascadeModeStop)
	assert.True(t, slice.GetID() != sliceID)
	assert.Equal(t, "", slice.forEachRules.GetID())

	mapRules := ForMap(func(struct{}) map[string]string { return nil })
	mapID := mapRules.GetID()
	assert.True(t, mapID != "")
	mapRules = mapRules.
		RulesForKeys(NewRule(func(string) error { return nil })).
		RulesForValues(NewRule(func(string) error { return nil })).
		RulesForItems(NewRule(func(MapItem[string, string]) error { return nil })).
		IncludeForKeys(New(For(GetSelf[string]()))).
		IncludeForValues(New(For(GetSelf[string]()))).
		IncludeForItems(New(For(GetSelf[MapItem[string, string]]()))).
		Cascade(CascadeModeStop)
	assert.True(t, mapRules.GetID() != mapID)
	for name, id := range map[string]string{
		"key":   mapRules.forKeyRules.GetID(),
		"value": mapRules.forValueRules.GetID(),
		"item":  mapRules.forItemRules.GetID(),
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, "", id)
		})
	}
}
