package cmd

import (
	"os"

	"github.com/ljcbaby/HDU-network-checker/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "checker",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Logger.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			log.Init(zap.DebugLevel)
		} else {
			log.Init(zap.InfoLevel)
		}
	}
}
