package parsers

import (
	"encoding/json"
	"fmt"
)

// ParseJSON decodes JSON data into a struct of type T.
func ParseJSON[T any](data []byte) (*T, error) {
	var result T
	err := json.Unmarshal(data, &result)

	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return &result, nil
}
