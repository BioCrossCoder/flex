package orderedcontainers

import (
	"encoding/json"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/linkedlist"
	"github.com/biocrosscoder/flex/typed/collections/set"
	"strings"
)

// OrderedChainSet represents a set data structure with maintained insertion order.
type OrderedChainSet[T comparable] struct {
	set.Set[T]
	sequence *linkedlist.LinkedList[T]
}

// NewOrderedChainSet creates a new OrderedChainSet with the given elements and returns a pointer to it.
func NewOrderedChainSet[T comparable](entries ...T) *OrderedChainSet[T] {
	elements := set.Of[T]()
	sequence := linkedlist.NewLinkedList[T]()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Append(entry)
	}
	return &OrderedChainSet[T]{elements, sequence}
}

// Add inserts the given element into the OrderedChainSet.
func (s *OrderedChainSet[T]) Add(element T) *OrderedChainSet[T] {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Append(element)
	}
	return s
}

// Discard removes the given element from the OrderedChainSet.
func (s *OrderedChainSet[T]) Discard(element T) bool {
	if s.Has(element) {
		_ = s.Set.Discard(element)
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

// Clear removes all elements from the OrderedChainSet.
func (s *OrderedChainSet[T]) Clear() *OrderedChainSet[T] {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

// Update adds all elements from another OrderedChainSet to the current OrderedChainSet.
func (s *OrderedChainSet[T]) Update(another OrderedChainSet[T]) *OrderedChainSet[T] {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

// Pop removes and returns the first element from the OrderedChainSet.
func (s *OrderedChainSet[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

// Elements returns a slice of all elements in the OrderedChainSet.
func (s *OrderedChainSet[T]) Elements() []T {
	return s.sequence.ToArray()
}

// Copy creates a shallow copy of the current OrderedChainSet and returns it.
func (s *OrderedChainSet[T]) Copy() OrderedChainSet[T] {
	newSeq := s.sequence.Copy()
	return OrderedChainSet[T]{
		s.Set.Copy(),
		&newSeq,
	}
}

// Equal checks if the current OrderedChainSet is equal to another OrderedChainSet.
func (s OrderedChainSet[T]) Equal(another OrderedChainSet[T]) bool {
	return s.sequence.Equal(*another.sequence)
}

// At returns the element at the specified index in the OrderedChainSet.
func (s OrderedChainSet[T]) At(index int) (T, error) {
	return s.sequence.At(index)
}

// IndexOf returns the index of the given element in the OrderedChainSet.
func (s OrderedChainSet[T]) IndexOf(element T) int {
	return s.sequence.IndexOf(element)
}

// ToList returns a copy of the sequence as a linked list.
func (s OrderedChainSet[T]) ToList() linkedlist.LinkedList[T] {
	return s.sequence.Copy()
}

// String returns the string representation of the OrderedChainSet.
func (s OrderedChainSet[T]) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(s.sequence.String())
}

// MarshalJSON returns the JSON encoding of the OrderedChainSet.
func (s OrderedChainSet[T]) MarshalJSON() ([]byte, error) {
	return s.sequence.MarshalJSON()
}

// UnmarshalJSON parses the JSON-encoded data and updates the OrderedChainSet with the parsed elements.
func (s *OrderedChainSet[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	s.Set = set.Of(arr...)
	s.sequence = linkedlist.NewLinkedList(arr...)
	return
}
