package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version bool
)

var rootCmd = &cobra.Command{
	Use:   "ghpm",
	Short: "ghpm is tool to manage privacy on github.",
	Args:  cobra.ExactArgs(0),
	Long:  `ghpm is tool to manage privacy on github. And quickly switch all repositories to private.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if version {
			versionCmd.Run(nil, nil)
			return nil
		}

		fmt.Println("Made with 💞 love 💞 for developers by a developer ❤️")

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "prints the version")
}
