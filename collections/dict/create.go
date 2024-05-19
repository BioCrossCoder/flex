package dict

import "github.com/biocrosscoder/flex/common"

// Copy method creates a shallow copy of the dictionary.
func (d Dict) Copy() Dict {
	backup := make(Dict, common.GetMapInitialCapacity(d.Size()))
	for k, v := range d {
		backup.Set(k, v)
	}
	return backup
}

// FromEntries method creates a new dictionary from the provided key-value pairs.
func FromEntries(entries ...[2]any) Dict {
	d := make(Dict, common.GetMapInitialCapacity(len(entries)))
	for _, entry := range entries {
		d.Set(entry[0], entry[1])
	}
	return d
}
