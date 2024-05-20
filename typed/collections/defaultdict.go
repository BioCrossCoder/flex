package collections

import (
	"encoding/json"
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/dict"
)

// DefaultDict is a wrapper around dict.Dict with a default value for keys which are not present.
type DefaultDict[K comparable, V any] struct {
	dict.Dict[K, V]
	builder func() V
}

// NewDefaultDict creates a new DefaultDict with the given items and default value.
func NewDefaultDict[K comparable, V any](items dict.Dict[K, V], defaultVal V) *DefaultDict[K, V] {
	if items == nil {
		items = make(dict.Dict[K, V])
	}
	return &DefaultDict[K, V]{
		items.Copy(),
		func() V {
			return defaultVal
		},
	}
}

// Get retrieves the value for the given key; creates and sets default value if key is not present.
func (d *DefaultDict[K, V]) Get(key K) (value V) {
	if d.Has(key) {
		return d.Dict.Get(key)
	}
	value = d.builder()
	_ = d.Set(key, value)
	return
}

// Pop removes the value for the given key and returns it; creates and sets default value if key is not present.
func (d *DefaultDict[K, V]) Pop(key K) (value V) {
	value, _ = d.Dict.Pop(key, d.builder())
	return
}

// Copy creates a copy of the DefaultDict.
func (d DefaultDict[K, V]) Copy() DefaultDict[K, V] {
	return DefaultDict[K, V]{
		d.Dict.Copy(),
		d.builder,
	}
}

// Clear removes all items from the DefaultDict.
func (d *DefaultDict[K, V]) Clear() *DefaultDict[K, V] {
	_ = d.Dict.Clear()
	return d
}

// Set sets the value for the given key in the DefaultDict.
func (d *DefaultDict[K, V]) Set(key K, value V) *DefaultDict[K, V] {
	_ = d.Dict.Set(key, value)
	return d
}

// Update updates the DefaultDict with the items from another DefaultDict.
func (d *DefaultDict[K, V]) Update(another DefaultDict[K, V]) *DefaultDict[K, V] {
	_ = d.Dict.Update(another.Dict)
	return d
}

// SetDefault sets the default value for keys not present in the DefaultDict.
func (d *DefaultDict[K, V]) SetDefault(value V) *DefaultDict[K, V] {
	d.builder = func() V {
		return value
	}
	return d
}

// Equal checks whether two DefaultDicts are equal.
func (d DefaultDict[K, V]) Equal(another DefaultDict[K, V]) bool {
	return d.Dict.Equal(another.Dict) && common.Equal(d.builder(), another.builder())
}

// String returns the string representation of the DefaultDict.
func (d DefaultDict[K, V]) String() string {
	return fmt.Sprint(d.Dict)
}

// MarshalJSON returns the JSON encoding of the DefaultDict.
func (d DefaultDict[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Dict)
}

// UnmarshalJSON sets the DefaultDict to the value represented by the JSON encoding.
func (d *DefaultDict[K, V]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &d.Dict)
}
