package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// parseCount parses the counts and returns the appropriate count value
func (l ArrayList) parseCount(counts ...int) int {
	return list.ParseCount(l.Len(), counts...)
}

// Remove removes the specified number of occurrences of the element from the list
func (l *ArrayList) Remove(element any, counts ...int) *ArrayList {
	count := l.parseCount(counts...)
	array := make(ArrayList, l.Len())
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

// RemoveRight removes the specified number of occurrences of the element from the end of the list
func (l *ArrayList) RemoveRight(element any, counts ...int) *ArrayList {
	count := l.parseCount(counts...)
	length := l.Len()
	array := make(ArrayList, length)
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
	*l = array[i+1 : end+1].Copy()
	return l
}

// Clear removes all elements from the list
func (l *ArrayList) Clear() *ArrayList {
	*l = make(ArrayList, 0)
	return l
}

// Push appends elements to the end of the list and returns the new length
func (l *ArrayList) Push(elements ...any) (length int) {
	length = l.Len() + len(elements)
	*l = l.Concat(ArrayList(elements))
	return
}

// Pop removes the element at the specified index (or the last element if no index is provided) and returns it
func (l *ArrayList) Pop(indexes ...int) (element any, err error) {
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
	*l = append((*l)[:index], (*l)[index+1:]...)
	return
}

// Unshift prepends elements to the beginning of the list and returns the new length
func (l *ArrayList) Unshift(elements ...any) (length int) {
	length = l.Len() + len(elements)
	*l = ArrayList(elements).Concat(*l)
	return
}

// Shift removes the first element from the list and returns it
func (l *ArrayList) Shift() (element any, err error) {
	return l.Pop(0)
}

// Insert inserts the specified element at the specified index
func (l *ArrayList) Insert(index int, element any) *ArrayList {
	length := l.Len()
	validIndex := l.parseIndex(index)
	*l = append(*l, nil)
	for i := length; i > validIndex; i-- {
		(*l)[i] = (*l)[i-1]
	}
	(*l)[validIndex] = element
	return l
}

// ForEach applies the specified action to each element of the list
func (l *ArrayList) ForEach(action func(any) any) *ArrayList {
	for i, item := range *l {
		(*l)[i] = action(item)
	}
	return l
}

// Replace replaces the specified number of occurrences of the old element with the new element
func (l *ArrayList) Replace(oldElement, newElement any, counts ...int) *ArrayList {
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

// ReplaceRight replaces the specified number of occurrences of the old element with the new element, starting from the end of the list
func (l *ArrayList) ReplaceRight(oldElement, newElement any, counts ...int) *ArrayList {
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

// Splice removes a section of the list and replaces it with the specified items
func (l *ArrayList) Splice(start, deleteCount int, items ...any) ArrayList {
	if deleteCount <= 0 {
		return make(ArrayList, 0)
	}
	start = l.parseIndex(start)
	endIndex := l.parseIndex(start + deleteCount)
	removed := (*l)[start:endIndex]
	head := (*l)[:start]
	tail := (*l)[endIndex:]
	insertCount := len(items)
	newList := make(ArrayList, l.Len()-removed.Len()+insertCount)
	copy(newList, head)
	copy(newList[start:], items)
	copy(newList[start+insertCount:], tail)
	*l = newList
	return removed.Copy()
}

// Fill fills a section of the list with the specified element
func (l *ArrayList) Fill(element any, area ...int) *ArrayList {
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

// Reverse reverses the order of elements in the list
func (l *ArrayList) Reverse() *ArrayList {
	for i, j := 0, l.Len()-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
	return l
}

// Set sets the element at the specified index to the provided value
func (l *ArrayList) Set(index int, element any) (err error) {
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

// RemoveIf removes elements from the list that satisfy the specified condition
func (l *ArrayList) RemoveIf(condition func(any) bool, counts ...int) ArrayList {
	count := l.parseCount(counts...)
	array := make(ArrayList, l.Len())
	i := 0
	removed := make(ArrayList, count)
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

// RemoveRightIf removes elements from the end of the list that satisfy the specified condition
func (l *ArrayList) RemoveRightIf(condition func(any) bool, counts ...int) ArrayList {
	count := l.parseCount(counts...)
	length := l.Len()
	array := make(ArrayList, length)
	end := length - 1
	i := end
	removed := make(ArrayList, count)
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
	*l = array[i+1 : end+1].Copy()
	return removed[:j:j]
}

// ReplaceIf replaces elements in the list that satisfy the specified condition with the new element
func (l *ArrayList) ReplaceIf(condition func(any) bool, newElement any, counts ...int) ArrayList {
	count := l.parseCount(counts...)
	replaced := make(ArrayList, count)
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

// ReplaceRightIf replaces elements in the list, starting from the end, that satisfy the specified condition with the new element
func (l *ArrayList) ReplaceRightIf(condition func(any) bool, newElement any, counts ...int) ArrayList {
	count := l.parseCount(counts...)
	replaced := make(ArrayList, count)
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
