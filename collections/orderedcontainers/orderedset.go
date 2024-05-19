package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/biocrosscoder/flex/collections/arraylist"
	"github.com/biocrosscoder/flex/collections/set"
	"github.com/biocrosscoder/flex/common"
	"strings"
)

// OrderedSet represents a set data structure that maintains the order of elements, which is based on a Set and an ArrayList.
type OrderedSet struct {
	set.Set
	sequence arraylist.ArrayList
}

// NewOrderedSet creates a new OrderedSet with the given initial entries
func NewOrderedSet(entries ...any) *OrderedSet {
	elements := set.Of()
	sequence := arraylist.Of()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Push(entry)
	}
	return &OrderedSet{elements, sequence}
}

// Add adds the given element to the OrderedSet
func (s *OrderedSet) Add(element any) *OrderedSet {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Push(element)
	}
	return s
}

// Discard removes the given element from the OrderedSet
func (s *OrderedSet) Discard(element any) bool {
	if s.Set.Discard(element) {
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

// Clear removes all elements from the OrderedSet
func (s *OrderedSet) Clear() *OrderedSet {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

// Update merges the elements from another OrderedSet into the current OrderedSet
func (s *OrderedSet) Update(another OrderedSet) *OrderedSet {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

// Pop removes and returns the last element from the OrderedSet
func (s *OrderedSet) Pop() (element any, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

// Elements returns a slice containing all the elements of the OrderedSet
func (s OrderedSet) Elements() []any {
	return s.sequence.Copy()
}

// Copy returns a copy of the current OrderedSet
func (s OrderedSet) Copy() OrderedSet {
	return OrderedSet{
		s.Set.Copy(),
		s.sequence.Copy(),
	}
}

// Equal checks if the current OrderedSet is equal to another OrderedSet
func (s OrderedSet) Equal(another OrderedSet) bool {
	return s.sequence.Equal(another.sequence)
}

// At returns the element at the specified index in the OrderedSet
func (s OrderedSet) At(index int) (any, error) {
	return s.sequence.At(index)
}

// IndexOf returns the index of the specified element in the OrderedSet
func (s OrderedSet) IndexOf(element any) int {
	return s.sequence.IndexOf(element)
}

// ToList returns a copy of the sequence as an ArrayList
func (s OrderedSet) ToList() arraylist.ArrayList {
	return s.sequence.Copy()
}

// String returns the string representation of the sequence in the OrderedSet
func (s OrderedSet) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.sequence))
}

// MarshalJSON converts the sequence in the OrderedSet to JSON
func (s OrderedSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.sequence)
}

// UnmarshalJSON decodes the JSON data to populate the OrderedSet sequence
func (s *OrderedSet) UnmarshalJSON(data []byte) (err error) {
	var arr []any
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	s.Set = set.Of(arr...)
	s.sequence = arraylist.Of(arr...)
	return
}
