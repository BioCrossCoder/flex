package itertools

import "flex/common"

type sliceIterator[T any] struct {
	entry   []T
	end     int
	pointer int
	value   T
	step    int
}

func NewSliceIterator[T any](entry []T, start, end, step int) ListIterator[T] {
	return &sliceIterator[T]{
		entry:   entry,
		end:     end,
		pointer: start,
		value:   *new(T),
		step:    step,
	}
}

func (iter *sliceIterator[T]) clear() {
	iter.value = *new(T)
	iter.entry = nil
}

func (iter *sliceIterator[T]) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer += iter.step
	return true
}

func (iter *sliceIterator[T]) Value() T {
	return iter.value
}

func (iter *sliceIterator[T]) Pour() []T {
	output := make([]T, 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}

func Slice[T any](entry []T, start, end, step int) (slice []T, err error) {
	length := len(entry)
	err = common.CheckRange(start, end, step, length)
	if err != nil {
		return
	}
	slice = NewSliceIterator[T](entry, common.ParseIndex(start, length), common.ParseIndex(end, length), step).Pour()
	return
}

func Reversed[T any](entry []T) ListIterator[T] {
	return NewSliceIterator(entry, len(entry)-1, 0, -1)
}
