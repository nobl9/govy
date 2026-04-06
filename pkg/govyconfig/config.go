package govyconfig

import (
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/logging"
)

var (
	includeTestFiles = false
	inferredPaths    = make(map[string]InferredPath)
	mu               sync.RWMutex
)

// InferredPath represents an inferred property path.
type InferredPath struct {
	// Path is the inferred property path.
	Path string
	// File is the relative path to the file where the [govy.PropertyRules.For] is detected.
	File string
	// Line is the line number in the File where the [govy.PropertyRules.For] is detected.
	Line int
}

func (g InferredPath) key() string {
	return getterLocationKey(g.File, g.Line)
}

// SetLogLevel sets the logging level for [slog.Logger] used by govy.
// It's safe to call this function concurrently.
func SetLogLevel(level slog.Level) {
	logging.SetLogLevel(level)
}

// SetInferredPath sets the inferred property path for the given file and line.
// Once it's registered it can be retrieved using [GetInferredPath].
// It is primarily exported for code generation utility of govy which runs in [InferPathModeGenerate].
func SetInferredPath(loc InferredPath) {
	mu.Lock()
	inferredPaths[loc.key()] = loc
	mu.Unlock()
}

// GetInferredPath returns the inferred property path for the given file and line.
// The path has to be first set using [SetInferredPath].
// It is primarily exported for govy to utilize when [InferPathModeGenerate] mode is set.
func GetInferredPath(file string, line int) string {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := inferredPaths[getterLocationKey(file, line)]
	if !ok {
		logging.Logger().Error(
			"inferred path was not found",
			slog.String("file", file),
			slog.Int("line", line),
		)
		return ""
	}
	return p.Path
}

// SetInferPathIncludeTestFiles sets whether to include test files in path inference mechanism.
func SetInferPathIncludeTestFiles(inc bool) {
	mu.Lock()
	includeTestFiles = inc
	mu.Unlock()
}

// GetInferPathIncludeTestFiles returns whether to include test files in path inference mechanism.
func GetInferPathIncludeTestFiles() bool {
	mu.RLock()
	defer mu.RUnlock()
	return includeTestFiles
}

func getterLocationKey(file string, line int) string {
	file = strings.TrimPrefix(file, internal.FindModuleRoot()+"/")
	return fmt.Sprintf("%s:%d", file, line)
}
