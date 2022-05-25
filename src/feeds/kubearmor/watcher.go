package kubearmor

import (
	"context"
	"encoding/json"

	logger "github.com/accuknox/observability/src/logger"
	"github.com/accuknox/observability/src/types"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
	kubearmor "github.com/kubearmor/KubeArmor/protobuf"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var log *zerolog.Logger = logger.GetInstance()

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

//GetWatchLogs - Watch Logs and Alerts
func GetWatchLogs() (kubearmor.LogService_WatchLogsClient, kubearmor.LogService_WatchAlertsClient, error) {
	var clientLog kubearmor.LogService_WatchLogsClient
	var clientAlert kubearmor.LogService_WatchAlertsClient
	//Connect KubeArmor Service
	connection, err := dialToKubeArmorService()
	if err != nil {
		return clientLog, clientAlert, err
	}
	// defer connection.Close()
	//Connect kubearmor log service client
	serviceClientLog := kubearmor.NewLogServiceClient(connection)
	//health check or try to connect until its connected
	for healthCheck := HealthCheck(serviceClientLog); !healthCheck; {
		log.Info().Msg("Trying to connect to kubearmor service and get a log service client.")
		serviceClientLog = kubearmor.NewLogServiceClient(connection)
		healthCheck = HealthCheck(serviceClientLog)
	}
	serviceClientAlert := kubearmor.NewLogServiceClient(connection)
	//health check or try to connect until its connected
	for healthCheck := HealthCheck(serviceClientAlert); !healthCheck; {
		log.Info().Msg("Trying to connect to kubearmor service and get a Alert service client.")
		serviceClientAlert = kubearmor.NewLogServiceClient(connection)
		healthCheck = HealthCheck(serviceClientAlert)
	}
	//Streaming the KubeArmor Logs
	clientLog, err = serviceClientLog.WatchLogs(context.Background(), &kubearmor.RequestMessage{Filter: "system"})
	if err != nil {
		log.Error().Msg("Error in watching kubearmor logs " + err.Error())
		return clientLog, clientAlert, err
	}
	clientAlert, err = serviceClientAlert.WatchAlerts(context.Background(), &kubearmor.RequestMessage{Filter: "policy"})
	if err != nil {
		log.Error().Msg("Error in watching kubearmor alerts " + err.Error())
		return clientLog, clientAlert, err
	}
	return clientLog, clientAlert, nil
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

//FetchAlert - Call FetchLogs
func FetchAlert(stream kubearmor.LogService_WatchAlertsClient) {
	for {
		var kubeArmorAlert types.KubeArmorLogAlert
		logs, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving kubearmor alert " + err.Error())
			return
		}

		jsonLog, _ := json.Marshal(logs)
		err = json.Unmarshal(jsonLog, &kubeArmorAlert)
		kubeArmorAlert.Category = "Alert"
		// fmt.Println("\n\nKubeArmor Alert ===>>> ", kubearmorLog)
		AggregateLogs(kubeArmorAlert)
	}
}

//FetchLogs - Convert Aggregate the Logs and Alert and Store into DB
func FetchLogs(stream kubearmor.LogService_WatchLogsClient) {
	for {
		var kubeArmorLog types.KubeArmorLogAlert
		logs, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving kubearmor log " + err.Error())
			return
		}
		jsonLog, _ := json.Marshal(logs)
		err = json.Unmarshal(jsonLog, &kubeArmorLog)
		kubeArmorLog.Action = "Allow"
		kubeArmorLog.Category = "Log"
		// fmt.Println("\n\nKubeArmor Logs ===>>> ", kubearmorLog)
		AggregateLogs(kubeArmorLog)
	}
}

func AggregateLogs(kubearmorLog types.KubeArmorLogAlert) {
	//Select Query to fetch ID
	row := database.ConnectDB().QueryRow(constants.SELECT_KUBEARMOR, kubearmorLog.ClusterName, kubearmorLog.HostName, kubearmorLog.NamespaceName, kubearmorLog.PodName, kubearmorLog.ContainerID, kubearmorLog.ContainerName,
		kubearmorLog.UID, kubearmorLog.Type, kubearmorLog.Source, kubearmorLog.Operation, kubearmorLog.Resource, kubearmorLog.Labels,
		kubearmorLog.Data, kubearmorLog.Category, kubearmorLog.Action, kubearmorLog.Result)
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
			kubearmorLog.UID, kubearmorLog.Type, kubearmorLog.Source, kubearmorLog.Operation, kubearmorLog.Resource, kubearmorLog.Labels,
			kubearmorLog.Data, kubearmorLog.Category, kubearmorLog.Action, kubearmorLog.Timestamp, kubearmorLog.Timestamp, kubearmorLog.Result)

		defer statement.Close()
	}
}
