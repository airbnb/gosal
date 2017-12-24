package cmd

import (
	"github.com/airbnb/gosal/version"
	"github.com/spf13/cobra"
)

var (
	fFull bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gosal",
	Long:  `Print the version number and build information of gosal`,
	Run: func(cmd *cobra.Command, args []string) {
		if fFull {
			version.PrintFull()
			return
		}
		version.Print()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVar(&fFull, "full", false, "print full version information")
}
