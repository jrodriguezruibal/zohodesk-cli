package main

import (
	"os"

	"github.com/jrodriguezruibal/zohodesk-cli/cmd"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	cmd.SetVersion(version, buildTime)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}