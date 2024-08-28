package govy

import (
	"fmt"
	"log/slog"

	"github.com/nobl9/govy/internal/nameinfer"
	"github.com/nobl9/govy/pkg/govyconfig"
)

func inferName() string {
	return inferNameWithMode(govyconfig.GetNameInferMode())
}

func inferNameWithMode(mode govyconfig.NameInferMode) string {
	switch mode {
	case govyconfig.NameInferModeDisable:
		return ""
	case govyconfig.NameInferModeGenerate:
		file, line := nameinfer.Frame(5)
		return govyconfig.GetInferredName(file, line)
	case govyconfig.NameInferModeRuntime:
		file, line := nameinfer.Frame(5)
		return nameinfer.InferName(file, line)
	default:
		slog.Error(fmt.Sprintf("unknown %T", mode))
		return ""
	}
}
