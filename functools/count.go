package functools

import (
	"flex/common"
	"reflect"
)

func Count(list, element any) (count int, err error) {
	err = common.IsIterable(list)
	if err != nil {
		return
	}
	value := reflect.ValueOf(list)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if reflect.DeepEqual(value.Index(i).Interface(), element) {
				count++
			}
		}
	case reflect.String:
		elements := []rune(list.(string))
		for i := 0; i < len(elements); i++ {
			if reflect.DeepEqual(string(elements[i]), element) {
				count++
			}
		}
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			if reflect.DeepEqual(iter.Value().Interface(), element) {
				count++
			}
		}
	}
	return
}
