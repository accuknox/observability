package types

//KubeArmor - Structure for KubeArmor Logs Flow
type KubeArmor struct {
	ClusterName   string `json:"cluster_name,omitempty"`
	HostName      string `json:"host_name,omitempty"`
	NamespaceName string `json:"namespace_name,omitempty"`
	PodName       string `json:"pod_name,omitempty"`
	ContainerID   string `json:"container_id,omitempty"`
	ContainerName string `json:"container_name,omitempty"`
	UID           int32  `json:"uid,omitempty"`
	Type          string `json:"type,omitempty"`
	Source        string `json:"source,omitempty"`
	Operation     string `json:"operation,omitempty"`
	Resource      string `json:"resource,omitempty"`
	Data          string `json:"data,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	UpdatedTime   int64  `json:"updated_time,omitempty"`
	Result        string `json:"result,omitempty"`
	Total         int64  `json:"total,omitempty"`
}

//Cilium - Structure for Hubble Log Flow
type Cilium struct {
	Verdict                     int32  `json:"verdict,omitempty"`
	EthernetSource              string `json:"ethernet_source,omitempty"`
	EthernetDestination         string `json:"ethernet_destination,omitempty"`
	IpSource                    string `json:"ip_source,omitempty"`
	IpDestination               string `json:"ip_destination,omitempty"`
	IpVersion                   int32  `json:"ip_version,omitempty"`
	IpEncrypted                 bool   `json:"ip_encrypted,omitempty"`
	L4TCPSourcePort             uint32 `json:"l4_tcp_source_port,omitempty"`
	L4TCPDestinationPort        uint32 `json:"l4_tcp_destination_port,omitempty"`
	L4UDPSourcePort             uint32 `json:"l4_udp_source_port,omitempty"`
	L4UDPDestinationPort        uint32 `json:"l4_udp_destination_port,omitempty"`
	L4ICMPv4Type                uint32 `json:"l4_icmpv4_type,omitempty"`
	L4ICMPv4Code                uint32 `json:"l4_icmpv4_code,omitempty"`
	L4ICMPv6Type                uint32 `json:"l4_icmpv6_type,omitempty"`
	L4ICMPv6Code                uint32 `json:"l4_icmpv6_code,omitempty"`
	SourceID                    uint32 `json:"source_id,omitempty"`
	SourceIdentity              uint32 `json:"source_identity,omitempty"`
	SourceNamespace             string `json:"source_namespace,omitempty"`
	SourceLabels                string `json:"source_labels,omitempty"`
	SourcePodName               string `json:"source_pod_name,omitempty"`
	DestinationID               uint32 `json:"destination_id,omitempty"`
	DestinationIdentity         uint32 `json:"destination_identity,omitempty"`
	DestinationNamespace        string `json:"destination_namespace,omitempty"`
	DestinationLabels           string `json:"destination_labels,omitempty"`
	DestinationPodName          string `json:"destination_pod_name,omitempty"`
	Type                        int32  `json:"type,omitempty"`
	NodeName                    string `json:"node_name,omitempty"`
	SourceNames                 string `json:"source_names,omitempty"`
	DestinationNames            string `json:"destination_names,omitempty"`
	L7Type                      int32  `json:"l7_type,omitempty"`
	L7LatencyNs                 uint64 `json:"l7_latency_ns,omitempty"`
	L7DnsQuery                  string `json:"l7_dns_query,omitempty"`
	L7DnsIps                    string `json:"l7_dns_ips,omitempty"`
	L7DnsTtl                    uint32 `json:"l7_dns_ttl,omitempty"`
	L7DnsCnames                 string `json:"l7_dns_cnames,omitempty"`
	L7DnsObservationsource      string `json:"l7_dns_observation_source,omitempty"`
	L7DnsRcode                  uint32 `json:"l7_dns_rcode,omitempty"`
	L7DnsQtypes                 string `json:"l7_dns_qtypes,omitempty"`
	L7DnsRrtypes                string `json:"l7_dns_rrtypes,omitempty"`
	L7HttpCode                  uint32 `json:"l7_http_code,omitempty"`
	L7HttpMethod                string `json:"l7_http_method,omitempty"`
	L7HttpUrl                   string `json:"l7_http_url,omitempty"`
	L7HttpProtocol              string `json:"l7_http_protocol,omitempty"`
	L7HttpHeaders               string `json:"l7_http_headers,omitempty"`
	L7KafkaErrorCode            int32  `json:"l7_kafka_error_code,omitempty"`
	L7KafkaApiVersion           int32  `json:"l7_kafka_api_version,omitempty"`
	L7KafkaApikey               string `json:"l7_kafka_api_key,omitempty"`
	L7KafkaCorrelationID        int32  `json:"l7_kafka_correlation_id,omitempty"`
	L7KafkaTopic                string `json:"l7_kafka_topic,omitempty"`
	EventTypeType               int32  `json:"event_type_type,omitempty"`
	EventTypeSubType            int32  `json:"event_type_sub_type,omitempty"`
	SourceServiceName           string `json:"source_service_name,omitempty"`
	SourceServiceNamespace      string `json:"source_service_namespace,omitempty"`
	DestinationServiceName      string `json:"destination_service_name,omitempty"`
	DestinationServiceNamespace string `json:"destination_service_namespace,omitempty"`
	TrafficDirection            int32  `json:"traffic_direction,omitempty"`
	PolicyMatchType             uint32 `json:"policy_match_type,omitempty"`
	TraceObservationPoint       int32  `json:"trace_observation_point,omitempty"`
	DropReasonDesc              int32  `json:"drop_reason_desc,omitempty"`
	IsReply                     bool   `json:"is_reply,omitempty"`
	DebugCapturePoint           int32  `json:"debug_capture_point,omitempty"`
	InterfaceIndex              uint32 `json:"interface_index,omitempty"`
	InterfaceName               string `json:"interface_name,omitempty"`
	ProxyPort                   uint32 `json:"proxy_port,omitempty"`
	StartTime                   int64  `json:"start_time,omitempty"`
	UpdatedTime                 int64  `json:"updated_time,omitempty"`
	Total                       int64  `json:"total,omitempty"`
}

type KubeArmorFilter struct {
	Operation []string `json:"operation"`
	Namespace string
}

type CiliumFilter struct {
	Type      string `json:"type"`
	Verdict   string `json:"verdict"`
	Direction string `json:"direction"`
}

type NetworkSummary struct {
	DestinationLabels string `json:"destination_labels"`
	TrafficDirection  string `json:"traffic_direction"`
}

type SystemSummery struct {
	Operation string `json:"operation"`
	Source    string `json:"source"`
	Resource  string `json:"resource"`
}
