package itertools

// GroupBy is a generic function that groups the elements of the input slice based on a given criterion function.
// It takes a slice of type T and a function that extracts a key of type U from each element, and returns a ListIterator of grouped elements.
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
