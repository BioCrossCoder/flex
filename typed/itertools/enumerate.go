package itertools

import "flex/common"

type Enumerator[T any] interface {
	Next() bool
	Value() *enumItem[T]
	Pour() []*enumItem[T]
}

type enumItem[T any] struct {
	Index int
	Value T
}

type enumerator[T any] struct {
	entry   []T
	end     int
	pointer int
	value   *enumItem[T]
	step    int
}

func Enumerate[T any](entry []T, start, end, step int) (iterator Enumerator[T], err error) {
	length := len(entry)
	err = common.CheckRange(start, end, step, length)
	if err != nil {
		return
	}
	iterator = &enumerator[T]{
		entry:   entry,
		end:     common.ParseIndex(end, length),
		pointer: common.ParseIndex(start, length),
		value:   nil,
		step:    step,
	}
	return
}

func (iter *enumerator[T]) clear() {
	iter.value = nil
	iter.entry = nil
}

func (iter *enumerator[T]) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = &enumItem[T]{
		Index: iter.pointer,
		Value: iter.entry[iter.pointer],
	}
	iter.pointer += iter.step
	return true
}

func (iter *enumerator[T]) Value() *enumItem[T] {
	return iter.value
}

func (iter *enumerator[T]) Pour() []*enumItem[T] {
	output := make([]*enumItem[T], 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}
