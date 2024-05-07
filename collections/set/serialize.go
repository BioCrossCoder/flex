package set

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (s Set) elements() []any {
	arr := make([]any, s.Size())
	i := 0
	for e := range s {
		arr[i] = e
		i++
	}
	return arr
}

func (s Set) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.elements()))
}

func (s Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.elements())
}

func (s *Set) UnmarshalJSON(data []byte) (err error) {
	var arr []any
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*s = Of(arr...)
	return
}
