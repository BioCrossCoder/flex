package itertools

import "github.com/biocrosscoder/flex/common"

type ZipIterator[T, U any] interface {
	Next() bool
	Value() *zipPair[T, U]
	Pour() []*zipPair[T, U]
}

type zipPairIterator[T, U any] struct {
	entry1  []T
	entry2  []U
	length  int
	pointer int
	value   *zipPair[T, U]
}

type zipPair[T, U any] struct {
	First  T
	Second U
}

func newZipPairIterator[T any, U any](entry1 []T, entry2 []U, length int) ZipIterator[T, U] {
	return &zipPairIterator[T, U]{entry1, entry2, length, 0, nil}
}

func (z *zipPairIterator[T, U]) clear() {
	z.entry1 = nil
	z.entry2 = nil
	z.value = nil
}

func (z *zipPairIterator[T, U]) getFirst() T {
	if z.pointer >= len(z.entry1) {
		return *new(T)
	}
	return z.entry1[z.pointer]
}

func (z *zipPairIterator[T, U]) getSecond() U {
	if z.pointer >= len(z.entry2) {
		return *new(U)
	}
	return z.entry2[z.pointer]
}

func (z *zipPairIterator[T, U]) Next() bool {
	if z.pointer == z.length {
		z.clear()
		return false
	}
	z.value = &zipPair[T, U]{
		z.getFirst(),
		z.getSecond(),
	}
	z.pointer++
	return true
}

func (z *zipPairIterator[T, U]) Value() *zipPair[T, U] {
	return z.value
}

func (z *zipPairIterator[T, U]) Pour() []*zipPair[T, U] {
	result := make([]*zipPair[T, U], z.length-z.pointer)
	i := 0
	for z.Next() {
		result[i] = z.Value()
		i++
	}
	return result
}

func ZipPair[T, U any](entry1 []T, entry2 []U) ZipIterator[T, U] {
	return newZipPairIterator(entry1, entry2, min(len(entry1), len(entry2)))
}

func ZipPairLongest[T, U any](entry1 []T, entry2 []U) ZipIterator[T, U] {
	return newZipPairIterator(entry1, entry2, max(len(entry1), len(entry2)))
}

type zipListIterator[T any] struct {
	entries [][]T
	length  int
	pointer int
	value   []T
}

func newZipListIterator[T any](entries [][]T, length int) ListIterator[[]T] {
	return &zipListIterator[T]{entries, length, 0, make([]T, len(entries))}
}

func (z *zipListIterator[T]) clear() {
	z.entries = nil
	z.value = nil
}

func (z *zipListIterator[T]) getValue(entryIndex int) T {
	if z.pointer >= len(z.entries[entryIndex]) {
		return *new(T)
	}
	return z.entries[entryIndex][z.pointer]
}

func (z *zipListIterator[T]) Next() bool {
	if z.pointer == z.length {
		z.clear()
		return false
	}
	z.value = make([]T, len(z.entries))
	for i := range z.value {
		z.value[i] = z.getValue(i)
	}
	z.pointer++
	return true
}

func (z *zipListIterator[T]) Value() []T {
	return z.value
}

func (z *zipListIterator[T]) Pour() [][]T {
	result := make([][]T, z.length-z.pointer)
	i := 0
	for z.Next() {
		result[i] = z.Value()
		i++
	}
	return result
}

func Zip[T any](entries ...[]T) (iterator ListIterator[[]T], err error) {
	entryCount := len(entries)
	if entryCount < 2 {
		err = common.ErrUnexpectedParamCount
		return
	}
	length := len(entries[0])
	for i := 1; i < entryCount; i++ {
		length = min(length, len(entries[i]))
	}
	iterator = newZipListIterator(entries, length)
	return
}

func ZipLongest[T any](entries ...[]T) (iterator ListIterator[[]T], err error) {
	entryCount := len(entries)
	if entryCount < 2 {
		err = common.ErrUnexpectedParamCount
		return
	}
	length := len(entries[0])
	for i := 1; i < entryCount; i++ {
		length = max(length, len(entries[i]))
	}
	iterator = newZipListIterator(entries, length)
	return
}
