package summary

import (
	"errors"
	"regexp"
	"strings"

	sum "github.com/accuknox/observability/src/proto/summary"
	"github.com/accuknox/observability/src/types"
	"github.com/accuknox/observability/utils/database"
	"github.com/rs/zerolog/log"
)

//GetSummaryLogs - Give Summary logs of Pod based on Label and Namespace Input
func GetSummaryLogs(pbRequest *sum.LogsRequest, stream sum.Summary_FetchLogsServer) error {
	log.Info().Msg("Get Summary Log Called")
	systemPods := make(map[string][]types.SystemSummery)
	networkPods := make(map[string][]types.NetworkSummary)
	//Fetch network Logs
	rows, err := database.ConnectDB().Query("select source_pod_name, verdict, destination_labels, destination_namespace, type, l4_tcp_destination_port, l4_udp_destination_port, l4_icmpv4_code, l4_icmpv6_code,l7_dns_cnames,l7_http_method, traffic_direction, updated_time, total from cilium_logs where source_labels like \"%"+pbRequest.Label+"%\" and source_namespace = ?", pbRequest.Namespace)
	if err != nil {
		log.Error().Msg("Error in Connection in Network Logs :" + err.Error())
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var netLog types.NetworkSummary
		var podName string
		if err := rows.Scan(&podName, &netLog.Verdict, &netLog.DestinationLabels, &netLog.DestinationNamespace, &netLog.Type, &netLog.L4TCPDestinationPort, &netLog.L4UDPDestinationPort, &netLog.L4ICMPv4Code, &netLog.L4ICMPv6Code, &netLog.L7DnsCnames, &netLog.L7HttpMethod, &netLog.TrafficDirection, &netLog.UpdatedTime, &netLog.Count); err != nil {
			log.Error().Msg("Error in Scan system Logs : " + err.Error())
			return err
		}
		networkPods[podName] = append(networkPods[podName], netLog)

	}

	//Fetch System Logs
	rows, err = database.ConnectDB().Query("select pod_name,operation,source,resource,action,updated_time,total from kubearmor_logs where labels like \"%"+pbRequest.Label+"%\" and namespace_name = ?", pbRequest.Namespace)
	if err != nil {
		log.Error().Msg("Error in Connection in System Logs :" + err.Error())
		return errors.New("error in Connecting system logs table")
	}
	defer rows.Close()
	for rows.Next() {
		var podName string
		var syslog types.SystemSummery
		if err := rows.Scan(&podName, &syslog.Operation, &syslog.Source, &syslog.Resource, &syslog.Action, &syslog.UpdatedTime, &syslog.Count); err != nil {
			log.Error().Msg("Error in Scan system Logs : " + err.Error())
			return errors.New("error in scanning system logs table")
		}
		systemPods[podName] = append(systemPods[podName], syslog)
	}

	for podName, sysLogs := range systemPods {

		var listOfFile, listOfProcess, listOfNetwork []*sum.ListOfSource
		//System Block
		fileSource := make(map[string][]*sum.ListOfDestination)
		processSource := make(map[string][]*sum.ListOfDestination)
		networkSource := make(map[string][]*sum.ListOfDestination)
		// source := make(map[string]int32)
		for _, sysLog := range sysLogs {
			source := strings.Split(sysLog.Source, " ")[0]
			//Checking System Operation that's File, Process and Network
			switch sysLog.Operation {
			case "File":
				fileSource[source] = convertListofDestination(fileSource[source], sysLog)
			case "Process":
				processSource[source] = convertListofDestination(processSource[source], sysLog)
			case "Network":
				protocol, _ := networkRegex(sysLog.Resource)
				if protocol != "" {
					networkSource[source] = convertListofDestination(networkSource[source], sysLog)
				}
			}
		}
		for source, resources := range fileSource {
			listOfFile = append(listOfFile, &sum.ListOfSource{
				Source:            source,
				ListOfDestination: resources,
			})
		}
		for source, resources := range processSource {
			listOfProcess = append(listOfProcess, &sum.ListOfSource{
				Source:            source,
				ListOfDestination: resources,
			})
		}

		for source, protocols := range networkSource {
			listOfNetwork = append(listOfNetwork, &sum.ListOfSource{
				Source:            source,
				ListOfDestination: protocols,
			})
		}
		var networkIngress, networkEgress []*sum.ListOfConnection
		//Network Block
		for _, netLog := range networkPods[podName] {
			//Check Traffic Direction is Ingress or Egress
			switch netLog.TrafficDirection {
			case "INGRESS":
				networkIngress = convertNetworkConnection(netLog, networkIngress)
			case "EGRESS":
				networkEgress = convertNetworkConnection(netLog, networkEgress)
			}
		}
		//Stream Block
		if err := stream.Send(&sum.LogsResponse{
			PodDetail:     podName,
			Namespace:     pbRequest.Namespace,
			ListOfFile:    listOfFile,
			ListOfProcess: listOfProcess,
			ListOfNetwork: listOfNetwork,
			Ingress:       networkIngress,
			Egress:        networkEgress,
		}); err != nil {
			log.Error().Msg("Error in Streaming Summary Logs : " + err.Error())
		}
	}

	return nil
}

