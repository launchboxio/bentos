package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bentos",
	Short: "BentOS, for self-hosted clouds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to BentOS!")
	},
}

func init() {

}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
