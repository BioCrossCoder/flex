package dict

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON converts the dictionary to JSON format.
func (d Dict[K, V]) MarshalJSON() ([]byte, error) {
	result := make(map[string]V)
	for k, v := range d {
		result[fmt.Sprint(k)] = v
	}
	return json.Marshal(result)
}

// UnmarshalJSON parses the JSON data into the dictionary.
func (d *Dict[K, V]) UnmarshalJSON(data []byte) (err error) {
	var m map[K]V
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
