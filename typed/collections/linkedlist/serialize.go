package linkedlist

import (
	"encoding/json"
	"fmt"
)

func (l LinkedList[T]) String() string {
	return fmt.Sprint(l.ToArray())
}

func (l LinkedList[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToArray())
}

func (l *LinkedList[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*l = *NewLinkedList(arr...)
	return
}
