package kubearmor

import (
	"encoding/json"

	"github.com/accuknox/observability/src/types"
	kubearmor "github.com/kubearmor/KubeArmor/protobuf"
)

//Convert Kubearmor Logs and Alert in Generic Log Struct
func convertLog(client interface{}) types.KubeArmorLogAlert {
	var kubeArmorLogAlert types.KubeArmorLogAlert
	switch stream := client.(type) {
	case kubearmor.LogService_WatchLogsClient:
		logs, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving kubearmor log " + err.Error())
			return kubeArmorLogAlert
		}
		jsonLog, _ := json.Marshal(logs)
		err = json.Unmarshal(jsonLog, &kubeArmorLogAlert)
		kubeArmorLogAlert.Action = "Allow"
		kubeArmorLogAlert.Category = "Log"

	case kubearmor.LogService_WatchAlertsClient:
		logs, err := stream.Recv()
		if err != nil {
			log.Error().Msg("Error in receiving kubearmor alert " + err.Error())
			return kubeArmorLogAlert
		}

		jsonLog, _ := json.Marshal(logs)
		err = json.Unmarshal(jsonLog, &kubeArmorLogAlert)
		kubeArmorLogAlert.Category = "Alert"

	}

	return kubeArmorLogAlert

}
