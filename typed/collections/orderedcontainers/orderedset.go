package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"github.com/biocrosscoder/flex/typed/collections/set"
	"strings"
)

// OrderedSet represents a set with elements in a specific order.
type OrderedSet[T comparable] struct {
	set.Set[T]
	sequence arraylist.ArrayList[T]
}

// NewOrderedSet creates a new OrderedSet with the provided elements, eliminating duplicates and maintaining the order of insertion.
func NewOrderedSet[T comparable](entries ...T) *OrderedSet[T] {
	elements := set.Of[T]()
	sequence := arraylist.Of[T]()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Push(entry)
	}
	return &OrderedSet[T]{elements, sequence}
}

// Add adds the specified element to the OrderedSet if it does not exist and returns the updated OrderedSet.
func (s *OrderedSet[T]) Add(element T) *OrderedSet[T] {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Push(element)
	}
	return s
}

// Discard removes the specified element from the OrderedSet and returns true if the removal was successful, false otherwise.
func (s *OrderedSet[T]) Discard(element T) bool {
	if s.Set.Discard(element) {
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

// Clear removes all elements from the OrderedSet and returns the updated empty set.
func (s *OrderedSet[T]) Clear() *OrderedSet[T] {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

// Update adds all elements from another OrderedSet to the current OrderedSet and returns the updated set.
func (s *OrderedSet[T]) Update(another OrderedSet[T]) *OrderedSet[T] {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

// Pop removes and returns the last element from the OrderedSet and any error encountered during the operation.
func (s *OrderedSet[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

// Elements returns a copy of the elements in the OrderedSet as a slice.
func (s OrderedSet[T]) Elements() []T {
	return s.sequence.Copy()
}

// Copy creates a new copy of the OrderedSet with the same elements and sequence.
func (s OrderedSet[T]) Copy() OrderedSet[T] {
	return OrderedSet[T]{
		s.Set.Copy(),
		s.sequence.Copy(),
	}
}

// Equal checks if the elements and their order in the current and another OrderedSet are the same.
func (s OrderedSet[T]) Equal(another OrderedSet[T]) bool {
	return s.sequence.Equal(another.sequence)
}

// At returns the element at the specified index in the sequence and any error encountered during the operation.
func (s OrderedSet[T]) At(index int) (T, error) {
	return s.sequence.At(index)
}

// IndexOf returns the first index of the specified element in the sequence, or -1 if the element is not found.
func (s OrderedSet[T]) IndexOf(element T) int {
	return s.sequence.IndexOf(element)
}

// ToList returns a copy of the sequence in the form of an ArrayList.
func (s OrderedSet[T]) ToList() arraylist.ArrayList[T] {
	return s.sequence.Copy()
}

// String returns the string representation of the sequence within the OrderedSet.
func (s OrderedSet[T]) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.sequence))
}

// MarshalJSON returns the JSON encoding of the sequence within the OrderedSet.
func (s OrderedSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.sequence)
}

// UnmarshalJSON sets the elements and sequence in the OrderedSet based on the provided JSON data.
func (s *OrderedSet[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	s.Set = set.Of(arr...)
	s.sequence = arraylist.Of(arr...)
	return
}
