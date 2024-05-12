// Package functools provides functional programming tools.
package functools

import "github.com/biocrosscoder/flex/itertools"

// Reduce applies a function to an input sequence and returns the accumulated result.
func Reduce(handler, entry any) (output any, err error) {
	iterator, err := itertools.Accumulate(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
