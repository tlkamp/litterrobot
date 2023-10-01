package main

import (
	"fmt"
	"os"

	"github.com/tlkamp/litterrobot/internal/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	versionString := fmt.Sprintf("%s - %s - %s", version, commit, date)
	if err := cmd.Execute(versionString); err != nil {
		os.Exit(1)
	}
}
