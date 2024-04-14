package functools

import "flex/itertools"

func Reduce(handler, entry any) (output any, err error) {
	iterator, err := itertools.Accumulate(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
