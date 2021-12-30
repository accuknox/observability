package kubearmor

import (
	"context"

	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
	kubearmor "github.com/kubearmor/KubeArmor/protobuf"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func dialToKubeArmorService() (*grpc.ClientConn, error) {
	//address to connect KubeArmor Service
	address := viper.GetString("kubeArmor.url") + ":" + viper.GetString("kubeArmor.port")
	//Connecting client on given target address
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting kubearmor-relay : " + err.Error())
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
	//Connect kubearmor log service client
	client := kubearmor.NewLogServiceClient(connection)
	//health check or try to connect until its connected
	for healthCheck := HealthCheck(client); !healthCheck; {
		log.Info().Msg("Trying to connect to kubearmor service and get a log service client.")
		client = kubearmor.NewLogServiceClient(connection)
		healthCheck = HealthCheck(client)
	}
	//Streaming the KubeArmor Logs
	stream, err := client.WatchLogs(context.Background(), &kubearmor.RequestMessage{})
	if err != nil {
		log.Error().Msg("Error in watching kubearmor logs " + err.Error())
		return
	}
	for {
		//Fetch the kubearmor Logs
		kubearmorLog, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving kubearmor log " + err.Error())
			return
		}
		// fmt.Println("\n\nKubeArmor Logs ===>>> ", kubearmorLog)
		//Check the logs which is Passed
		if kubearmorLog.Result == constants.STATUS {
			//Select Query to fetch ID
			row := database.ConnectDB().QueryRow(constants.SELECT_KUBEARMOR, kubearmorLog.ClusterName, kubearmorLog.HostName, kubearmorLog.NamespaceName, kubearmorLog.PodName, kubearmorLog.ContainerID, kubearmorLog.ContainerName,
				kubearmorLog.UID, kubearmorLog.Type, kubearmorLog.Source, kubearmorLog.Operation, kubearmorLog.Resource,
				kubearmorLog.Data)
			if err != nil {
				log.Error().Msg("Error in Select Query from KubeArmor Log Table : " + err.Error())
				return
			}
			//Store the ID
			var id int
			//Scan ID
			row.Scan(&id)
			//Check record is unique or not
			if id != 0 {
				//Prepare the update query statement
				statement, err := database.ConnectDB().Prepare(constants.UPDATE_KUBEARMOR)
				if err != nil {
					log.Error().Msg("Error in Prepare Update KubeArmor statement: " + err.Error())
					return
				}
				//Execute the update query statement
				statement.Exec(kubearmorLog.Timestamp, id)
				defer statement.Close()

			} else {
				//Prepare the insert query statement
				statement, err := database.ConnectDB().Prepare(constants.INSERT_KUBEARMOR)
				if err != nil {
					log.Error().Msg("Error in Prepare Insert KubeArmor statement: " + err.Error())
					return
				}
				//Execute the insert query statement
				statement.Exec(kubearmorLog.ClusterName, kubearmorLog.HostName, kubearmorLog.NamespaceName, kubearmorLog.PodName, kubearmorLog.ContainerID, kubearmorLog.ContainerName,
					kubearmorLog.UID, kubearmorLog.Type, kubearmorLog.Source, kubearmorLog.Operation, kubearmorLog.Resource,
					kubearmorLog.Data, kubearmorLog.Timestamp, kubearmorLog.Timestamp)

				defer statement.Close()
			}
		}
	}

}

//HealthCheck - Health check of connection
func HealthCheck(client kubearmor.LogServiceClient) bool {
	log.Info().Msg("Kubearmor-relay HealthCheck method starts ")

	value := int32(57684)
	arg := &kubearmor.NonceMessage{
		Nonce: value,
	}
	//Checking client healthcheck
	healthCheck, err := client.HealthCheck(context.Background(), arg)
	if err != nil {
		log.Error().Msg("Kubearmor-relay connection is not successful!")
		log.Error().Err(err)
	}

	if healthCheck != nil && healthCheck.Retval == value {
		log.Info().Msg("Kubearmor-relay connection is successful!")
		return true
	}

	return false
}
