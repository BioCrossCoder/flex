package dict

import "github.com/biocrosscoder/flex/common"

// Clear method removes all key-value pairs from the dictionary and returns the empty dictionary.
func (d *Dict) Clear() *Dict {
	*d = make(Dict)
	return d
}

// Set method sets the value for the given key in the dictionary and returns the updated dictionary.
func (d *Dict) Set(key, value any) *Dict {
	(*d)[key] = value
	return d
}

// Delete method removes the key-value pair for the given key from the dictionary and returns true if the key existed, false otherwise.
func (d *Dict) Delete(key any) bool {
	ok := d.Has(key)
	if ok {
		delete(*d, key)
	}
	return ok
}

// Pop method removes the key and its value from the dictionary. It returns the value for the given key, and an error if the key doesn't exist and no default value is provided.
func (d *Dict) Pop(key any, args ...any) (value any, err error) {
	var defaultVal any
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

// PopItem method removes and returns an arbitrary key-value pair from the dictionary, along with an error if the dictionary is empty.
func (d *Dict) PopItem() (key, value any, err error) {
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

// Update method updates the dictionary with key-value pairs from another dictionary and returns the updated dictionary.
func (d *Dict) Update(another Dict) *Dict {
	count1 := d.Size()
	count2 := another.Size()
	if common.WillReHash(count1, count2) {
		capacity := common.GetMapInitialCapacity(count1 + count2)
		newDict := make(Dict, capacity)
		for k, v := range *d {
			newDict.Set(k, v)
		}
		*d = newDict
	}
	for k, v := range another {
		d.Set(k, v)
	}
	return d
}
