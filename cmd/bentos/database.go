package main

import "github.com/spf13/cobra"

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Interact with the BentOS database",
}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
