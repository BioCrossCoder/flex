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
