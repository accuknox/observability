package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "knox",
	Short: "A CLI Utility to help manage Observability",
	Long: `CLI Utility to help manage Observability
	
Observability is based on Cilium and Kubearmor Logs. 
Using this we can identify behavior of container, vm, pod, node at the system level.
	  `,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}
