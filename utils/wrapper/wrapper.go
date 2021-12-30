package wrapper

import "strings"

//ConvertArrayToString - Convert Array of string to String
func ConvertArrayToString(arr []string) string {
	return strings.Join(arr, ", ")

}

//ConvertStringToArray - Convert String to Array of string
func ConvertStringToArray(str string) []string {
	return strings.Split(str, ",")
}
