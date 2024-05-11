package arraylist

import (
	"github.com/biocrosscoder/flex/common"
	"slices"
)

func (l ArrayList[T]) IndexOf(element T) (index int) {
	index = -1
	for i, item := range l {
		if common.Equal(item, element) {
			index = i
			break
		}
	}
	return
}

func (l ArrayList[T]) LastIndexOf(element T) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if common.Equal(l[i], element) {
			index = i
			break
		}
	}
	return
}

func (l ArrayList[T]) At(index int) (element T, err error) {
	if index < 0 {
		index += l.Len()
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	element = l[index]
	return
}

func (l ArrayList[T]) Find(by func(T) bool) (element T, found bool) {
	index := l.FindIndex(by)
	if index != -1 {
		found = true
		element = l[index]
	}
	return
}

func (l ArrayList[T]) FindIndex(by func(T) bool) (index int) {
	return slices.IndexFunc(l, by)
}

func (l ArrayList[T]) FindLast(by func(T) bool) (element T, found bool) {
	index := l.FindLastIndex(by)
	if index != -1 {
		found = true
		element = l[index]
	}
	return
}

func (l ArrayList[T]) FindLastIndex(by func(T) bool) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if by(l[i]) {
			index = i
			break
		}
	}
	return
}

func (l ArrayList[T]) Head() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[0]
	return
}

func (l ArrayList[T]) Tail() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[l.Len()-1]
	return
}

func (l ArrayList[T]) FindIndexes(by func(T) bool, counts ...int) (indexes []int) {
	count := l.searchCount(counts...)
	indexes = make([]int, 0)
	for i, item := range l {
		if count == 0 {
			break
		}
		if by(item) {
			indexes = append(indexes, i)
			count--
		}
	}
	return
}

func (l ArrayList[T]) FindLastIndexes(by func(T) bool, counts ...int) (indexes []int) {
	count := l.searchCount(counts...)
	indexes = make([]int, 0)
	for i := l.Len() - 1; i >= 0; i-- {
		if count == 0 {
			break
		}
		if by(l[i]) {
			indexes = append(indexes, i)
			count--
		}
	}
	return
}

func (l ArrayList[T]) Finds(by func(T) bool, counts ...int) (elements []T) {
	count := l.searchCount(counts...)
	elements = make([]T, 0)
	for _, item := range l {
		if count == 0 {
			break
		}
		if by(item) {
			elements = append(elements, item)
			count--
		}
	}
	return
}

func (l ArrayList[T]) FindLasts(by func(T) bool, counts ...int) (elements []T) {
	count := l.searchCount(counts...)
	elements = make([]T, 0)
	for i := l.Len() - 1; i >= 0; i-- {
		if count == 0 {
			break
		}
		if by(l[i]) {
			elements = append(elements, l[i])
			count--
		}
	}
	return
}

func (l ArrayList[T]) searchCount(counts ...int) int {
	if len(counts) == 0 || counts[0] <= 0 {
		return l.Len()
	}
	return counts[0]
}
