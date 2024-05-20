// Package dict provides a flexible hash map implementation.
package dict

import (
	"github.com/biocrosscoder/flex/common"
	"maps"
)

// Dict is a generic type representing a dictionary with keys of type K and values of type V.
type Dict[K comparable, V any] map[K]V

// Size returns the number of key-value pairs in the dictionary.
func (d Dict[K, V]) Size() int {
	return len(d)
}

// Get retrieves the value associated with the specified key, and returns a default value if the key is not present.
func (d Dict[K, V]) Get(key K, defaultValue ...V) (value V) {
	if d.Has(key) {
		value = d[key]
	} else if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return
}

// Keys returns a slice containing all the keys in the dictionary.
func (d Dict[K, V]) Keys() []K {
	keys := make([]K, d.Size())
	i := 0
	for k := range d {
		keys[i] = k
		i++
	}
	return keys
}

// Values returns a slice containing all the values in the dictionary.
func (d Dict[K, V]) Values() []V {
	values := make([]V, d.Size())
	i := 0
	for _, v := range d {
		values[i] = v
		i++
	}
	return values
}

// DictItem represents a key-value pair in the dictionary.
type DictItem[K comparable, V any] struct {
	Key   K
	Value V
}

// Items returns a slice of pointers to DictItem, each representing a key-value pair in the dictionary.
func (d Dict[K, V]) Items() []*DictItem[K, V] {
	items := make([]*DictItem[K, V], d.Size())
	i := 0
	for k, v := range d {
		items[i] = &DictItem[K, V]{k, v}
		i++
	}
	return items
}

// Copy creates a shallow copy of the dictionary.
func (d Dict[K, V]) Copy() Dict[K, V] {
	return maps.Clone(d)
}

// Has checks if the dictionary contains the specified key.
func (d Dict[K, V]) Has(key K) bool {
	_, ok := d[key]
	return ok
}

// Empty checks if the dictionary is empty.
func (d Dict[K, V]) Empty() bool {
	return d.Size() == 0
}

// Equal compares two dictionaries for equality based on their keys and values.
func (d Dict[K, V]) Equal(another Dict[K, V]) bool {
	return maps.EqualFunc(d, another, func(a, b V) bool {
		return common.Equal(a, b)
	})
}
