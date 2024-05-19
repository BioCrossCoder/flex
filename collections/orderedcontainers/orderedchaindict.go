// Package orderedcontainers provides ordered Dict and Set implementations.
package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/biocrosscoder/flex/collections/dict"
	"github.com/biocrosscoder/flex/collections/linkedlist"
	"github.com/biocrosscoder/flex/common"
	"strings"
)

// OrderedChainDict represents a dictionary with ordered keys, which is based on a Dict and a LinkedList.
type OrderedChainDict struct {
	dict.Dict
	sequence *linkedlist.LinkedList
}

// NewOrderedChainDict creates a new OrderedChainDict with the given key-value entries.
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

// Clear clears the OrderedChainDict and returns the modified instance.
func (d *OrderedChainDict) Clear() *OrderedChainDict {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

// Set sets the key-value pair in the OrderedChainDict and returns the modified instance.
func (d *OrderedChainDict) Set(key, value any) *OrderedChainDict {
	if !d.Has(key) {
		_ = d.sequence.Append(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

// Delete deletes the key from the OrderedChainDict and returns true if successful, false otherwise.
func (d *OrderedChainDict) Delete(key any) bool {
	if d.Has(key) {
		_ = d.Dict.Delete(key)
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

// Pop removes the key and returns the corresponding value from the OrderedChainDict.
func (d *OrderedChainDict) Pop(key any, args ...any) (value any, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

// PopItem removes and returns the last key-value pair from the OrderedChainDict.
func (d *OrderedChainDict) PopItem() (key, value any, err error) {
	if d.Empty() {
		err = common.ErrEmptyDict
		return
	}
	key, _ = d.sequence.Tail()
	value, _ = d.Pop(key)
	return
}

// Update merges an OrderedChainDict into the current OrderedChainDict and returns the modified instance.
func (d *OrderedChainDict) Update(another OrderedChainDict) *OrderedChainDict {
	for _, key := range another.Keys() {
		_ = d.Set(key, another.Get(key))
	}
	return d
}

// Keys returns the keys of the OrderedChainDict as an array.
func (d OrderedChainDict) Keys() []any {
	return d.sequence.ToArray()
}

// Values returns the values of the OrderedChainDict as an array.
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

// Items returns the key-value pairs of the OrderedChainDict as an array of tuples.
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

// Copy creates a deep copy of the OrderedChainDict.
func (d OrderedChainDict) Copy() OrderedChainDict {
	newSeq := d.sequence.Copy()
	return OrderedChainDict{
		d.Dict.Copy(),
		&newSeq,
	}
}

// Equal checks if two OrderedChainDict instances are equal.
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
		if !common.Equal(d.Get(key1), another.Get(key2)) {
			return false
		}
	}
	return true
}

// KeyAt returns the key at the specified index in the OrderedChainDict.
func (d OrderedChainDict) KeyAt(index int) (any, error) {
	return d.sequence.At(index)
}

// IndexOf returns the index of the specified key in the OrderedChainDict.
func (d OrderedChainDict) IndexOf(key any) int {
	return d.sequence.IndexOf(key)
}

// String returns the string representation of the OrderedChainDict.
func (d OrderedChainDict) String() string {
	items := make([]string, d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key any) any {
		items[i] = fmt.Sprintf("%v:%v", key, d.Get(key))
		i++
		return key
	})
	return "map[" + strings.Join(items, " ") + "]"
}

// MarshalJSON returns the JSON encoding of the OrderedChainDict.
func (d OrderedChainDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Items())
}

// UnmarshalJSON sets the OrderedChainDict from its JSON representation.
func (d *OrderedChainDict) UnmarshalJSON(data []byte) (err error) {
	var items [][2]any
	err = json.Unmarshal(data, &items)
	if err != nil {
		return
	}
	*d = *NewOrderedChainDict(items...)
	return
}
