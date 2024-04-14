package functools

import "flex/itertools"

func Map(handler, entry any) (output any, err error) {
	iterator, err := itertools.Map(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
