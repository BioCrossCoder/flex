// Package sortedcontainers provides sotred data structures.
package sortedcontainers

import (
	"cmp"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
)

// SortedDict is a data structure that represents a dictionary with sorted keys.
type SortedDict[K cmp.Ordered, V any] struct {
	dict.Dict[K, V]
	sequence sortedlist.SortedList[K]
}

// NewSortedDict creates a new SortedDict instance with the provided comparison function
// and initializes it with the given dictionary or an empty one if not provided.
func NewSortedDict[K cmp.Ordered, V any](f func(a, b K) int, src dict.Dict[K, V]) *SortedDict[K, V] {
	if src == nil {
		src = make(dict.Dict[K, V])
	}
	if f == nil {
		f = sortedlist.AscendOrder
	}
	return &SortedDict[K, V]{
		src,
		*sortedlist.NewSortedList(f, src.Keys()...),
	}
}

// Clear removes all key-value pairs from the SortedDict.
func (d *SortedDict[K, V]) Clear() *SortedDict[K, V] {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

// Set adds or updates a key-value pair in the SortedDict and ensures that the key is maintained in the sorting sequence.
func (d *SortedDict[K, V]) Set(key K, value V) *SortedDict[K, V] {
	if !d.Has(key) {
		_ = d.sequence.Insert(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

// Delete removes the key-value pair corresponding to the specified key from the SortedDict.
func (d *SortedDict[K, V]) Delete(key K) bool {
	if d.Dict.Delete(key) {
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

// Pop removes and returns the value associated with the specified key from the SortedDict.
func (d *SortedDict[K, V]) Pop(key K, args ...V) (value V, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

// PopItem removes and returns the last key-value pair from the SortedDict based on the sorted keys.
func (d *SortedDict[K, V]) PopItem() (key K, value V, err error) {
	key, err = d.sequence.Tail()
	if err != nil {
		err = common.ErrEmptyDict
		return
	}
	value, _ = d.Pop(key)
	return
}

// Update adds or updates multiple key-value pairs from another dictionary into the SortedDict.
func (d *SortedDict[K, V]) Update(another dict.Dict[K, V]) *SortedDict[K, V] {
	for k, v := range another {
		_ = d.Set(k, v)
	}
	return d
}

// Keys returns a slice containing all the keys in the SortedDict in sorted order.
func (d SortedDict[K, V]) Keys() []K {
	return d.sequence.ToArray()
}

// Values returns a slice containing all the values in the SortedDict in the corresponding key order.
func (d SortedDict[K, V]) Values() []V {
	length := d.Size()
	values := make([]V, length)
	for i := 0; i < length; i++ {
		key, _ := d.KeyAt(i)
		values[i] = d.Get(key)
	}
	return values
}

// Items returns an array containing all the key-value pairs in the SortedDict in the sorted key order.
func (d SortedDict[K, V]) Items() []*dict.DictItem[K, V] {
	length := d.Size()
	items := make([]*dict.DictItem[K, V], length)
	for i := 0; i < length; i++ {
		key, _ := d.KeyAt(i)
		value := d.Get(key)
		items[i] = &dict.DictItem[K, V]{key, value}
	}
	return items
}

// Copy creates and returns a new SortedDict instance that is an exact copy of the original SortedDict.
func (d SortedDict[K, V]) Copy() SortedDict[K, V] {
	return SortedDict[K, V]{
		d.Dict.Copy(),
		d.sequence.Copy(),
	}
}

// Equal compares the SortedDict with another SortedDict to check if they are equal in terms of contents and order.
func (d SortedDict[K, V]) Equal(another SortedDict[K, V]) bool {
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
		if !common.Equal(d.Get(key1), another.Get(key2)) {
			return false
		}
	}
	return true
}

// KeyAt returns the key at the specified index in the sorted sequence of keys stored in the SortedDict.
func (d SortedDict[K, V]) KeyAt(index int) (K, error) {
	return d.sequence.At(index)
}

// IndexOf returns the index of the specified key in the sorted sequence of keys stored in the SortedDict.
func (d SortedDict[K, V]) IndexOf(key K) int {
	return d.sequence.IndexOf(key)
}
