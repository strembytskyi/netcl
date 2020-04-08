package cmd

import (
	"errors"
	"fmt"
	"net"
	"netcl/configs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nameServer)
}

var configNs = configs.GetNSConfig()

var nameServer = &cobra.Command{
	Use:   configNs.Name,
	Short: configNs.Short,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No host entered")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getNameServers(args[0])
	},
}

func getNameServers(host string) error {
	ns, err := net.LookupNS(host)
	if err != nil {
		return err
	}

	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}
	return nil
}
