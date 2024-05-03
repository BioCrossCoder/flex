package orderedcontainers

import (
	"flex/collections/arraylist"
	"flex/collections/dict"
	"flex/common"
)

type OrderedDict struct {
	dict.Dict
	sequence arraylist.ArrayList
}

func NewOrderedDict(entries ...[2]any) *OrderedDict {
	elements := dict.FromEntries()
	sequence := arraylist.Of()
	for _, entry := range entries {
		key := entry[0]
		value := entry[1]
		if elements.Has(key) {
			_ = sequence.Remove(key)
		}
		_ = elements.Set(key, value)
		_ = sequence.Push(key)
	}
	return &OrderedDict{elements, sequence}
}

func (d *OrderedDict) Clear() *OrderedDict {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

func (d *OrderedDict) Set(key, value any) *OrderedDict {
	if !d.Has(key) {
		_ = d.sequence.Push(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

func (d *OrderedDict) Delete(key any) bool {
	if d.Has(key) {
		_ = d.Dict.Delete(key)
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

func (d *OrderedDict) Pop(key any, args ...any) (value any, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

func (d *OrderedDict) PopItem() (key, value any, err error) {
	if d.Empty() {
		err = common.ErrEmptyDict
		return
	}
	key, _ = d.sequence.Tail()
	value, _ = d.Pop(key)
	return
}

func (d *OrderedDict) Update(another OrderedDict) *OrderedDict {
	for _, key := range another.Keys() {
		_ = d.Set(key, another.Get(key))
	}
	return d
}

func (d OrderedDict) Keys() arraylist.ArrayList {
	return d.sequence
}

func (d OrderedDict) Values() arraylist.ArrayList {
	values := make(arraylist.ArrayList, d.Size())
	for i, key := range d.sequence {
		values[i] = d.Get(key)
	}
	return values
}

func (d OrderedDict) Items() arraylist.ArrayList {
	items := make(arraylist.ArrayList, d.Size())
	for i, key := range d.sequence {
		items[i] = [2]any{key, d.Get(key)}
	}
	return items
}

func (d OrderedDict) Copy() OrderedDict {
	return OrderedDict{
		d.Dict.Copy(),
		d.sequence.Copy(),
	}
}

func (d OrderedDict) Equal(another OrderedDict) bool {
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

func (d OrderedDict) KeyAt(index int) (any, error) {
	return d.sequence.At(index)
}

func (d OrderedDict) IndexOf(key any) int {
	return d.sequence.IndexOf(key)
}
