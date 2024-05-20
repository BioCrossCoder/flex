package arraylist

// Copy creates a shallow copy of the ArrayList.
func (l ArrayList[T]) Copy() ArrayList[T] {
	backup := make(ArrayList[T], l.Len())
	copy(backup, l)
	return backup
}

// Concat concatenates the current ArrayList with another ArrayList.
func (l ArrayList[T]) Concat(another ArrayList[T]) ArrayList[T] {
	mid := l.Len()
	linkedList := make(ArrayList[T], mid+another.Len())
	copy(linkedList, l)
	copy(linkedList[mid:], another)
	return linkedList
}

// Slice returns a new ArrayList by slicing the current ArrayList based on the provided arguments.
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
		start = l.sliceIndex(args[0], true)
	}
	if argsCount >= 2 {
		end = l.sliceIndex(args[1], false)
	}
	if argsCount >= 3 {
		step = args[2]
	}
	if (start < end && step < 0) || (start > end && step > 0) || (start == end) || (step == 0) {
		return make(ArrayList[T], 0)
	}
	i := 0
	var list ArrayList[T]
	if step < 0 {
		list = make(ArrayList[T], (start-end-step-1)/(-step))
		for j := start; j > end; j += step {
			list[i] = l[j]
			i++
		}
	} else {
		list = make(ArrayList[T], (end-start+step-1)/step)
		for j := start; j < end; j += step {
			list[i] = l[j]
			i++
		}
	}
	return list
}

// ToSpliced returns a new ArrayList after applying the splice operation to the current ArrayList.
func (l ArrayList[T]) ToSpliced(start, deleteCount int, items ...T) ArrayList[T] {
	list := l.Copy()
	_ = list.Splice(start, deleteCount, items...)
	return list
}

// ToReversed returns a new ArrayList with its elements reversed compared to the current ArrayList.
func (l ArrayList[T]) ToReversed() ArrayList[T] {
	list := l.Copy()
	_ = list.Reverse()
	return list
}

// With returns a new ArrayList after modifying an element at the specified index in the current ArrayList.
func (l ArrayList[T]) With(index int, element T) ArrayList[T] {
	list := l.Copy()
	_ = list.Set(index, element)
	return list
}

// Of creates a new ArrayList with the provided elements.
func Of[T any](elements ...T) ArrayList[T] {
	return ArrayList[T](elements)
}

// Repeat creates a new ArrayList by repeating the specified element for the given count.
func Repeat[T any](element T, count int) ArrayList[T] {
	list := make(ArrayList[T], count)
	for i := 0; i < count; i++ {
		list[i] = element
	}
	return list
}
