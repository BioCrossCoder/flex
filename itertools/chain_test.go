package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChain(t *testing.T) {
	convey.Convey("link iterable entries together", t, func() {
		entry1 := []int{1, 2, 3}
		entry2 := [3]any{1, "a", true}
		entry3 := "hello"
		result, err := Chain(entry1, entry2, entry3)
		assert.Nil(t, err)
		assert.Equal(t, []any{1, 2, 3, 1, "a", true, "h", "e", "l", "l", "o"}, result.Pour())
	})
}
