package hubble

import (
	"context"
	"time"

	logger "github.com/accuknox/observability/src/logger"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
	"github.com/accuknox/observability/utils/wrapper"
	"github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/api/v1/observer"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var log *zerolog.Logger = logger.GetInstance()

func dialToHubbleService() (*grpc.ClientConn, error) {
	//address to connect Cilium Service
	address := viper.GetString("cilium.url") + ":" + viper.GetString("cilium.port")
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting to grpc " + err.Error())
		return connection, err
	}
	return connection, nil
}

//GetWatchLogs - Watch Logs
func GetWatchLogs() (observer.Observer_GetFlowsClient, error) {

	var stream observer.Observer_GetFlowsClient
	//Connect Hubble Service
	connection, err := dialToHubbleService()
	if err != nil {
		return stream, err
	}
	// defer connection.Close()
	//Connect hubble client observer
	client := observer.NewObserverClient(connection)
	//health check or  try to connect until its connected
	for uptime, _ := HealthCheck(client); uptime == 0; {
		log.Info().Msg("Trying to reconnect to hubble relay server and get a observer client.")
		//Connect hubble client observer
		client = observer.NewObserverClient(connection)
		uptime, _ = HealthCheck(client)
		time.Sleep(1 * time.Second)
	}
	//Streaming the Cilium Logs
	stream, err = client.GetFlows(context.Background(), &observer.GetFlowsRequest{
		Follow: true,
		Whitelist: []*flow.FlowFilter{
			{
				TcpFlags: []*flow.TCPFlags{
					{SYN: true},
					{FIN: true},
					{RST: true},
					{NS: true},
					{ECE: true},
				},
			},
			{
				Protocol: []string{"udp", "icmp", "http", "kafka"},
			},
		},
	})
	if err != nil {
		log.Error().Msg("Error in watching logs " + err.Error())
		return stream, err
	}
	log.Info().Msg("Hubble-relay connection is successful!")
	return stream, nil
}

//HealthCheck - Health check of connection
func HealthCheck(client observer.ObserverClient) (uint64, error) {
	log.Info().Msgf("Hubble-relay HealthCheck method starts ")
	//Checking hubble observer server status
	status, err := client.ServerStatus(context.Background(), &observer.ServerStatusRequest{})
	if err != nil {
		log.Error().Msg("Error in Hubble Server Status : " + err.Error())
		return 0, err
	}

	return status.GetUptimeNs(), nil
}

