package utils

import (
	"encoding/json"
	"fmt"
)

// ToObject parses a JSON string into the given target object (pointer)
func ToObject(jsonStr string, target interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), target); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	return nil
}

// ToString converts a Go object into a JSON string
func ToString(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("failed to serialize object: %w", err)
	}
	return string(bytes), nil
}
