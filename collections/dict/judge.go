package dict

import (
	"flex/collections/set"
	"flex/common"
)

func (d Dict) Has(key any) bool {
	_, ok := d[key]
	return ok
}

func (d Dict) Empty() bool {
	return d.Size() == 0
}

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
