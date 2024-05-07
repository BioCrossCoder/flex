package linkedlist

import (
	"encoding/json"
	"fmt"
)

func (l LinkedList) String() string {
	return fmt.Sprint(l.ToArray())
}

func (l LinkedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToArray())
}

func (l *LinkedList) UnmarshalJSON(data []byte) (err error) {
	var arr []any
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*l = *NewLinkedList(arr...)
	return
}
