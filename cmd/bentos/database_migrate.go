package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var databaseMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Execute database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running database migrations")
	},
}

func init() {
	databaseCmd.AddCommand(databaseMigrateCmd)
}
