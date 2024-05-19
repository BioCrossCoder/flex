package dict

import (
	"github.com/biocrosscoder/flex/collections/set"
	"github.com/biocrosscoder/flex/common"
)

// Has checks if the key exists in the dictionary and returns true if found, otherwise returns false.
func (d Dict) Has(key any) bool {
	_, ok := d[key]
	return ok
}

// Empty returns true if the dictionary is empty, otherwise returns false.
func (d Dict) Empty() bool {
	return d.Size() == 0
}

// Equal checks if the current dictionary is equal to the provided dictionary and returns true if they are equal, otherwise returns false.
func (d Dict) Equal(another Dict) bool {
	keys1 := set.Of(d.Keys()...)
	keys2 := set.Of(another.Keys()...)
	if !keys1.Equal(keys2) {
		return false
	}
	for k, v := range d {
		if !common.Equal(v, another[k]) {
			return false
		}
	}
	return true
}
