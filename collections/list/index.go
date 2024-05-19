package list

import "github.com/biocrosscoder/flex/common"

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

func IsIndexValid(index, length int) (err error) {
	if index < 0 || index >= length {
		err = common.ErrOutOfRange
	}
	return
}
