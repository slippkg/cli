package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dryRun bool

var syncCmd = &cobra.Command{
	Use:		"sync",
	Short:	"Synchronize system state with slip.lua",
	Run:		func(cmd *cobra.Command, args []string) {
		if dryRun {
			fmt.Println("Dry run: showing changes if applied")
			// TODO: implement dry-run diff logic
		} else {
			fmt.Println()
			// TODO: enforce stat from slip.lua
		}
	},
}

func init() {
	syncCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would change without applying")
	rootCmd.AddCommand(syncCmd)
}
