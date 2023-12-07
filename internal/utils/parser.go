package utils

import "encoding/json"

func ParseEntity[T any](data string) (*T, error) {
	var response T

	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
