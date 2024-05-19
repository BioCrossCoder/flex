package list

import "github.com/biocrosscoder/flex/common"

// SliceIndex returns the valid index within the slice based on the given index, length, and accessibility flag.
func SliceIndex(index, length int, accessible bool) int {
	if index < 0 {
		index += length
	}
	if index < 0 {
		index = -1
		if accessible {
			index++
		}
	}
	if index >= length {
		index = length
		if accessible {
			index--
		}
	}
	return index
}

// ParseIndex returns the valid index within the slice based on the given index and length.
func ParseIndex(index, length int) int {
	if index < 0 {
		index += length
		if index < 0 {
			return 0
		}
	} else if index > length {
		return length
	}
	return index
}

// IsIndexValid checks if the index is valid within the slice based on the given index and length.
func IsIndexValid(index, length int) (err error) {
	if index < 0 || index >= length {
		err = common.ErrOutOfRange
	}
	return
}
