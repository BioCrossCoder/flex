package linkedlist

import (
	"encoding/json"
	"fmt"
)

// String returns the string representation of the linked list
func (l LinkedList) String() string {
	return fmt.Sprint(l.ToArray())
}

// MarshalJSON converts the linked list to JSON format
func (l LinkedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToArray())
}

// UnmarshalJSON populates the linked list from JSON-formatted data
func (l *LinkedList) UnmarshalJSON(data []byte) (err error) {
	var arr []any
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	*l = *NewLinkedList(arr...)
	return
}
