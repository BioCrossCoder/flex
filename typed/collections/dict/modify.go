package dict

import (
	"github.com/biocrosscoder/flex/common"
	"maps"
)

// Clear removes all key-value pairs from the dictionary.
func (d *Dict[K, V]) Clear() *Dict[K, V] {
	*d = make(Dict[K, V])
	return d
}

// Set adds or updates the key with the specified value in the dictionary.
func (d *Dict[K, V]) Set(key K, value V) *Dict[K, V] {
	(*d)[key] = value
	return d
}

// Delete removes the key-value pair with the specified key from the dictionary.
// It returns true if the key exists in the dictionary and false otherwise.
func (d *Dict[K, V]) Delete(key K) bool {
	ok := d.Has(key)
	if ok {
		delete(*d, key)
	}
	return ok
}

// Pop removes the key and returns its value.
// If the key exists, its value is returned. If it doesn't exist and a default value is provided, the default value is returned.
// If the key doesn't exist and no default value is provided, an error is returned.
func (d *Dict[K, V]) Pop(key K, args ...V) (value V, err error) {
	var defaultVal V
	argCount := len(args)
	if !d.Has(key) && argCount == 0 {
		err = common.ErrKeyNotFound
		return
	}
	if argCount >= 1 {
		defaultVal = args[0]
	}
	value = d.Get(key, defaultVal)
	_ = d.Delete(key)
	return
}

// PopItem removes and returns an arbitrary key-value pair from the dictionary.
func (d *Dict[K, V]) PopItem() (key K, value V, err error) {
	if d.Empty() {
		err = common.ErrEmptyDict
		return
	}
	for k, v := range *d {
		key = k
		value = v
		_ = d.Delete(k)
		break
	}
	return
}

// Update merges the key-value pairs from another dictionary into the current dictionary.
func (d *Dict[K, V]) Update(another Dict[K, V]) *Dict[K, V] {
	maps.Copy(*d, another)
	return d
}
