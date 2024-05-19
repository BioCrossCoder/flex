package dict

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON converts the dictionary to JSON format.
func (d Dict) MarshalJSON() ([]byte, error) {
	result := make(map[string]any)
	for k, v := range d {
		result[fmt.Sprint(k)] = v
	}
	return json.Marshal(result)
}

// UnmarshalJSON converts the JSON data to a dictionary.
func (d *Dict) UnmarshalJSON(data []byte) (err error) {
	var m map[string]any
	err = json.Unmarshal(data, &m)
	if err != nil {
		return
	}
	_ = d.Clear()
	for k, v := range m {
		_ = d.Set(k, v)
	}
	return
}
