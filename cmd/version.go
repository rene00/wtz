package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [command]",
	Short: "show version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(version)
		return nil
	},
}

var version string

// SetVersion sets the global version string so it can be called from main and exposed to cmd.
func SetVersion(v string) {
	version = v
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
