package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var rootCmd = &cobra.Command{
	Use:   "netpol",
	Short: "Interacts with network policies",
	Long:  `Interacts with network policies`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Error("You must specify a sub-command.")
		err := cmd.Help()
		if err != nil {
			klog.Error(err)
		}
		os.Exit(1)
	},
}

// Execute the stuff
func Execute(VERSION string, COMMIT string) {
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
