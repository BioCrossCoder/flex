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
