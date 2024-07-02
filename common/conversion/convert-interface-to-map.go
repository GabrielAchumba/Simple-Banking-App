package conversion

import (
	"encoding/json"
)

func ConvertInterfaceToMap(data interface{}) (map[string]interface{}, error) {
	// Attempt type assertion to map[string]interface{}
	if m, ok := data.(map[string]interface{}); ok {
		return m, nil
	}
	// If not already a map, attempt to unmarshal JSON data to a map
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
