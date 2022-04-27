package aggregator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	agg "github.com/accuknox/observability/src/proto/aggregator"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetNetworkLogs(pbNetworkRequest *agg.NetworkLogsRequest, stream agg.Aggregator_FetchNetworkLogsServer) error {
	var count int64
	var query string
	var filterQuery []string

	//Check Direction Filter exist
	if pbNetworkRequest.Direction != "" {
		filterQuery = append(filterQuery, " traffic_direction =  \""+strings.ToUpper(pbNetworkRequest.Direction)+"\"")
	}

	//Check Type Filter exist
	if pbNetworkRequest.Type != "" {
		filterQuery = append(filterQuery, " type = \""+strings.ToUpper(pbNetworkRequest.Type)+"\"")
	}

	//Check Verdict Filter exist
	if len(pbNetworkRequest.Verdict) != 0 {
		filterQuery = append(filterQuery, " verdict in (\""+strings.ToUpper(strings.Join(pbNetworkRequest.Verdict, "\",\""))+"\")")
	}

	//Check Protocol Filter exist
	if len(pbNetworkRequest.Protocol) != 0 {

		switch pbNetworkRequest.Protocol {
		case "TCP":
			filterQuery = append(filterQuery, " l4_tcp_source_port != 0")
		case "UDP":
			filterQuery = append(filterQuery, " l4_udp_source_port != 0")
		case "ICMPv4":
			filterQuery = append(filterQuery, " l4_icmpv4_type != 0")
		case "ICMPv6":
			filterQuery = append(filterQuery, " l4_icmpv6_type != 0")
		default:
			log.Error().Msg("Invalid protocol filter Value : " + pbNetworkRequest.Protocol)
			return status.Errorf(codes.InvalidArgument, "Error in Protocol Filter")
		}

	}
	// Check L7 exist
	if len(pbNetworkRequest.L7) != 0 {

		switch pbNetworkRequest.L7 {
		case "DNS":
			filterQuery = append(filterQuery, " l7_dns_cnames != \"\"")
		case "HTTP":
			filterQuery = append(filterQuery, " l7_http_url != \"\"")
		default:
			log.Error().Msg("Invalid L7 Protocol filter Value : " + pbNetworkRequest.L7)
			return status.Errorf(codes.InvalidArgument, "Error in L7 Protocol Filter")
		}

	}

	//Check Source Pod Filter exist
	if len(pbNetworkRequest.SourcePod) != 0 {
		filterQuery = append(filterQuery, " source_pod_name in (\""+strings.Join(pbNetworkRequest.SourcePod, "\",\"")+"\")")
	}
	//Check Source Namespace Filter exist
	if len(pbNetworkRequest.SourceNamespace) != 0 {
		filterQuery = append(filterQuery, " source_namespace in (\""+strings.Join(pbNetworkRequest.SourceNamespace, "\",\"")+"\")")
	}
	//Check Destination Pod Filter exist
	if len(pbNetworkRequest.DestinationPod) != 0 {
		filterQuery = append(filterQuery, " destination_pod_name in (\""+strings.Join(pbNetworkRequest.DestinationPod, "\",\"")+"\")")
	}
	//Check Destination Namespace Filter exist
	if len(pbNetworkRequest.DestinationNamespace) != 0 {
		filterQuery = append(filterQuery, " destination_namespace in (\""+strings.Join(pbNetworkRequest.DestinationNamespace, "\",\"")+"\")")
	}
	//Check Node Filter exist
	if len(pbNetworkRequest.Node) != 0 {
		filterQuery = append(filterQuery, " node_name in (\""+strings.Join(pbNetworkRequest.Node, "\",\"")+"\")")
	}
	//Check SourceLabel Filter exist
	if pbNetworkRequest.SourceLabel != "" {
		filterQuery = append(filterQuery, " source_labels like \"%"+pbNetworkRequest.SourceLabel+"%\"")
	}
	//Check DestinationLabel Filter exist
	if pbNetworkRequest.DestinationLabel != "" {
		filterQuery = append(filterQuery, " destination_labels like \"%"+pbNetworkRequest.DestinationLabel+"%\"")
	}
	// Check Since Filter exist
	if pbNetworkRequest.Since != "" {

		currentTime := time.Now().UTC().Unix()

		givenTime, err := strconv.ParseInt(pbNetworkRequest.Since[:len(pbNetworkRequest.Since)-1], 10, 64)
		if err != nil {
			log.Error().Msg("invalid Since filter value : " + pbNetworkRequest.Since)
			return status.Errorf(codes.InvalidArgument, "Error in Since Filter")
		}

		switch pbNetworkRequest.Since[len(pbNetworkRequest.Since)-1:] {
		case "d":
			filterQuery = append(filterQuery, " updated_time > "+fmt.Sprint(currentTime-int64(givenTime)*24*60*60))
		case "h":
			filterQuery = append(filterQuery, " updated_time > "+fmt.Sprint(currentTime-int64(givenTime)*60*60))
		case "m":
			filterQuery = append(filterQuery, " updated_time > "+fmt.Sprint(currentTime-int64(givenTime)*60))
		case "s":
			filterQuery = append(filterQuery, " updated_time > "+fmt.Sprint(currentTime-int64(givenTime)))
		default:
			log.Error().Msg("invalid Since filter value : " + pbNetworkRequest.Since[len(pbNetworkRequest.Since)-1:])
			return status.Errorf(codes.InvalidArgument, "Error in Since Filter")
		}
	}

	//Check Any filter exist
	if len(filterQuery) != 0 {
		query = " where" + strings.Join(filterQuery, " and")
	}

	//Check User want log or count of log
	if pbNetworkRequest.Count {
		query = constants.SELECT_COUNT_CILIUM + query

		//Fetch rows
		row := database.ConnectDB().QueryRow(query)
		if err := row.Scan(&count); err != nil {
			log.Error().Msg("Error in Connection in Network Logs :" + err.Error())
			return errors.New("error in Connecting network logs table")
		}
		if err := stream.Send(&agg.NetworkLogsResponse{Count: count}); err != nil {
			log.Error().Msg("Error in Streaming Network Count : " + err.Error())
			return err
		}
	} else {
		query = constants.SELECT_ALL_CILIUM + query + constants.ORDER_BY_UPDATED_TIME
		//Check limit exist
		if pbNetworkRequest.Limit != 0 {
			//query to fetch all logs with limit
			query = query + constants.LIMIT + strconv.FormatInt(pbNetworkRequest.Limit, 10)
		}
		//Fetch rows
		rows, err := database.ConnectDB().Query(query)
		if err != nil {
			log.Error().Msg("Error in Connection in System Logs : " + err.Error())
			return errors.New("error in Connecting system logs table")
		}
		defer rows.Close()
		for rows.Next() {
			var netlog agg.NetworkLog
			//Scan logs
			if err := rows.Scan(&netlog.Verdict,
				&netlog.IpSource, &netlog.IpDestination, &netlog.IpVersion, &netlog.IpEncrypted,
				&netlog.L4TcpSourcePort, &netlog.L4TcpDestinationPort, &netlog.L4UdpSourcePort, &netlog.L4UdpDestinationPort,
				&netlog.L4Icmpv4Type, &netlog.L4Icmpv4Code, &netlog.L4Icmpv6Type, &netlog.L4Icmpv6Code,
				&netlog.SourceNamespace, &netlog.SourceLabels, &netlog.SourcePodName,
				&netlog.DestinationNamespace, &netlog.DestinationLabels, &netlog.DestinationPodName,
				&netlog.Type, &netlog.NodeName, &netlog.L7Type,
				&netlog.L7DnsCnames, &netlog.L7DnsObservationSource,
				&netlog.L7HttpCode, &netlog.L7HttpMethod, &netlog.L7HttpUrl, &netlog.L7HttpProtocol, &netlog.L7HttpHeaders,
				&netlog.EventTypeType, &netlog.EventTypeSubType, &netlog.SourceServiceName, &netlog.SourceServiceNamespace, &netlog.DestinationServiceName, &netlog.DestinationServiceNamespace,
				&netlog.TrafficDirection, &netlog.TraceObservationPoint, &netlog.DropReasonDesc,
				&netlog.IsReply,
				&netlog.StartTime, &netlog.UpdatedTime, &netlog.Total); err != nil {
				log.Error().Msg("Error in Scan network Logs : " + err.Error())
				return status.Errorf(codes.InvalidArgument, "Error in scanning network logs table")
			}
			if err := stream.Send(&agg.NetworkLogsResponse{Logs: &netlog}); err != nil {
				log.Error().Msg("Error in Streaming Network Logs : " + err.Error())
				return err
			}
		}

	}
	return nil
}
