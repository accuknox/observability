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
	rows, err = database.ConnectDB().Query("select pod_name,operation,source,resource from kubearmor_logs where labels like \"%"+pbRequest.Label+"%\" and namespace_name = ?", pbRequest.Namespace)
	if err != nil {
		log.Error().Msg("Error in Connection in System Logs :" + err.Error())
		return errors.New("error in Connecting system logs table")
	}
	defer rows.Close()
	for rows.Next() {
		var podName string
		var syslog types.SystemSummery
		if err := rows.Scan(&podName, &syslog.Operation, &syslog.Source, &syslog.Resource); err != nil {
			log.Error().Msg("Error in Scan system Logs : " + err.Error())
			return errors.New("error in scanning system logs table")
		}
		systemPods[podName] = append(systemPods[podName], syslog)
	}

	for podName, sysLogs := range systemPods {

		var listOfFile, listOfProcess, listOfNetwork []*sum.ListOfSource
		var ingressIn, ingressOut, egressIn, egressOut int32
		//System Block
		fileSource := make(map[string][]string)
		processSource := make(map[string][]string)
		networkSource := make(map[string][]string)

		for _, sysLog := range sysLogs {
			source := strings.Split(sysLog.Source, " ")[0]
			resource := strings.Split(sysLog.Resource, " ")[0]
			switch sysLog.Operation {
			case "File":
				if !isExist(fileSource[source], resource) {
					fileSource[source] = append(fileSource[source], resource)
				}
			case "Process":
				if !isExist(processSource[source], resource) {
					processSource[source] = append(processSource[source], resource)
				}
			case "Network":
				protocol, _ := networkRegex(sysLog.Resource)
				if protocol != "" {
					if !isExist(networkSource[source], protocol) {
						networkSource[source] = append(networkSource[source], protocol)
					}
				}
			}
		}
		for source, resources := range fileSource {
			listOfFile = append(listOfFile, &sum.ListOfSource{
				Source:   source,
				Resource: resources,
			})
		}
		for source, resources := range processSource {
			listOfProcess = append(listOfProcess, &sum.ListOfSource{
				Source:   source,
				Resource: resources,
			})
		}

		for source, protocols := range networkSource {
			listOfNetwork = append(listOfNetwork, &sum.ListOfSource{
				Source:   source,
				Resource: protocols,
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
		return "tcp", nil
	}
	reudp, err = regexp.Compile("domain=.*type=SOCK_DGRAM")
	if err != nil {
		log.Error().Msgf("failed udp regexp compile err=%s", err.Error())
		return "", err
	}
	if reudp.MatchString(str) {
		return "udp", nil
	}
	reicmp, err = regexp.Compile(`domain=.*protocol=(\b58\b|\b1\b)`) //1=icmp, 58=icmp6
	if err != nil {
		log.Error().Msgf("failed icmp regexp compile err=%s", err.Error())
		return "", err
	}
	if reicmp.MatchString(str) {
		return "icmp", nil
	}
	reraw, err = regexp.Compile("domain=.*type=SOCK_RAW")
	if err != nil {
		log.Error().Msgf("failed raw regexp compile err=%s", err.Error())
		return "", err
	}
	if reraw.MatchString(str) {
		return "raw", nil
	}
	return "", nil
}

//isExist - To find string is exist in the array or not
func isExist(arr []string, str string) bool {
	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}
