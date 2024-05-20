package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
	"slices"
)

// parseCount parses the count argument and returns the actual count to be used.
func (l ArrayList[T]) parseCount(counts ...int) int {
	return list.ParseCount(l.Len(), counts...)
}

// Remove removes the occurrences of the specified element from the list.
// If the count argument is provided, only the first count occurrences will be removed.
// Returns the modified list.
func (l *ArrayList[T]) Remove(element T, counts ...int) *ArrayList[T] {
	count := l.parseCount(counts...)
	array := make(ArrayList[T], l.Len())
	i := 0
	for _, v := range *l {
		if count > 0 && common.Equal(v, element) {
			count--
			continue
		}
		array[i] = v
		i++
	}
	*l = array[:i:i]
	return l
}

// RemoveRight removes the occurrences of the specified element from the right end of the list.
// If the count argument is provided, only the first count occurrences will be removed.
// Returns the modified list.
func (l *ArrayList[T]) RemoveRight(element T, counts ...int) *ArrayList[T] {
	count := l.parseCount(counts...)
	length := l.Len()
	array := make(ArrayList[T], length)
	end := length - 1
	i := end
	for j := end; j >= 0; j-- {
		v := (*l)[j]
		if count > 0 && common.Equal(v, element) {
			count--
			continue
		}
		array[i] = v
		i--
	}
	*l = slices.Clip(array[i+1 : end+1])
	return l
}

// Clear removes all elements from the list.
func (l *ArrayList[T]) Clear() *ArrayList[T] {
	*l = make(ArrayList[T], 0)
	return l
}

// Push adds one or more elements to the end of the list.
func (l *ArrayList[T]) Push(elements ...T) (length int) {
	length = l.Len() + len(elements)
	*l = l.Concat(ArrayList[T](elements))
	return
}

