package govyconfig

import (
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/pkg/jsonpath"
)

var (
	includeTestFiles = false
	inferredPaths    = make(map[string]InferredPath)
	mu               sync.RWMutex
)

// InferredPath represents an inferred relative property path for a specific getter call site.
type InferredPath struct {
	// Path is the inferred relative property path fragment.
	// It does not include a leading `$` segment.
	Path jsonpath.Path
	// File is the path, relative to the module root, to the file where a govy property
	// constructor call was detected.
	File string
	// Line is the line number in File where that constructor call was detected.
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

// SetInferredPath sets the inferred relative property path for the given file and line.
// Once it's registered it can be retrieved using [GetInferredPath].
// It is primarily exported for code generation utility of govy which runs in [InferPathModeGenerate].
func SetInferredPath(loc InferredPath) {
	mu.Lock()
	inferredPaths[loc.key()] = loc
	mu.Unlock()
}

// GetInferredPath returns the inferred relative property path for the given file and line.
// The path has to be first set using [SetInferredPath].
// It is primarily exported for govy to utilize when [InferPathModeGenerate] mode is set.
func GetInferredPath(file string, line int) jsonpath.Path {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := inferredPaths[getterLocationKey(file, line)]
	if !ok {
		logging.Logger().Error(
			"inferred path was not found",
			slog.String("file", file),
			slog.Int("line", line),
		)
		return jsonpath.Path{}
	}
	return p.Path
}

// SetInferPathIncludeTestFiles sets whether test files participate in path inference.
func SetInferPathIncludeTestFiles(inc bool) {
	mu.Lock()
	includeTestFiles = inc
	mu.Unlock()
}

// GetInferPathIncludeTestFiles returns whether test files participate in path inference.
func GetInferPathIncludeTestFiles() bool {
	mu.RLock()
	defer mu.RUnlock()
	return includeTestFiles
}

func getterLocationKey(file string, line int) string {
	file = strings.TrimPrefix(file, internal.FindModuleRoot()+"/")
	return fmt.Sprintf("%s:%d", file, line)
}
