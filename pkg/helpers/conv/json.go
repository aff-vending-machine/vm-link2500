package conv

import (
	"encoding/json"
	"fmt"
)

// ToJSONStr converts a struct to a JSON string
func ToJSONStr(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("unable to marshal data: %w", err)
	}

	return string(b), nil
}

// ToMap converts a struct to a map
func ToMap(data interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal data: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal data into map: %w", err)
	}

	return result, nil
}

// ToArray converts a struct to an array
func ToArray[T any](data interface{}) ([]T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal data: %w", err)
	}

	var result []T
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal data into array: %w", err)
	}

	return result, nil
}

// ToStruct converts a struct to another struct
func ToStruct[T any](data interface{}) (*T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal data: %w", err)
	}

	var result T
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal data into target struct: %w", err)
	}

	return &result, nil
}
