package dict

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	convey.Convey("jsonify and stringify", t, func() {
		d := Dict[string, int]{"a": 1, "b": 2, "c": 3}
		data, err := json.Marshal(d)
		assert.Nil(t, err)
		d2 := Dict[string, int]{}
		assert.Nil(t, json.Unmarshal(data, &d2))
		assert.Equal(t, fmt.Sprint(d), fmt.Sprint(d2))
	})
}

func ExampleDict_MarshalJSON() {
	dict := Dict[string, any]{"name": "John", "age": 30}
	jsonBytes, _ := json.Marshal(dict)
	fmt.Println(string(jsonBytes))
	// Output: {"age":30,"name":"John"}
}

func ExampleDict_UnmarshalJSON() {
	jsonStr := []byte(`{"age":30,"name":"John"}`)
	var dict Dict[string, any]
	_ = json.Unmarshal(jsonStr, &dict)
	fmt.Println(dict)
	// Output: map[age:30 name:John]
}
