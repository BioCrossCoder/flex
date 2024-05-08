package itertools

func GroupBy[T any, U comparable](entry []T, by func(T) U) ListIterator[[]T] {
	groupsMap := make(map[U][]T)
	orders := make(map[int]U)
	count := 0
	for _, v := range entry {
		key := by(v)
		group, ok := groupsMap[key]
		if !ok {
			group = make([]T, 0)
			orders[count] = key
			count++
		}
		groupsMap[key] = append(group, v)
	}
	groupsList := make([][]T, count)
	for i := 0; i < count; i++ {
		groupsList[i] = groupsMap[orders[i]]
	}
	return NewListIterator(groupsList)
}
