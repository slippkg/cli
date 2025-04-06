package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:		"slip",
	Short: 	"slip is a stateless system config enforcer",
	Long:		`slip enforces system state based on a declarative config file (slip.lua).`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
