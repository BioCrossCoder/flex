package functools

import "reflect"

// CountBy counts the number of elements in the given entry that satisfy the given condition, and it will return -1 if the entry is not array, slice or string.
func CountBy(entry any, condition func(any) bool) (count int) {
	list := reflect.ValueOf(entry)
	switch list.Kind() { //nolint
	case reflect.String:
		for _, char := range list.String() {
			if condition(string(char)) {
				count++
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			if condition(list.Index(i).Interface()) {
				count++
			}
		}
	default:
		count = -1
	}
	return
}
