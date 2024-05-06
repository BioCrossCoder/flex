package sortedlist

import (
	"flex/common"
	"slices"
)

func (l SortedList[T]) parseCount(counts ...int) int {
	if len(counts) == 0 {
		return 1
	}
	if counts[0] <= 0 {
		return l.Len()
	}
	return counts[0]
}

func (l *SortedList[T]) Remove(element T, counts ...int) *SortedList[T] {
	count := l.parseCount(counts...)
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return l
	}
	for index >= 0 && l.elements[index] == element {
		index--
	}
	index++
	i := index
	for count > 0 && i < l.Len() && common.Equal(l.elements[i], element) {
		count--
		i++
	}
	_ = l.RemoveRange(index, i)
	return l
}

func (l *SortedList[T]) RemoveRange(start, end int) SortedList[T] {
	return SortedList[T]{
		l.elements.Splice(start, end-start),
		l.cmp,
	}
}

func (l *SortedList[T]) Clear() *SortedList[T] {
	_ = l.elements.Clear()
	return l
}

func (l *SortedList[T]) Pop(indexes ...int) (element T, err error) {
	return l.elements.Pop(indexes...)
}

func (l *SortedList[T]) Insert(element T) *SortedList[T] {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if exist {
		for index < l.Len() && l.elements[index] == element {
			index++
		}
	}
	_ = l.elements.Insert(index, element)
	return l
}

func (l *SortedList[T]) Reverse() *SortedList[T] {
	_ = l.elements.Reverse()
	l.cmp = func(a, b T) int {
		return -l.cmp(a, b)
	}
	return l
}