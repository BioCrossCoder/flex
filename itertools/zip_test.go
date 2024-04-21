package itertools

import (
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
