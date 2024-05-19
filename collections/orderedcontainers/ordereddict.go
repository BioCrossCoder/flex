package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/biocrosscoder/flex/collections/arraylist"
	"github.com/biocrosscoder/flex/collections/dict"
	"github.com/biocrosscoder/flex/common"
	"strings"
)

// OrderedDict represents an ordered dictionary, which is a collection of key-value pairs
// with a specific order based on the sequence of insertion, which is based on a Dict and an ArrayList.
type OrderedDict struct {
	dict.Dict
	sequence arraylist.ArrayList
}

// NewOrderedDict creates a new ordered dictionary with the given key-value entries.
// It returns a pointer to the created OrderedDict.
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

// Clear clears all the elements in the ordered dictionary and the insertion order sequence.
// It returns the updated OrderedDict.
func (d *OrderedDict) Clear() *OrderedDict {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

// Set sets the key-value pair in the ordered dictionary.
// It maintains the insertion order of the keys and returns the updated OrderedDict.
func (d *OrderedDict) Set(key, value any) *OrderedDict {
	if !d.Has(key) {
		_ = d.sequence.Push(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

// Delete deletes the key-value pair with the specified key from the ordered dictionary.
// It returns true if the key existed and was deleted; otherwise, it returns false.
func (d *OrderedDict) Delete(key any) bool {
	if d.Dict.Delete(key) {
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

// Pop removes the key-value pair with the specified key from the ordered dictionary.
// It also returns the value of the removed key-value pair and an error if applicable.
func (d *OrderedDict) Pop(key any, args ...any) (value any, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

// PopItem removes and returns the last key-value pair from the ordered dictionary.
// It also returns an error if the dictionary is empty.
func (d *OrderedDict) PopItem() (key, value any, err error) {
	key, err = d.sequence.Tail()
	if err != nil {
		err = common.ErrEmptyDict
		return
	}
	value, _ = d.Pop(key)
	return
}

// Update updates the ordered dictionary with the key-value pairs from another OrderedDict.
// It returns the updated OrderedDict.
func (d *OrderedDict) Update(another OrderedDict) *OrderedDict {
	for _, key := range another.Keys() {
		_ = d.Set(key, another.Get(key))
	}
	return d
}

// Keys returns a copy of all keys in the ordered dictionary based on the insertion order.
func (d OrderedDict) Keys() []any {
	return d.sequence.Copy()
}

// Values returns a slice of all values in the ordered dictionary based on the insertion order of keys.
func (d OrderedDict) Values() []any {
	values := make([]any, d.Size())
	for i, key := range d.sequence {
		values[i] = d.Get(key)
	}
	return values
}

// Items returns a slice of all key-value pairs in the ordered dictionary based on the insertion order of keys.
func (d OrderedDict) Items() [][2]any {
	items := make([][2]any, d.Size())
	for i, key := range d.sequence {
		items[i] = [2]any{key, d.Get(key)}
	}
	return items
}

// Copy returns a copy of the ordered dictionary with the same key-value pairs and insertion order.
func (d OrderedDict) Copy() OrderedDict {
	return OrderedDict{
		d.Dict.Copy(),
		d.sequence.Copy(),
	}
}

// Equal checks if the ordered dictionary is equal to another OrderedDict in terms of key-value pairs and insertion order.
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
		if !common.Equal(d.Get(key1), another.Get(key2)) {
			return false
		}
	}
	return true
}

// KeyAt returns the key at the specified index in the insertion order sequence.
func (d OrderedDict) KeyAt(index int) (any, error) {
	return d.sequence.At(index)
}

// IndexOf returns the index of the specified key in the insertion order sequence.
func (d OrderedDict) IndexOf(key any) int {
	return d.sequence.IndexOf(key)
}

// String returns a string representation of the ordered dictionary in the format of a map.
func (d OrderedDict) String() string {
	items := make([]string, d.Size())
	for i, key := range d.sequence {
		items[i] = fmt.Sprintf("%v:%v", key, d.Get(key))
	}
	return "map[" + strings.Join(items, " ") + "]"
}

// MarshalJSON returns the JSON encoding of the ordered dictionary.
func (d OrderedDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Items())
}

// UnmarshalJSON sets the ordered dictionary from its JSON encoding.
func (d *OrderedDict) UnmarshalJSON(data []byte) (err error) {
	var items [][2]any
	err = json.Unmarshal(data, &items)
	if err != nil {
		return
	}
	*d = *NewOrderedDict(items...)
	return
}
