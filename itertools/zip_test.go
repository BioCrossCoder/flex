package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZip(t *testing.T) {
	convey.Convey("Call Zip on [string|array|slice]", t, func() {
		arr := [3]int{1, 2, 3}
		str := "hello"
		sli := []any{1, "a", true}
		expected := []any{
			[]any{1, "h", 1},
			[]any{2, "e", "a"},
			[]any{3, "l", true},
		}
		result, err := Zip(arr, str, sli)
		assert.Nil(t, err)
		assert.Equal(t, expected, result.Pour())
	})
}

func TestZipLongest(t *testing.T) {
	convey.Convey("Call ZipLongest on [string|array|slice]", t, func() {
		arr := [3]int{1, 2, 3}
		str := "hello"
		sli := []any{1, "a", true}
		expected := []any{
			[]any{1, "h", 1},
			[]any{2, "e", "a"},
			[]any{3, "l", true},
			[]any{nil, "l", nil},
			[]any{nil, "o", nil},
		}
		result, err := ZipLongest(arr, str, sli)
		assert.Nil(t, err)
		assert.Equal(t, expected, result.Pour())
	})
}

func ExampleZip() {
	arr := [3]int{1, 2, 3}
	str := "hello"
	sli := []any{1, "a", true}
	iter, _ := Zip(arr, str, sli)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [1 h 1]
	// [2 e a]
	// [3 l true]
}

func ExampleZipLongest() {
	arr := [3]int{1, 2, 3}
	str := "hello"
	sli := []any{1, "a", true}
	iter, _ := ZipLongest(arr, str, sli)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [1 h 1]
	// [2 e a]
	// [3 l true]
	// [<nil> l <nil>]
	// [<nil> o <nil>]
}
