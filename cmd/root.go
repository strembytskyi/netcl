package cmd

import (
	"fmt"
	"netcl/configs"
	"os"

	"github.com/spf13/cobra"
)

var configRoot = configs.GetRootConfig()

var rootCmd = &cobra.Command{
	Use:          configRoot.Name,
	Short:        configRoot.Short,
	SilenceUsage: configRoot.Silence,
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
