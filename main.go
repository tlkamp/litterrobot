package main

import (
	"fmt"

	"github.com/tlkamp/litterrobot/internal/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	versionString := fmt.Sprintf("%s - %s - %s", version, commit, date)
	cmd.Execute(versionString) //nolint:errcheck
}
