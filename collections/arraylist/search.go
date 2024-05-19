package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// IndexOf returns the index of the first occurrence of the specified element in the list, or -1 if the element is not found.
func (l ArrayList) IndexOf(element any) (index int) {
	index = -1
	for i, item := range l {
		if common.Equal(item, element) {
			index = i
			break
		}
	}
	return
}

// LastIndexOf returns the index of the last occurrence of the specified element in the list, or -1 if the element is not found.
func (l ArrayList) LastIndexOf(element any) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if common.Equal(l[i], element) {
			index = i
			break
		}
	}
	return
}

// At returns the element at the specified index in the list, or an error if the index is out of range.
func (l ArrayList) At(index int) (element any, err error) {
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

// Find returns the first element in the list that satisfies the provided function.
func (l ArrayList) Find(by func(any) bool) (element any, found bool) {
	index := l.FindIndex(by)
	if index != -1 {
		found = true
		element = l[index]
	}
	return
}

// FindIndex returns the index of the first element in the list that satisfies the provided function, or -1 if no such element is found.
func (l ArrayList) FindIndex(by func(any) bool) (index int) {
	index = -1
	for i, item := range l {
		if by(item) {
			index = i
			break
		}
	}
	return
}

// FindLast returns the last element in the list that satisfies the provided function.
func (l ArrayList) FindLast(by func(any) bool) (element any, found bool) {
	index := l.FindLastIndex(by)
	if index != -1 {
		found = true
		element = l[index]
	}
	return
}

// FindLastIndex returns the index of the last element in the list that satisfies the provided function, or -1 if no such element is found.
func (l ArrayList) FindLastIndex(by func(any) bool) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if by(l[i]) {
			index = i
			break
		}
	}
	return
}

// Head returns the first element of the list, or an error if the list is empty.
func (l ArrayList) Head() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[0]
	return
}

// Tail returns the last element of the list, or an error if the list is empty.
func (l ArrayList) Tail() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[l.Len()-1]
	return
}

// FindIndexes returns the indexes of elements in the list that satisfy the provided function, with a maximum count for each occurrence.
func (l ArrayList) FindIndexes(by func(any) bool, counts ...int) (indexes []int) {
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

// FindLastIndexes returns the indexes of elements in the list that satisfy the provided function, with a maximum count for each occurrence, in reverse order.
func (l ArrayList) FindLastIndexes(by func(any) bool, counts ...int) (indexes []int) {
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

// Finds returns the elements in the list that satisfy the provided function, with a maximum count for each occurrence.
func (l ArrayList) Finds(by func(any) bool, counts ...int) (elements []any) {
	count := l.searchCount(counts...)
	elements = make([]any, 0)
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

// FindLasts returns the elements in the list that satisfy the provided function, with a maximum count for each occurrence, in reverse order.
func (l ArrayList) FindLasts(by func(any) bool, counts ...int) (elements []any) {
	count := l.searchCount(counts...)
	elements = make([]any, 0)
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

// searchCount calculates the maximum count for each occurrence based on the given counts and the length of the list.
func (l ArrayList) searchCount(counts ...int) int {
	return list.SearchCount(l.Len(), counts...)
}
