package govyconfig

import (
	"log/slog"
	"sync"

	"github.com/nobl9/govy/internal/logging"
)

var (
	includeTestFiles = false
	mu               sync.RWMutex
)

// SetLogLevel sets the logging level for [slog.Logger] used by govy.
// It's safe to call this function concurrently.
func SetLogLevel(level slog.Level) {
	logging.SetLogLevel(level)
}

// SetNameInferIncludeTestFiles sets whether to include test files in name inference mechanism.
func SetNameInferIncludeTestFiles(inc bool) {
	mu.Lock()
	includeTestFiles = inc
	mu.Unlock()
}

// GetNameInferIncludeTestFiles returns whether to include test files in name inference mechanism.
func GetNameInferIncludeTestFiles() bool {
	mu.RLock()
	defer mu.RUnlock()
	return includeTestFiles
}
