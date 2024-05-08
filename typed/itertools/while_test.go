package itertools

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestWhile(t *testing.T) {
	f := func(x int) bool {
		return x > 0
	}
	entry := []int{1, 2, 3, 0, 4, 5, 6}
	convey.Convey("drop while", t, func() {
		iter := DropWhile(f, entry)
		assert.Equal(t, []int{0, 4, 5, 6}, iter.Pour())
	})
	convey.Convey("take while", t, func() {
		iter := TakeWhile(f, entry)
		assert.Equal(t, []int{1, 2, 3}, iter.Pour())
	})
}
