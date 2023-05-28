package conv

import "encoding/json"

func ToJson(data interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
