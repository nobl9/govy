package govy

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/nobl9/govy/internal/inferpath"
	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/pkg/govyconfig"
)

// InferPathMode defines a mode of property path inference.
type InferPathMode int

const (
	// InferPathModeDisable disables property path inference.
	// It is the default mode.
	InferPathModeDisable InferPathMode = iota
	// InferPathModeRuntime infers property paths during runtime,
	// whenever For, ForSlice, ForPointer or ForMap are created.
	// If you're not reusing these [govy.PropertyRules], but rather creating them dynamically,
	// beware of significant performance cost of the inference mechanism.
	InferPathModeRuntime
	// InferPathModeGenerate does the heavy lifting of inferring property paths
	// in a separate step which involves code generation.
	// When creating new [govy.PropertyRules], the only performance hit is due to the
	// usage of [runtime] package which helps us get the caller frame details.
	InferPathModeGenerate
)

// InferPathFunc is a function blueprint for inferring property paths.
// It is only called for struct fields.
// Tag value is the raw value of the struct tag, it needs to be parsed with [reflect.StructTag].
type InferPathFunc func(fieldName, tagValue string) string

// InferPathDefaultFunc is the default function for inferring field paths from struct tags.
// It looks for json and yaml tags, preferring json if both are set.
func InferPathDefaultFunc(fieldName, tagValue string) string {
	return inferpath.InferPathDefaultFunc(fieldName, tagValue)
}

type internalInferPathFunc func(mode InferPathMode) Path

// getInferPathFunc is a closure which returns an [internalInferPathFunc].
// It captures the inferred path once and caches it.
// It is safe to call this function concurrently.
func getInferPathFunc(callers int, pc []uintptr) internalInferPathFunc {
	var (
		once sync.Once
		path Path
	)
	return func(mode InferPathMode) Path {
		once.Do(func() {
			if callers < 1 {
				return
			}
			frame, _ := runtime.CallersFrames(pc).Next()
			if frame.File == "" || frame.Line == 0 {
				logging.Logger().Error(
					"invalid frame captured for path inference",
					"file", frame.File,
					"line", frame.Line,
				)
				return
			}
			switch mode {
			case InferPathModeGenerate:
				path = ParsePath(govyconfig.GetInferredPath(frame.File, frame.Line))
			case InferPathModeRuntime:
				path = inferpath.InferPath(frame.File, frame.Line)
			case InferPathModeDisable:
			default:
				logging.Logger().Error(fmt.Sprintf("unknown %T", mode), "mode", int(mode))
			}
		})
		return path
	}
}

// getCallersAndProgramCounter returns number of callers and program counters
// of function invocations on the calling goroutine's stack.
// Its results are intended to be passed directly to [getInferPathFunc].
func getCallersAndProgramCounter(skipFrames int) (callers int, pc []uintptr) {
	pc = make([]uintptr, 1)
	callers = runtime.Callers(skipFrames, pc)
	return callers, pc
}
