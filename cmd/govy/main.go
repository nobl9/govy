package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/nobl9/govy/pkg/govyconfig"
)

const (
	govyCmdName      = "govy"
	nameInferCmdName = "nameinfer"
)

var subcommands = []string{
	nameInferCmdName,
}

func main() {
	govyconfig.SetLogLevel(slog.LevelDebug)

	rootCmd := flag.NewFlagSet(govyCmdName, flag.ExitOnError)
	rootCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", govyCmdName)
		fmt.Fprintf(os.Stderr, "  %s <subcommand> [flags]\n", govyCmdName)
		fmt.Fprintf(os.Stderr, "Subcommands:\n")
		for _, cmd := range subcommands {
			fmt.Fprintf(os.Stderr, "  %s\n", cmd)
		}
	}

	if len(os.Args) < 2 {
		rootCmd.Usage()
		os.Exit(1)
	}

	var cmd interface{ Run() error }
	switch os.Args[1] {
	case nameInferCmdName:
		cmd = newNameInferCommand()
	default:
		errFatalWithUsage(
			rootCmd,
			"'%s' is not a valid subcommand, try: %s",
			os.Args[1],
			strings.Join(subcommands, ", "),
		)
		return
	}
	if err := cmd.Run(); err != nil {
		errFatal(err.Error())
	}
}

func errFatalWithUsage(cmd *flag.FlagSet, f string, a ...any) {
	f = "Error: " + f
	if len(a) == 0 {
		fmt.Fprintln(os.Stderr, f)
	} else {
		fmt.Fprintf(os.Stderr, f+"\n", a...)
	}
	cmd.Usage()
	os.Exit(1)
}

func errFatal(f string, a ...any) {
	f = "Error: " + f
	if len(a) == 0 {
		fmt.Fprintln(os.Stderr, f)
	} else {
		fmt.Fprintf(os.Stderr, f+"\n", a...)
	}
	os.Exit(1)
}
