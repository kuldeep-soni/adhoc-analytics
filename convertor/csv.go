package convertor

import (
	"encoding/json"
	"reflect"
	"strings"
)

func ToCSV(data interface{}) string {
	v := reflect.ValueOf(data)
	sb := strings.Builder{}
	for i := 0; i < v.Type().NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		value := v.Field(i).Interface()
		vB , _ := json.Marshal(value)
		sb.WriteString(fieldName)
		sb.WriteString(",")
		sb.Write(vB)
		sb.WriteString(",")
	}
	return sb.String()
}
