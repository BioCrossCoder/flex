package list

func (l List) Copy() List {
	backup := make(List, l.Len())
	copy(backup, l)
	return backup
}

func (l List) Concat(another List) List {
	mid := l.Len()
	linkedList := make(List, mid+another.Len())
	copy(linkedList, l)
	copy(linkedList[mid:], another)
	return linkedList
}

func (l List) Slice(args ...int) List {
	argsCount := len(args)
	if argsCount == 0 {
		return l.Copy()
	}
	start := 0
	end := l.Len()
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
	list := make(List, 0)
	if (start < end && step < 0) || (start > end && step > 0) {
		return list
	}
	condition := func(start, end, step int) bool {
		if step > 0 {
			return start < end
		} else {
			return start > end
		}
	}
	for i := start; condition(i, end, step); i += step {
		list = append(list, l[i])
	}
	return list
}

func (l List) ToSpliced(start, deleteCount int, items ...any) List {
	list := l.Copy()
	_ = list.Splice(start, deleteCount, items...)
	return list
}

func (l List) ToReversed() List {
	list := l.Copy()
	_ = list.Reverse()
	return list
}

func (l List) With(index int, element any) List {
	list := l.Copy()
	list[l.parseIndex(index)] = element
	return list
}

func Of(elements ...any) List {
	return List(elements)
}

func Repeat(element any, count int) List {
	list := make(List, count)
	for i := 0; i < count; i++ {
		list[i] = element
	}
	return list
}
