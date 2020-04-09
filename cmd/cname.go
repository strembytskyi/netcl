package cmd

import (
	"errors"
	"fmt"
	"net"
	"netcl/configs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cname)
}

var configCname = configs.GetCnameConfig()

var cname = &cobra.Command{
	Use:   configCname.Name,
	Short: configCname.Short,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No host entered")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getCname(args[0])
	},
}

func getCname(host string) error {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return err
	}

	fmt.Println(cname)
	return nil
}
