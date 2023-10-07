package main

import (
	"github.com/launchboxio/bentos/internal/agent"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start a BentOS agent",
	Run: func(cmd *cobra.Command, args []string) {
		agent.Run()
	},
}

func init() {
	agentCmd.Flags().String("driver", "docker", "The runtime adapter to use for managing instances")
	rootCmd.AddCommand(agentCmd)
}
