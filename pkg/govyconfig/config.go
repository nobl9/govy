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
	inferredNames    = make(map[string]InferredName)
	mu               sync.RWMutex
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

// SetLogLevel sets the logging level for [slog.Logger] used by govy.
// It's safe to call this function concurrently.
func SetLogLevel(level slog.Level) {
	logging.SetLogLevel(level)
}

// SetInferredName sets the inferred property name for the given file and line.
// Once it's registered it can be retrieved using [GetInferredName].
// It is primarily exported for code generation utility of govy which runs in [InferNameModeGenerate].
func SetInferredName(loc InferredName) {
	mu.Lock()
	inferredNames[loc.key()] = loc
	mu.Unlock()
}

// GetInferredName returns the inferred property name for the given file and line.
// The name has to be first set using [SetInferredName].
// It is primarily exported for govy to utilize when [InferNameModeGenerate] mode is set.
func GetInferredName(file string, line int) string {
	mu.RLock()
	defer mu.RUnlock()
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

// SetInferNameIncludeTestFiles sets whether to include test files in name inference mechanism.
func SetInferNameIncludeTestFiles(inc bool) {
	mu.Lock()
	includeTestFiles = inc
	mu.Unlock()
}

// GetInferNameIncludeTestFiles returns whether to include test files in name inference mechanism.
func GetInferNameIncludeTestFiles() bool {
	mu.RLock()
	defer mu.RUnlock()
	return includeTestFiles
}

func getterLocationKey(file string, line int) string {
	file = strings.TrimPrefix(file, internal.FindModuleRoot()+"/")
	return fmt.Sprintf("%s:%d", file, line)
}
