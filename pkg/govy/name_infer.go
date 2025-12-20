package govy

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/nobl9/govy/internal/infername"
	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/pkg/govyconfig"
)

// InferNameMode defines a mode of property name's inference.
type InferNameMode int

const (
	// InferNameModeDisable disables property name's inference.
	// It is the default mode.
	InferNameModeDisable InferNameMode = iota
	// InferNameModeRuntime infers property names' during runtime,
	// whenever For, ForSlice, ForPointer or ForMap are created.
	// If you're not reusing these [govy.PropertyRules], but rather creating them dynamically,
	// beware of significant performance cost of the inference mechanism.
	InferNameModeRuntime
	// InferNameModeGenerate does the heavy lifting of inferring property names
	// in a separate step which involves code generation.
	// When creating new [govy.PropertyRules], the only performance hit is due to the
	// usage of [runtime] package which helps us get the caller frame details.
	InferNameModeGenerate
)

// InferNameFunc is a function blueprint for inferring property names.
// It is only called for struct fields.
// Tag value is the raw value of the struct tag, it needs to be parsed with [reflect.StructTag].
type InferNameFunc func(fieldName, tagValue string) string

// InferNameDefaultFunc is the default function for inferring field names from struct tags.
// It looks for json and yaml tags, preferring json if both are set.
func InferNameDefaultFunc(fieldName, tagValue string) string {
	return infername.InferNameDefaultFunc(fieldName, tagValue)
}

type internalInferNameFunc func(mode InferNameMode) string

// getInferNameFunc is a closure which returns an [internalInferNameFunc].
// It captures the inferred name once and caches it.
// It is safe to call this function concurrently.
func getInferNameFunc(callers int, pc []uintptr) internalInferNameFunc {
	var (
		once sync.Once
		name string
	)
	return func(mode InferNameMode) string {
		once.Do(func() {
			if callers < 1 {
				return
			}
			frame, _ := runtime.CallersFrames(pc).Next()
			switch mode {
			case InferNameModeGenerate:
				name = govyconfig.GetInferredName(frame.File, frame.Line)
			case InferNameModeRuntime:
				name = infername.InferName(frame.File, frame.Line)
			case InferNameModeDisable:
			default:
				logging.Logger().Error(fmt.Sprintf("unknown %T", mode))
			}
		})
		return name
	}
}

// getCallersAndProgramCounter returns number of callers and program counters
// of function invocations on the calling goroutine's stack.
// Its results are intended to be passed directly to [getInferNameFunc].
func getCallersAndProgramCounter(skipFrames int) (callers int, pc []uintptr) {
	pc = make([]uintptr, 1)
	callers = runtime.Callers(skipFrames, pc)
	return callers, pc
}
