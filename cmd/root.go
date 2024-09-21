package cmd

import (
	"os"

	"github.com/ljcbaby/HDU-network-checker/checker"
	"github.com/ljcbaby/HDU-network-checker/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "checker",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger.Sugar().Info("Complete by ljcbaby, Version ", Version)
		checker.BasicCheck()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Logger.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.MousetrapHelpText = ""
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			log.Init(zap.DebugLevel)
			log.Logger.Debug("Verbose output enabled")
		} else {
			log.Init(zap.InfoLevel)
		}
	}
}
