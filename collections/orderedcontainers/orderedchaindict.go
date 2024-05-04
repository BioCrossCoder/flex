package orderedcontainers

import (
	"flex/collections/dict"
	"flex/collections/linkedlist"
	"flex/common"
)

type OrderedChainDict struct {
	dict.Dict
	sequence *linkedlist.LinkedList
}

func NewOrderedChainDict(entries ...[2]any) *OrderedChainDict {
	elements := dict.FromEntries()
	sequence := linkedlist.NewLinkedList()
	for _, entry := range entries {
		key := entry[0]
		value := entry[1]
		if elements.Has(key) {
			_ = sequence.Remove(key)
		}
		_ = elements.Set(key, value)
		_ = sequence.Append(key)
	}
	return &OrderedChainDict{elements, sequence}
}

func (d *OrderedChainDict) Clear() *OrderedChainDict {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

func (d *OrderedChainDict) Set(key, value any) *OrderedChainDict {
	if !d.Has(key) {
		_ = d.sequence.Append(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

func (d *OrderedChainDict) Delete(key any) bool {
	if d.Has(key) {
		_ = d.Dict.Delete(key)
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

func (d *OrderedChainDict) Pop(key any, args ...any) (value any, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

func (d *OrderedChainDict) PopItem() (key, value any, err error) {
	if d.Empty() {
		err = common.ErrEmptyDict
		return
	}
	key, _ = d.sequence.Tail()
	value, _ = d.Pop(key)
	return
}

func (d *OrderedChainDict) Update(another OrderedChainDict) *OrderedChainDict {
	for _, key := range another.Keys() {
		_ = d.Set(key, another.Get(key))
	}
	return d
}

func (d OrderedChainDict) Keys() []any {
	return d.sequence.ToArray()
}

func (d OrderedChainDict) Values() []any {
	values := make([]any, d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key any) any {
		values[i] = d.Get(key)
		i++
		return key
	})
	return values
}

func (d OrderedChainDict) Items() [][2]any {
	items := make([][2]any, d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key any) any {
		items[i] = [2]any{key, d.Get(key)}
		i++
		return key
	})
	return items
}

func (d OrderedChainDict) Copy() OrderedChainDict {
	newSeq := d.sequence.Copy()
	return OrderedChainDict{
		d.Dict.Copy(),
		&newSeq,
	}
}

func (d OrderedChainDict) Equal(another OrderedChainDict) bool {
	if d.Size() != another.Size() {
		return false
	}
	keys1 := d.Keys()
	keys2 := another.Keys()
	for i := 0; i < d.Size(); i++ {
		key1 := keys1[i]
		key2 := keys2[i]
		if key1 != key2 {
			return false
		}
		if d.Get(key1) != another.Get(key2) {
			return false
		}
	}
	return true
}

func (d OrderedChainDict) KeyAt(index int) (any, error) {
	return d.sequence.At(index)
}

func (d OrderedChainDict) IndexOf(key any) int {
	return d.sequence.IndexOf(key)
}
