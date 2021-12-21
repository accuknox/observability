package hubble

import (
	"context"
	"fmt"
	"time"

	"github.com/cilium/cilium/api/v1/observer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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
func GetWatchLogs() {
	//Connect Hubble Service
	connection, err := dialToHubbleService()
	if err != nil {
		return
	}
	defer connection.Close()
	//Connect hubble client observer
	client := observer.NewObserverClient(connection)
	//health check or  try to connect until its connected
	for uptime, _ := checkServerUptime(client); uptime == 0; {
		log.Info().Msg("Trying to reconnect to hubble relay server and get a observerclient.")
		//Connect hubble client observer
		client = observer.NewObserverClient(connection)
		uptime, _ = checkServerUptime(client)
		time.Sleep(1 * time.Second)
	}
	//Streaming the Cilium Logs
	stream, err := client.GetFlows(context.Background(), &observer.GetFlowsRequest{Follow: true})
	if err != nil {
		log.Error().Msg("Error in watching logs " + err.Error())
		return
	}
	for {
		//Receiving logs from stream
		hubbleLog, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving hubble log " + err.Error())
			return
		}
		fmt.Println("\n\nHubble Logs ===>>> ", hubbleLog)
	}

}

//checkServerUptime - Health check of connection
func checkServerUptime(client observer.ObserverClient) (uint64, error) {
	log.Info().Msg("CheckServerUptime  method starts ")
	//Checking hubble observer server status
	status, err := client.ServerStatus(context.Background(), &observer.ServerStatusRequest{})
	log.Info().Msg("Hubble server Status  : " + fmt.Sprint(status))
	if err != nil {
		log.Error().Msg("Error in Hubble Server Status : " + err.Error())
		return 0, err
	}

	return status.GetUptimeNs(), nil
}
