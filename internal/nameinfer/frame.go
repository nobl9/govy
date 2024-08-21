package nameinfer

import "runtime"

// Frame returns the file and line number of the caller.
// It's intended to be used in the context of [govy.For] and similar functions.
func Frame(skipFrames int) (file string, line int) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipFrames, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, frame.Line
}
