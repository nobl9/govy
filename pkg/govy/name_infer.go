package govy

import (
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"sync"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/internal/nameinfer"
)

var (
	inferredNames = make(map[string]InferredName)
	nameInferMu   sync.RWMutex
)

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

// NameInferFunc is a function blueprint for inferring property names.
// It is only called for struct fields.
// Tag value is the raw value of the struct tag, it needs to be parsed with [reflect.StructTag].
type NameInferFunc func(fieldName, tagValue string) string

// NameInferDefaultFunc is the default function for inferring field names from struct tags.
// It looks for json and yaml tags, preferring json if both are set.
func NameInferDefaultFunc(fieldName, tagValue string) string {
	return nameinfer.NameInferDefaultFunc(fieldName, tagValue)
}

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
// It is primarily exported for code generation utility of govy which runs in [NameInferModeGenerate].
func SetInferredName(loc InferredName) {
	nameInferMu.Lock()
	inferredNames[loc.key()] = loc
	nameInferMu.Unlock()
}

// GetInferredName returns the inferred property name for the given file and line.
// The name has to be first set using [SetInferredName].
// It is primarily exported for govy to utilize when [NameInferModeGenerate] mode is set.
func GetInferredName(file string, line int) string {
	nameInferMu.RLock()
	defer nameInferMu.RUnlock()
	name, ok := inferredNames[getterLocationKey(file, line)]
	if !ok {
		logging.Logger().Error(
			"inferred name was not found",
			slog.String("file", file),
			slog.Int("line", line),
		)
		return ""
	}
	return name.Name
}

type internalNameInferFunc func(mode NameInferMode) string

// getNameInferFunc is a closure which returns an [internalNameInferFunc].
// It captures the inferred name once and caches it.
// It is safe to call this function concurrently.
func getNameInferFunc(callers int, pc []uintptr) internalNameInferFunc {
	var (
		once sync.Once
		name string
	)
	return func(mode NameInferMode) string {
		once.Do(func() {
			if callers < 1 {
				return
			}
			frame, _ := runtime.CallersFrames(pc).Next()
			switch mode {
			case NameInferModeGenerate:
				name = GetInferredName(frame.File, frame.Line)
			case NameInferModeRuntime:
				name = nameinfer.InferName(frame.File, frame.Line)
			case NameInferModeDisable:
			default:
				logging.Logger().Error(fmt.Sprintf("unknown %T", mode))
			}
		})
		return name
	}
}

// getCallersAndProgramCounter returns number of callers and program counters
// of function invocations on the calling goroutine's stack.
// Its results are intended to be passed directly to [getInferredNameFunc].
func getCallersAndProgramCounter(skipFrames int) (callers int, pc []uintptr) {
	pc = make([]uintptr, 1)
	callers = runtime.Callers(skipFrames, pc)
	return callers, pc
}
