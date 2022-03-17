package cmd

import "github.com/spf13/cobra"

var logCmd = &cobra.Command{
	Use:   "logs",
	Short: "To Get the logs based on Network or System",
	Long:  `To Get the logs based on Network(Hubble Relay) and System(Kubearmor Relay).`,
}

func init() {
	rootCmd.AddCommand(logCmd)
}
