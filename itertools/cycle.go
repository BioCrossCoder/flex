// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// cycler is an iterator that cycles through a slice of values.
type cycler struct {
	entry   []any
	length  int
	pointer int
	value   any
}

// NewCycler creates a new cycler iterator from a slice.
func NewCycler(entry []any) *cycler {
	return &cycler{
		entry:   entry,
		length:  len(entry),
		pointer: -1,
		value:   nil,
	}
}

// Next returns the next value in the cycler.
func (iter *cycler) Next() any {
	iter.pointer = (iter.pointer + 1) % iter.length
	iter.value = iter.entry[iter.pointer]
	return iter.value
}

// Cycle creates an iterator that cycles through a sequence of values.
func Cycle(entry any) (iterator *cycler, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	switch value.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		iterator = NewCycler(common.CopyList(value, value.Len()))
	case reflect.String:
		iterator = NewCycler(common.ConvertStringToList(entry.(string)))
	}
	return
}
