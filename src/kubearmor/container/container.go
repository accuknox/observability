package container

import (
	"strconv"
	"strings"

	"github.com/accuknox/observability/src/types"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
)

//FilterByContainerID - To fetch the kubearmor aggregated logs based on Specific Container ID(s)
func FilterByContainerID(id []string, limit int) ([]types.KubeArmor, error) {

	var result []types.KubeArmor

	//query to fetch logs based on container-id(s)
	query := constants.SELECT_Container_ID_KUBEARMOR
	if limit != 0 {
		//query to fetch logs based on container-id(s) with limit
		query = query + constants.LIMIT + strconv.Itoa(limit)
	}
	//Fetch rows
	rows, err := database.ConnectDB().Query(query, strings.Join(id, ", "))
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

//FilterByContainerName - To fetch the kubearmor aggregated logs based on Specific Container Name(s)
func FilterByContainerName(name []string, limit int) ([]types.KubeArmor, error) {

	var result []types.KubeArmor

	//query to fetch logs based on container-name(s)
	query := constants.SELECT_Container_Name_KUBEARMOR
	if limit != 0 {
		//query to fetch logs based on container-name(s) with limit
		query = query + constants.LIMIT + strconv.Itoa(limit)
	}
	//Fetch rows
	rows, err := database.ConnectDB().Query(query, strings.Join(name, ", "))
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
