// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
	"unicode/utf8"
)

// Reversed returns a reversed copy of the input sequence.
func Reversed(entry any) (output any, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	var length int
	var iterator Iterator
	switch value.Kind() { //nolint
	case reflect.Slice, reflect.Array:
		length = value.Len()
		iterator = NewSliceIterator(common.CopyList(value, length), length-1, 0, -1)
		output = iterator.Pour()
	case reflect.String:
		raw := value.String()
		length = utf8.RuneCountInString(raw)
		iterator = NewSliceIterator(common.ConvertStringToList(raw), length-1, 0, -1)
		output = ""
		for _, c := range iterator.Pour().([]any) {
			output = output.(string) + c.(string)
		}
	}
	return
}
