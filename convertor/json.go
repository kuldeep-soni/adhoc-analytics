package convertor

import "encoding/json"

func ToJson(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}
