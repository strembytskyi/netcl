package cmd

import (
	"errors"
	"fmt"
	"net"
	"netcl/configs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipAddresses)
}

var configIP = configs.GetIPConfig()

var ipAddresses = &cobra.Command{
	Use:   configIP.Name,
	Short: configIP.Short,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No host entered")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getIPAddresses(args[0])
	},
}

func getIPAddresses(host string) error {
	ip, err := net.LookupIP(host)
	if err != nil {
		return err
	}

	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}
	return nil
}
