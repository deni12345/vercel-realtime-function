package utils

import (
	"encoding/json"
	"fmt"
)

func ToString(value interface{}) string {
	indentedJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Sprintf("error marshaling to JSON: %v", err)
	}
	return string(indentedJSON)
}
