package itertools

import "flex/common"

type Iterator interface {
	Next() bool
	Value() any
	Pour() any
}

type listConvertor struct {
	entry   []any
	length  int
	handler func(any) any
	pointer int
	value   any
}

func NewListIterator(entry []any, handler func(any) any) Iterator {
	return &listConvertor{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 0,
		value:   nil,
	}
}

func (iter *listConvertor) clear() {
	iter.value = nil
	iter.entry = nil
	iter.handler = nil
}

func (iter *listConvertor) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entry[iter.pointer])
	iter.pointer++
	return true
}

func (iter *listConvertor) Value() any {
	return iter.value
}

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

type listFilter struct {
	listConvertor
}

func NewListFilter(entry []any, handler func(any) bool) Iterator {
	iter := new(listFilter)
	iter.entry = entry
	iter.length = len(entry)
	iter.handler = func(p any) any {
		return handler(p)
	}
	iter.pointer = 0
	iter.value = false
	return iter
}

func (iter *listFilter) Pour() any {
	output := make([]any, 0)
	for iter.Next() {
		if iter.Value().(bool) {
			value := iter.entry[iter.pointer-1]
			output = append(output, value)
		}
	}
	return output
}

type mapConvertor struct {
	entryKeys   []any
	entryValues []any
	length      int
	handler     func(any) any
	pointer     int
	value       any
}

func NewMapIterator(entry map[any]any, handler func(any) any) Iterator {
	keys, values, length := common.ConvertMapToLists(entry)
	return &mapConvertor{
		entryKeys:   keys,
		entryValues: values,
		length:      length,
		handler:     handler,
		pointer:     0,
		value:       nil,
	}
}

func (iter *mapConvertor) clear() {
	iter.value = nil
	iter.entryKeys = nil
	iter.entryValues = nil
	iter.handler = nil
}

func (iter *mapConvertor) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entryValues[iter.pointer])
	iter.pointer++
	return true
}

func (iter *mapConvertor) Value() any {
	return iter.value
}

func (iter *mapConvertor) Pour() any {
	length := iter.length - iter.pointer
	output := make(map[any]any, length)
	for iter.Next() {
		key := iter.entryKeys[iter.pointer-1]
		output[key] = iter.Value()
	}
	return output
}

type mapFilter struct {
	mapConvertor
}

func NewMapFilter(entry map[any]any, handler func(any) bool) Iterator {
	keys, values, length := common.ConvertMapToLists(entry)
	iter := new(mapFilter)
	iter.entryKeys = keys
	iter.entryValues = values
	iter.length = length
	iter.handler = func(p any) any {
		return handler(p)
	}
	iter.pointer = 0
	iter.value = false
	return iter
}

func (iter *mapFilter) Pour() any {
	output := make(map[any]any)
	for iter.Next() {
		if iter.Value().(bool) {
			key := iter.entryKeys[iter.pointer-1]
			value := iter.entryValues[iter.pointer-1]
			output[key] = value
		}
	}
	return output
}

type accumulator struct {
	entry   []any
	length  int
	handler func(any, any) any
	pointer int
	value   any
}

func NewAccumulator(entry []any, handler func(any, any) any) Iterator {
	return &accumulator{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 1,
		value:   entry[0],
	}
}

func (iter *accumulator) clear() {
	iter.value = nil
	iter.entry = nil
	iter.handler = nil
}

func (iter *accumulator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.value, iter.entry[iter.pointer])
	iter.pointer++
	return true
}

func (iter *accumulator) Value() any {
	return iter.value
}

func (iter *accumulator) Pour() any {
	result := iter.Value()
	for iter.Next() {
		result = iter.Value()
	}
	return result
}

type zipIterator struct {
	entry1  []any
	entry2  []any
	length  int
	pointer int
	value   any
}

func NewZipIterator(entry1, entry2 []any, length int) Iterator {
	return &zipIterator{
		entry1:  entry1,
		entry2:  entry2,
		length:  length,
		pointer: 0,
		value:   nil,
	}
}

func (iter *zipIterator) clear() {
	iter.entry1 = nil
	iter.entry2 = nil
	iter.value = nil
}

func (iter *zipIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = [2]any{iter.entry1[iter.pointer], iter.entry2[iter.pointer]}
	iter.pointer++
	return true
}

func (iter *zipIterator) Value() any {
	return iter.value
}

func (iter *zipIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([][2]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value().([2]any)
		i++
	}
	return output
}
