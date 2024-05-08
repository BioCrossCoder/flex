package sortedlist

import (
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
	for index >= 0 && l.cmp(l.elements[index], element) == 0 {
		index--
	}
	index++
	i := index
	for count > 0 && i < l.Len() && l.cmp(l.elements[i], element) == 0 {
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
		for index < l.Len() && l.cmp(l.elements[index], element) == 0 {
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

func (l *SortedList[T]) RemoveIf(condition func(T) bool, counts ...int) SortedList[T] {
	return SortedList[T]{
		l.elements.RemoveIf(condition, counts...),
		l.cmp,
	}
}

func (l *SortedList[T]) RemoveRightIf(condition func(T) bool, counts ...int) SortedList[T] {
	return SortedList[T]{
		l.elements.RemoveRightIf(condition, counts...),
		l.cmp,
	}
}
