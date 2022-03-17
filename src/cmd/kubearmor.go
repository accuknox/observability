package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/accuknox/observability/src/get"
	"github.com/accuknox/observability/src/kubearmor/all"
	"github.com/accuknox/observability/src/kubearmor/container"
	"github.com/accuknox/observability/src/kubearmor/host"
	"github.com/accuknox/observability/src/kubearmor/namespace"
	"github.com/accuknox/observability/src/kubearmor/pod"
	"github.com/accuknox/observability/src/types"
	"github.com/spf13/cobra"
)

var limit int

var option types.KubeArmorFilter

//kaCmd represents the kubearmor command
var kaCmd = &cobra.Command{
	Use:   "system",
	Short: "System commends for get aggregated kubearmor logs",
	Long:  `System commands give all the aggregated kubearmor logs for observability. `,
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs
		logs, err := get.GetKubearmor(option, limit)
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

//kaAllCmd represents the kubearmor command for fetch all logs
var kaAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "Get all the aggregated logs",
	Long:    `Get all the aggregated logs`,
	Example: `knox kubearmor all`,
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs
		logs, err := all.All(option, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println("Data : ", log)
		}

		return nil
	},
}

//kaHostCmd represents the kubearmor command for fetch all logs based on host name(s)
var kaHostCmd = &cobra.Command{
	Use:     "host",
	Short:   "Get the logs based on host-name",
	Long:    `Get the aggregated logs based on host-name`,
	Example: "knox kubearmor host <host-name-1> <host-name-2>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a host-name as argument")
		}
		return nil
	},
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs based on host name(s)
		logs, err := host.FilterByHostName(args, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println("Data : ", log)

		}

		return nil
	},
}

//kaNamespaceCmd represents the kubearmor command for fetch all logs based on namespace(s)
var kaNamespaceCmd = &cobra.Command{
	Use:     "namespace",
	Short:   "Get the logs based on namespace",
	Long:    `Get the aggregated logs based on namespace`,
	Example: "knox kubearmor namespace <namespace-1> <namespace-2>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a namespace-name as argument")
		}
		return nil
	},
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs based on namespace(s)
		logs, err := namespace.FilterByNamespace(args, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println("Data : ", log)
		}
		return nil
	},
}

//kaPodCmd represents the kubearmor command for fetch all logs based on pod(s)
var kaPodCmd = &cobra.Command{
	Use:     "pod",
	Short:   "Get the logs based on pod",
	Long:    `Get the aggregated logs based on pod`,
	Example: "knox kubearmor pod <pod-1> <pod-2>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a pod-name as argument")
		}
		return nil
	},
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs based on pod(s)
		logs, err := pod.FilterByPod(args, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println("Data : ", log)
		}
		return nil
	},
}

//kaContainerIDCmd represents the kubearmor command for fetch all logs based on container ID(s)
var kaContainerIDCmd = &cobra.Command{
	Use:     "container-id",
	Short:   "Get the logs based on container-id",
	Long:    `Get the aggregated logs based on container-id`,
	Example: "knox kubearmor container-id <container-id-1> <container-id-2>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a container-id as argument")
		}
		return nil
	},
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs based on container-id(s)
		logs, err := container.FilterByContainerID(args, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println("Data : ", log)
		}
		return nil
	},
}

//kaContainerNameCmd represents the kubearmor command for fetch all logs based on container Name(s)
var kaContainerNameCmd = &cobra.Command{
	Use:     "container-name",
	Short:   "Get the logs based on container-name",
	Long:    `Get the aggregated logs based on container-name`,
	Example: "knox kubearmor container-name <container-name-1> <container-name-2>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a container-name as argument")
		}
		return nil
	},
	RunE: func(cm *cobra.Command, args []string) error {
		//fetch all the logs based on container name(s)
		logs, err := container.FilterByContainerName(args, limit)
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println("Data : ", log)
		}
		return nil
	},
}

func init() {

	logCmd.AddCommand(kaCmd)
	// kaCmd.AddCommand(kaAllCmd)
	// kaCmd.AddCommand(kaHostCmd)
	// kaCmd.AddCommand(kaNamespaceCmd)
	// kaCmd.AddCommand(kaPodCmd)
	// kaCmd.AddCommand(kaContainerIDCmd)
	// kaCmd.AddCommand(kaContainerNameCmd)

	kaCmd.Flags().IntVarP(&limit, "limit", "l", 0, "fetch limited value")
	kaCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "", "fetch by namespace")
	// kaAllCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Fetch Limited Value")
	// kaHostCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Fetch Limited Value")
	// kaNamespaceCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Fetch Limited Value")
	// kaPodCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Fetch Limited Value")
	// kaContainerIDCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Fetch Limited Value")
	// kaContainerNameCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Fetch Limited Value")
	// kaAllCmd.Flags().StringArrayVarP(&option.Operation, "operation", "o", nil, "Filter by Operation")
}
