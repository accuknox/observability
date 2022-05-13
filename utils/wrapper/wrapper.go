package wrapper

import (
	"fmt"
	"strings"

	"github.com/accuknox/observability/utils/constants"
)

//ConvertArrayToString - Convert Array of string to String
func ConvertArrayToString(arr []string) string {
	var str string
	for _, label := range arr {
		if !strings.HasPrefix(label, "k8s:io.cilium.") {
			if !strings.HasPrefix(label, "k8s:io.kubernetes.") {
				tstr := strings.TrimPrefix(label, "k8s:")
				if str != "" {
					str += constants.COMMA
				}
				str += tstr
			}
		}
	}
	return str

}

//ConvertStringToArray - Convert String to Array of string
func ConvertStringToArray(str string) []string {
	return strings.Split(str, ",")
}

func ConvertFilterString(filter []string) string {
	var query string
	//Create the filter query
	for i, value := range filter {
		query = query + constants.BACK_SLASH + fmt.Sprint(value) + constants.BACK_SLASH
		if len(filter) > i+1 {
			query += constants.COMMA
		}
	}
	return query
}
