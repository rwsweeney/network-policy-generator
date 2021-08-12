package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generate)

}

var generate = &cobra.Command{
	Use:   "generate",
	Short: "Generates a network policy",
	Long:  `Generates a network policy`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}
