package arraylist

import "flex/common"

func (l ArrayList[T]) parseCount(counts ...int) int {
	if len(counts) == 0 {
		return 1
	}
	if counts[0] <= 0 {
		return l.Len()
	}
	return counts[0]
}

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
	*l = array[i+1 : end+1].Copy()
	return l
}

func (l *ArrayList[T]) Clear() *ArrayList[T] {
	*l = make(ArrayList[T], 0)
	return l
}

func (l *ArrayList[T]) Push(elements ...T) (length int) {
	length = l.Len() + len(elements)
	*l = l.Concat(ArrayList[T](elements))
	return
}

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
	*l = append((*l)[:index], (*l)[index+1:]...)
	return
}

func (l *ArrayList[T]) Unshift(elements ...T) (length int) {
	length = l.Len() + len(elements)
	*l = ArrayList[T](elements).Concat(*l)
	return
}

func (l *ArrayList[T]) Shift() (element T, err error) {
	return l.Pop(0)
}

func (l *ArrayList[T]) Insert(index int, element T) *ArrayList[T] {
	length := l.Len()
	validIndex := l.parseIndex(index)
	*l = append(*l, element)
	for i := length; i > validIndex; i-- {
		(*l)[i] = (*l)[i-1]
	}
	(*l)[validIndex] = element
	return l
}

func (l *ArrayList[T]) ForEach(action func(T) T) *ArrayList[T] {
	for i, item := range *l {
		(*l)[i] = action(item)
	}
	return l
}

func (l *ArrayList[T]) Replace(oldElement, newElement T, counts ...int) *ArrayList[T] {
	if common.Equal(oldElement, newElement) {
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

func (l *ArrayList[T]) ReplaceRight(oldElement, newElement T, counts ...int) *ArrayList[T] {
	if common.Equal(oldElement, newElement) {
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

func (l *ArrayList[T]) Reverse() *ArrayList[T] {
	for i, j := 0, l.Len()-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
	return l
}

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
