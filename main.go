package easyjson

import (
	"encoding/json"
	"fmt"
)

// ParseString parses a JSON string and returns a map[string]interface{}.
func ParseString(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetString gets a string value from a JSON map.
func GetString(data map[string]interface{}, key string) (string, error) {
	value, ok := data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	strValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("value for key %s is not a string", key)
	}
	return strValue, nil
}

// GetInt gets an integer value from a JSON map.
func GetInt(data map[string]interface{}, key string) (int, error) {
	value, ok := data[key]
	if !ok {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	intValue, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("value for key %s is not an integer", key)
	}
	return int(intValue), nil
}

// GetArray gets an array value from a JSON map.
func GetArray(data map[string]interface{}, key string) ([]interface{}, error) {
	value, ok := data[key]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	arrayValue, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("value for key %s is not an array", key)
	}
	return arrayValue, nil
}

// GetObject gets an object value from a JSON map.
func GetObject(data map[string]interface{}, key string) (map[string]interface{}, error) {
	value, ok := data[key]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	objectValue, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("value for key %s is not an object", key)
	}
	return objectValue, nil
}
