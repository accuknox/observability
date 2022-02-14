package output

import (
	"fmt"

	"github.com/accuknox/observability/src/types"
)

//NetworkOutput - Output of Network(Cilium) logs
func NetworkOutput(network types.Cilium) {

	//Check for Verdict
	switch network.Verdict {
	case 1:
		fmt.Print("\nVerdict : FORWARDED,")
	case 2:
		fmt.Print("\nVerdict : DROPPED,")
	case 3:
		fmt.Print("\nVerdict : ERROR,")
	case 4:
		fmt.Print("\nVerdict : AUDIT,")
	default:
		fmt.Print("\nVerdict : VERDICT_UNKNOWN,")
	}
	//Check Ethernet exist
	if network.EthernetSource != "" || network.EthernetDestination != "" {
		fmt.Print(" Ethernet : {Source : ", network.EthernetSource, ", Destination : ", network.EthernetDestination, "},")
	}
	//Check IP version
	var ipVersion string
	switch network.IpVersion {
	case 1:
		ipVersion = "IPv4"
	case 2:
		ipVersion = "IPv6"
	default:
		ipVersion = "IP_NOT_USED"
	}
	//Print IP address
	fmt.Print(" IP : {Source : ", network.IpSource, ", Destination : ", network.IpDestination, ", Version : ", ipVersion, ", Encrypted : ", network.IpEncrypted, "},")

	//Check L4 TCP exist
	if network.L4TCPSourcePort != 0 || network.L4TCPDestinationPort != 0 {
		fmt.Print(" L4: {DNS: {SourcePort : ", network.L4TCPSourcePort, ", DestinationPort : ", network.L4TCPDestinationPort, "}},")
	}
	//Check L4 UDP exist
	if network.L4UDPSourcePort != 0 || network.L4UDPDestinationPort != 0 {
		fmt.Print(" L4 UDP: {SourcePort : ", network.L4UDPSourcePort, ", DestinationPort : ", network.L4UDPDestinationPort, "},")
	}
	//Check L4 ICMPv4 exist
	if network.L4ICMPv4Type != 0 || network.L4ICMPv4Code != 0 {
		fmt.Print(" L4 ICMPv4: {Type : ", network.L4ICMPv4Type, ", Code : ", network.L4ICMPv4Code, "},")
	}
	//Check L4 ICMPv6 exist
	if network.L4ICMPv6Type != 0 || network.L4ICMPv6Code != 0 {
		fmt.Print(" L4 ICMPv6: {Type : ", network.L4ICMPv6Type, ", Code : ", network.L4ICMPv6Code, "},")
	}
	//Source
	fmt.Print(" Source : {")
	//Check Source ID exist
	if network.SourceID != 0 {
		fmt.Print("ID : ", network.SourceID)
	}
	fmt.Print(" Identity : ", network.SourceIdentity)
	if network.SourceNamespace != "" {
		fmt.Print(" Namespace : ", network.SourceNamespace)
	}
	fmt.Print(" Labels : ", network.SourceLabels)
	if network.SourcePodName != "" {
		fmt.Printf(" PodName : %s", network.SourcePodName)
	}
	fmt.Print("}")
}
