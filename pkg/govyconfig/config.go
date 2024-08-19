package govyconfig

import (
	"fmt"
	"log/slog"
	"reflect"
	"strings"
	"sync"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/logging"
)

var (
	inferredNames                  = make(map[string]InferredName)
	nameInferFunc    NameInferFunc = NameInferDefaultRule
	nameInferMode                  = NameInferModeDisable
	includeTestFiles               = false

	mu sync.RWMutex
)

// InferredName represents an inferred property name.
type InferredName struct {
	// Name is the inferred property name.
	Name string
	// File is the relative path to the file where the [govy.PropertyRules.For] is detected.
	File string
	// Line is the line number in the File where the [govy.PropertyRules.For] is detected.
	Line int
}

func (g InferredName) key() string {
	return getterLocationKey(g.File, g.Line)
}

func getterLocationKey(file string, line int) string {
	file = strings.TrimPrefix(file, internal.FindModuleRoot()+"/")
	return fmt.Sprintf("%s:%d", file, line)
}

// SetInferredName sets the inferred property name for the given file and line.
// Once it's registered it can be retrieved using [GetInferredName].
// It is primarily exported for code generation utility of govy which runs in NameInferModeGenerate.
func SetInferredName(loc InferredName) {
	mu.Lock()
	inferredNames[loc.key()] = loc
	mu.Unlock()
}

// GetInferredName returns the inferred property name for the given file and line.
// The name has to be first set using [SetInferredName].
// It is primarily exported for [govy] to utilize when NameInferModeGenerate mode is set.
func GetInferredName(file string, line int) string {
	mu.RLock()
	defer mu.RUnlock()
	name, ok := inferredNames[getterLocationKey(file, line)]
	if !ok {
		slog.Debug("")
	}
	return name.Name
}

// SetLogLevel sets the logging level for [slog.Logger] used by [govy].
// It's safe to call this function concurrently.
func SetLogLevel(level slog.Level) {
	logging.SetLogLevel(level)
}

// NameInferMode defines a mode of property names' inference.
type NameInferMode int

const (
	// NameInferModeDisable disables property names' inference.
	// It is the default mode.
	NameInferModeDisable NameInferMode = iota
	// NameInferModeRuntime infers property names' during runtime,
	// whenever For, ForSlice, ForPointer or ForMap are created.
	// If you're not reusing these [govy.PropertyRules], but rather creating them dynamically,
	// beware of significant performance cost of the inference mechanism.
	NameInferModeRuntime
	// NameInferModeGenerate does the heavy lifting of inferring property names
	// in a separate step which involves code generation.
	// When creating new [govy.PropertyRules], the only performance hit is due to the
	// usage of [runtime] package which helps us get the caller frame details.
	NameInferModeGenerate
)

// SetNameInferMode sets the mode of property names' inference.
// It overrides the default mode [NameInferModeDisable].
// It's safe to call this function concurrently.
func SetNameInferMode(mode NameInferMode) {
	mu.Lock()
	nameInferMode = mode
	mu.Unlock()
}

func GetNameInferMode() NameInferMode {
	mu.RLock()
	defer mu.RUnlock()
	return nameInferMode
}

// SetNameInferFunc sets the rule for inferring field names from struct tags.
// It overrides the default rule [NameInferDefaultRule].
// It's safe to call this function concurrently.
func SetNameInferFunc(rule NameInferFunc) {
	mu.Lock()
	nameInferFunc = rule
	mu.Unlock()
}

func GetNameInferFunc() NameInferFunc {
	mu.RLock()
	defer mu.RUnlock()
	return nameInferFunc
}

// NameInferFunc is a function blueprint for inferring property names.
// It is only called for struct fields.
// Tag value is the raw value of the struct tag, it needs to be parsed with [reflect.StructTag].
type NameInferFunc func(fieldName, tagValue string) string

// NameInferDefaultRule is the default rule for inferring field names from struct tags,
// it looks for json and yaml tags, preferring json if both are set.
func NameInferDefaultRule(fieldName, tagValue string) string {
	for _, tagKey := range []string{"json", "yaml"} {
		tagValues := strings.Split(
			reflect.StructTag(strings.Trim(tagValue, "`")).Get(tagKey),
			",",
		)
		if len(tagValues) > 0 && tagValues[0] != "" {
			fieldName = tagValues[0]
			break
		}
	}
	return fieldName
}

// SetNameInferIncludeTestFiles sets whether to include test files in name inference mechanism.
func SetNameInferIncludeTestFiles(inc bool) {
	mu.Lock()
	includeTestFiles = inc
	mu.Unlock()
}

// GetNameInferIncludeTestFiles returns whether to include test files in name inference mechanism.
func GetNameInferIncludeTestFiles() bool {
	mu.RLock()
	defer mu.RUnlock()
	return includeTestFiles
}
