package set

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (s Set[T]) elements() []T {
	arr := make([]T, s.Size())
	i := 0
	for e := range s {
		arr[i] = e
		i++
	}
	return arr
}

func (s Set[T]) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.elements()))
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.elements())
}

func (s *Set[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*s = Of(arr...)
	return
}
