package functools

import "github.com/biocrosscoder/flex/itertools"

// Filter applies a filter function to an entry and the returned output only contains the entries that pass the filter function.
func Filter(filter, entry any) (output any, err error) {
	iterator, err := itertools.Filter(filter, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
