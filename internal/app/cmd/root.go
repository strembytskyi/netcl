package cmd

import (
	"fmt"
	"netcl/configs"
	"os"

	"github.com/spf13/cobra"
)

var config = configs.GetRootConfig()

var rootCmd = &cobra.Command{
	Use:          config.Name,
	Short:        config.Short,
	SilenceUsage: config.Silence,
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
