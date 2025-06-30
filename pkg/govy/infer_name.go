package govy

import (
	"fmt"

	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/internal/nameinfer"
	"github.com/nobl9/govy/pkg/govyconfig"
)

func inferName(skipFrames int) string {
	return inferNameWithMode(govyconfig.GetNameInferMode(), skipFrames)
}

func inferNameWithMode(mode govyconfig.NameInferMode, skipFrames int) string {
	switch mode {
	case govyconfig.NameInferModeDisable:
		return ""
	case govyconfig.NameInferModeGenerate:
		file, line := nameinfer.Frame(skipFrames)
		return govyconfig.GetInferredName(file, line)
	case govyconfig.NameInferModeRuntime:
		file, line := nameinfer.Frame(skipFrames)
		return nameinfer.InferName(file, line)
	default:
		logging.Logger().Error(fmt.Sprintf("unknown %T", mode))
		return ""
	}
}
