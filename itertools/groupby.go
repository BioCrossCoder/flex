package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// GroupBy groups the elements of the given list by the given key function and returns an iterator of groups.
func GroupBy(entry any, by func(any) any) (iterator Iterator, err error) {
	err = common.IsList(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	groupsMap := make(map[any]any)
	orders := make(map[int]any)
	count := 0
	for _, v := range common.CopyList(value, value.Len()) {
		key := by(v)
		group, ok := groupsMap[key]
		if !ok {
			group = make([]any, 0)
			orders[count] = key
			count++
		}
		group = append(group.([]any), v)
		groupsMap[key] = group
	}
	groupsList := make([]any, count)
	for i := 0; i < count; i++ {
		groupsList[i] = groupsMap[orders[i]]
	}
	iterator = NewListIterator(groupsList)
	return
}
