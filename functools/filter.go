package functools

import "flex/itertools"

func Filter(handler, entry any) (output any, err error) {
	iterator, err := itertools.Filter(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
