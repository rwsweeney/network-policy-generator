package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkConfig)

}

var checkConfig = &cobra.Command{
	Use:   "check-config",
	Short: "Checks the config file.",
	Long:  `Checks the config file to see if it is in N'sync with everything else.`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}
