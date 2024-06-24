package runner

import (
	"fmt"
	"os"
)

// Checking version
func CheckVersion() {
	version := fmt.Sprintf(Succfix + "Current crtfinder version: " + g.Bold(g.Green("v1.0")))
	fmt.Fprintln(os.Stderr, version)
}