package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.1.7"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version number of ghpm, then exits.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("'%s' (linux/amd64, linux/x86_64, linux/arm64, darwin/arm64, darwin/amd64)\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
