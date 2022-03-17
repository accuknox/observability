package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/accuknox/observability/src/get"
	"github.com/accuknox/observability/src/types"
	"github.com/spf13/cobra"
)

var networkOption types.CiliumFilter

//kaCmd represents the kubearmor command
var nwCmd = &cobra.Command{
	Use:   "network",
	Short: "Network commends for get aggregated cilium logs",
	Long:  `Network commands give all the aggregated cilium logs for observability. `,
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs
		logs, err := get.GetCilium(networkOption, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			output, _ := json.Marshal(log)
			fmt.Println(string(output))
		}
		return nil
	},
}

func init() {
	logCmd.AddCommand(nwCmd)
	nwCmd.Flags().IntVarP(&limit, "limit", "l", 0, "fetch limited value")
	nwCmd.Flags().StringVarP(&networkOption.Type, "type", "t", "", "fetch by type <L3_L4> or <L7>")
	nwCmd.Flags().StringVarP(&networkOption.Verdict, "verdict", "v", "", "fetch by verdict <FORWARDED> or <DROPPED> or <ERROR> or <AUDIT>")
	nwCmd.Flags().StringVarP(&networkOption.Direction, "direction", "d", "", "fetch by traffic direction <INGRESS> or <EGRESS>")
}
