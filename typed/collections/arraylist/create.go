package arraylist

func (l ArrayList[T]) Copy() ArrayList[T] {
	backup := make(ArrayList[T], l.Len())
	copy(backup, l)
	return backup
}

func (l ArrayList[T]) Concat(another ArrayList[T]) ArrayList[T] {
	mid := l.Len()
	linkedList := make(ArrayList[T], mid+another.Len())
	copy(linkedList, l)
	copy(linkedList[mid:], another)
	return linkedList
}

func (l ArrayList[T]) Slice(args ...int) ArrayList[T] {
	argsCount := len(args)
	if argsCount == 0 {
		return l.Copy()
	}
	srcListLength := l.Len()
	start := 0
	end := srcListLength
	step := 1
	if argsCount >= 1 {
		start = l.parseIndex(args[0])
	}
	if argsCount >= 2 {
		end = l.parseIndex(args[1])
	}
	if argsCount >= 3 {
		step = args[2]
	}
	if (start < end && step < 0) || (start > end && step > 0) || (start == end) || (step == 0) {
		return make(ArrayList[T], 0)
	}
	sliceListLength := 0
	list := make(ArrayList[T], srcListLength)
	condition := func(start, end, step int) bool {
		if step > 0 {
			return start < end
		} else {
			return start > end
		}
	}
	for i := start; condition(i, end, step); i += step {
		list[sliceListLength] = l[i]
		sliceListLength++
	}
	return list[:sliceListLength].Copy()
}

func (l ArrayList[T]) ToSpliced(start, deleteCount int, items ...T) ArrayList[T] {
	list := l.Copy()
	_ = list.Splice(start, deleteCount, items...)
	return list
}

func (l ArrayList[T]) ToReversed() ArrayList[T] {
	list := l.Copy()
	_ = list.Reverse()
	return list
}

func (l ArrayList[T]) With(index int, element T) ArrayList[T] {
	list := l.Copy()
	_ = list.Set(index, element)
	return list
}

func Of[T any](elements ...T) ArrayList[T] {
	return ArrayList[T](elements)
}

func Repeat[T any](element T, count int) ArrayList[T] {
	list := make(ArrayList[T], count)
	for i := 0; i < count; i++ {
		list[i] = element
	}
	return list
}
