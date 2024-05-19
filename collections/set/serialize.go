package set

import (
	"encoding/json"
	"fmt"
	"strings"
)

// elements returns all the elements in the set as a slice.
func (s Set) elements() []any {
	arr := make([]any, s.Size())
	i := 0
	for e := range s {
		arr[i] = e
		i++
	}
	return arr
}

// String returns the string representation of the set.
func (s Set) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.elements()))
}

// MarshalJSON customizes the JSON encoding for the set.
func (s Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.elements())
}

// UnmarshalJSON customizes the JSON decoding for the set.
func (s *Set) UnmarshalJSON(data []byte) (err error) {
	var arr []any
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*s = Of(arr...)
	return
}
