package arraylist

func (l ArrayList) Copy() ArrayList {
	backup := make(ArrayList, l.Len())
	copy(backup, l)
	return backup
}

func (l ArrayList) Concat(another ArrayList) ArrayList {
	mid := l.Len()
	linkedList := make(ArrayList, mid+another.Len())
	copy(linkedList, l)
	copy(linkedList[mid:], another)
	return linkedList
}

func (l ArrayList) Slice(args ...int) ArrayList {
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
		return make(ArrayList, 0)
	}
	i := 0
	var list ArrayList
	if step < 0 {
		list = make(ArrayList, (start-end-step-1)/(-step))
		for j := start; j > end; j += step {
			list[i] = l[j]
			i++
		}
	} else {
		list = make(ArrayList, (end-start+step-1)/step)
		for j := start; j < end; j += step {
			list[i] = l[j]
			i++
		}
	}
	return list
}

func (l ArrayList) ToSpliced(start, deleteCount int, items ...any) ArrayList {
	list := l.Copy()
	_ = list.Splice(start, deleteCount, items...)
	return list
}

func (l ArrayList) ToReversed() ArrayList {
	list := l.Copy()
	_ = list.Reverse()
	return list
}

func (l ArrayList) With(index int, element any) ArrayList {
	list := l.Copy()
	_ = list.Set(index, element)
	return list
}

func Of(elements ...any) ArrayList {
	return ArrayList(elements)
}

func Repeat(element any, count int) ArrayList {
	list := make(ArrayList, count)
	for i := 0; i < count; i++ {
		list[i] = element
	}
	return list
}
