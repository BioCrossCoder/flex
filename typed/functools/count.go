package functools

import (
	"github.com/biocrosscoder/flex/common"
)

func Count[T any](entry []T, value T) (count int) {
	for _, v := range entry {
		if common.Equal(v, value) {
			count++
		}
	}
	return
}

func CountBy[T any](entry []T, condition func(T) bool) (count int) {
	for _, v := range entry {
		if condition(v) {
			count++
		}
	}
	return
}