//networkRegex - To Get the Protocol using Regex
func networkRegex(str string) (string, error) {
	var retcp, reudp, reicmp, reraw *regexp.Regexp

	retcp, err := regexp.Compile("domain=.*type=SOCK_STREAM")
	if err != nil {
		log.Error().Msgf("failed tcp regexp compile err=%s", err.Error())
		return "", err
	}
	if retcp.MatchString(str) {
		return "TCP", nil
	}
	reudp, err = regexp.Compile("domain=.*type=SOCK_DGRAM")
	if err != nil {
		log.Error().Msgf("failed udp regexp compile err=%s", err.Error())
		return "", err
	}
	if reudp.MatchString(str) {
		return "UDP", nil
	}
	reicmp, err = regexp.Compile(`domain=.*protocol=(\b58\b|\b1\b)`) //1=icmp, 58=icmp6
	if err != nil {
		log.Error().Msgf("failed icmp regexp compile err=%s", err.Error())
		return "", err
	}
	if reicmp.MatchString(str) {
		return "ICMP", nil
	}
	reraw, err = regexp.Compile("domain=.*type=SOCK_RAW")
	if err != nil {
		log.Error().Msgf("failed raw regexp compile err=%s", err.Error())
		return "", err
	}
	if reraw.MatchString(str) {
		return "RAW", nil
	}
	return "", nil
}

//convertListofDestination - Create the mapping between Source and Destination/Resource/Protocol
func convertListofDestination(arr []*sum.ListOfDestination, sysLog types.SystemSummery) []*sum.ListOfDestination {
	destination := aggregateFolder(sysLog.Resource)
	//Check Operation is Network
	if sysLog.Operation == "Network" {
		destination, _ = networkRegex(sysLog.Resource)
	}
	for _, value := range arr {
		if value.Destination == destination {
			value.Count += sysLog.Count
			value.LastUpdatedTime = sysLog.UpdatedTime
			return arr
		}
	}
	arr = append(arr, &sum.ListOfDestination{
		Destination:     destination,
		Count:           sysLog.Count,
		Status:          strings.ToUpper(sysLog.Action),
		LastUpdatedTime: sysLog.UpdatedTime,
	})
	return arr
}

/* aggregateFolder - Aggreagte the Folder or File path with Parent name
For Example - Folder Name is /abc/bin/1234 or /abc/xyz.txt --> convert this into /abc/*
*/
func aggregateFolder(str string) string {

	switch str {
	case "":
		return str
	case "/":
		return str
	default:
		if strings.HasPrefix(str, "/") {
			s := strings.SplitAfterN(str, "/", -1)[1]
			if strings.HasSuffix(s, "/") {
				return "/" + s + "*"
			}

			return "/" + s
		}
		return str
	}
}

func convertNetworkConnection(netLog types.NetworkSummary, list []*sum.ListOfConnection) []*sum.ListOfConnection {

	var listOfConn sum.ListOfConnection

	listOfConn.DestinationLabels = netLog.DestinationLabels
	listOfConn.DestinationNamespace = netLog.DestinationNamespace

	if netLog.L4TCPDestinationPort != 0 {
		listOfConn.Port = netLog.L4TCPDestinationPort
		if netLog.L7HttpMethod != "" {
			listOfConn.Protocol = "HTTP"
		} else {
			listOfConn.Protocol = "TCP"
		}
	} else if netLog.L4UDPDestinationPort != 0 {
		listOfConn.Port = netLog.L4UDPDestinationPort
		if netLog.L7DnsCnames != "" {
			listOfConn.Protocol = "DNS"
		} else {
			listOfConn.Protocol = "TCP"
		}
	} else if netLog.L4ICMPv4Code != 0 {
		listOfConn.Protocol = "ICMPv4"
	} else {
		listOfConn.Protocol = "ICMPv6"
	}

	//Find Status
	switch netLog.Verdict {
	case "FORWARDED", "REDIRECTED":
		listOfConn.Status = "ALLOW"
	case "DROPPED", "ERROR":
		listOfConn.Status = "DENY"
	case "AUDIT":
		listOfConn.Status = "AUDIT"
	}

	for _, value := range list {

		if value.DestinationLabels == listOfConn.DestinationLabels && value.DestinationNamespace == listOfConn.DestinationNamespace &&
			value.Protocol == listOfConn.Protocol && value.Port == listOfConn.Port && value.Status == listOfConn.Status {
			value.Count += netLog.Count
			value.LastUpdatedTime = netLog.UpdatedTime
			return list
		}
	}
	listOfConn.Count = netLog.Count
	listOfConn.LastUpdatedTime = netLog.UpdatedTime
	list = append(list, &listOfConn)
	return list
}
