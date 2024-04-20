package itertools

import "flex/common"

type listIterator struct {
	entry   []any
	length  int
	pointer int
	value   any
}

func NewListIterator(entry []any) common.Iterator {
	return &listIterator{
		entry:   entry,
		length:  len(entry),
		pointer: 0,
		value:   nil,
	}
}

func (iter *listIterator) clear() {
	iter.value = nil
	iter.entry = nil
}

func (iter *listIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer++
	return true
}

func (iter *listIterator) Value() any {
	return iter.value
}

func (iter *listIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

type mapIterator struct {
	entryKeys   []any
	entryValues []any
	length      int
	pointer     int
	value       any
}

func NewMapIterator(entry map[any]any) common.Iterator {
	keys, value, length := common.ConvertMapToLists(entry)
	return &mapIterator{
		entryKeys:   keys,
		entryValues: value,
		length:      length,
		pointer:     0,
		value:       nil,
	}
}

func (iter *mapIterator) clear() {
	iter.value = nil
	iter.entryKeys = nil
	iter.entryValues = nil
}

func (iter *mapIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entryValues[iter.pointer]
	iter.pointer++
	return true
}

func (iter *mapIterator) Value() any {
	return iter.value
}

func (iter *mapIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make(map[any]any, length)
	for iter.Next() {
		key := iter.entryKeys[iter.pointer-1]
		output[key] = iter.Value()
	}
	return output
}
