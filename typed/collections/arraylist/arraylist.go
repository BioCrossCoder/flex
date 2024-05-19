package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

type ArrayList[T any] []T

func (l ArrayList[T]) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

func (l ArrayList[T]) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

func (l ArrayList[T]) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
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

func (l ArrayList[T]) Equal(another ArrayList[T]) bool {
	if l.Len() != another.Len() {
		return false
	}
	for i := 0; i < l.Len(); i++ {
		if !common.Equal(l[i], another[i]) {
			return false
		}
	}
	return true
}
