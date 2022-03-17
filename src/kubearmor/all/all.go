package all

import (
	"strconv"

	"github.com/accuknox/observability/src/types"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
	"github.com/accuknox/observability/utils/wrapper"
)

//All - To fetch all the kubearmor aggregated logs
func All(option types.KubeArmorFilter, limit int) ([]types.KubeArmor, error) {

	var result []types.KubeArmor
	//query to fetch all logs
	query := constants.SELECT_ALL_KUBEARMOR
	//Check option exist
	if len(option.Operation) != 0 {
		query = query + ` where operation in (` + wrapper.ConvertFilterString(option.Operation) + `)`
	}
	//Check limit exist
	if limit != 0 {
		//query to fetch all logs with limit
		query = query + constants.LIMIT + strconv.Itoa(limit)
	}
	//Fetch rows
	rows, err := database.ConnectDB().Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var kubeArmor types.KubeArmor
		//Scan the record
		if err := rows.Scan(&kubeArmor.ClusterName, &kubeArmor.HostName,
			&kubeArmor.NamespaceName, &kubeArmor.PodName, &kubeArmor.ContainerID, &kubeArmor.ContainerName, &kubeArmor.Total); err != nil {
			return result, err
		}
		//append record
		result = append(result, kubeArmor)
	}
	return result, nil
}
