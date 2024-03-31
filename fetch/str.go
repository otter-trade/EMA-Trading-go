package fetch

import (
	"encoding/json"
)

func MapToJson(param map[string]any) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
