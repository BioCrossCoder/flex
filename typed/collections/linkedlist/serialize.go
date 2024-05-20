package linkedlist

import (
	"encoding/json"
	"fmt"
)

// String returns a string representation of the linked list
func (l LinkedList[T]) String() string {
	return fmt.Sprint(l.ToArray())
}

// MarshalJSON converts the linked list to JSON
func (l LinkedList[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToArray())
}

// UnmarshalJSON populates the linked list from JSON data
func (l *LinkedList[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*l = *NewLinkedList(arr...)
	return
}
