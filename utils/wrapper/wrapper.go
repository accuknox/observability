package wrapper

import (
	"fmt"
	"strings"

	"github.com/accuknox/observability/utils/constants"
)

//ConvertArrayToString - Convert Array of string to String
func ConvertArrayToString(arr []string) string {
	return strings.Join(arr, ", ")

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
