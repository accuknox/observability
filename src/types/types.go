package types

//KubeArmor - Structure for KubeArmor Logs Flow
type KubeArmor struct {
	ClusterName   string
	HostName      string
	NamespaceName string
	PodName       string
	ContainerID   string
	ContainerName string
	UID           int32
	Type          string
	Source        string
	Operation     string
	Resource      string
	Data          string
}
