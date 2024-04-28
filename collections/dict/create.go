package dict

func (d Dict) Copy() Dict {
	backup := make(Dict, d.Size())
	for k, v := range d {
		backup.Set(k, v)
	}
	return backup
}

func FromEntries(entries ...[2]any) Dict {
	d := make(Dict, len(entries))
	for _, entry := range entries {
		d.Set(entry[0], entry[1])
	}
	return d
}
