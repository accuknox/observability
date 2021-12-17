package kubearmor

import (
	"context"
	"fmt"

	kubearmor "github.com/kubearmor/KubeArmor/protobuf"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func dialToKubeArmorService() (*grpc.ClientConn, error) {
	//address to connect KubeArmor Service
	address := viper.GetString("kubeArmor.url") + ":" + viper.GetString("kubeArmor.port")
	fmt.Println("Address : ", address)
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting to grpc " + err.Error())
		return connection, err
	}
	return connection, nil
}

//GetWatchLogs - Watch Logs
func GetWatchLogs() {
	//Connect KubeArmor Service
	connection, err := dialToKubeArmorService()
	if err != nil {
		return
	}
	defer connection.Close()

	client := kubearmor.NewLogServiceClient(connection)
	//health check or  try to connect until its connected
	for healthCheck := HealthCheck(client); !healthCheck; {
		log.Info().Msg("Trying to connect to kubearmor service and get a log service client.")
		client = kubearmor.NewLogServiceClient(connection)
		healthCheck = HealthCheck(client)
	}
	//Streaming the KubeArmor Logs
	stream, err := client.WatchLogs(context.Background(), &kubearmor.RequestMessage{})
	if err != nil {
		log.Error().Msg("Error in watching logs " + err.Error())
		return
	}
	for {
		//Fetch the kubearmor Logs
		kubearmorLog, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving kubearmor log " + err.Error())
			return
		}

		fmt.Println("\n\nKubeArmor Logs ===>>> ", kubearmorLog)
	}

}

//HealthCheck - Health check of connection
func HealthCheck(client kubearmor.LogServiceClient) bool {
	log.Info().Msg(" IsLogServiceServerHealthy method starts ")

	value := int32(57684)
	arg := &kubearmor.NonceMessage{
		Nonce: value,
	}

	healthCheck, err := client.HealthCheck(context.Background(), arg)
	if err != nil {
		log.Error().Msg("The connection is not successful!")
		log.Error().Err(err)
	}

	if healthCheck != nil && healthCheck.Retval == value {
		log.Info().Msg(" The connection is successful!")
		return true
	}

	return false
}
