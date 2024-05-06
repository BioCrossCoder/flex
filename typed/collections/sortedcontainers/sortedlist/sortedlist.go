package sortedlist

import (
	"cmp"
	"flex/common"
	"flex/typed/collections/arraylist"
	"slices"
)

type SortedList[T cmp.Ordered] struct {
	elements arraylist.ArrayList[T]
	cmp      func(a, b T) int
}

func AscendOrder[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

func DescendOrder[T cmp.Ordered](a, b T) int {
	return -cmp.Compare(a, b)
}

func NewSortedList[T cmp.Ordered](cmp func(a, b T) int, elements ...T) *SortedList[T] {
	arr := arraylist.Of(elements...)
	if cmp == nil {
		cmp = AscendOrder[T]
	}
	slices.SortFunc(arr, cmp)
	return &SortedList[T]{arr, cmp}
}

func (l SortedList[T]) Len() int {
	return l.elements.Len()
}

func (l SortedList[T]) Count(element T) int {
	return l.elements.Count(element)
}

func (l SortedList[T]) Includes(element T) bool {
	return l.elements.Includes(element)
}

func (l SortedList[T]) Empty() bool {
	return l.elements.Empty()
}

func (l SortedList[T]) Copy() SortedList[T] {
	return SortedList[T]{l.elements.Copy(), l.cmp}
}

func (l SortedList[T]) Slice(args ...int) SortedList[T] {
	f := l.cmp
	if len(args) >= 3 && args[2] < 0 {
		f = func(a, b T) int {
			return -l.cmp(a, b)
		}
	}
	return SortedList[T]{l.elements.Slice(args...), f}
}

func (l SortedList[T]) ToReversed() SortedList[T] {
	list := l.Copy()
	_ = list.Reverse()
	return list
}

func (l SortedList[T]) ToArray() []T {
	return l.elements.Copy()
}

func (l SortedList[T]) Equal(another SortedList[T]) bool {
	for i, v := range l.elements {
		if !common.Equal(v, another.elements[i]) {
			return false
		}
	}
	return true
}

func (l SortedList[T]) ToList() arraylist.ArrayList[T] {
	return l.elements.Copy()
}
