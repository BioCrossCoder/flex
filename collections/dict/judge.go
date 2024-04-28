package dict

func (d Dict) Has(key any) bool {
	_, ok := d[key]
	return ok
}

func (d Dict) Empty() bool {
	return d.Size() == 0
}