// Pop removes and returns the last element of the list.
// If the indexes argument is provided, the element at the specified index will be removed.
// Returns the removed element and an error if the index is out of range.
func (l *ArrayList[T]) Pop(indexes ...int) (element T, err error) {
	argCount := len(indexes)
	if argCount >= 2 {
		err = common.ErrTooManyArguments
		return
	}
	var index int
	if argCount == 0 {
		index = l.Len() - 1
		element = (*l)[index]
		*l = (*l)[:index]
		return
	}
	index = indexes[0]
	if index < 0 {
		index += l.Len()
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	element = (*l)[index]
	*l = slices.Delete(*l, index, index+1)
	return
}

// Unshift adds one or more elements to the beginning of the list.
func (l *ArrayList[T]) Unshift(elements ...T) (length int) {
	length = l.Len() + len(elements)
	*l = ArrayList[T](elements).Concat(*l)
	return
}

// Shift removes and returns the first element of the list.
func (l *ArrayList[T]) Shift() (element T, err error) {
	return l.Pop(0)
}

// Insert inserts the specified element at the specified index.
// If the index is negative, the index will be calculated from the end of the list.
// Returns the modified list.
func (l *ArrayList[T]) Insert(index int, element T) *ArrayList[T] {
	*l = slices.Insert(*l, l.parseIndex(index), element)
	return l
}

// ForEach applies the specified action to each element of the list.
// Returns the modified list.
func (l *ArrayList[T]) ForEach(action func(T) T) *ArrayList[T] {
	for i, item := range *l {
		(*l)[i] = action(item)
	}
	return l
}

// Replace replaces occurrences of the specified old element with the specified new element.
// If the count argument is not provided, only the first count occurrences will be replaced.
// Returns the modified list.
func (l *ArrayList[T]) Replace(oldElement, newElement T, counts ...int) *ArrayList[T] {
	if common.Equal(oldElement, newElement) {
		return l
	}
	count := l.parseCount(counts...)
	for i, v := range *l {
		if count == 0 {
			break
		}
		if common.Equal(v, oldElement) {
			(*l)[i] = newElement
			count--
		}
	}
	return l
}

// ReplaceRight replaces occurrences of the specified old element with the specified new element from the right end of the list.
// If the count argument is not provided, only the first count occurrences will be replaced.
// Returns the modified list.
func (l *ArrayList[T]) ReplaceRight(oldElement, newElement T, counts ...int) *ArrayList[T] {
	if common.Equal(oldElement, newElement) {
		return l
	}
	count := l.parseCount(counts...)
	for i := l.Len() - 1; i >= 0; i-- {
		if count == 0 {
			break
		}
		if common.Equal((*l)[i], oldElement) {
			(*l)[i] = newElement
			count--
		}
	}
	return l
}

// Splice replaces the specified elements with the specified new elements.
func (l *ArrayList[T]) Splice(start, deleteCount int, elements ...T) ArrayList[T] {
	if deleteCount <= 0 {
		return make(ArrayList[T], 0)
	}
	start = l.parseIndex(start)
	endIndex := l.parseIndex(start + deleteCount)
	removed := (*l)[start:endIndex]
	head := (*l)[:start]
	tail := (*l)[endIndex:]
	insertCount := len(elements)
	newList := make(ArrayList[T], l.Len()-removed.Len()+insertCount)
	copy(newList, head)
	copy(newList[start:], elements)
	copy(newList[start+insertCount:], tail)
	*l = newList
	return removed.Copy()
}

// Fill fills the specified area with the specified element.
// If the area argument is not provided, the entire list will be filled with the specified element.
func (l *ArrayList[T]) Fill(element T, area ...int) *ArrayList[T] {
	argCount := len(area)
	start := 0
	end := l.Len()
	if argCount >= 1 {
		start = l.parseIndex(area[0])
	}
	if argCount >= 2 {
		end = l.parseIndex(area[1])
	}
	for i := start; i < end; i++ {
		(*l)[i] = element
	}
	return l
}

// Reverse reverses the order of the elements in the list.
func (l *ArrayList[T]) Reverse() *ArrayList[T] {
	slices.Reverse(*l)
	return l
}

// Set sets the element at the specified index to the specified element.
// If the index is negative, the index will be calculated from the end of the list.
// Returns an error if the index is out of range.
func (l *ArrayList[T]) Set(index int, element T) (err error) {
	if index < 0 {
		index += l.Len()
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	(*l)[index] = element
	return
}

// RemoveIf removes the elements that satisfy the specified condition.
// If the count argument is not provided, only the first count occurrences will be removed.
// Returns the removed elements.
func (l *ArrayList[T]) RemoveIf(condition func(T) bool, counts ...int) ArrayList[T] {
	count := l.parseCount(counts...)
	array := make(ArrayList[T], l.Len())
	i := 0
	removed := make(ArrayList[T], count)
	j := 0
	for _, v := range *l {
		if count > 0 && condition(v) {
			count--
			removed[j] = v
			j++
			continue
		}
		array[i] = v
		i++
	}
	*l = array[:i:i]
	return removed[:j:j]
}

// RemoveRightIf removes the elements that satisfy the specified condition from the right end of the list.
// If the count argument is not provided, only the first count occurrences will be removed.
// Returns the removed elements.
func (l *ArrayList[T]) RemoveRightIf(condition func(T) bool, counts ...int) ArrayList[T] {
	count := l.parseCount(counts...)
	length := l.Len()
	array := make(ArrayList[T], length)
	end := length - 1
	i := end
	removed := make(ArrayList[T], count)
	j := 0
	for k := end; k >= 0; k-- {
		v := (*l)[k]
		if count > 0 && condition(v) {
			count--
			removed[j] = v
			j++
			continue
		}
		array[i] = v
		i--
	}
	*l = slices.Clip(array[i+1 : end+1])
	return removed[:j:j]
}

// ReplaceIf replaces the elements that satisfy the specified condition with the specified new element.
// If the count argument is not provided, only the first count occurrences will be replaced.
// Returns the replaced elements.
func (l *ArrayList[T]) ReplaceIf(condition func(T) bool, newElement T, counts ...int) ArrayList[T] {
	count := l.parseCount(counts...)
	replaced := make(ArrayList[T], count)
	j := 0
	for i, v := range *l {
		if count == 0 {
			break
		}
		if condition(v) {
			replaced[j] = v
			j++
			(*l)[i] = newElement
			count--
		}
	}
	return replaced[:j:j]
}

// ReplaceRightIf replaces the elements that satisfy the specified condition with the specified new element from the right end of the list.
// If the count argument is not provided, only the first count occurrences will be replaced.
// Returns the replaced elements.
func (l *ArrayList[T]) ReplaceRightIf(condition func(T) bool, newElement T, counts ...int) ArrayList[T] {
	count := l.parseCount(counts...)
	replaced := make(ArrayList[T], count)
	j := 0
	for i := l.Len() - 1; i >= 0; i-- {
		if count == 0 {
			break
		}
		v := (*l)[i]
		if condition(v) {
			replaced[j] = v
			j++
			(*l)[i] = newElement
			count--
		}
	}
	return replaced[:j:j]
}
