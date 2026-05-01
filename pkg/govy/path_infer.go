package govy

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/nobl9/govy/internal/inferpath"
	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/jsonpath"
)

// InferPathMode defines how govy infers relative property paths from getter expressions.
type InferPathMode int

const (
	// InferPathModeDisable disables property path inference.
	// It is the default mode.
	InferPathModeDisable InferPathMode = iota
	// InferPathModeRuntime infers relative property paths from getter expressions at runtime.
	// The inferred path does not include `$`; parent validators and collection rules prepend
	// their own segments when composing nested paths.
	// If you're not reusing these [PropertyRules], but rather creating them dynamically,
	// beware of significant performance cost of the inference mechanism.
	InferPathModeRuntime
	// InferPathModeGenerate does the heavy lifting of inferring relative property paths
	// in a separate step which involves code generation.
	// When creating new [PropertyRules], the only performance hit is due to the
	// usage of [runtime] package which helps us get the caller frame details.
	InferPathModeGenerate
)

type inferPathFunc = func(mode InferPathMode) jsonpath.Path

// getInferPathFunc returns a closure which infers and caches a property path.
// It captures the inferred path once and caches it.
// It is safe to call this function concurrently.
func getInferPathFunc(callers int, pc []uintptr) inferPathFunc {
	var (
		once sync.Once
		path jsonpath.Path
	)
	return func(mode InferPathMode) jsonpath.Path {
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
				path = govyconfig.GetInferredPath(frame.File, frame.Line)
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
