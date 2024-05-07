package dict

import (
	"encoding/json"
	"fmt"
)

func (d Dict[K, V]) MarshalJSON() ([]byte, error) {
	result := make(map[string]V)
	for k, v := range d {
		result[fmt.Sprint(k)] = v
	}
	return json.Marshal(result)
}

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
