package arraylist

import (
	"flex/common"
)

type ArrayList[T any] []T

func (l ArrayList[T]) parseIndex(index int) int {
	length := l.Len()
	if index < 0 {
		index += length
		if index < 0 {
			return 0
		}
	} else if index > length {
		return length
	}
	return index
}

func (l ArrayList[T]) isIndexValid(index int) (err error) {
	if index < 0 || index >= l.Len() {
		err = common.ErrOutOfRange
	}
	return
}

func (l ArrayList[T]) Len() int {
	return len(l)
}

func (l ArrayList[T]) Count(element T) (count int) {
	for _, item := range l {
		if common.Equal(item, element) {
			count++
		}
	}
	return
}

func (l ArrayList[T]) Includes(element T) bool {
	return l.IndexOf(element) != -1
}

func (l ArrayList[T]) Empty() bool {
	return l.Len() == 0
}