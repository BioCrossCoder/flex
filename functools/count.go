package functools

import (
	"reflect"
)

func CountBy(entry any, condition func(any) bool) (count int) {
	list := reflect.ValueOf(entry)
	switch list.Kind() {
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
	case reflect.Map:
		iter := list.MapRange()
		for iter.Next() {
			if condition(iter.Value().Interface()) {
				count++
			}
		}
	default:
		count = -1
	}
	return
}
