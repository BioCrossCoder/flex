package sortedlist

import "slices"

func (l *SortedList[T]) Remove(element T, counts ...int) *SortedList[T] {
	_ = l.elements.Remove(element, counts...)
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
