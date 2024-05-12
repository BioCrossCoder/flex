// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// listConvertor is an iterator that converts elements of a list by provided handler function.
type listConvertor struct {
	entry   []any
	length  int
	handler func(any) any
	pointer int
	value   any
}

// NewListConvertor creates a new listConvertor iterator.
func NewListConvertor(entry []any, handler func(any) any) Iterator {
	return &listConvertor{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 0,
		value:   nil,
	}
}

// clear releases the resources used by the iterator.
func (iter *listConvertor) clear() {
	iter.value = nil
	iter.entry = nil
	iter.handler = nil
}

// Next moves the pointer to the next element and returns true if there is a next element, or just returns false otherwise.
func (iter *listConvertor) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entry[iter.pointer])
	iter.pointer++
	return true
}

// Value returns the element converted by the handler function from the element in the list pointed by the pointer.
func (iter *listConvertor) Value() any {
	return iter.value
}

// Pour returns all the unvisited elements of the list converted by the handler function.
func (iter *listConvertor) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

// Map applies the provided handler function to each element of the input sequence and returns an iterator that produces the results.
func Map(handler, entry any) (iterator Iterator, err error) {
	err = common.IsInputFuncValid(handler, 1, 1)
	if err != nil {
		return
	}
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	iterHandler := func(a any) any {
		params := []reflect.Value{reflect.ValueOf(a)}
		return reflect.ValueOf(handler).Call(params)[0].Interface()
	}
	switch value.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		iterator = NewListConvertor(common.CopyList(value, length), iterHandler)
	case reflect.String:
		iterator = NewListConvertor(common.ConvertStringToList(value.String()), iterHandler)
	}
	return
}
