package utils

import (
	"encoding/json"
)

// Using json tag to move data from struct to struct
func StructToStruct(fromStruct interface{}, toStruct interface{}) error {
	fromStructJson, err := json.Marshal(fromStruct)
	if err != nil {
		return err
	}
	json.Unmarshal(fromStructJson, toStruct)
	return nil
}
