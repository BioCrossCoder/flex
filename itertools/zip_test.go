package itertools

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	convey.Convey("Call Zip on [string|array]", t, func() {
		str := "hello"
		arr := []int{1, 2, 3}
		expected := [][2]any{
			{1, "h"},
			{2, "e"},
			{3, "l"},
		}
		result, err := ZipResult(arr, str)
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})
}

func TestZipIterator(t *testing.T) {
	convey.Convey("zip all", t, func() {
		entry1 := []any{1, 2, 3}
		entry2 := []any{"a", "b", "c"}
		expected := [][2]any{
			{1, "a"},
			{2, "b"},
		}
		iterator := NewZipIterator(entry1, entry2, 2)
		result := iterator.Pour()
		assert.Equal(t, expected, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}
