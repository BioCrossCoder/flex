package list

import "flex/common"

func (l *List) Remove(element any, counts ...int) *List {
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = l.Count(element)
	}
	for i := 0; i < count; i++ {
		index := l.IndexOf(element)
		_, _ = l.Pop(index)
	}
	return l
}

func (l *List) RemoveRight(element any, counts ...int) *List {
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = l.Count(element)
	}
	for i := 0; i < count; i++ {
		index := l.LastIndexOf(element)
		_, _ = l.Pop(index)
	}
	return l
}

func (l *List) Clear() *List {
	*l = make(List, 0)
	return l
}

func (l *List) Push(elements ...any) (length int) {
	length = l.Len() + len(elements)
	*l = l.Concat(List(elements))
	return
}

func (l *List) Pop(indexes ...int) (element any, err error) {
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

func (l *List) Unshift(elements ...any) (length int) {
	length = l.Len() + len(elements)
	*l = List(elements).Concat(*l)
	return
}

func (l *List) Shift() (element any, err error) {
	return l.Pop(0)
}

func (l *List) Insert(index int, element any) *List {
	length := l.Len()
	validIndex := l.parseIndex(index)
	*l = append(*l, nil)
	for i := length; i > validIndex; i-- {
		(*l)[i] = (*l)[i-1]
	}
	(*l)[validIndex] = element
	return l
}

func (l *List) ForEach(action func(any) any) *List {
	for i, item := range *l {
		(*l)[i] = action(item)
	}
	return l
}

func (l *List) Replace(oldElement, newElement any, counts ...int) *List {
	if oldElement == newElement {
		return l
	}
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = l.Count(oldElement)
	}
	for i := 0; i < count; i++ {
		index := l.IndexOf(oldElement)
		(*l)[index] = newElement
	}
	return l
}

func (l *List) ReplaceRight(oldElement, newElement any, counts ...int) *List {
	if oldElement == newElement {
		return l
	}
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = l.Count(oldElement)
	}
	for i := 0; i < count; i++ {
		index := l.LastIndexOf(oldElement)
		(*l)[index] = newElement
	}
	return l
}

func (l *List) Splice(start, deleteCount int, items ...any) List {
	if deleteCount <= 0 {
		return make(List, 0)
	}
	start = l.parseIndex(start)
	endIndex := l.parseIndex(start + deleteCount)
	removed := (*l)[start:endIndex]
	head := (*l)[:start]
	tail := (*l)[endIndex:]
	insertCount := len(items)
	newList := make(List, l.Len()-removed.Len()+insertCount)
	copy(newList, head)
	copy(newList[start:], items)
	copy(newList[start+insertCount:], tail)
	*l = newList
	return removed.Copy()
}

func (l *List) Fill(element any, area ...int) *List {
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

func (l *List) Reverse() *List {
	for i, j := 0, l.Len()-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
	return l
}

func (l *List) Set(index int, element any) (err error) {
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
