package main

import (
	"github.com/launchboxio/bentos/internal/api"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the BentOS server component",
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		server, err := api.New()
		if err != nil {
			panic(err)
		}
		err = server.Run(address)
	},
}

func init() {
	serverCmd.Flags().String("address", ":8080", "The server address")
	rootCmd.AddCommand(serverCmd)
}
