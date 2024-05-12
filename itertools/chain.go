// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// Chain creates an iterator that chains multiple sequences together.
func Chain(entries ...any) (iterator Iterator, err error) {
	elements := make([]any, 0)
	for _, entry := range entries {
		err = common.IsSequence(entry)
		if err != nil {
			return
		}
		value := reflect.ValueOf(entry)
		switch value.Kind() { //nolint
		case reflect.Array, reflect.Slice:
			elements = append(elements, common.CopyList(value, value.Len())...)
		case reflect.String:
			elements = append(elements, common.ConvertStringToList(value.String())...)
		}
	}
	return NewListIterator(elements), nil
}
