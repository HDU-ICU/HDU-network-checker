package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "-dev"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of checker",
	Long:  `All software has versions. This is checker's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("checker v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
