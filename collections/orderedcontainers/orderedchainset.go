package orderedcontainers

import (
	"encoding/json"
	"github.com/biocrosscoder/flex/collections/linkedlist"
	"github.com/biocrosscoder/flex/collections/set"
	"github.com/biocrosscoder/flex/common"
	"strings"
)

// OrderedChainSet represents a set data structure with elements kept in insertion order, which is based on a Set and a LinkedList.
type OrderedChainSet struct {
	set.Set
	sequence *linkedlist.LinkedList
}

// NewOrderedChainSet creates and initializes a new OrderedChainSet with the specified elements.
func NewOrderedChainSet(entries ...any) *OrderedChainSet {
	elements := set.Of()
	sequence := linkedlist.NewLinkedList()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Append(entry)
	}
	return &OrderedChainSet{elements, sequence}
}

// Add adds the specified element to the OrderedChainSet if it is not already present.
func (s *OrderedChainSet) Add(element any) *OrderedChainSet {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Append(element)
	}
	return s
}

// Discard removes the specified element from the OrderedChainSet.
func (s *OrderedChainSet) Discard(element any) bool {
	if s.Has(element) {
		_ = s.Set.Discard(element)
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

// Clear removes all elements from the OrderedChainSet.
func (s *OrderedChainSet) Clear() *OrderedChainSet {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

// Update adds all elements from another OrderedChainSet to this OrderedChainSet.
func (s *OrderedChainSet) Update(another OrderedChainSet) *OrderedChainSet {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

// Pop removes and returns an element from the end of the OrderedChainSet.
func (s *OrderedChainSet) Pop() (element any, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

// Elements returns all elements of the OrderedChainSet in insertion order.
func (s *OrderedChainSet) Elements() []any {
	return s.sequence.ToArray()
}

// Copy returns a new copy of the OrderedChainSet.
func (s *OrderedChainSet) Copy() OrderedChainSet {
	newSeq := s.sequence.Copy()
	return OrderedChainSet{
		s.Set.Copy(),
		&newSeq,
	}
}

// Equal checks if this OrderedChainSet is equal to another OrderedChainSet.
func (s OrderedChainSet) Equal(another OrderedChainSet) bool {
	return s.sequence.Equal(*another.sequence)
}

// At returns the element at the specified index in the insertion order.
func (s OrderedChainSet) At(index int) (any, error) {
	return s.sequence.At(index)
}

// IndexOf returns the index of the specified element in the insertion order.
func (s OrderedChainSet) IndexOf(element any) int {
	return s.sequence.IndexOf(element)
}

// ToList returns a linkedlist.LinkedList containing the elements in the insertion order.
func (s OrderedChainSet) ToList() linkedlist.LinkedList {
	return s.sequence.Copy()
}

// String returns a string representation of the OrderedChainSet.
func (s OrderedChainSet) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(s.sequence.String())
}

// MarshalJSON converts the OrderedChainSet to JSON.
func (s OrderedChainSet) MarshalJSON() ([]byte, error) {
	return s.sequence.MarshalJSON()
}

// UnmarshalJSON parses JSON data and initializes the OrderedChainSet.
func (s *OrderedChainSet) UnmarshalJSON(data []byte) (err error) {
	var arr []any
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	s.Set = set.Of(arr...)
	s.sequence = linkedlist.NewLinkedList(arr...)
	return
}
