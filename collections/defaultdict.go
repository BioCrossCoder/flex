package collections

import (
	"encoding/json"
	"fmt"
	"github.com/biocrosscoder/flex/collections/dict"
	"github.com/biocrosscoder/flex/common"
)

// DefaultDict is a Dict with a default value.
type DefaultDict struct {
	dict.Dict
	builder func() any
}

// NewDefaultDict creates a new DefaultDict with the given items and default value.
func NewDefaultDict(items dict.Dict, defaultVal any) *DefaultDict {
	if items == nil {
		items = make(dict.Dict)
	}
	return &DefaultDict{
		items.Copy(),
		func() any {
			return defaultVal
		},
	}
}

// Get retrieves the value for the specified key. If the key does not exist, it returns the default value.
func (d *DefaultDict) Get(key any) (value any) {
	if d.Has(key) {
		return d.Dict.Get(key)
	}
	value = d.builder()
	_ = d.Set(key, value)
	return
}

// Pop removes the specified key and returns its corresponding value. If the key does not exist, it returns the default value.
func (d *DefaultDict) Pop(key any) (value any) {
	value, _ = d.Dict.Pop(key, d.builder())
	return
}

// Copy creates a copy of the DefaultDict.
func (d DefaultDict) Copy() DefaultDict {
	return DefaultDict{
		d.Dict.Copy(),
		d.builder,
	}
}

// Clear removes all items from the DefaultDict.
func (d *DefaultDict) Clear() *DefaultDict {
	_ = d.Dict.Clear()
	return d
}

// Set sets the value for the specified key in the DefaultDict.
func (d *DefaultDict) Set(key, value any) *DefaultDict {
	_ = d.Dict.Set(key, value)
	return d
}

// Update updates the DefaultDict with the key-value pairs from another DefaultDict.
func (d *DefaultDict) Update(another DefaultDict) *DefaultDict {
	_ = d.Dict.Update(another.Dict)
	return d
}

// SetDefault sets the default value for the DefaultDict.
func (d *DefaultDict) SetDefault(value any) *DefaultDict {
	d.builder = func() any {
		return value
	}
	return d
}

// Equal checks if two DefaultDicts are equal.
func (d DefaultDict) Equal(another DefaultDict) bool {
	return d.Dict.Equal(another.Dict) && common.Equal(d.builder(), another.builder())
}

// String returns the string representation of the DefaultDict.
func (d DefaultDict) String() string {
	return fmt.Sprint(d.Dict)
}

// MarshalJSON returns the JSON encoding of the DefaultDict.
func (d DefaultDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Dict)
}

// UnmarshalJSON sets the value of the DefaultDict based on the input JSON data.
func (d *DefaultDict) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &d.Dict)
}
