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

	systemPods := make(map[string][]types.SystemSummery)
	networkPods := make(map[string][]types.NetworkSummary)
	//Fetch network Logs
	rows, err := database.ConnectDB().Query("select source_pod_name, destination_labels, traffic_direction from cilium_logs where source_labels like \"%"+pbRequest.Label+"%\" and source_namespace = ?", pbRequest.Namespace)
	if err != nil {
		log.Error().Msg("Error in Connection in Network Logs :" + err.Error())
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var netLog types.NetworkSummary
		var podName string
		if err := rows.Scan(&podName, &netLog.DestinationLabels, &netLog.TrafficDirection); err != nil {
			log.Error().Msg("Error in Scan system Logs : " + err.Error())
			return err
		}
		networkPods[podName] = append(networkPods[podName], netLog)

	}

	//Fetch System Logs
	rows, err = database.ConnectDB().Query("select pod_name,operation,source,resource,total from kubearmor_logs where labels like \"%"+pbRequest.Label+"%\" and namespace_name = ?", pbRequest.Namespace)
	if err != nil {
		log.Error().Msg("Error in Connection in System Logs :" + err.Error())
		return errors.New("error in Connecting system logs table")
	}
	defer rows.Close()
	for rows.Next() {
		var podName string
		var syslog types.SystemSummery
		if err := rows.Scan(&podName, &syslog.Operation, &syslog.Source, &syslog.Resource, &syslog.Count); err != nil {
			log.Error().Msg("Error in Scan system Logs : " + err.Error())
			return errors.New("error in scanning system logs table")
		}
		systemPods[podName] = append(systemPods[podName], syslog)
	}

	for podName, sysLogs := range systemPods {

		var listOfFile, listOfProcess, listOfNetwork []*sum.ListOfSource
		var ingressIn, ingressOut, egressIn, egressOut int32
		//System Block
		fileSource := make(map[string][]*sum.ListOfDestination)
		processSource := make(map[string][]*sum.ListOfDestination)
		networkSource := make(map[string][]*sum.ListOfDestination)
		// source := make(map[string]int32)
		for _, sysLog := range sysLogs {
			source := aggregateFolder(sysLog.Source)
			resource := aggregateFolder(sysLog.Resource)
			switch sysLog.Operation {
			case "File":
				fileSource[source] = convertListofDestination(fileSource[source], resource, sysLog.Count)
			case "Process":
				processSource[source] = convertListofDestination(processSource[source], resource, sysLog.Count)
			case "Network":
				protocol, _ := networkRegex(sysLog.Resource)
				if protocol != "" {
					networkSource[source] = convertListofDestination(networkSource[source], protocol, sysLog.Count)
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

		//Network Block
		for _, netLog := range networkPods[podName] {
			switch netLog.TrafficDirection {
			case "INGRESS":
				if netLog.DestinationLabels == "reserved:world" {
					ingressOut++
				} else {
					ingressIn++
				}
			case "EGRESS":
				if netLog.DestinationLabels == "reserved:world" {
					egressOut++
				} else {
					egressIn++
				}
			}
		}
		//Stream Block
		if err := stream.Send(&sum.LogsResponse{
			PodDetail:     podName,
			Namespace:     pbRequest.Namespace,
			ListOfFile:    listOfFile,
			ListOfProcess: listOfProcess,
			ListOfNetwork: listOfNetwork,
			Ingress: &sum.ListOfConnection{
				In:  ingressIn,
				Out: ingressOut,
			},
			Egress: &sum.ListOfConnection{
				In:  egressIn,
				Out: egressOut,
			},
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
func convertListofDestination(arr []*sum.ListOfDestination, destination string, count int32) []*sum.ListOfDestination {
	for _, value := range arr {
		if value.Destination == destination {
			value.Count += count
			return arr
		}
	}
	arr = append(arr, &sum.ListOfDestination{
		Destination: destination,
		Count:       count,
		Status:      "ALLOW",
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
