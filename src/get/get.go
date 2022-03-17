package get

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/accuknox/observability/src/types"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
)

//GetKubearmor - To fetch all the kubearmor aggregated logs
func GetKubearmor(option types.KubeArmorFilter, limit int) ([]types.KubeArmor, error) {

	var result []types.KubeArmor
	//query to fetch all logs
	query := constants.SELECT_ALL_KUBEARMOR
	//Check option exist
	if option.Namespace != "" {
		query = query + constants.WHERE_NAMESPACE_NAME + option.Namespace + constants.QUOTATION
	}

	//Fetch based on lastest updated
	query = query + constants.ORDER_BY_UPDATED_TIME
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
			&kubeArmor.NamespaceName, &kubeArmor.PodName, &kubeArmor.ContainerID, &kubeArmor.ContainerName,
			&kubeArmor.UID, &kubeArmor.Type, &kubeArmor.Source, &kubeArmor.Operation, &kubeArmor.Resource,
			&kubeArmor.Data, &kubeArmor.StartTime, &kubeArmor.UpdatedTime, &kubeArmor.Total); err != nil {
			return result, err
		}
		//append record
		result = append(result, kubeArmor)
	}
	return result, nil
}

//GetCilium - To fetch all the cilium aggregated logs
func GetCilium(option types.CiliumFilter, limit int) ([]types.Cilium, error) {

	var result []types.Cilium
	//query to fetch all logs
	query := constants.SELECT_ALL_CILIUM
	if !reflect.DeepEqual(option, (types.CiliumFilter{})) {
		var flowType, verdict, direction int
		var optionQuery string
		//Check Type exist
		if option.Type != "" {
			switch strings.ToUpper(option.Type) {
			case constants.L3_L4:
				flowType = 1
			case constants.L7:
				flowType = 2
			default:
				return nil, errors.New(constants.INCORRECT_TYPE)
			}
			optionQuery = constants.TYPE + fmt.Sprint(flowType)
		}
		//Check Verdict exist
		if option.Verdict != "" {
			switch strings.ToUpper(option.Verdict) {
			case constants.FORWARDED:
				verdict = 1
			case constants.DROPPED:
				verdict = 2
			case constants.ERROR:
				verdict = 3
			case constants.AUDIT:
				verdict = 4
			default:
				return nil, errors.New(constants.INCORRECT_VERDICT)
			}
			//Check option query exist
			if optionQuery != "" {
				optionQuery = optionQuery + constants.AND + constants.VERDICT + fmt.Sprint(verdict)
			} else {
				optionQuery = constants.VERDICT + fmt.Sprint(verdict)
			}
		}
		//Check Direction exist
		if option.Direction != "" {
			switch strings.ToUpper(option.Direction) {
			case constants.INGRESS:
				direction = 1
			case constants.EGRESS:
				direction = 2
			default:
				return nil, errors.New(constants.INCORRECT_DIRECTION)
			}
			//Check option query exist
			if optionQuery != "" {
				optionQuery = optionQuery + constants.AND + constants.TRAFFIC_DIRECTION + fmt.Sprint(direction)
			} else {
				optionQuery = constants.TRAFFIC_DIRECTION + fmt.Sprint(direction)
			}
		}
		query = query + constants.WHERE + optionQuery

	}
	//Fetch based on lastest updated
	query = query + constants.ORDER_BY_UPDATED_TIME
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
		var cilium types.Cilium
		//Scan the record
		if err := rows.Scan(&cilium.Verdict, &cilium.EthernetSource, &cilium.EthernetDestination,
			&cilium.IpSource, &cilium.IpDestination, &cilium.IpVersion, &cilium.IpEncrypted,
			&cilium.L4TCPSourcePort, &cilium.L4TCPDestinationPort, &cilium.L4UDPSourcePort, &cilium.L4UDPDestinationPort,
			&cilium.L4ICMPv4Type, &cilium.L4ICMPv4Code, &cilium.L4ICMPv6Type, &cilium.L4ICMPv6Code,
			&cilium.SourceID, &cilium.SourceIdentity, &cilium.SourceNamespace, &cilium.SourceLabels, &cilium.SourcePodName,
			&cilium.DestinationID, &cilium.DestinationIdentity, &cilium.DestinationNamespace, &cilium.DestinationLabels, &cilium.DestinationPodName,
			&cilium.Type, &cilium.NodeName, &cilium.SourceNames, &cilium.DestinationNames, &cilium.L7Type, &cilium.L7LatencyNs,
			&cilium.L7DnsQuery, &cilium.L7DnsIps, &cilium.L7DnsTtl, &cilium.L7DnsCnames, &cilium.L7DnsObservationsource, &cilium.L7DnsRcode, &cilium.L7DnsQtypes, &cilium.L7DnsRrtypes,
			&cilium.L7HttpCode, &cilium.L7HttpMethod, &cilium.L7HttpUrl, &cilium.L7HttpProtocol, &cilium.L7HttpHeaders,
			&cilium.L7KafkaErrorCode, &cilium.L7KafkaApiVersion, &cilium.L7KafkaApikey, &cilium.L7KafkaCorrelationID, &cilium.L7KafkaTopic,
			&cilium.EventTypeType, &cilium.EventTypeSubType, &cilium.SourceServiceName, &cilium.SourceServiceNamespace, &cilium.DestinationServiceName, &cilium.DestinationServiceNamespace,
			&cilium.TrafficDirection, &cilium.PolicyMatchType, &cilium.TraceObservationPoint, &cilium.DropReasonDesc,
			&cilium.IsReply, &cilium.DebugCapturePoint, &cilium.InterfaceIndex, &cilium.InterfaceName, &cilium.ProxyPort,
			&cilium.StartTime, &cilium.UpdatedTime, &cilium.Total); err != nil {
			return result, err
		}
		//append record
		result = append(result, cilium)
	}
	return result, nil
}
