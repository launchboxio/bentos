package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start a BentOS agent",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the BentOS agent")
	},
}

func init() {
	agentCmd.Flags().String("driver", "docker", "The runtime adapter to use for managing instances")
	rootCmd.AddCommand(agentCmd)
}
