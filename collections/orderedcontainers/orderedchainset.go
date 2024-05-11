package orderedcontainers

import (
	"encoding/json"
	"github.com/biocrosscoder/flex/collections/linkedlist"
	"github.com/biocrosscoder/flex/collections/set"
	"github.com/biocrosscoder/flex/common"
	"strings"
)

type OrderedChainSet struct {
	set.Set
	sequence *linkedlist.LinkedList
}

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

func (s *OrderedChainSet) Add(element any) *OrderedChainSet {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Append(element)
	}
	return s
}

func (s *OrderedChainSet) Discard(element any) bool {
	if s.Has(element) {
		_ = s.Set.Discard(element)
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

func (s *OrderedChainSet) Clear() *OrderedChainSet {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

func (s *OrderedChainSet) Update(another OrderedChainSet) *OrderedChainSet {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

func (s *OrderedChainSet) Pop() (element any, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

func (s *OrderedChainSet) Elements() []any {
	return s.sequence.ToArray()
}

func (s *OrderedChainSet) Copy() OrderedChainSet {
	newSeq := s.sequence.Copy()
	return OrderedChainSet{
		s.Set.Copy(),
		&newSeq,
	}
}

func (s OrderedChainSet) Equal(another OrderedChainSet) bool {
	return s.sequence.Equal(*another.sequence)
}

func (s OrderedChainSet) At(index int) (any, error) {
	return s.sequence.At(index)
}

func (s OrderedChainSet) IndexOf(element any) int {
	return s.sequence.IndexOf(element)
}

func (s OrderedChainSet) ToList() linkedlist.LinkedList {
	return s.sequence.Copy()
}

func (s OrderedChainSet) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(s.sequence.String())
}

func (s OrderedChainSet) MarshalJSON() ([]byte, error) {
	return s.sequence.MarshalJSON()
}

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
