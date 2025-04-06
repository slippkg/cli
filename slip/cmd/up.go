package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:		"up",
	Short: 	"Check for available updates to installed packages",
	Run:		func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for updates...")
		// TODO: compare installed versions vs latest
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
