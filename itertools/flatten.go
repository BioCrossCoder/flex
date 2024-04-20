package itertools

import (
	"flex/common"
	"reflect"
)

func Flatten(entry any) (output []any, err error) {
	err = common.IsList(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	output = make([]any, 0)
	for _, v := range common.CopyList(value, length) {
		err := common.IsList(v)
		if err != nil {
			output = append(output, v)
		} else {
			elements, _ := Flatten(v)
			output = append(output, elements...)
		}
	}
	return
}
