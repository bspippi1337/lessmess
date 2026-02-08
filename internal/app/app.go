package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/bspippi1337/restless/internal/gui"
	"github.com/bspippi1337/restless/internal/tui"
)

func Run(args []string, stdin, stdout, stderr *os.File) int {
	fs := flag.NewFlagSet("restless", flag.ContinueOnError)
	fs.SetOutput(stderr)

	var mode string
	fs.StringVar(&mode, "mode", "", "Mode: auto|tui|gui (default: auto)")
	fs.StringVar(&mode, "m", "", "Alias for --mode")

	var quiet bool
	fs.BoolVar(&quiet, "quiet", false, "Disable animations and extra flair")

	var version bool
	fs.BoolVar(&version, "version", false, "Print version and exit")

	if err := fs.Parse(args); err != nil {
		return 2
	}
	if version {
		fmt.Fprintln(stdout, "restless v0.0.0-alpha (skeleton)")
		return 0
	}

	// Auto mode: if stdout is a TTY => TUI, else GUI (for now, GUI wraps the same TUI).
	if mode == "" || mode == "auto" {
		if isTTY(stdout) {
			mode = "tui"
		} else {
			mode = "gui"
		}
	}

	switch mode {
	case "tui":
		if err := tui.Run(stdin, stdout, quiet); err != nil {
			fmt.Fprintln(stderr, "restless:", err)
			return 1
		}
		return 0
	case "gui":
		if err := gui.Run(stdin, stdout, stderr, quiet); err != nil {
			fmt.Fprintln(stderr, "restless:", err)
			return 1
		}
		return 0
	default:
		fmt.Fprintf(stderr, "Unknown mode: %s\n", mode)
		return 2
	}
}

func isTTY(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}
