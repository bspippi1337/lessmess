package gui

import (
	"fmt"
	"os"

	"github.com/bspippi1337/restless/internal/tui"
)

// Run is a placeholder for the "native wrapper around the TUI".
// In v1 it will spawn a real OS window with a terminal-like surface.
// For now: we print a short note (only if not quiet) and run the same TUI.
func Run(stdin, stdout, stderr *os.File, quiet bool) error {
	if !quiet {
		fmt.Fprintln(stderr, "[gui] skeleton: running TUI core. native window wrapper comes next.")
	}
	return tui.Run(stdin, stdout, quiet)
}
