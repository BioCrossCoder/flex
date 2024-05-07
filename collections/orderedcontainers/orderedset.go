package orderedcontainers

import (
	"encoding/json"
	"flex/collections/arraylist"
	"flex/collections/set"
	"flex/common"
	"fmt"
	"strings"
)

type OrderedSet struct {
	set.Set
	sequence arraylist.ArrayList
}

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

func (s *OrderedSet) Add(element any) *OrderedSet {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Push(element)
	}
	return s
}

func (s *OrderedSet) Discard(element any) bool {
	if s.Set.Discard(element) {
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

func (s *OrderedSet) Clear() *OrderedSet {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

func (s *OrderedSet) Update(another OrderedSet) *OrderedSet {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

func (s *OrderedSet) Pop() (element any, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

func (s OrderedSet) Elements() []any {
	return s.sequence.Copy()
}

func (s OrderedSet) Copy() OrderedSet {
	return OrderedSet{
		s.Set.Copy(),
		s.sequence.Copy(),
	}
}

func (s OrderedSet) Equal(another OrderedSet) bool {
	return s.sequence.Equal(another.sequence)
}

func (s OrderedSet) At(index int) (any, error) {
	return s.sequence.At(index)
}

func (s OrderedSet) IndexOf(element any) int {
	return s.sequence.IndexOf(element)
}

func (s OrderedSet) ToList() arraylist.ArrayList {
	return s.sequence.Copy()
}

func (s OrderedSet) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.sequence))
}

func (s OrderedSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.sequence)
}

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
