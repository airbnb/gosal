package cmd

import (
	"fmt"
	"os"

	"github.com/airbnb/gosal/sal"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gosal",
	Short: "gosal uploads machine details to sal",
	Long: `Gosal is intended to be a multi platform client for sal.

Complete documentation is available at https://github.com/airbnb/gosal/.`,
	Run: func(cmd *cobra.Command, args []string) {
		sal.SendCheckin()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
