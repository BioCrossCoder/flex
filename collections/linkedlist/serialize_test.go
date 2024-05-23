package linkedlist

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	convey.Convey("jsonify and stringify", t, func() {
		l := NewLinkedList(1, 2, 3)
		data, err := json.Marshal(l)
		assert.Nil(t, err)
		l2 := NewLinkedList()
		assert.Nil(t, json.Unmarshal(data, l2))
		assert.Equal(t, fmt.Sprint(l), fmt.Sprint(l2))
	})
}

func ExampleLinkedList_String() {
	l := NewLinkedList(1, 2, 3)
	fmt.Println(l)
	// Output: [1 2 3]
}

func ExampleLinkedList_MarshalJSON() {
	l := NewLinkedList(1, 2, 3)
	jsonData, _ := json.Marshal(l)
	fmt.Println(string(jsonData))
	// Output: [1,2,3]
}

func ExampleLinkedList_UnmarshalJSON() {
	jsonData := []byte("[4,5,6]")
	var l LinkedList
	_ = json.Unmarshal(jsonData, &l)
	fmt.Println(l)
	// Output: [4 5 6]
}
