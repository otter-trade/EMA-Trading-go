package fetch

import (
	"encoding/json"
	"fmt"
)

func MapToJson(param map[string]any) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func ToStr(p any) string {
	returnStr := ""
	switch p := p.(type) {
	case []int32:
		returnStr = string(p)
	case []uint8:
		returnStr = string(p)
	case int32:
		returnStr = string(p)
	case uint8:
		returnStr = string(p)
	default:
		returnStr = fmt.Sprintf("%+v", p)
	}

	return returnStr
}
