package sortedlist

import (
	"github.com/biocrosscoder/flex/common"
	"slices"
)

func (l SortedList[T]) IndexOf(element T) int {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return -1
	}
	for index >= 0 && l.cmp(l.elements[index], element) == 0 {
		index--
	}
	return index + 1
}

func (l SortedList[T]) LastIndexOf(element T) int {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return -1
	}
	for index < l.Len() && l.cmp(l.elements[index], element) == 0 {
		index++
	}
	return index - 1
}

func (l SortedList[T]) At(index int) (T, error) {
	return l.elements.At(index)
}

func (l SortedList[T]) Find(by func(T) bool) (T, bool) {
	return l.elements.Find(by)
}

func (l SortedList[T]) FindIndex(by func(T) bool) int {
	return l.elements.FindIndex(by)
}

func (l SortedList[T]) FindLast(by func(T) bool) (T, bool) {
	return l.elements.FindLast(by)
}

func (l SortedList[T]) FindLastIndex(by func(T) bool) int {
	return l.elements.FindLastIndex(by)
}

func (l SortedList[T]) Head() (T, error) {
	return l.elements.Head()
}

func (l SortedList[T]) Tail() (T, error) {
	return l.elements.Tail()
}

func (l SortedList[T]) Max() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = slices.MaxFunc([]T{l.elements[0], l.elements[l.Len()-1]}, l.cmp)
	return
}

func (l SortedList[T]) Min() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = slices.MinFunc([]T{l.elements[0], l.elements[l.Len()-1]}, l.cmp)
	return
}
