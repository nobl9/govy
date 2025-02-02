package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync/atomic"
)

const defaultLogLevel = slog.LevelError

var (
	logger   atomic.Pointer[slog.Logger]
	logLevel *slog.LevelVar
)

func Logger() *slog.Logger {
	return logger.Load()
}

func SetLogLevel(level slog.Level) {
	logLevel.Set(level)
}

func init() {
	logLevel = new(slog.LevelVar)
	if logLevelStr := os.Getenv("GOVY_LOG_LEVEL"); logLevelStr != "" {
		level := new(slog.Level)
		if err := level.UnmarshalText([]byte(logLevelStr)); err != nil {
			fmt.Fprintf(os.Stderr, "invalid log level %q: %v, defaulting to %s\n", logLevelStr, err, logLevel)
		}
		logLevel.Set(*level)
	}
	if logLevel.Level() == 0 {
		logLevel.Set(defaultLogLevel)
	}
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		// We're using our own source handler.
		AddSource: false,
		Level:     logLevel,
	})
	// Source handler should always be the first in the chain
	// in order to keep the number of frames it has to skip consistent.
	handler := sourceHandler{Handler: jsonHandler}
	defaultLogger := slog.New(contextHandler{Handler: handler})
	logger.Swap(defaultLogger)
}

type logContextAttrKey struct{}

// contextHandler is a [slog.Handler] that adds contextual attributes
// to the [slog.Record] before calling the underlying handler.
type contextHandler struct{ slog.Handler }

// Handle adds contextual attributes to the Record before calling the underlying handler.
func (h contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(logContextAttrKey{}).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}
	return h.Handler.Handle(ctx, r)
}

// sourceHandler is a [slog.Handler] that adds [slog.Source] information to the [slog.Record].
type sourceHandler struct{ slog.Handler }

// Handle adds [slog.Source] information to the [slog.Record]
// before calling the underlying handler.
func (h sourceHandler) Handle(ctx context.Context, r slog.Record) error {
	f, ok := runtime.CallersFrames([]uintptr{r.PC}).Next()
	if !ok {
		r.AddAttrs(slog.Attr{
			Key: slog.SourceKey,
			Value: slog.AnyValue(&slog.Source{
				Function: f.Function,
				File:     f.File,
				Line:     f.Line,
			}),
		})
	}
	return h.Handler.Handle(ctx, r)
}
