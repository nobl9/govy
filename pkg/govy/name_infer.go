package govy

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/internal/nameinfer"
	"github.com/nobl9/govy/pkg/govyconfig"
)

// getInferredNameFunc is a closure which returns a [nameFunc].
// It captures the inferred name once and caches it.
// It is safe to call this function concurrently.
func getInferredNameFunc(callers int, pc []uintptr) nameFunc {
	var (
		once sync.Once
		name string
	)
	return func() string {
		once.Do(func() {
			if callers < 1 {
				return
			}
			frame, _ := runtime.CallersFrames(pc).Next()
			mode := govyconfig.GetNameInferMode()
			switch mode {
			case govyconfig.NameInferModeGenerate:
				name = govyconfig.GetInferredName(frame.File, frame.Line)
			case govyconfig.NameInferModeRuntime:
				name = nameinfer.InferName(frame.File, frame.Line)
			case govyconfig.NameInferModeDisable:
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
