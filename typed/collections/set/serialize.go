package set

import (
	"encoding/json"
	"fmt"
	"strings"
)

// elements returns a slice containing all elements in the set.
func (s Set[T]) elements() []T {
	arr := make([]T, s.Size())
	i := 0
	for e := range s {
		arr[i] = e
		i++
	}
	return arr
}

// String returns a string representation of the set.
func (s Set[T]) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.elements()))
}

// MarshalJSON converts the set into JSON format.
func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.elements())
}

// UnmarshalJSON populates the set with data from JSON format.
func (s *Set[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*s = Of(arr...)
	return
}
