package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print the version and build information for zohodesk-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("zohodesk-cli %s\n", cfgVersion)
		fmt.Printf("  Build time: %s\n", cfgBuildTime)
		fmt.Printf("  Go version: %s\n", runtime.Version())
		fmt.Printf("  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}