//FetchLogs - Fetch Logs from Hubble Relay
func FetchLogs(stream observer.Observer_GetFlowsClient) {
	for {
		// Receiving logs from stream
		hubbleLog, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving hubble log " + err.Error())
			return
		}
		// fmt.Println("\n\nHubble Logs ===>>> ", hubbleLog)
		var getFlow *flow.Flow = hubbleLog.GetFlow()
		if getFlow != nil {
			// l3
			var ip flow.IP
			//Check l3 exist
			if getFlow.IP != nil {
				ip = *getFlow.IP
			}
			// l4
			var l4TCP flow.TCP
			var l4UDP flow.UDP
			var l4ICMPv4 flow.ICMPv4
			var l4ICMPv6 flow.ICMPv6
			//Check l4 exist
			if getFlow.L4 != nil {
				//Check TCP exist
				if getFlow.L4.GetTCP() != nil {
					l4TCP = *getFlow.L4.GetTCP()
				}
				//Check UDP exist
				if getFlow.L4.GetUDP() != nil {
					l4UDP = *getFlow.L4.GetUDP()
				}
				//Check ICMPv4 exist
				if getFlow.L4.GetICMPv4() != nil {
					l4ICMPv4 = *getFlow.L4.GetICMPv4()
				}
				//Check ICMPv6 exist
				if getFlow.L4.GetICMPv6() != nil {
					l4ICMPv6 = *getFlow.L4.GetICMPv6()
				}
			}
			//Endpoint for source and destination
			var source, destination flow.Endpoint
			//Check Source Endpoint exist
			if getFlow.Source != nil {
				source = *getFlow.Source
			}
			//Check Destination Endpoint exist
			if getFlow.Destination != nil {
				destination = *getFlow.Destination
			}

			//l7
			var l7 flow.Layer7
			var l7Type string
			var l7DNS flow.DNS
			var l7HTTP flow.HTTP
			var l7HTTPHeaders string
			//Check l7 exist
			if getFlow.L7 != nil {
				l7 = *getFlow.GetL7()
				l7Type = l7.GetType().Enum().String()
				//Check DNS exist
				if l7.GetDns() != nil {
					l7DNS = *l7.GetDns()
				}
				//Check HTTP exist
				if l7.GetHttp() != nil {
					l7HTTP = *l7.GetHttp()
					var headers []string
					//Check Headers exist
					if l7HTTP.GetHeaders() != nil {
						//convert headers in key=value format.
						for _, header := range l7HTTP.Headers {
							headers = append(headers, header.Key+"="+header.Value)
						}
						//convert http Header into string format
						l7HTTPHeaders = wrapper.ConvertArrayToString(headers)
					}
				}
			}

			//EventType
			var eventType, eventSubType int32
			if getFlow.EventType != nil {
				eventType = getFlow.EventType.GetType()
				eventSubType = getFlow.EventType.GetSubType()
			}

			//Service Name for source and destination
			var sourceService, destinationService flow.Service
			//Check Service Source exist
			if getFlow.SourceService != nil {
				sourceService = *getFlow.GetSourceService()
			}
			//Check Service Destination exist
			if getFlow.DestinationService != nil {
				destinationService = *getFlow.GetDestinationService()
			}

			var isReply wrappers.BoolValue
			//Check IsReply exist
			if getFlow.IsReply != nil {
				isReply = *getFlow.IsReply
			}

			var dropReason string
			//Check Verdict is Dropped
			if getFlow.GetVerdict().Enum().String() == "DROPPED" {
				dropReason = getFlow.GetDropReasonDesc().Enum().String()
			}
			//Select Query to fetch ID
			row := database.ConnectDB().QueryRow(constants.SELECT_CILIUM, getFlow.GetVerdict().Enum().String(),
				ip.Source, ip.Destination, ip.GetIpVersion().Enum().String(), ip.Encrypted,
				l4TCP.SourcePort,
				l4TCP.DestinationPort,
				l4UDP.SourcePort,
				l4UDP.DestinationPort,
				l4ICMPv4.Type,
				l4ICMPv4.Code,
				l4ICMPv6.Type,
				l4ICMPv6.Code,
				source.Namespace, wrapper.ConvertArrayToString(source.Labels), source.PodName,
				destination.Namespace, wrapper.ConvertArrayToString(destination.Labels), destination.PodName,
				getFlow.GetType().Enum().String(),
				getFlow.NodeName,
				l7Type,
				wrapper.ConvertArrayToString(l7DNS.Cnames),
				l7DNS.ObservationSource,
				l7HTTP.Code,
				l7HTTP.Method,
				l7HTTP.Url,
				l7HTTP.Protocol,
				l7HTTPHeaders,
				eventType, eventSubType,
				sourceService.Name, sourceService.Namespace,
				destinationService.Name, destinationService.Namespace,
				getFlow.GetTrafficDirection().Enum().String(),
				getFlow.GetTraceObservationPoint().Enum().String(),
				dropReason,
				isReply.Value)
			if row.Err() != nil {
				log.Error().Msg("Error in Select Query from Cilium Log Table : " + row.Err().Error())
			}
			//Store the ID
			var id int
			//Scan ID
			row.Scan(&id)
			//Check record is unique or not
			if id != 0 {
				//Prepare the update query statement
				statement, err := database.ConnectDB().Prepare(constants.UPDATE_CILIUM)
				if err != nil {
					log.Error().Msg("Error in Prepare Update Cilium statement: " + err.Error())
					return
				}
				//Execute the update query statement
				statement.Exec(getFlow.Time.Seconds, id)
				defer statement.Close()

			} else {
				//Prepare the insert query statement
				statement, err := database.ConnectDB().Prepare(constants.INSERT_CILIUM)
				if err != nil {
					log.Error().Msg("Error in Prepare Cilium Insert statement: " + err.Error())
					return
				}

				//Execute the insert query statement
				statement.Exec(getFlow.GetVerdict().Enum().String(),
					ip.Source, ip.Destination, ip.GetIpVersion().Enum().String(), ip.Encrypted,
					l4TCP.SourcePort,
					l4TCP.DestinationPort,
					l4UDP.SourcePort,
					l4UDP.DestinationPort,
					l4ICMPv4.Type,
					l4ICMPv4.Code,
					l4ICMPv6.Type,
					l4ICMPv6.Code,
					source.Namespace, wrapper.ConvertArrayToString(source.Labels), source.PodName,
					destination.Namespace, wrapper.ConvertArrayToString(destination.Labels), destination.PodName,
					getFlow.GetType().Enum().String(),
					getFlow.NodeName,
					l7Type,
					wrapper.ConvertArrayToString(l7DNS.Cnames),
					l7DNS.ObservationSource,
					l7HTTP.Code,
					l7HTTP.Method,
					l7HTTP.Url,
					l7HTTP.Protocol,
					l7HTTPHeaders,
					eventType, eventSubType,
					sourceService.Name, sourceService.Namespace,
					destinationService.Name, destinationService.Namespace,
					getFlow.GetTrafficDirection().Enum().String(),
					getFlow.GetTraceObservationPoint().Enum().String(),
					dropReason,
					isReply.Value,
					getFlow.Time.Seconds, getFlow.Time.Seconds)

				defer statement.Close()

			}
		}
	}
}
